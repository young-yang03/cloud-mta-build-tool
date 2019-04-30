package artifacts

import (
	"fmt"
	"os"

	"github.com/SAP/cloud-mta-build-tool/internal/commands"
	"github.com/SAP/cloud-mta/mta"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {

	var _ = Describe("ExecuteProjectBuild", func() {
		It("Sanity - post phase", func() {
			err := ExecuteProjectBuild(getTestPath("mtahtml5"), "dev", "post", os.Getwd)
			Ω(err).Should(BeNil())
		})
		It("wrong phase", func() {
			err := ExecuteProjectBuild(getTestPath("mta"), "dev", "wrong phase", os.Getwd)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal(`the "wrong phase" phase of mta project build is invalid; supported phases: "pre", "post"`))
		})
		It("wrong location", func() {
			err := ExecuteProjectBuild(getTestPath("mta"), "xx", "pre", func() (string, error) {
				return "", fmt.Errorf("error")
			})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("failed to initialize the location when validating descriptor:"))
		})
		It("mta.yaml not found", func() {
			err := ExecuteProjectBuild(getTestPath("mta1"), "dev", "pre", os.Getwd)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("failed to read"))
		})
		It("Sanity - custom builder", func() {
			err := ExecuteProjectBuild(getTestPath("mta"), "dev", "pre", os.Getwd)
			Ω(err.Error()).Should(ContainSubstring(`"command1"`))
			Ω(err.Error()).Should(ContainSubstring("failed"))
		})
	})

	var _ = Describe("getProjectBuilderCommands", func() {
		It("Builder and commands defined", func() {
			projectBuild := mta.ProjectBuilder{
				Builder:  "npm",
				Commands: []string{"abc"},
			}

			cmds, err := getProjectBuilderCommands(projectBuild)
			Ω(err).Should(Succeed())
			Ω(len(cmds.Command)).Should(Equal(2))
			Ω(cmds.Command[0]).Should(Equal("npm install"))
			Ω(cmds.Command[1]).Should(Equal("npm prune --production"))
		})
	})

	var _ = Describe("execProjectBuilders", func() {
		It("Before Defined with nothing to execute", func() {
			builders := []mta.ProjectBuilder{}
			projectBuild := mta.ProjectBuild{
				BeforeAll: builders,
			}
			oMta := mta.MTA{
				BuildParams: &projectBuild,
			}
			Ω(execProjectBuilders(&oMta, "pre")).Should(Succeed())
		})
		It("After Defined with nothing to execute", func() {
			builders := []mta.ProjectBuilder{}
			projectBuild := mta.ProjectBuild{
				AfterAll: builders,
			}
			oMta := mta.MTA{
				BuildParams: &projectBuild,
			}
			Ω(execProjectBuilders(&oMta, "post")).Should(Succeed())
		})
		It("Before Defined with wrong builder", func() {
			builders := []mta.ProjectBuilder{
				{
					Builder: "xxx",
				},
			}
			projectBuild := mta.ProjectBuild{
				BeforeAll: builders,
			}
			oMta := mta.MTA{
				BuildParams: &projectBuild,
			}
			Ω(execProjectBuilders(&oMta, "pre")).Should(HaveOccurred())
		})
		It("After Defined with wrong builder", func() {
			builders := []mta.ProjectBuilder{
				{
					Builder: "xxx",
				},
			}
			projectBuild := mta.ProjectBuild{
				AfterAll: builders,
			}
			oMta := mta.MTA{
				BuildParams: &projectBuild,
			}
			Ω(execProjectBuilders(&oMta, "post")).Should(HaveOccurred())
		})
	})

	var _ = Describe("runBuilder", func() {
		It("Sanity", func() {
			buildersCfg := commands.BuilderTypeConfig
			commands.BuilderTypeConfig =

				[]byte(`
builders:
- name: testbuilder
  info: "installing module dependencies & remove dev dependencies"
  path: "path to config file which override the following default commands"
  commands:
  - command: go version
`)
			builder := mta.ProjectBuilder{
				Builder: "testbuilder",
			}
			Ω(execProjectBuilder([]mta.ProjectBuilder{builder})).Should(Succeed())
			commands.BuilderTypeConfig = buildersCfg
		})
		It("Builder does not exist", func() {
			builder := mta.ProjectBuilder{
				Builder: "testbuilder",
			}
			Ω(execProjectBuilder([]mta.ProjectBuilder{builder})).Should(HaveOccurred())
		})
		It("Custom builder", func() {
			builder := mta.ProjectBuilder{Builder: "custom", Commands: []string{`echo "aaa"`}}
			Ω(execProjectBuilder([]mta.ProjectBuilder{builder})).Should(Succeed())
		})

		It("Fails on command execution", func() {
			buildersCfg := commands.BuilderTypeConfig
			commands.BuilderTypeConfig =

				[]byte(`
builders:
- name: testbuilder
  info: "installing module dependencies & remove dev dependencies"
  path: "path to config file which override the following default commands"
  commands:
  - command: go test unknown_test.go
`)
			builder := mta.ProjectBuilder{
				Builder: "testbuilder",
			}
			Ω(execProjectBuilder([]mta.ProjectBuilder{builder})).Should(HaveOccurred())
			commands.BuilderTypeConfig = buildersCfg
		})
	})

})

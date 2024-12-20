package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/SAP/cloud-mta-build-tool/internal/artifacts"
)

// meta command flags
var metaCmdSrc string
var metaCmdMtaYamlFilename string
var metaCmdTrg string
var metaCmdDesc string
var metaCmdExtensions []string
var metaCmdPlatform string

// mtar command flags
var mtarCmdSrc string
var mtarCmdMtaYamlFilename string
var mtarCmdTrg string
var mtarCmdDesc string
var mtarCmdTrgProvided string
var mtarCmdExtensions []string
var mtarCmdMtarName string

// init - inits flags of init command
func init() {

	// set flags of meta command
	metaCmd.Flags().StringVarP(&metaCmdSrc, "source", "s", "",
		"The path to the MTA project; the current path is set as default")
	metaCmd.Flags().StringVarP(&metaCmdMtaYamlFilename, "filename", "f", "",
		"The mta yaml filename of the MTA project; the mta.yaml is set as default")
	metaCmd.Flags().StringVarP(&metaCmdTrg, "target", "t", "",
		"The path to the folder in which a temporary folder with generated metadata is created; the current path is set as default")
	metaCmd.Flags().StringVarP(&metaCmdDesc, "desc", "d", "",
		`The MTA descriptor; supported values: "dev" (development descriptor, default value) and "dep" (deployment descriptor)`)
	metaCmd.Flags().StringSliceVarP(&metaCmdExtensions, "extensions", "e", nil,
		"The MTA extension descriptors")
	metaCmd.Flags().StringVarP(&metaCmdPlatform, "platform", "p", "cf",
		`The deployment platform; supported platforms: "cf", "xsa", "neo"`)
	metaCmd.Flags().BoolP("help", "h", false, `Displays detailed information about the "meta" command`)

	// set flags of mtar command
	mtarCmd.Flags().StringVarP(&mtarCmdSrc, "source", "s", "",
		"The path to the MTA project; the current path is set as default")
	mtarCmd.Flags().StringVarP(&mtarCmdMtaYamlFilename, "filename", "f", "",
		"The mta yaml filename of the MTA project; the mta.yaml is set as default")
	mtarCmd.Flags().StringVarP(&mtarCmdTrg, "target", "t", "",
		`The path to the folder in which the MTAR file is created; the path to the "mta_archives" subfolder of the current folder is set as default`)
	mtarCmd.Flags().StringVarP(&mtarCmdDesc, "desc", "d", "",
		`The MTA descriptor; supported values: "dev" (development descriptor, default value) and "dep" (deployment descriptor)`)
	mtarCmd.Flags().StringSliceVarP(&mtarCmdExtensions, "extensions", "e", nil,
		"The MTA extension descriptors")
	mtarCmd.Flags().StringVarP(&mtarCmdMtarName, "mtar", "m", "*",
		"The archive name")
	mtarCmd.Flags().StringVarP(&mtarCmdTrgProvided, "target_provided", "", "",
		"The MTA target provided indicator; supported values: true, false")
	_ = mtarCmd.Flags().MarkHidden("target_provided")
	mtarCmd.Flags().BoolP("help", "h", false, `Displays detailed information about the "mtar" command`)

}

// Generate metadata info from deployment
var metaCmd = &cobra.Command{
	Use:   "meta",
	Short: "Generates the META-INF folder",
	Long:  "Generates META-INF folder with manifest and MTAD files",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteGenMeta(metaCmdSrc, metaCmdMtaYamlFilename, metaCmdTrg, metaCmdDesc, metaCmdExtensions, metaCmdPlatform, os.Getwd)
		logError(err)
		return err
	},
	Hidden:        true,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Generate mtar from build artifacts
var mtarCmd = &cobra.Command{
	Use:   "mtar",
	Short: "Generates MTA archive",
	Long:  "Generates MTA archive from the folder with all artifacts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := artifacts.ExecuteGenMtar(mtarCmdSrc, mtarCmdMtaYamlFilename, mtarCmdTrg, mtarCmdTrgProvided, mtarCmdDesc, mtarCmdExtensions, mtarCmdMtarName, os.Getwd)
		logError(err)
		return err
	},
	Hidden:        true,
	SilenceUsage:  true,
	SilenceErrors: true,
}

package dir

import (
	"path/filepath"

	"github.com/SAP/cloud-mta/mta"
)

// ModuleLoc - module location type that provides services for stand alone module build command
type ModuleLoc struct {
	loc               *Loc
	targetPathDefined bool
}

// GetTarget - gets the target path
func (ep *ModuleLoc) GetTarget() string {
	return ep.loc.GetTarget()
}

// GetTargetTmpRoot - gets the target root path
func (ep *ModuleLoc) GetTargetTmpRoot() string {
	if ep.targetPathDefined {
		return ep.loc.GetTarget()
	}
	// default target folder for module build results is defined under the temp folder
	return filepath.Dir(ep.loc.GetTarget())
}

// GetSourceModuleDir - gets the absolute path to the module
func (ep *ModuleLoc) GetSourceModuleDir(modulePath string) string {
	return ep.loc.GetSourceModuleDir(modulePath)
}

// GetSourceModuleArtifactRelPath - gets the relative path to the module artifact
// The ModuleLoc type is used in context of stand alone module build command and as opposed to the module build command in the context
// of Makefile saves its build result directly under the target (temporary or specific) without considering the original artifact path in the source folder
func (ep *ModuleLoc) GetSourceModuleArtifactRelPath(modulePath, artifactPath string) (string, error) {
	return "", nil
}

// GetTargetModuleDir - gets the to module build results
func (ep *ModuleLoc) GetTargetModuleDir(moduleName string) string {
	return ep.loc.GetTarget()
}

// ParseFile returns a reference to the MTA object resulting from the given mta.yaml file merged with the extension descriptors.
func (ep *ModuleLoc) ParseFile() (*mta.MTA, error) {
	return ep.loc.ParseFile()
}

func (ep *ModuleLoc) SetStrictParmeter(strick bool) bool {
	return ep.loc.SetStrictParmeter(strick)
}

// ModuleLocation - provides target location of stand alone MTA module build result
func ModuleLocation(loc *Loc, targetPathDefined bool) *ModuleLoc {
	return &ModuleLoc{loc: loc, targetPathDefined: targetPathDefined}
}

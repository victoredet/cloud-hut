// Package projects
package projects

import (
	"fmt"

	"mobileHost/utils"
)

func CreateProject(projectName string, projectType string) {
	if !utils.ValidateAddProjectType(projectType) {
		fmt.Println("Invalid project type")
		return
	}
}

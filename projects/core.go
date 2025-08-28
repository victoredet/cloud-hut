// Package projects
package projects

import (
	"fmt"

	"mobileHost/database"
	"mobileHost/models"
	"mobileHost/utils"
)

func CreateProject(projectName string, projectType string) {
	if !utils.ValidateAddProjectType(projectType) {
		fmt.Println("Invalid project type")
		return
	}

	project := models.Project{Name: projectName, Type: projectType}
	database.DB.Create(&project)
}

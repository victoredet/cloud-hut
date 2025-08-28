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

func ListProjects() {
	var projects []models.Project
	result := database.DB.Find(&projects)

	if result.Error != nil {
		fmt.Println("Error fetching projects:", result.Error)
		return
	}

	if len(projects) == 0 {
		fmt.Println("No projects found. ")
		fmt.Println("Use `cloud-hut project` to create one.")
		return
	}

	fmt.Println("Projects:")
	for _, project := range projects {
		fmt.Printf("- %s (%s)\n", project.Name, project.Type)
	}
}

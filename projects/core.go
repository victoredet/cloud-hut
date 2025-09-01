// Package projects
package projects

import (
	"fmt"
	"os/exec"
	"path/filepath"

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

func createLaravelProject(projectName string, gitRepo string) {
	fmt.Println("Creating Laravel project:", projectName)

	keyPath := filepath.Join("/", "var", "www")

	projectDir := exec.Command("mkdir", "-p", keyPath, projectName)
	err := projectDir.Run()
	if err != nil {
		fmt.Println("Error creating project directory:", err)
		return
	}

	gitCloneCmd := exec.Command("git", "clone", gitRepo, filepath.Join(keyPath, projectName))
	err = gitCloneCmd.Run()
	if err != nil {
		fmt.Println("Error cloning repository. Please make sure you have all the right permissions.")
		fmt.Println("Error:", err)
		return
	}
}

func exposeProjectToInternet(projectName string) {
}

// Package utils
package utils

import "fmt"

func ValidateAddProjectType(projectType string) bool {
	fmt.Println("project type", projectType)
	// add flask and fastapi
	return projectType == "laravel" || projectType == "php" || projectType == "node" || projectType == "flask" || projectType == "fastapi"
}

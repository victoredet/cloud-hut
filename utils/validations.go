// Package utils
package utils

import "fmt"

func ValidateAddProjectType(projectType string) bool {
	fmt.Println("project type", projectType)
	return projectType == "laravel" ||
		projectType == "php" ||
		projectType == "node" ||
		projectType == "flask" ||
		projectType == "fastapi"
}

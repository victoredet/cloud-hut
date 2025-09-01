// Package utils
package utils

import "fmt"

func ValidateAddProjectType(projectType string) bool {
	fmt.Println("project type", projectType)
	return projectType == "laravel" ||
		projectType == "php" ||
		projectType == "react" ||
		projectType == "vue" ||
		projectType == "next" ||
		projectType == "nuxt" ||
		projectType == "node" ||
		projectType == "flask" ||
		projectType == "fastapi"
}

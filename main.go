package main

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cloud-hut",
		Short: "cloud Hut",
	}

	apacheCmd := &cobra.Command{
		Use:   "apache",
		Short: "install and set up apache",
		Run: func(cmd *cobra.Command, args []string) {
			var name string
			prompt := &survey.Input{
				Message: "What is your name?",
			}
			survey.AskOne(prompt, &name)

			fmt.Printf("Hello, %s!\n", name)
		},
	}

	gitSetUp := &cobra.Command{
		Use:   "git",
		Short: "set up git",
		Run: func(cmd *cobra.Command, args []string) {
			var email string
			var name string
			emailH := &survey.Input{
				Message: "Enter your git email address: ",
			}

			survey.AskOne(emailH, &email)

			nameH := &survey.Input{
				Message: "Enter your name: ",
			}

			survey.AskOne(nameH, &name)

			// fmt.Println("Installing git...")
			// gitInstaller := exec.Command("pkg", "install", "git")
			// gitInstallerOutput, err := gitInstaller.Output()
			//    if err != nil {
			//      fmt.Println("Error:", err)
			//      return
			//    }
			//    fmt.Println(string(gitInstallerOutput))

			// set user.email
			emailCmd := exec.Command("git", "config", "--global", "user.email", email)
			if out, err := emailCmd.CombinedOutput(); err != nil {
				fmt.Println("Error setting email:", err)
				fmt.Println("Output:", string(out))
				return
			}

			// Set user.name
			nameCmd := exec.Command("git", "config", "--global", "user.name", name)
			if out, err := nameCmd.CombinedOutput(); err != nil {
				fmt.Println("Error setting name:", err)
				fmt.Println("Output:", string(out))
				return
			}





		},
	}

	rootCmd.AddCommand(gitSetUp)
	rootCmd.AddCommand(apacheCmd)

	rootCmd.Execute()
}

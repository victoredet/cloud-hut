package main 

import (
	"fmt"
"os/exec"

	"github.com/spf13/cobra"
  "github.com/AlecAivazis/survey/v2"
)



func main() {
	var rootCmd = &cobra.Command{
		Use:   "cloud-hut",
		Short: "cloud Hut",
	}

	var apacheCmd = &cobra.Command{
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



	var gitSetUp = &cobra.Command{
    Use:   "git",
    Short: "set up git",
    Run: func(cmd *cobra.Command, args []string) {
		var email string
		var name string
		emailH := &survey.Input{
				Message:"Enter your git email address: ",
		}

    survey.AskOne(emailH, &email)

    nameH := &survey.Input{
				Message:"Enter your name: ",
    }

    survey.AskOne(nameH, &name)

    //commandH := exec.Command("ls", "-la") // or any other command
			
		//install git 
			gitInstaller := exec.Command("pkg", "install", "git")	
			gitInstallerOutput, err := gitInstaller.Output()
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      fmt.Println(string(gitInstallerOutput))








			//
		// commandH := exec.Command("git", "config", "--global", "user.email", email, "&&",   "git", "config", "--global", "user.name", name)
		// out, err := commandH.Output()
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	return
		// }
		//
		//  fmt.Println(string(out))
		//
  	//create ssh key and return it as cat 
    },
  }
  
	rootCmd.AddCommand(gitSetUp)
	rootCmd.AddCommand(apacheCmd)

	rootCmd.Execute()
}


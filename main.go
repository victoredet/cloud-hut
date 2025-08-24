package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
			fmt.Println("Installing apache...")
			apacheInstaller := exec.Command("pkg", "install", "apache2")
			apacheInstallerOutput, err := apacheInstaller.Output()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(string(apacheInstallerOutput))
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

			fmt.Println("Installing git...")
			gitInstaller := exec.Command("pkg", "install", "git")
			gitInstallerOutput, err := gitInstaller.Output()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(string(gitInstallerOutput))

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

			// generate ssh key
			fmt.Println("Generating ssh key...")
			sshCmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email)
			if out, err := sshCmd.CombinedOutput(); err != nil {
				fmt.Println("Error generating ssh key:", err)
				fmt.Println("Output:", string(out))
				return
			}

			// cat out the ssh key
			fmt.Println("Your ssh key is:")

			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error finding home directory:", err)
				return
			}

			sshKeyPath := filepath.Join(home, ".ssh", "id_ed25519.pub")
			sshCatCmd := exec.Command("cat", sshKeyPath)

			if out, err := sshCatCmd.CombinedOutput(); err != nil {
				fmt.Println("Error generating ssh key:", err)
				fmt.Println("Output:", string(out))
				return
			}
		},
	}

	rootCmd.Version = "0.0.1"
	rootCmd.AddCommand(gitSetUp)
	rootCmd.AddCommand(apacheCmd)

	rootCmd.Execute()
}

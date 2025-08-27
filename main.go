package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
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
			// check apache already installed
			fmt.Println("Checking if apache is already installed...")
			apacheCmd := exec.Command("pacman", "-Q", "apache")
			apacheCmdOutput, err := apacheCmd.Output()
			if err == nil {
				fmt.Println("Apache is already set up for use")
				fmt.Println(string(apacheCmdOutput))
				return
			}

			fmt.Println("Installing apache...")
			apacheInstaller := exec.Command("sudo", "pacman", "--noconfirm", "-S", "apache")
			apacheInstallerOutput, err := apacheInstaller.CombinedOutput()
			fmt.Println(string(apacheInstallerOutput))
			if err != nil {
				fmt.Println("Error:", err)
				// tell server what has happened here
				fmt.Println("Please allow us send diagnostic information for help")
			}
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
			// check git already installed
			checkGitCmd := exec.Command("pacman", "-Q", "git")
			checkGitCmdOutput, err := checkGitCmd.Output()
			if err == nil {
				fmt.Println("Git is already set up for use")
				fmt.Println(string(checkGitCmdOutput))
			} else {
				gitInstaller := exec.Command("sudo", "pacman", "-S", "--noconfirm", "git")
				gitInstallerOutput, err := gitInstaller.Output()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Println(string(gitInstallerOutput))

			}

			fmt.Println("Setting up git...")

			emailCmd := exec.Command("git", "config", "--global", "user.email", email)
			if out, err := emailCmd.CombinedOutput(); err != nil {
				fmt.Println("Error setting email:", err)
				fmt.Println("Output:", string(out))
				return
			}

			nameCmd := exec.Command("git", "config", "--global", "user.name", name)
			if out, err := nameCmd.CombinedOutput(); err != nil {
				fmt.Println("Error setting name:", err)
				fmt.Println("Output:", string(out))
				return
			}

			fmt.Println("Generating ssh key...")
			sshCmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25518"))
			if out, err := sshCmd.CombinedOutput(); err != nil {
				fmt.Println("Error generating ssh key:", err)
				fmt.Println("Output:", string(out))
				return
			}

			fmt.Println("SSH key generated successfully.")
			fmt.Println("Copy this to GitHub: ")
			keyPath := filepath.Join(os.Getenv("HOME"), ".ssh", "id_ed25518")

			sshKeyDisplayCmd := keyPath + ".pub"
			pubKey, err := os.ReadFile(sshKeyDisplayCmd)
			if err != nil {
				fmt.Println("Error reading public key:", err)
				return
			}

			fmt.Println("Copy this to GitHub:")
			fmt.Println(string(pubKey))

			pubKeyStr := string(pubKey)
			err = clipboard.WriteAll(pubKeyStr)
			if err != nil {
				fmt.Println("Failed to copy to clipboard:", err)
			} else {
				fmt.Println("✅ SSH public key copied to clipboard.")
			}

			exec.Command("xdg-open", "https://github.com/settings/ssh/new")
		},
	}

	composerSetupCmd := &cobra.Command{
		Use:   "composer",
		Short: "install and set up composer",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Installing composer...")
			composerInstaller := exec.Command("sudo", "pacman", "--noconfirm", "-S", "composer")
			composerInstallerOutput, err := composerInstaller.Output()
			fmt.Println(string(composerInstallerOutput))
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Please allow us send diagnostic information for help")
			}
		},
	}

	// node, yarn and npm
	nodeSetupCmd := &cobra.Command{
		Use:   "node",
		Short: "install and set up node",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Installing node...")
			nodeInstaller := exec.Command("sudo", "pacman", "--noconfirm", "-S", "npm")
			nodeInstallerOutput, err := nodeInstaller.Output()
			fmt.Println(string(nodeInstallerOutput))
			if err != nil {
				fmt.Println("Error:", err)
				fmt.Println("Please allow us send diagnostic information for help")
			}
		},
	}

	globalSetupCmd := &cobra.Command{
		Use:   "init",
		Short: "set up the system",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to cloud-hut! Give us a few moments to set up ")

			fmt.Println("Checking if sudo is installed...")
			sudoCmd := exec.Command("pacman", "-Q", "sudo")
			_, err := sudoCmd.Output()
			if err == nil {
				fmt.Println("Permissions are already set up for use")
			}
			exec.Command("sudo", "pacman", "Sy")

			apacheCmd.Run(cmd, args)
			gitSetUp.Run(cmd, args)
			composerSetupCmd.Run(cmd, args)
			nodeSetupCmd.Run(cmd, args)
		},
	}

	rootCmd.Version = "0.0.1"
	rootCmd.AddCommand(gitSetUp)
	rootCmd.AddCommand(apacheCmd)
	rootCmd.AddCommand(globalSetupCmd)
	rootCmd.AddCommand(composerSetupCmd)
	rootCmd.AddCommand(nodeSetupCmd)
	rootCmd.Execute()

	fmt.Println("Press Ctrl+C to exit...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Block here until Ctrl+C
	fmt.Println("\nExiting.")
}

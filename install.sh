#!/bin/bash

# Function to detect the OS
get_os() {
  case "$(uname -s)" in
  Linux*) os=linux ;;
  Darwin*) os=macos ;;
  CYGWIN* | MINGW*) os=windows ;;
  *) os="unknown" ;;
  esac
  echo "$os"
}

# Set your GitHub release URL
REPO_OWNER="victoredet"
REPO_NAME="cloud-hut"
RELEASE_TAG="Oak" # Update this with the latest version
# RELEASE_TAG="v1.0.0" # Update this with the latest version

# Get the OS
OS=$(get_os)

echo "Detected OS: $OS"

# Function to download and install the binary
download_and_install() {
  case $OS in
  linux)
    URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$RELEASE_TAG/installapp-linux"
    DEST="/usr/local/bin/cloud-hut"
    ;;
  macos)
    URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$RELEASE_TAG/installapp-macos"
    DEST="/usr/local/bin/cloud-hut"
    ;;
  windows)
    URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$RELEASE_TAG/installapp.exe"
    DEST="$HOME/cloud-hut.exe"
    ;;
  *)
    echo "Unsupported OS. Only Linux, macOS, and Windows are supported."
    exit 1
    ;;
  esac

  echo "Downloading $URL..."
  curl -L $URL -o $DEST

  # Make it executable (for Linux and macOS)
  if [[ $OS == "linux" || $OS == "macos" ]]; then
    echo "Making the binary executable..."
    chmod +x $DEST
  fi

  echo "Installation complete!"
  echo "You can now run the app using: $DEST"
  if [[ $OS == "linux" || $OS == "macos" ]]; then
    echo "Adding alias to your shell configuration..."

    # Determine the shell type and add the alias accordingly
    if [ -n "$BASH_VERSION" ]; then
      # If the user is using bash
      echo "alias cloud-hut='$DEST'" >>~/.bashrc
      source ~/.bashrc
      echo "Alias added to ~/.bashrc"
    elif [ -n "$ZSH_VERSION" ]; then
      # If the user is using zsh
      echo "alias cloud-hut='$DEST'" >>~/.zshrc
      source ~/.zshrc
      echo "Alias added to ~/.zshrc"
    else
      echo "No suitable shell configuration found. Please manually add the alias."
    fi
  fi
}

# Download and install the binary
download_and_install

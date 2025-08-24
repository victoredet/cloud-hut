#!/bin/bash

set -e

REPO_URL="https://github.com/victoredet/cloud-hut"
BIN_NAME="cloud-hut"
INSTALL_DIR="/usr/local/bin"
ARCH=$(uname -m)
OS=$(uname | tr '[:upper:]' '[:lower:]')

# Map machine architecture to Go arch naming
if [ "$ARCH" == "x86_64" ]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
fi

BINARY_URL="$REPO_URL/releases/latest/download/${BIN_NAME}-${OS}-${ARCH}"

echo "Downloading $BIN_NAME for $OS/$ARCH from $BINARY_URL..."
curl -L "$BINARY_URL" -o "$BIN_NAME"
chmod +x "$BIN_NAME"
sudo mv "$BIN_NAME" "$INSTALL_DIR/"

echo "✅ $BIN_NAME installed to $INSTALL_DIR"
echo "Run '$BIN_NAME --help' to get started."

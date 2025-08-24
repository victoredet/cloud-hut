#!/bin/bash

echo "Installing cloud-hut..."
curl -L https://example.com/cloud-hut-linux -o /usr/local/bin/cloud-hut
chmod +x /usr/local/bin/cloud-hut
echo "Installed! Run 'cloud-hut --help' to get started."

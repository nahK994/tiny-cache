#!/bin/bash

# Define download URLs for server and client binaries (use actual URLs from your release)
SERVER_BINARY_URL="https://github.com/nahK994/TinyCache/releases/download/bin/tinycache-server"
CLIENT_BINARY_URL="https://github.com/nahK994/TinyCache/releases/download/bin/tinycache"
# Define installation paths
SERVER_INSTALL_PATH="/usr/local/sbin/tinycache-server"
CLIENT_INSTALL_PATH="/usr/local/bin/tinycache"
SERVICE_FILE_PATH="/etc/systemd/system/tinycache.service"

# Step 1: Download the server binary
echo "Downloading the server binary from $SERVER_BINARY_URL..."
sudo curl -L $SERVER_BINARY_URL -o $SERVER_INSTALL_PATH
sudo chmod +x $SERVER_INSTALL_PATH

# Step 2: Download the client binary
echo "Downloading the client binary from $CLIENT_BINARY_URL..."
sudo curl -L $CLIENT_BINARY_URL -o $CLIENT_INSTALL_PATH
sudo chmod +x $CLIENT_INSTALL_PATH

# Step 3: Install the server as a systemd service
echo "Installing the server as a service..."

# Create the systemd service file
sudo tee $SERVICE_FILE_PATH > /dev/null <<EOL
[Unit]
Description=TinyCache Server
After=network.target

[Service]
ExecStart=$SERVER_INSTALL_PATH
Restart=always
User=$(whoami)
WorkingDirectory=/usr/local/sbin

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd to recognize the new service
sudo systemctl daemon-reload

# Enable the service to start at boot and start it immediately
sudo systemctl enable tinycache.service
sudo systemctl start tinycache.service

echo "TinyCache server has been installed and started as a service."
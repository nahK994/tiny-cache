#!/bin/bash

# Define installation paths
SERVER_INSTALL_PATH="/usr/local/sbin/tinycache-server"
CLIENT_INSTALL_PATH="/usr/local/bin/tinycache"
SERVICE_FILE_PATH="/etc/systemd/system/tinycache.service"

# Step 1: Stop and disable the systemd service
echo "Stopping the TinyCache service..."
sudo systemctl stop tinycache.service

echo "Disabling the TinyCache service..."
sudo systemctl disable tinycache.service

# Step 2: Remove the systemd service file
if [[ -f $SERVICE_FILE_PATH ]]; then
    echo "Removing the service file at $SERVICE_FILE_PATH..."
    sudo rm $SERVICE_FILE_PATH
else
    echo "Service file not found."
fi

# Step 3: Reload systemd daemon to apply changes
echo "Reloading systemd daemon..."
sudo systemctl daemon-reload

# Step 4: Remove the server and client binaries
if [[ -f $SERVER_INSTALL_PATH ]]; then
    echo "Removing the server binary at $SERVER_INSTALL_PATH..."
    sudo rm $SERVER_INSTALL_PATH
else
    echo "Server binary not found."
fi

if [[ -f $CLIENT_INSTALL_PATH ]]; then
    echo "Removing the client binary at $CLIENT_INSTALL_PATH..."
    sudo rm $CLIENT_INSTALL_PATH
else
    echo "Client binary not found."
fi

echo "TinyCache has been successfully uninstalled."

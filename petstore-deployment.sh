#!/bin/bash

# App Variables
APP_NAME="petstore"
GO_BINARY="main"
APP_DIR="/home/ec2-user/petstore"
SYSTEMD_SERVICE="petstore.service"

# Update the package list and install required packages
sudo yum update -y
sudo yum install -y git

# Install Go
sudo yum install -y golang

# Set Go environment variables (adjust as needed)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

# Verify Go and Git installations
go version
git --version

# Clone and build petstore repo
git clone https://github.com/rafliputraa/petstore.git

cd petstore

go build cmd/main.go

# Update the credentials .env file according to the owned user's RDS
PG_USERNAME=rafliputraa
PG_PASSWORD=1234567890
PG_HOST=localhost
PG_PORT=5432

# Create a systemd service file
sudo tee "/etc/systemd/system/$SYSTEMD_SERVICE" > /dev/null <<EOL
[Unit]
Description=$APP_NAME

[Service]
ExecStart=$APP_DIR/$GO_BINARY
WorkingDirectory=$APP_DIR

[Install]
WantedBy=multi-user.target
EOL

# Enable and start the service
sudo systemctl enable "$SYSTEMD_SERVICE"
sudo systemctl start "$SYSTEMD_SERVICE"
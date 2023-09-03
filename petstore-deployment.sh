#!/bin/bash

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

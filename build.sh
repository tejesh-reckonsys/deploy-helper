#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Configuration
PROJECT_NAME="deploy-helper"
VERSION=$(git describe --tags --abbrev=0) # Get the latest tag as the version
BUILD_DIR="build"
RELEASE_ASSET_DESC="Builds for macOS and Linux"

# Ensure version is set
if [ -z "$VERSION" ]; then
    echo "No git tags found. Please create a tag first."
    exit 1
fi

# Clean and create build directory
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

# Build for macOS (amd64 and arm64)
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o "$BUILD_DIR/${PROJECT_NAME}_macOS_amd64" .
GOOS=darwin GOARCH=arm64 go build -o "$BUILD_DIR/${PROJECT_NAME}_macOS_arm64" .

# Build for Linux (amd64 and arm64)
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$BUILD_DIR/${PROJECT_NAME}_linux_amd64" .
GOOS=linux GOARCH=arm64 go build -o "$BUILD_DIR/${PROJECT_NAME}_linux_arm64" .

echo "Build process complete."


#!/bin/bash
if [ "$(id -u)" -ne 0 ]; then
    echo "This script must be run as root or with sudo."
    exit 1
fi

REPO_URL="https://api.github.com/repos/CosmicPredator/chibi-cli/releases/latest"
BINARY_NAME="chibi"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"
CONFIG_DIR="/home/$SUDO_USER/.config/chibi"
OS_ARCH="linux_amd64"

install_binary() {
    echo "Fetching the latest release information..."

    if command -v curl > /dev/null 2>&1; then
        RELEASE_INFO=$(curl -s "$REPO_URL")
    elif command -v wget > /dev/null 2>&1; then
        RELEASE_INFO=$(wget -q -O- "$REPO_URL")
    else
        echo "Error: Neither curl nor wget is installed."
        exit 1
    fi

    DOWNLOAD_URL=$(echo "$RELEASE_INFO" | jq -r ".assets[] | select(.name | test(\"$OS_ARCH\")) | .browser_download_url")

    if [ -z "$DOWNLOAD_URL" ]; then
        echo "Error: No binary found for $OS_ARCH in the latest release."
        exit 1
    fi

    echo "Downloading $BINARY_NAME from $DOWNLOAD_URL..."
    wget -O "$BINARY_NAME" "$DOWNLOAD_URL"

    if [ ! -f "$BINARY_NAME" ]; then
        echo "Error: Failed to download $BINARY_NAME."
        exit 1
    fi

    echo "Installing $BINARY_NAME to $INSTALL_PATH..."
    mv "$BINARY_NAME" "$INSTALL_PATH"

    chmod +x "$INSTALL_PATH"

    if [ -x "$INSTALL_PATH" ]; then
        echo "$BINARY_NAME installed successfully at $INSTALL_PATH"
    else
        echo "Error: Failed to install $BINARY_NAME."
        exit 1
    fi
}

uninstall_binary() {
    if [ -f "$INSTALL_PATH" ]; then
        echo "Uninstalling $BINARY_NAME from $INSTALL_PATH..."
        rm -f "$INSTALL_PATH"

        if [ ! -f "$INSTALL_PATH" ]; then
            echo "$BINARY_NAME uninstalled successfully."
        else
            echo "Error: Failed to uninstall $BINARY_NAME."
            exit 1
        fi
    else
        echo "Error: $BINARY_NAME is not installed."
        exit 1
    fi

    if [ -d "$CONFIG_DIR" ]; then
        echo "Removing $CONFIG_DIR directory..."
        rm -rf "$CONFIG_DIR"
        
        if [ ! -d "$CONFIG_DIR" ]; then
            echo "$CONFIG_DIR removed successfully."
        else
            echo "Error: Failed to remove $CONFIG_DIR."
            exit 1
        fi
    else
        echo "No configuration directory found at $CONFIG_DIR."
    fi
}

if [ "$1" == "--uninstall" ]; then
    uninstall_binary
else
    install_binary
fi

exit 0

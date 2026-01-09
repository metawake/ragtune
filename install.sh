#!/bin/bash
# RagTune installer script
# Usage: curl -sSL https://raw.githubusercontent.com/metawake/ragtune/main/install.sh | bash

set -e

REPO="metawake/ragtune"
BINARY="ragtune"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

case "$OS" in
    linux|darwin) ;;
    mingw*|msys*|cygwin*) OS="windows" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Get latest release version
echo "Fetching latest release..."
LATEST=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST" ]; then
    echo "Error: Could not fetch latest release"
    echo "Try installing with Go instead:"
    echo "  go install github.com/$REPO/cmd/ragtune@latest"
    exit 1
fi

VERSION="${LATEST#v}"
FILENAME="${BINARY}_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$LATEST/$FILENAME"

echo "Downloading $BINARY $LATEST for $OS/$ARCH..."
echo "URL: $URL"

# Create temp directory
TMP_DIR=$(mktemp -d)
trap "rm -rf $TMP_DIR" EXIT

# Download and extract
curl -sL "$URL" -o "$TMP_DIR/$FILENAME"
tar -xzf "$TMP_DIR/$FILENAME" -C "$TMP_DIR"

# Install
if [ -w "$INSTALL_DIR" ]; then
    mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/"
else
    echo "Installing to $INSTALL_DIR (requires sudo)..."
    sudo mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/"
fi

chmod +x "$INSTALL_DIR/$BINARY"

echo ""
echo "✓ Installed $BINARY $LATEST to $INSTALL_DIR/$BINARY"
echo ""
echo "Get started:"
echo "  $BINARY --help"
echo ""
echo "Quick start:"
echo "  docker run -d -p 6333:6333 -p 6334:6334 qdrant/qdrant"
echo "  $BINARY ingest ./docs --collection demo --embedder ollama"
echo "  $BINARY explain \"your query\" --collection demo"



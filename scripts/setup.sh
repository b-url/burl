#!/usr/bin/env zsh

# Exit when a command fails.
set -euo pipefail

BREW_PACKAGES=(go-task jq golangci-lint git-cliff ariga/tap/atlas)
NODE_VERSION="20.17.0"

BLACK="\033[30m"
RED="\033[31m"
GREEN="\033[32m"
YELLOW="\033[33m"
BLUE="\033[34m"
MAGENTA="\033[35m"
CYAN="\033[36m"
WHITE="\033[37m"
DEFAULT="\033[39m"
RESET="\033[0m"
BOLD="\033[1m"
UNDERLINE="\033[4m"
REVERSED="\033[7m"

pretty_print() {
    local message="$1"
    local color="${2:-$DEFAULT}"
    local style="${3:-}"

    echo -e "${color}${style}$message${RESET}"
}

header() {
    pretty_print "\n========================================================================"
    pretty_print "$1" $BLUE $BOLD
    pretty_print "========================================================================\n"
}

success() {
    pretty_print ""
    pretty_print "✅ $1" $GREEN
}

install_go () {
    header "Setting up Go 🐹"

    if command -v go; then
        echo "Go is already installed"
        go version
        return
    else
        curl -L https://git.io/vQhTU | bash -s -- --version 1.23.1
    fi

    success "Finished setting up Go!"
}

# Install Homebrew, a package manager for macOS
install_brew() {
    header "Setting up Homebrew 🍺"

    if command -v brew; then
        echo "Homebrew is already installed"
        eval "$(brew shellenv)"
    else
        /bin/zsh -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        INSTALL_DIR="/opt/homebrew/bin"

        if read -q "confirm? -> add $INSTALL_DIR to path in ~/.zprofile? [y/N] "; then
            echo 'eval "$('$INSTALL_DIR'/brew shellenv)"' >>~/.zprofile
        fi

        echo
        eval "$($INSTALL_DIR/brew shellenv)"

        if ! command -v brew >/dev/null; then
            echo "Brew installation failed!"
            exit 1
        fi
    fi

    success "Finished setting up Homebrew!"
}

# Install Volta, a JavaScript tool manager
install_volta() {
    header "Setting up Volta ⚡️"

    if command -v volta; then
        echo "Volta is already installed"
    else
        curl https://get.volta.sh | bash
    fi

    success "Finished setting up Volta!"
}

# Install Node.js using Volta
install_node() {
    header "Installing Node.js 🚀"

    if command -v node; then
        echo "Node.js is already installed"
        node --version
        return
    else 
        volta install node@${NODE_VERSION}
        # Run node version to verify installation
        node --version
    fi

    success "Finished installing Node.js!"
}

# Install or upgrade a package with Homebrew
brew_install() {
    if brew ls --versions "$1"; then
        brew upgrade "$1"
    else
        brew install "$1"
    fi
}

# Install packages with Homebrew
install_packages() {
    header "Installing packages with Homebrew 🍺"

    for package in "${BREW_PACKAGES[@]}"; do
        echo "Installing or upgrading $package..."
        brew_install "$package"
        echo "Installed or upgraded $package"
    done

    success "Finished installing packages with Homebrew!"
}

# Setup git hooks.
setup_git_hooks() {
    header "Setting up git hooks 🎣"

    task setup-git-hooks
    success "Finished setting up git hooks!"
}

echo "Setting up your development environment..."

install_go
install_brew
install_volta
install_node
install_packages

setup_git_hooks

pretty_print "\nFinished setting up your development environment! 🎉" $GREEN $BOLD
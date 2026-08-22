# Justfile for yell (Fyne Toast Notification App)

os_type := os()

# Default recipe when running 'just'
default: build-local

# Install prerequisites for cross-compilation and building
setup:
	@echo "==> Installing fyne-cross"
	go install github.com/fyne-io/fyne-cross@latest
	@echo "==> Installing fyne-cli"
	go install fyne.io/fyne/v2/cmd/fyne@latest

# Run the app locally in development mode
run:
	go run main.go --title="Woooah! Love Alert" --message="you love me bro?" --icon="❤️‍🔥"

# Run the app locally with custom CLI flags
run-args TITLE="Break Time" MSG="Go to bed right now" ICON="🫩":
	go run main.go -title "{{TITLE}}" -message "{{MSG}}" -icon "{{ICON}}"

# Build binary locally for current platform (macOS/Linux/Windows)
build-local:
	@echo "==> Building local binary"
	go build -ldflags "-s -w" -o yell main.go

# Cross-compile for Windows (AMD64) using fyne-cross
build-windows:
	@echo "==> Cross-compiling for Windows via fyne-cross"
	fyne-cross windows -arch=amd64

# Cross-compile for Linux (AMD64) using fyne-cross
build-linux:
	@echo "==> Cross-compiling for Linux via fyne-cross"
	fyne-cross linux -arch=amd64,arm64

# Cross-compile for macOS (Darwin AMD64 & ARM64) using fyne-cross
build-darwin:
	@echo "==> Cross-compiling for macOS via fyne-cross"
	fyne-cross darwin -arch=arm64 -app-id="yell.app"

# Build container image using Dockerfile.build
docker-build:
	@echo "==> Building ARM64 Docker container"
	docker build -t yell-builder -f Dockerfile.build .

# Run the built Windows binary via PowerShell (from project root)
run-windows-ps MSG="Take a break!":
	powershell -Command ".\fyne-cross\bin\windows-amd64\yell.exe -message '{{MSG}}'"

install:
	@if [ "{{os_type}}" = "macos" ]; then \
		just install-macos; \
	elif [ "{{os_type}}" = "linux" ]; then \
		just install-linux; \
	else \
		echo "Unsupported host OS for installation: {{os_type}}"; \
		exit 1; \
	fi

# Internal recipe: macOS symlink deployment
install-macos: build-darwin
	@echo "==> Deploying yell.app from fyne-cross build..."
	@if [ -w /Applications ]; then \
		rm -rf /Applications/yell.app; \
		cp -R fyne-cross/dist/darwin-arm64/yell.app /Applications/; \
	else \
		sudo rm -rf /Applications/yell.app; \
		sudo cp -R fyne-cross/dist/darwin-arm64/yell.app /Applications/; \
	fi
	@echo "==> Creating symlink to /usr/local/bin/yell..."
	@if [ -w /usr/local/bin ]; then \
		ln -sf "/Applications/yell.app/Contents/MacOS/yell" /usr/local/bin/yell; \
	else \
		sudo ln -sf "/Applications/yell.app/Contents/MacOS/yell" /usr/local/bin/yell; \
	fi
	@echo "==> Successfully installed macOS app bundle and symlink!"

# Internal recipe: Linux binary deployment
install-linux: build-linux
	@echo "==> Deploying Linux binary to /usr/local/bin..."
	@if [ -w /usr/local/bin ]; then \
		cp fyne-cross/dist/linux-amd64/yell /usr/local/bin/yell; \
	else \
		sudo cp fyne-cross/dist/linux-amd64/yell /usr/local/bin/yell; \
	fi
	@echo "==> Successfully installed Linux binary!"

# OS-Specific Uninstallation
uninstall:
	@if [ "{{os_type}}" = "macos" ]; then \
		echo "==> Removing macOS yell.app and symlink..."; \
		if [ -w /Applications/yell.app ]; then rm -rf /Applications/yell.app; else sudo rm -rf /Applications/yell.app; fi; \
		if [ -w /usr/local/bin/yell ]; then rm -f /usr/local/bin/yell; else sudo rm -f /usr/local/bin/yell; fi; \
	elif [ "{{os_type}}" = "linux" ]; then \
		echo "==> Removing Linux yell binary..."; \
		if [ -w /usr/local/bin/yell ]; then rm -f /usr/local/bin/yell; else sudo rm -f /usr/local/bin/yell; fi; \
	else \
		echo "Unsupported host OS for uninstallation: {{os_type}}"; \
		exit 1; \
	fi
	@echo "==> Successfully uninstalled."

# Clean up built binaries and generated fyne-cross artifacts
clean:
	@echo "==> Cleaning build artifacts"
	rm -rf yell yell.exe fyne-cross/
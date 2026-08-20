# Justfile for yell (Fyne Toast Notification App)

# Default recipe when running 'just'
default: build-local

# Install prerequisites for cross-compilation and building
setup:
	@echo "==> Installing fyne-cross"
	go install github.com/fyne-io/fyne-cross@latest

# Run the app locally in development mode
run:
	go run main.go --title="Woooah! Love Alert" --message="you love me bro?" --icon="❤️‍🔥" --sound=happy

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
	fyne-cross linux -arch=amd64,arm64 -tags oto_purego,ebitengine_purego

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

# Install the binary locally into /usr/local/bin (macOS/Darwin)
install: build-local
	@echo "==> Installing yell to /usr/local/bin..."
	@if [ -w /usr/local/bin ]; then \
		cp yell /usr/local/bin/yell; \
	else \
		sudo cp yell /usr/local/bin/yell; \
	fi
	@echo "==> Successfully installed! You can now run 'yell' from anywhere."

# Uninstall the binary from /usr/local/bin
uninstall:
	@echo "==> Removing yell from /usr/local/bin..."
	@if [ -w /usr/local/bin/yell ]; then \
		rm -f /usr/local/bin/yell; \
	else \
		sudo rm -f /usr/local/bin/yell; \
	fi

# Clean up built binaries and generated fyne-cross artifacts
clean:
	@echo "==> Cleaning build artifacts"
	rm -rf yell yell.exe fyne-cross/
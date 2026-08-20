FROM fyneio/fyne-cross-images:linux

# Add support for both multi-arch targets
RUN dpkg --add-architecture amd64 && \
    dpkg --add-architecture arm64 && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    libasound2-dev:amd64 \
    libasound2-dev:arm64 \
    pkg-config && \
    rm -rf /var/lib/apt/lists/*
# Start with the official TinyGo development image
FROM tinygo/tinygo-dev:latest

# Install additional development tools
RUN apt-get update && apt-get install -y \
    git \
    make \
    gcc \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Install Go tools
RUN go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# Set working directory
WORKDIR /app

# Add convenient aliases and environment variables
RUN echo 'alias tg="tinygo"' >> ~/.bashrc \
    && echo 'alias tgb="tinygo build"' >> ~/.bashrc \
    && echo 'alias tgr="tinygo run"' >> ~/.bashrc

# Set default command
CMD ["/bin/bash"]
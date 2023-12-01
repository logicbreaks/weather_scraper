#!/bin/bash

# Download and install Go if it's not already installed
if [ -z "$(which go)" ]; then
  curl -sSL https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz | tar -xzv -C /tmp/go
  export PATH=/tmp/go/bin:$PATH
fi

# Clone the weather_scraper repository from GitHub
git clone https://github.com/logicbreaks/weather_scraper /tmp/weather_scraper

# Build the weather_scraper executable
cd /tmp/weather_scraper
go build -o weather_scraper

# Move the executable to the user's bin directory
mkdir -p ~/.local/bin
mv weather_scraper ~/.local/bin/

# Add the executable to the user's PATH environment variable
echo "export PATH=\$PATH:~/.local/bin" >> ~/.profile

# Source the ~/.profile file to make the PATH changes take effect
source ~/.profile

# Remove the temporary directories
rm -rf /tmp/go
rm -rf /tmp/weather_scraper

echo "weather_scraper installation complete. You can now run the weather_scraper command from anywhere on your system."

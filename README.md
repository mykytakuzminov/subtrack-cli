# Subtrack CLI

A simple CLI tool to track your subscriptions, built with Go.

> This is my first Go project, built to practice the language after coming from Python.

## Installation

```bash
git clone https://github.com/mykytakuzminov/subtrack-cli.git
cd subtrack-cli
go build -o subtrack .
sudo mkdir -p /usr/local/bin
sudo mv subtrack /usr/local/bin/
```

## Usage

```bash
# Add a subscription
subtrack add --name Netflix --price 9.99 --cycle monthly

# List all subscriptions
subtrack list

# Delete a subscription
subtrack delete <id>
```

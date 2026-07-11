<div align="center">

# 🚀 Subtrack CLI

CLI tool to manage your subscriptions

[![CI](https://github.com/mykytakuzminov/subtrack-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/mykytakuzminov/subtrack-cli/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

</div>

---

## Features

- **Subscription** - create, list and delete

## Getting Started

### Installation

```bash
git clone https://github.com/mykytakuzminov/subtrack-cli.git
cd subtrack-cli
go build -o subtrack .
sudo mkdir -p /usr/local/bin
sudo mv subtrack /usr/local/bin/
```

### Usage

```bash
# Add a subscription
subtrack add --name Netflix --price 9.99 --cycle monthly

# List all subscriptions
subtrack list

# Delete a subscription
subtrack delete <id>
```

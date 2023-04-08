#!/bin/bash
# Install pre commit 
echo "Install pre-commit"
pip3 install pre-commit

# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2

# Install pre commit to .git
pre-commit install --install-hooks --overwrite

# Run against all the files
pre-commit run --all-files
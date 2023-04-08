#!/bin/bash
# Install pre commit 
echo "Install pre-commit"
pip3 install pre-commit

# Install go tools
echo "Install go tools"
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
# Install golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2

# Install pre commit to .git
echo "Pre-commit Install to git hooks"
pre-commit install --install-hooks -f --overwrite

# Run against all the files
echo "Run hooks"
pre-commit run --all-files
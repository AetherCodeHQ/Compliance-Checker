# Compliance Checker

![CI](https://github.com/Qyroxen/Compliance-Checker/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Compliance-Checker/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Compliance-Checker?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Compliance-Checker)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Compliance-Checker)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Compliance-Checker?style=social)](https://github.com/Qyroxen/Compliance-Checker/stargazers)

## What is it?

Compliance Checker is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Compliance-Checker.git
cd Compliance-Checker
go build -o compliancechecker .

# Run
./compliancechecker --help
```

## CLI Usage

```bash
# Basic usage
./compliancechecker

# With flags
./compliancechecker --verbose --output json

# Get help
./compliancechecker --help
```

## Examples

```bash
# Example 1
./compliancechecker example1

# Example 2
./compliancechecker example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o compliancechecker .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Compliance-Checker/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Compliance-Checker?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Compliance-Checker/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Compliance-Checker?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Compliance-Checker/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Compliance-Checker" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Compliance-Checker/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Compliance-Checker" alt="Pull Requests">
  </a>
</p>

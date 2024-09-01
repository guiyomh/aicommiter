# Contributing to AICommitter

First off, thank you for considering contributing to AICommitter! Your help is essential to making this tool better for everyone. Below are the guidelines for contributing to the project.

## Table of Contents

- [Contributing to AICommitter](#contributing-to-aicommitter)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Code of Conduct](#code-of-conduct)
    - [How to Contribute](#how-to-contribute)
  - [Development Environment Setup](#development-environment-setup)
    - [Prerequisites](#prerequisites)
    - [Cloning the Repository](#cloning-the-repository)
    - [Building the Project](#building-the-project)
  - [Submitting Changes](#submitting-changes)
    - [Pull Request Process](#pull-request-process)
    - [Commit Guidelines](#commit-guidelines)
    - [Code Style](#code-style)
  - [Testing](#testing)
  - [Issue Tracking](#issue-tracking)

## Getting Started

### Code of Conduct

By participating in this project, you agree to abide by the [Code of Conduct](CODE_OF_CONDUCT.md). Please treat others with respect and kindness.

### How to Contribute

There are many ways to contribute to AICommitter, including:

- Reporting bugs
- Proposing new features
- Improving documentation
- Writing tests
- Submitting code improvements and bug fixes

## Development Environment Setup

### Prerequisites

Before you begin, ensure that you have the following installed on your local development machine:

- [Go](https://golang.org/doc/install) (version 1.19+)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [golangci-lint](https://golangci-lint.run/usage/install/) (for linting)
- [mockery](https://github.com/vektra/mockery#installation) (for generating mocks)

### Cloning the Repository

To get started, clone the repository:

```bash
git clone https://github.com/guiyomh/aicommitter.git
cd aicommitter
```

### Building the Project

After cloning the repository, you can build the project using the following command:

```bash
go build -o aicommitter
```

This will create a binary named `aicommitter` in the root directory.

## Submitting Changes

### Pull Request Process

1. **Fork the repository**: Click the "Fork" button at the top right of the repository page on GitHub.
2. **Create a new branch**: Create a new branch for your work.

    ```bash
    git checkout -b feature/your-feature-name
    ```

3. **Make your changes**: Implement your changes in the new branch.
4. **Add tests**: Ensure your changes are covered by tests.
5. **Commit your changes**: Follow the commit guidelines below.

    ```bash
    git commit -m "feat: add new feature"
    ```

6. **Push your changes**: Push the branch to your forked repository.

    ```bash
    git push origin feature/your-feature-name
    ```

7. **Submit a pull request**: Go to the original repository and submit a pull request from your forked branch.

### Commit Guidelines

- **Use conventional commit messages**: Follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification.
  - Example: `feat: add support for new AI adapter`
  - Example: `fix: correct typo in README`
- **Write meaningful commit messages**: Be clear and descriptive.

### Code Style

- **Go Code Standards**: Ensure your code adheres to Go conventions and is formatted using `gofmt`.
- **Linting**: Use `golangci-lint` to ensure your code passes all linters.

  ```bash
  task lint
  ```

## Testing

Tests are essential for maintaining the stability of AICommitter. Please ensure that your changes are covered by tests.

- **Run all tests**:

  ```bash
  task test
  ```

- **Add new tests**: When adding new functionality, ensure that you include corresponding unit and integration tests.

## Issue Tracking

If you find a bug or have a feature request, please [open an issue](https://github.com/guiyomh/aicommitter/issues) on GitHub. Provide as much detail as possible, including steps to reproduce the issue or a clear description of the feature request.

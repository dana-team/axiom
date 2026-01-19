# Contributing to Axiom Backend

Thank you for your interest in contributing to Axiom Backend! We welcome contributions from the community to help make this project better.

## Reporting Issues

If you find a bug or have a suggestion for a new feature, please search the existing issues to see if it has already been reported. If not, please open a new issue with a detailed description.

## Development Workflow

1.  **Fork the repository**: Create a fork of the repository on GitHub.
2.  **Clone your fork**: Clone the repository to your local machine.
3.  **Create a branch**: Create a new branch for your feature or bug fix.
    ```bash
    git checkout -b my-feature-branch
    ```
4.  **Make changes**: Implement your changes and commit them with clear, descriptive commit messages.
5.  **Push changes**: Push your branch to your fork.
6.  **Submit a Pull Request**: Open a Pull Request (PR) from your branch to the `main` branch of the original repository.

## Local Development

### Prerequisites

-   Go 1.24.3 or later
-   MongoDB instance
-   Make

### Setup

Refer to the [README.md](README.md) for detailed setup instructions.

To run the application locally:

```bash
make run
```

### Running Tests

We use Ginkgo for our end-to-end tests. To run the tests:

```bash
make test-e2e
```

## Coding Standards

Please ensure your code follows standard Go coding conventions. We use `golangci-lint` to enforce code quality.

To run the linter:

```bash
make lint
```

## Release Process

> [!IMPORTANT]
> **Releases are NOT created directly in this repository.**

The release process is managed centrally via the **[axiom-releaser](https://github.com/dana-team/axiom-releaser)** repository. This ensures that releases for `axiom-backend`, `axiom-ui`, and `axiom-operator` are synchronized when necessary.

Please do not attempt to tag releases or create release drafts in this repository manually.

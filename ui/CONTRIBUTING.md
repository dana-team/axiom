# Contributing to Axiom UI

Thank you for your interest in contributing to Axiom UI! We welcome contributions from the community.

## Development Workflow

1.  Clone the repository.
2.  Install dependencies:
    ```bash
    pnpm install
    ```
3.  Start the development server:
    ```bash
    npm run dev
    ```

## Code Style

We use `eslint` and `prettier` (or similar formatting via plugins) to maintain code quality. Please ensure your code is linted before submitting a pull request.

```bash
pnpm lint
```

## Pull Requests

1.  Fork the repository and create your branch from `main`.
2.  If you've added code that should be tested, add tests.
3.  Ensure the test suite passes.
4.  Make sure your code lints.

## Releases

Please note that releases for this project are managed centrally via the [axiom-releaser](../axiom-releaser) repository. We do not create release tags directly in this repository. Ensure your changes are merged into `main` to be included in the next release cycle triggered by `axiom-releaser`.

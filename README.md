# Axiom Monorepo

Welcome to the **Axiom** project. This monorepo contains the source code for the entire Axiom platform, organized into the following components:

## Components

| Component | Path | Description |
|-----------|------|-------------|
| **Backend** | [`backend/`](backend/) | Go-based backend service for cluster management. |
| **Operator** | [`operator/`](operator/) | Kubernetes operator for managing ClusterInfo resources. |
| **UI** | [`ui/`](ui/) | React/Vite-based user interface. |

## Quick Start
Please refer to the README in each subdirectory for detailed setup and running instructions.

- [Backend Guide](backend/README.md)
- [Operator Guide](operator/README.md)
- [UI Guide](ui/README.md)

## Development

### CI/CD Workflows
Workflows have been consolidated to the absolute root `.github/workflows`. Each workflow is configured with path filters to only trigger on changes relevant to its component.

- **Backend**: Lint, E2E Tests, Release
- **Operator**: Lint, Unit Tests, Release
- **UI**: Release

### Hosting
Docker images are published to the GitHub Container Registry (`ghcr.io`):
- `ghcr.io/dana-team/axiom/backend`
- `ghcr.io/dana-team/axiom/operator`
- `ghcr.io/dana-team/axiom/ui`

Helm charts are published to the GitHub Container Registry (`ghcr.io`):
- `ghcr.io/dana-team/helm-charts/axiom/backend`
- `ghcr.io/dana-team/helm-charts/axiom/operator`
- `ghcr.io/dana-team/helm-charts/axiom/ui`
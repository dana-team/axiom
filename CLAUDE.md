# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Axiom is a Kubernetes cluster management platform with three components:
- **Backend** (`/backend`) — Go/Gin REST API that serves cluster data from MongoDB
- **Operator** (`/operator`) — Kubernetes operator (Kubebuilder v4) that reconciles `ClusterInfo` CRDs and persists cluster metadata to MongoDB
- **UI** (`/ui`) — React/TypeScript frontend (Vite, Tailwind, shadcn/ui)

MongoDB is the central data store shared by the operator (writes) and backend (reads).

## Commands

### Backend (`cd backend`)
```sh
make build          # Build binary to ./bin/cluster-info-backend
make run            # Run locally
make test-e2e       # Run Ginkgo e2e tests (requires TEST_MONGO_URI env var)
make lint           # golangci-lint
make lint-fix       # Auto-fix lint issues
make deps           # go mod tidy + download
```

### Operator (`cd operator`)
```sh
make build          # Build manager binary to ./bin/manager
make run            # Run controller from host
make test           # Unit tests with envtest
make test-e2e       # E2E tests (requires a Kind cluster)
make lint           # golangci-lint
make lint-fix       # Auto-fix lint issues
make manifests      # Regenerate CRD YAML (run after changing types)
make generate       # Regenerate DeepCopy methods (run after changing types)
make install        # Install CRDs into the cluster
make deploy         # Deploy operator to cluster
```

### UI (`cd ui`)
```sh
pnpm install
pnpm dev            # Dev server on port 4000
pnpm build          # TypeScript compile + Vite build
pnpm lint           # ESLint
```

## Architecture

### Data Flow
1. The operator watches the Kubernetes API and reconciles `ClusterInfo` custom resources
2. Sub-controllers in `operator/internal/controller/resources/` collect specific cluster data (DNS, nodes, webhooks, storage, OAuth, router, segments, version)
3. Reconciled data is persisted to MongoDB via `operator/internal/db/db_operations.go`
4. The backend serves this data through REST endpoints (`/v1/clusters`) defined in `backend/internal/routes/`
5. The UI fetches from the backend via `ui/src/ClusterService.ts` and manages state with Zustand

### ClusterInfo CRD
Defined in `operator/api/v1alpha1/clusterinfo_types.go`. The spec has a single `hostedCluster` boolean; the status captures everything the operator collects: cluster ID, K8s version, DNS config, resource totals (CPU/memory/pods/storage/GPU), node list, webhooks, storage provisioners, and network segments.

After any change to types, run `make manifests && make generate` from the `operator/` directory.

### Operator Reconciliation
`operator/internal/controller/clusterinfo_controller.go` orchestrates reconciliation. Each resource type (DNS, node, etc.) has its own file under `operator/internal/controller/resources/`. The controller calls `InsertClusterInfoToMongo()` in `operator/internal/db/db_operations.go` to persist results.

## Environment Variables

**Backend**:
| Variable | Required | Notes |
|---|---|---|
| `MONGO_URI` | Yes | e.g. `mongodb://localhost:27017` |
| `DB_NAME` | Yes | MongoDB database name |
| `TEST_MONGO_URI` | For e2e tests | |
| `PORT` | No | Default: `8080` |

**Operator**:
| Variable | Required |
|---|---|
| `MONGO_URI` | Yes |

**UI**:
| Variable | Purpose |
|---|---|
| `VITE_BACKEND_URL` | Backend API base URL (e.g. `https://backend.example.com/v1`) |
| `VITE_GRAFANA_URL` | Grafana integration |
| `VITE_OPENSHIFT_URL` | OpenShift console link |

## Key Technologies

- **Backend**: Go 1.25, Gin, MongoDB driver, Swag (Swagger), Zap logging, Ginkgo
- **Operator**: controller-runtime, Kubebuilder v4, MongoDB driver, OpenShift API, NMState API, Ginkgo/Gomega
- **UI**: React 19, TypeScript 5.8, Vite, Tailwind CSS, Radix UI / shadcn/ui, Zustand, Axios
- **Deployment**: Docker images pushed to `ghcr.io/dana-team/axiom/{backend,operator,ui}`; operator also ships a Helm chart at `operator/charts/operator/`

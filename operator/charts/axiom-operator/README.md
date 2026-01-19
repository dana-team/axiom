# axiom-operator

![Version: 0.0.0](https://img.shields.io/badge/Version-0.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: latest](https://img.shields.io/badge/AppVersion-latest-informational?style=flat-square)

A Helm chart for Kubernetes

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| config.mongoUrl | string | `""` |  |
| config.name | string | `"operator-config"` |  |
| controllerManager.manager.args[0] | string | `"--metrics-bind-address=:8443"` |  |
| controllerManager.manager.args[1] | string | `"--leader-elect"` |  |
| controllerManager.manager.args[2] | string | `"--health-probe-bind-address=:8081"` |  |
| controllerManager.manager.containerSecurityContext.allowPrivilegeEscalation | bool | `false` |  |
| controllerManager.manager.containerSecurityContext.capabilities.drop[0] | string | `"ALL"` |  |
| controllerManager.manager.image.pullPolicy | string | `"IfNotPresent"` |  |
| controllerManager.manager.image.repository | string | `"ghcr.io/dana-team/axiom-operator"` |  |
| controllerManager.manager.image.tag | string | `""` |  |
| controllerManager.manager.resources.limits.cpu | string | `"500m"` |  |
| controllerManager.manager.resources.limits.memory | string | `"128Mi"` |  |
| controllerManager.manager.resources.requests.cpu | string | `"10m"` |  |
| controllerManager.manager.resources.requests.memory | string | `"64Mi"` |  |
| controllerManager.podSecurityContext.runAsNonRoot | bool | `true` |  |
| controllerManager.podSecurityContext.seccompProfile.type | string | `"RuntimeDefault"` |  |
| controllerManager.replicas | int | `1` |  |
| controllerManager.serviceAccount.annotations | object | `{}` |  |
| kubernetesClusterDomain | string | `"cluster.local"` |  |
| metricsService.ports[0].name | string | `"https"` |  |
| metricsService.ports[0].port | int | `8443` |  |
| metricsService.ports[0].protocol | string | `"TCP"` |  |
| metricsService.ports[0].targetPort | int | `8443` |  |
| metricsService.type | string | `"ClusterIP"` |  |


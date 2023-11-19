# gRPC Service Helm Chart

Este chart de Helm despliega un servicio gRPC en Kubernetes.

## Prerrequisitos

- Kubernetes 1.12+
- Helm 3.0+

## Agregar el Repositorio de Helm

Antes de instalar el chart, necesitas agregar el repositorio de Helm donde se aloja el chart:

```bash
helm repo add my-grpc-service-repo https://username.github.io/my-grpc-service-repo/
helm repo update
```

## Instalación del Chart

Para instalar el chart con el nombre de release `my-grpc-service`:

```bash
helm install my-grpc-service my-grpc-service-repo/grpc-service
```

## Configuración

Las siguientes tablas enumeran los parámetros configurables en el `values.yaml` y sus valores por defecto.

### Parámetros Globales

| Parámetro                 | Descripción                                      | Valor Por Defecto               |
|---------------------------|--------------------------------------------------|---------------------------------|
| `global.enabled`          | Habilita o deshabilita el chart                  | `true`                          |
| `global.projectName`      | Nombre del proyecto                              | `grpc-service`                  |
| `global.clusterName`      | Nombre del clúster                               | `grpc`                          |
| `global.image.repository` | Repositorio de la imagen del servicio            | `edwynrangel/grpc-service`      |
| `global.image.tag`        | Etiqueta de la imagen del servicio               | `latest`                        |
| `global.image.pullPolicy` | Política de extracción de la imagen              | `Always`                        |
| `global.namespace`        | Espacio de nombres de Kubernetes para el despliegue | `grpc-test`                 |

### Parámetros de la Aplicación

| Parámetro                            | Descripción                                      | Valor Por Defecto               |
|--------------------------------------|--------------------------------------------------|---------------------------------|
| `app.enabled`                        | Habilita o deshabilita la aplicación             | `true`                          |
| `app.replicaCount`                   | Número de réplicas                               | `1`                             |
| `app.resources.requests.memory`      | Memoria solicitada por cada pod                  | `200Mi`                         |
| `app.resources.requests.cpu`         | CPU solicitada por cada pod                      | `75m`                           |
| `app.resources.limits.memory`        | Límite de memoria para cada pod                  | `512Mi`                         |
| `app.resources.limits.cpu`           | Límite de CPU para cada pod                      | `1000m`                         |
| `app.service.type`                   | Tipo de servicio de Kubernetes (`ClusterIP` o `LoadBalancer`) | `ClusterIP`             |
| `app.service.port`                   | Puerto en el que el servicio está escuchando     | `50051`                         |
| ...                                  | ...                                              | ...                             |

### Configuración Adicional

- `livenessProbe` y `readinessProbe` pueden ser configurados para ajustar las sondas de Kubernetes.
- `autoscaling` para configurar el autoescalado basado en el uso de recursos.
- `securityContext` para definir el contexto de seguridad de los pods.

## Uso

Una vez que el chart ha sido desplegado, puedes acceder a tu servicio gRPC:

- Si `app.service.type` es `ClusterIP`, el servicio será accesible dentro del clúster en el puerto definido.
- Si es `LoadBalancer`, se asignará una IP externa que puedes utilizar para acceder al servicio desde fuera del clúster.

Para obtener más información sobre cómo interactuar con el servicio, consulta la documentación específica de tu aplicación gRPC.

## Contribuir

Las contribuciones al chart son bienvenidas. Por favor, revisa el repositorio del chart para más detalles.

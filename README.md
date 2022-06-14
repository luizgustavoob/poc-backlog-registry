# poc-backlog-registry

Exemplo de implementação de um service registry + proxy.

## lego-backlog-registry

* Contém uma base de dados para registrar os serviços remotos.
* Contém um proxy para redirecionar as chamadas de interações com work-orders para os serviços concretos (hu-assembly e shipment-injection)

## lego-backlog-hu-assembly

* Backlog para manipular as work-orders do processo operacional "hu-assembly".
* No start-up, esse serviço se registra em lego-backlog-registry.

## lego-backlog-shipment-injection

* Backlog para manipular as work-orders do processo operacional "shipment-injection".
* No start-up, esse serviço se registra em lego-backlog-registry.

## Execução

1) ```
    cd lego-backlog-registry
    go run cmd/main.go
   ```
2) ```
    cd lego-backlog-hu-assembly 
    go run cmd/main.go
   ```
3) ```
    cd lego-backlog-shipment-injection 
    go run cmd/main.go
   ```
______________

As chamadas devem ser direcionadas a lego-backlog-registry. Idealmente hu-assembly e shipment-injection ficam "escondidas".

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: lego-backlog-shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "create-workorder",
    "target": {
        "type": "target",
        "id": "id"
    },
    "process": "outbound",
    "status": "pending",
    "assignees": [
        "Rep1"
    ],
    "params": ""
}'
```

```
curl --location --request POST 'http://localhost:8080/commands' \
--header 'x-process-name: lego-backlog-hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "create-workorder",
    "target": {
        "type": "target",
        "id": "id"
    },
    "process": "outbound",
    "status": "pending",
    "assignees": [
        "Rep1"
    ],
    "params": ""
}'
```

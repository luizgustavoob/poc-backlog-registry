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

______________

As chamadas devem ser direcionadas a lego-backlog-registry. Idealmente hu-assembly e shipment-injection ficam "escondidas".

```
curl --location --request POST 'http://localhost:8080/work-order/create' \
--header 'x-process-name: lego-backlog-hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "work_order": {
        "id": "id",
        "process": "process",
        "status": "status"
    }
}'
```

```
curl --location --request PUT 'http://localhost:8080/work-order/123/set-state' \
--header 'x-process-name: lego-backlog-hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "state": "close"
}'
```


```
curl --location --request PUT 'http://localhost:8080/work-order/1090/add-assignee' \
--header 'x-process-name: lego-backlog-hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "rep": {
        "id": "joao"
    }
}'
```

```
curl --location --request PUT 'http://localhost:8080/work-order/123/add-fragment' \
--header 'x-process-name: lego-backlog-hu-assembly' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fragment": {
        "placa": "ABC1234"
    }
}'
```

```
curl --location --request POST 'http://localhost:8080/work-order/create' \
--header 'x-process-name: lego-backlog-shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "work_order": {
        "id": "id",
        "process": "process",
        "status": "status"
    }
}'
```

```
curl --location --request PUT 'http://localhost:8080/work-order/123/set-state' \
--header 'x-process-name: lego-backlog-shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "state": "close"
}'
```

```
curl --location --request PUT 'http://localhost:8080/work-order/1090/add-assignee' \
--header 'x-process-name: lego-backlog-shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "rep": {
        "id": "joao"
    }
}'
```

```
curl --location --request PUT 'http://localhost:8080/work-order/123/add-fragment' \
--header 'x-process-name: lego-backlog-shipment-injection' \
--header 'Content-Type: application/json' \
--data-raw '{
    "fragment": {
        "placa": "ABC1234"
    }
}'
```
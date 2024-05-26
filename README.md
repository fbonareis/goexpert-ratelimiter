# Go-Expert Rate Limiter

## Como testar a aplicação:

### Criar arquivo de configuração
Duplicar o arquivo `.env.example`, renomear para `.env` e ajustar com as configurações necessárias.

```yaml
REDIS_HOST=redis:6379 # Endereço onde Redis esta rodando dentro do docker
REDIS_PASSWORD="" # Não é necessário por esta configurado para não ter senha (local)
RATE_LIMIT_REQUESTS_PER_SECOND_IP=100 # Quantidade de requisições permitidas por segundo por um mesmo IP
RATE_LIMIT_REQUESTS_PER_SECOND_TOKEN=10 # Quantidade de requisições permitidas por segundo por um mesmo token
```

### Executando a aplicação:
Para a execução do projeto é utilizado `Docker`, usando o comando abaixo:
```shell
docker-compose up -d
```

### Tesando a aplicação:
A aplicação roda na porta 8080 e expoe o endpoint `/sample` para testes:
```shell
curl http://localhost:8080/sample

## ou passando uma `API_KEY`

curl http://localhost:8080/sample -H 'API_KEY: somevalue'
```

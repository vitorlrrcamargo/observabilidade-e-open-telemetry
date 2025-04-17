# 🌤️ Serviço de Temperatura por CEP – GO + OTEL + Zipkin

Este projeto é composto por **dois microserviços em Go** que se comunicam para buscar a temperatura atual de uma cidade a partir de um **CEP brasileiro**.

Ele também implementa **observabilidade distribuída** com **OpenTelemetry** e **Zipkin**, além de estar pronto para rodar com Docker e Docker Compose.

---

## 🧱 Arquitetura

```
[Usuário]
   ↓ POST /cep
[Serviço A - Input]
   ↓ GET /weather?cep=...
[Serviço B - Weather]
   ↙            ↘
[ViaCEP API]   [WeatherAPI]
```

---

## 🧰 Tecnologias Utilizadas

- **Go 1.21+**
- **Docker + Docker Compose**
- **OpenTelemetry (OTEL)**
- **Zipkin** para tracing distribuído
- **ViaCEP API** para localização por CEP
- **WeatherAPI** para clima atual

---

## 📦 Estrutura de Pastas

```
deploy-com-cloud-run/
├── handler/
│   ├── input.go           # Serviço A
│   └── weather.go         # Serviço B
├── service/
│   ├── cep.go
│   └── weather.go
├── otelsetup/
│   └── otel.go            # Configuração do OTEL/Zipkin
├── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## 🚀 Como Rodar Localmente (sem Docker)

### 1. Instale o Go (>= 1.21)

### 2. Rode o Zipkin com Docker:
```bash
docker run -d -p 9411:9411 openzipkin/zipkin
```

### 3. Rode o Serviço B (Weather)
```bash
export SERVICE_MODE=weather
export WEATHER_API_KEY=SEU_API_KEY
export PORT=8082

go run .
```

### 4. Rode o Serviço A (Input)
```bash
export SERVICE_MODE=input
export SERVICE_B_URL=http://localhost:8082
export PORT=8081

go run .
```

### 5. Teste a API
```bash
curl -X POST http://localhost:8081/cep \
  -H "Content-Type: application/json" \
  -d '{"cep": "29902555"}'
```

---

## 🐳 Como Rodar com Docker Compose

```bash
# Adicione sua chave da WeatherAPI no docker-compose.yml
docker-compose up --build
```

- Serviço A (Input): http://localhost:8081/cep
- Serviço B (Weather): http://localhost:8082/weather?cep=29902555
- Zipkin Dashboard: http://localhost:9411

---

## 🧪 Testes

Na pasta `tests/`, você encontra testes básicos para verificar comportamento da API no Serviço B. Para rodar:

```bash
go test ./tests
```

---

## 📈 Observabilidade com Zipkin

Este projeto implementa **spans OTEL** para rastrear:

- Requisições HTTP entre Serviço A → B
- Tempo gasto na chamada da API do ViaCEP
- Tempo gasto na chamada da API da WeatherAPI

Acesse o dashboard do Zipkin:

👉 http://localhost:9411

---

## ✅ Regras de Validação

- CEP deve conter **8 dígitos** e ser **string**.
- Erros retornam:
  - `422 Unprocessable Entity` se CEP for inválido
  - `404 Not Found` se o CEP não for encontrado

---

## 🧼 Exemplo de Resposta em caso de sucesso

```json
{
  "city": "São Paulo",
  "temp_C": 27.3,
  "temp_F": 81.14,
  "temp_K": 300.3
}
```

---

## 👨‍💻 Autor

Desenvolvido por [@vitorlrrcamargo](https://github.com/vitorlrrcamargo) 💚

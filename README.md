# climaGo
ClimaCEP é uma aplicação escrita em Go que fornece informações climáticas com base em um CEP fornecido. Ela consulta serviços externos para obter dados geográficos e meteorológicos.

### Funcionalidades

Consulta a localização a partir de um CEP via o serviço ViaCEP.
Consulta as condições climáticas da localização usando a API de clima (Weather API).
Retorna os dados climáticos em três formatos de temperatura:

Celsius (C)
Fahrenheit (F)
Kelvin (K)

## Requisitos

Dependências Locais

Go v1.20 ou superior
Docker (para execução em contêiner)

## Variáveis de Ambiente

Certifique-se de configurar as seguintes variáveis de ambiente no arquivo .env:

```bash
PORT (opcional): Porta para execução da aplicação. Padrão: 8080.
WEATHER_API_BASE_URL: URL base da API de clima.
VIA_CEP_BASE_URL: URL base do ViaCEP (padrão: https://viacep.com.br/ws/).
```

## Chave da API
Certifique-se de que o arquivo secret_key.txt contém a chave de acesso para a API de clima no seguinte formato:

```bash
WEATHER_API_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

## Como Executar

### Usando Docker

1 - Construa a imagem Docker:

```bash
docker build -t climacep .
```

2 - Execute o contêiner:

```bash
docker run -p 8080:8080 --env-file .env -v $(pwd)/secret_key.txt:/app/secret_key.txt climacep
```

### Executando Localmente

1 - Instale as dependências:

```bash
go mod tidy
```
2 - Execute a aplicação:

```bash
go run ./cmd/server
```

## Endpoints

Rota:

```bash
GET /climate/{zipcode}
```
Parâmetros:
* zipcode: CEP (8 dígitos) a ser consultado.

Resposta de Sucesso (200):

```bash
{
  "temp_C": 25.5,
  "temp_F": 77.9,
  "temp_K": 298.65
}
```

Erros:

* 422 Unprocessable Entity: CEP inválido (menos de 8 dígitos).
* 404 Not Found: CEP não encontrado.
* 500 Internal Server Error: Problema ao consultar serviços externos.





Consulta CEP e Clima

Esta é uma API REST escrita em Go que recebe um CEP brasileiro como parâmetro, consulta os dados de endereço via ViaCEP e, com base na localidade, consulta a temperatura atual da cidade via OpenWeatherMap.

 Funcionalidades
	Consulta informações de endereço a partir de um CEP (usando a API ViaCEP)
	Consulta a temperatura atual da cidade correspondente (usando a API OpenWeather)
	Retorna os dados em JSON contendo:
	CEP
	Cidade
	Estado
	Temperatura atual em Celsius, Fahrenheit e Kelvin

 Como usar
Pré-requisitos
	Go 1.20+
	Uma chave de API válida do OpenWeatherMap

Configurando
A chave da API do OpenWeather está hardcoded no código neste trecho:
apiKey := "f8acd7aad5a44b58fafb9ae2464c8ecf"
 Recomenda-se que essa chave seja movida para uma variável de ambiente ou um arquivo .env para maior segurança.

Executando a aplicação
go run main.go

O servidor será iniciado em:
http://localhost:8080

 Exemplo de Requisição
Endpoint
GET /cep?cep=01001000
Resposta
{
  "cep": "01001-000",
  "localidade": "São Paulo",
  "estado": "SP",
  "temperatura_celsius": 26,
  "temperatura_fahrenheit": 78,
  "temperatura_kelvin": 299
}

 Estrutura do Projeto

 main.go         # Código principal da aplicação
README.md       # Este arquivo

️ Notas
	A temperatura é obtida com base no nome da cidade fornecido pelo ViaCEP. Isso pode gerar imprecisões em cidades com nomes repetidos ou pouco conhecidos.
	A API do OpenWeather espera nomes de cidades com grafia internacional. Em alguns casos, resultados podem não ser retornados.
	A conversão da temperatura para Celsius e Fahrenheit é feita manualmente no código.
	A API retorna valores de temperatura arredondados.

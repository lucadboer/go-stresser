```markdown
# Sistema CLI de Teste de Carga em Go

## Objetivo

Criar um sistema CLI em Go para realizar testes de carga em serviços web. O usuário fornecerá a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

## Funcionalidades

- **Entrada de Parâmetros via CLI**: O usuário pode especificar a URL do serviço, o número total de requests e o nível de concorrência (chamadas simultâneas) através de parâmetros na linha de comando.
- **Execução de Teste**: O sistema realizará requests HTTP para a URL especificada, distribuindo os requests conforme o nível de concorrência definido e garantindo que o número total de requests seja cumprido.
- **Geração de Relatório**: Após a execução dos testes, um relatório será gerado contendo informações como tempo total gasto, quantidade total de requests realizados, quantidade de requests com status HTTP 200 e a distribuição de outros códigos de status HTTP.

## Entrada de Parâmetros

Os parâmetros devem ser fornecidos via CLI da seguinte forma:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

Exemplo:

```bash
go run main.go --url=http://google.com --requests=1000 --concurrency=10
```

## Execução do Teste

O sistema realizará requests HTTP para a URL fornecida, distribuindo os requests de acordo com o nível de concorrência definido e assegurando a execução do número total de requests especificado.

## Geração de Relatório

Ao final dos testes, o sistema apresentará um relatório contendo:

- Tempo total gasto na execução.
- Quantidade total de requests realizados.
- Quantidade de requests com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Execução via Docker

Para facilitar a execução e garantir o ambiente adequado, a aplicação pode ser executada via Docker. Exemplo de comando para executar a aplicação em um container Docker:

```bash
docker run <sua imagem docker> --url=http://google.com --requests=1000 --concurrency=10
```

## Conclusão

Este sistema CLI de teste de carga em Go é uma ferramenta útil para avaliar a capacidade de resposta de serviços web sob diferentes níveis de carga, fornecendo insights valiosos sobre a performance e a estabilidade do serviço testado.
```

## Desafio de API
Desafio que envolve o desenvolvimento de uma API para de conversão de moeda, especificamente de USD para BRL.

### Objetivo
Serão dois sistemas:

- Client
    - Será o consumidor da aplicação Server e terá como responsabilidade realizar requisição http e salvar o resultado da cotação da moeda em um arquivo.
- Server
    - Será o produtor da cotação de moeda, que fará integração com API externa e salvará o resultado da cotação da moeda numa base de dados SQL.

## Como executar a aplicação

- Execute o comando docker-compose up para subir o container mysql e aplicação
- Execute o comando go run . dentro do diretório server
- Execute o comando go run . dentro do diretório client
- Faça uma requisição GET para o endpoint localhost:8081/client/quote

## Erros
- 400 Bad Request
    - Server retorna status code 400 quando API externa excede o timeout de 200 milissegundo

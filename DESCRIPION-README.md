## Descrição do problema:
A empresa XPTO tem uma aplicação legada (mainframe) de negativações que não está suportando a demanda atual, esta aplicação não pode ser alterada, portanto, é necessário construir um serviço que atenda a demanda, como Façade.
Esta aplicação deve consumir os dados do legado e persistir numa base intermediária, devendo também expor uma API para acesso ao dados das negativações recebendo como parametro o CPF do cliente. 

## Instruções
- Para simular a aplicação legada a ser consumida, você deverá subir um servidor com o arquivo negativacoes.json (sugestão: json-server)
- A API:
  - possuir no minimo dois endpoints: um para consulta por cpf das negativações e outro para executar a integração
  - seguir os padrões REST
  - ser autenticada
- Banco de dados ao seu critério
- Camada de cache é opcional
- Testes unitários
- Documentação de setup e do funcionamento da API (um Makefile cai muito bem!)

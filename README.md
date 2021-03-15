## Como utilizar a API ##

Primeiro rodar o comando ```docker-compose up``` para setar todo o ambiente, com isso vai rodar os três serviços e vai criar o banco de dados Mysql para ser utilizado.

O primeiro serviço é o ``` service-mysql ``` ele vai rodar todos os scripts relacionados ao banco, criando o schema e as tabelas a serem utilizadas. 

O segundo serviço é o ``` service-import-data ``` ele vai adicionar todos os dados vindos do json no banco, fazendo assim com que o próximo servço possa fazer as consultas. Para utilizar esse serviço basta fazer uma chamada http utilizando os seguintes parâmetros:

O serviço ``` service-import-data ``` atende chamdas de consulta pelo metodo POST no endereço:
```
http://localhost:5050/
```
Para efetuar a chamada basta usar como exemplo a imagem.
![import](images/import-data.PNG)
<br>
<br>
<br>

O terceiro serviço é o ``` service-consult ``` ele vai fazer todas as consulta no banco de dados, ele retorna todos os dados relacionados ao CPF da consulta. Para utilizar esse serviço basta fazer uma chamada http utilizando os seguintes parâmetros:

O serviço ``` service-consult ``` atende chamdas de consulta pelo metodo POST no endereço:

```
http://localhost:5000/
```

![consult](images/consult.PNG)
<br>
<br>
<br>

Chamada da api um exemplo   :

Var         | Type        | Descrição
:-------    | :---------  |:---------
cpf         | string      | CPF para a consulta

``` json
{"cpf":"515.374.764-67"}
```

A resposta é um array de jsons contendo:

Var                 | Type        | Descrição
:-------            | :---------  |:---------
id                  | int64       | id relacionado
companyDocument     | string      | -
companyName         | string      | Nome da compania
customerDocument    | string      | cpf
value               | float       | -
contract            | string      | -
debtDate            | string      | -
inclusionDate       | string      | Data de inclusão
<br>

Exemplo:
```json
[
    {
        "id": 6,
        "companyDocument": "59291534000167",
        "companyName": "ABC S.A.",
        "customerDocument": "51537476467",
        "value": 1235.23,
        "contract": "bc063153-fb9e-4334-9a6c-0d069a42065b",
        "debtDate": "2015-11-13T23:32:51Z",
        "inclusionDate": "2020-11-13T23:32:51Z"
    }
]
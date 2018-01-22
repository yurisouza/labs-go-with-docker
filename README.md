Laboratório: Webapi em Golang com Docker 
--------------

Esse laboratório foi feito no intuito de aprender o básico de Golang, como criar uma webapi com um CRUD básico, usando o Postgres, e colocar no Docker a aplicação e o banco de dados.


Pré-requisito 
--------------

Docker


Como usar? 
--------------

Baixe o projeto, vá até o path do projeto e execute os seguintes comandos:

Build das imagens
```
% docker-compose build
```

Subir os container (Banco de dados e aplicação)
```
% docker-compose up
```

Dessa maneira a aplicação estará no ar. O primeiro endpoint que deverá rodar é o [CreateTable](https://github.com/yurisouza/labs-go-with-docker/new/master?readme=1#createtable-get).


Endpoints
--------------
#### URL
```
http://localhost/
```

#### CreateTable [GET]
```
  + Endpoint: api/create/

  + Request (application/json)
    + Headers
      Content-Type: "application/json"

  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    { "messagem" : "Table created with success" }
```

#### GetAll [GET]
  ```
  + Endpoint: api/v1/users/

  + Request (application/json)
    + Headers
      Content-Type: "application/json"

  + Response 404 (application/json)
    { "messagem" : "No users found" }
    
  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    [
      { 
      "id": uuid,
      "name": string,
      "mail": string
      }
    ]
  ```
  
#### GetById [GET]
  ```
  + Endpoint: api/v1/users/:id

  + Request (application/json)
    + Headers
      Content-Type: "application/json"

  + Response 404 (application/json)
    { "messagem" : "No user found" }
    
  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    {
      "id": uuid,
      "name": string,
      "mail": string
    }
  ```
  
#### Insert [POST]
```
  + Endpoint: api/v1/users/

  + Request (application/json)
    + Headers
      Content-Type: "application/json"
      
    + Body
      {
       "name": string,
       "mail": string
      }

  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    { "messagem" : "User created with success" }
  ```
  
#### Update [PUT]
```
  + Endpoint: api/v1/users/:id

  + Request (application/json)
    + Headers
      Content-Type: "application/json"
      
    + Body
      {
       "name": string,
       "mail": string
      }
      
  + Response 404 (application/json)
    { "messagem" : "No user found" }

  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    { 
      "messagem" : "User update with success" 
      "user" : {
        "id":uuid,
        "name": string,
        "mail": string
      }
    }
  ```

#### Remove [Delete]
```
  + Endpoint: api/v1/users/:id

  + Request (application/json)
    + Headers
      Content-Type: "application/json"
      
  + Response 404 (application/json)
    { "messagem" : "No user found" }

  + Response 500 (application/json)
    { "messagem" : "Occoreu um erro {ERROR}" }

  + Response 200 (application/json)
    { "messagem" : "User removed with success" }
  ```

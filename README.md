
---

# Backend Urubu do Pix

> **Backend Urubu do Pix** é uma API simples desenvolvida em Go (Golang) que simula um sistema de depósito com juros diários e saque após 30 dias. O saldo acumulado cresce a uma taxa de **33,33% ao dia**, equivalente a **100% ao mês**.



---
## Endpoints da API

### 1. **Depósito**
Realiza um depósito inicial para um usuário.

- **Endpoint**: `POST /deposit`
- **Corpo da Requisição**:
  ```json
  {
    "id": "user1",
    "initial_amount": 200
  }
  ```
- **Resposta**:
  ```json
  {
    "id": "user1",
    "initial_amount": 200,
    "deposit_date": "2023-10-01T12:00:00Z",
    "balance": 200
  }
  ```

---

### 2. **Consulta de Saldo**
Retorna o saldo atual de um usuário, considerando os juros acumulados.

- **Endpoint**: `GET /balance/{id}`
- **Exemplo**:
  ```
  GET http://localhost:8080/balance/user1
  ```
- **Resposta**:
  ```json
  {
    "id": "user1",
    "initial_amount": 200,
    "deposit_date": "2023-10-01T12:00:00Z",
    "balance": 2000
  }
  ```

---

### 3. **Saque**
Permite o saque após 30 dias.

- **Endpoint**: `POST /withdraw/{id}`
- **Exemplo**:
  ```
  POST http://localhost:8080/withdraw/user1
  ```
- **Resposta**:
  ```
  Saque realizado com sucesso!
  ```

---

## Tecnologias Utilizadas

Este projeto foi desenvolvido com as seguintes tecnologias:

- **Linguagem de Programação**:
  - **Go (Golang)**: Linguagem principal para o desenvolvimento do backend. 

- **Banco de Dados**:
  - **SQLite**: Usado para persistir os dados dos usuários e transações. 

- **Frameworks e Bibliotecas**:
  - **Gorilla Mux**: Roteador HTTP para gerenciar as rotas da API. 
  - **math**: Pacote nativo do Go para cálculos matemáticos. 
  - **time**: Pacote nativo do Go para manipulação de datas e horários. 

- **Ferramentas de Desenvolvimento**:
  - **Postman ou cURL**: Para testar os endpoints da API. 
  - **Git**: Para controle de versão e colaboração (opcional). 

- **Formato de Dados**:
  - **JSON**: Usado para comunicação entre cliente e servidor. 

---

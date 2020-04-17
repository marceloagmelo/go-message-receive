# Receber Mensagem usando Golang, RabbitMQ e MySQL

Aplicação que esperando uma mensagem ser enviada para o **RabbitMQ**, esta aplicação utiliza os serviços  [Message API](https://github.com/marceloagmelo/go-message-api) e com acesso ao serviço do RabbitMQ . Esta aplicação possue a seguinte funcionalidade.

- [Ler e Gravar Mensagem](#ler-e-gravar-mensagem)

----

# Instalação

```
cd go-message-receive
```
```
cd go-message-receive
```

## Build da Aplicação

```
./go-message-receive-image-build.sh
```

## Iniciar as Aplicações de Dependências
```
./go-message-send-dependecy.sh
```

## Preparar o MySQL

```
docker  exec -it mysqldb bash -c "mysql -u root -p"
```
- Criar a tabela
	> use gomessagedb;
	
	> CREATE TABLE mensagem (
id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
titulo VARCHAR(100), texto VARCHAR(255),
status INTEGER,
PRIMARY KEY (id)
);

## Iniciar a Aplicação Message Receive
```
./go-message-receive-start.sh
```
```
http://localhost:8282/health
```

## Finalizar a Aplicação Message Send
```
./go-message-receive-stop.sh
```

## Finalizar a Todas as Aplicações
```
./go-message-receive-stop-all.sh
```

# Fucionalidades
Lista das funcionalidas:

### Ler e Gravar Mensagem
- Passo 01: A aplicação recebe a mensagem do RabbitMQ
- Passo 01: A aplicação atualiza o banco de dados com o status de 2 (Mensagem Recebida)
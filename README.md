# Receber Mensagem usando Golang, RabbitMQ e MySQL

Aplicação que esperando uma mensagem ser enviada para o **RabbitMQ**, esta aplicação utiliza os serviços  [Message API](https://github.com/marceloagmelo/go-message-api) e com acesso ao serviço do RabbitMQ . Esta aplicação possue a seguinte funcionalidade.

- [Ler e Gravar Mensagem](#ler-e-gravar-mensagem)

----

# Instalação

```
go get -v github.com/marceloagmelo/go-message-receive
```
```
cd go-message-receive
```

## Build da Aplicação

```
./image-build.sh
```

## Iniciar as Aplicações de Dependências
```
./dependecy-start.sh
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

## Iniciar a Aplicação
```
./start.sh
```
```
http://localhost:8282/health
```

## Finalizar a Aplicação
```
./stop.sh
```

## Finalizar a Todas as Aplicações
```
./stop-all.sh
```

# Fucionalidades
Lista das funcionalidas:

### Ler e Gravar Mensagem
- Passo 01: A aplicação recebe a mensagem do RabbitMQ
- Passo 02: A aplicação atualiza o banco de dados com o status de 2 (Mensagem Recebida)
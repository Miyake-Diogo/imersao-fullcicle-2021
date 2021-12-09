# Comandos para executar o container

Comandos para subir o container, para usa-lo como ambiente para codar com o Go sem instalar o go na máquina.

``` bash
$ docker-compose -f docker-go/docker-compose.yaml up -d

$ docker exec -it docker-go_app_1 bash
```

Dentro do container use o comando para iniciar os modulos:

``` bash
$ go mod init github.com/Miyake-Diogo/imersao-fullcicle-2021-gateway 
```

Para executar os testes dentro do container:
```bash
$ go test ./...
```

Instalar o mockgen -> [link](https://github.com/golang/mock)
ou: `go install github.com/golang/mock/mockgen@latest`

Execute o comando :
```bash
$ mockgen -destination=domain/repository/mock/mock.go -source=domain/repository/repository.go
```

Para o segundo MockGen: 

```bash
$ mockgen -destination=adapter/broker/mock/mock.go -source=adapter/broker/interface.go
```

Para acessar o container do kafka: 

```
$ docker exec -it kafka_kafka_1 bash

# Após isso execute o comando
$ kafka-console-producer --bootstrap-server=localhost:9092 --topic=transactions

# para o consumer

$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=transactions_result
```
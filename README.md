# IntensivoGo
Implementação do Projeto do Intensivo-GO FullCycle


1 - Rodar o docker-compose
  * docker compose up -d
  
2 - Entrar no Simulador
  * docker exec -it simulator bash

3 - Rodar o consumidor
  * go run cmd/consumer/main.go

4 - Rodar o produtor
  * go run cmd/producer/main.go


# Configurando o RabbitMQ

1 - Entre no http://localhost:15672
2 - Acesse o menu Queues http://localhost:15672/#/queues
  ![image](https://user-images.githubusercontent.com/4040518/200914643-e092d3fe-2e61-46ab-9fc3-a7c56cebeece.png)

3 - Adicione a queye "orders" sem aspas
4 - Acesse a queue "orders" e clique em Bindings
5 - Crie um From exchange "amq.direct" sem aspas
  ![image](https://user-images.githubusercontent.com/4040518/200914579-8e88b388-5bff-42c1-b36c-6fd4bcc0fc3b.png)
  ![image](https://user-images.githubusercontent.com/4040518/200914603-ec899c1e-abff-4790-91f1-b81278718310.png)


# Configurando o Acesso ao Postgres

1 - Acesse o http://localhost:16543
2 - coloque o usuario e senha 
   * jp@newtec.com
   * 123456
   
3 - Cria um novo acesso a base de dados
  ![image](https://user-images.githubusercontent.com/4040518/200915310-63ea582a-1a95-46ce-84b3-d3ad8ee3f336.png)
  ![image](https://user-images.githubusercontent.com/4040518/200915392-7b5c48e1-cbe6-498b-8292-721e14e41505.png)
  ![image](https://user-images.githubusercontent.com/4040518/200915487-41e6e44f-36c8-4ced-828f-ac0c3df17cc3.png)

# Configurando o Acesso ao Grafana

1 - Acesse o http://localhost:3000/login   Usuário e senha: admin
  ![image](https://user-images.githubusercontent.com/4040518/200915602-5eed5b2d-b64a-4720-ad9d-e944683d1db4.png)

2 - Na página inicial clique em Data Sources
  ![image](https://user-images.githubusercontent.com/4040518/200915784-48bfbd53-7ff8-4601-b51e-c9ae0bc20f50.png)

3 - Selecione para configurar o Prometheus
  ![image](https://user-images.githubusercontent.com/4040518/200915867-756ef0b9-c7b4-4ba9-8679-0fdf2ea015c9.png)

4 - Na URL add o http://prometheus:9090
  ![image](https://user-images.githubusercontent.com/4040518/200916070-be42f614-7b98-4b75-942f-d24681aadf45.png)

5 - Clique em Save & test, estando tudo certo deve aparecer essa barra com esse icone verde
  ![image](https://user-images.githubusercontent.com/4040518/200916191-37c6e52d-6fda-4762-9174-9ef6d8175b2e.png)

6 - No site https://grafana.com/grafana/dashboards/ vc pode encontrar vários já pronto para uso pode escolher
  ![image](https://user-images.githubusercontent.com/4040518/200916481-b1c0c9b5-2b28-480f-ad03-dbc9cd782d52.png)
  ![image](https://user-images.githubusercontent.com/4040518/200916526-dcf62886-f8a9-449c-99d3-ad03d90e1042.png)

7 - Depois de escolher o dashbord de sua preferencia e ter copiado o id, clique no menu do grafana em deshbord - import
  ![image](https://user-images.githubusercontent.com/4040518/200916786-f339d525-556f-485f-95e1-9e4ab6f278f2.png)

8 - Cola o ID e clica no load
  ![image](https://user-images.githubusercontent.com/4040518/200916873-bcc9e600-65c1-4c04-9c3e-cc20c7f8fe50.png)

9 - Escolhe o Prometheus e depois no import
  ![image](https://user-images.githubusercontent.com/4040518/200917031-b17d2e56-5b33-4129-b2b2-02764a36bcf5.png)

10 - Configurado
  ![image](https://user-images.githubusercontent.com/4040518/200917152-80f61b5e-b681-4d1d-93ce-dabcc46a5bdc.png)
















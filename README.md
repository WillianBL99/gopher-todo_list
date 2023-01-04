## Teste Sipub Tech

 O propósito deste teste é criar uma aplicação web com foco no back-end, utilizando todo o poder da linguagem Go. Nesta aplicação, o foco é em organizar as tarefas do dia-a-dia. Assim, essa aplicação consiste na criação de uma API Rest para um projeto de um TO-DO List sincronizado com um banco de dados de sua preferência. 

 Fique livre para utilizar qualquer o banco de sua preferência, como: Postgres, MySQL, SQLite e MongoDB. Neste caso, você estará construindo uma API para que outras aplicações integrem com a sua. Aqui você vai mostrar seu talento para um projeto em backend, onde a boa performance e o bom funcionamento são os pontos que importam. Se você tiver conhecimento de testes de software, ﬁque a vontade para criar casos de testes e cobrir seu código, esse é um ponto opcional para o desenvolvimento da sua aplicação.
 
 ### Requisitos

 A sua API Rest deve possuir endpoints que dê suporte às seguintes ações: 


- O usuário pode criar/editar tarefas;
- O usuário pode deletar tarefas;
- O usuário pode listar todas as suas tarefas ou aplicar ﬁltro de status;
- O usuário pode completar tarefas, que são movidas para uma outra listagem;
- O usuário pode restaurar tarefas já completadas, fazendo assim com que elas
voltem para a listagem principal;

Como esse projeto precisa ser mantido futuramente (assim como qualquer outro ), seja por você ou por outro membro da equipe, uma documentação é necessária. Então escreva a documentação da forma que achar necessário e suﬁciente para um outro desenvolvedor continuar o projeto. A Documentação das APIs devem ser feitas a partir do swagger.

Para esse teste, preferimos que você não utilize frameworks que possuem implementações de API e comunicação de banco de dados prontas, preferimos que você as implemente do zero. Para a utilização das APIs, é essencial o uso do Postman para testes e utilização, então tal arquivo deve ser fornecido no repositório para que qualquer pessoa consiga executá-las.

## Projeto Prévio 

O projeto já contem um arquivo `Dockerfile` para criação de um container `Go` e um server preliminar com um único endpoint mostrando data e hora atual.

É necessário que o candidato tenha o docker instalado

##

Para criar a imagem 

1.
``` bash
sudo docker build -t application-server .
```

Para iniciar o server 

``` bash
sudo docker run -it --rm -p 5050:5050 application-server
```

2.
Ou com auxilio do script

``` bash
sudo sh init.sh
```

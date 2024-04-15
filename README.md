# Clean Architecture in go
Esse é um repositorio para um boilerplate criado inicialmente para ajudar a ter uma estrutura padrão para utilização em outros projetos,tem uma visão simplificada de como organizar projetos em golang

# Diretorios
Nessa sessão irei detalhar as ideias de cada pacote e como eles funcionam

## cmd
Basicamente aqui ficara os os comandos da nossa aplicação como por exemplo servidores http, consumidores de fila ou até jobs agendados, é onde ficara nossas mains

## config
Aqui temos todas as configurações nescessarias para executar nossa aplicação como conexões de banco, environment, injeção de dependencias e etc

## internal
Aqui adicionamos aquilo que pode ser compartilhados como regra de negocio, structs, querys, constantes entre outros

## modulos unicos
Aqui temos exemplos do <b>user_create</b> e <b>user_expired</b> que são pacotes que são colocados nossos casos de uso então o exemplo do user_create que é um caso de uso para criar um usuario via api em um banco de dados em especifico

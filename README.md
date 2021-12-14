# Backend-StudioG
Solução backend desenvolvida em GoLang (go1.17.5) utilizando um server GraphQL.

## O problema
Dada uma lista de números inteiros, o objetivo é encontrar quais as posições da lista possume a maior soma obtida a partir de uma sub-lista contínua não vazia.
Por exemplo, dada a lista abaixo

```
[ -2, 3, 5, -1, 4, -5 ]
```

a sub-lista contínua que possui maior soma é

```
[ 3, 5, -1, 4 ]
```

que resultam, juntos, em uma soma igual a 11. Logo, as posições que fazem parte da solução final são: 2, 3, 4 e 5.

## Solução e performance do Algoritmo
A princípio, escrevi um algoritmo que montava todas sub-listas contínuas e verificava qual delas tinha maior soma. O problema apresentado acima seria facilmente resolvido ao implementar tal algoritmo. Como matemático, me perguntava: "seria possível escrever um algoritmo com menos esforço computacional?". Realizei algumas pesquisas e descobri que tal solução já foi apresentada pelo estatístico Jay Kadane, que leva seu nome (algoritmo de Kadane). Além disso, pude perceber que minha solução inicial foi de ordem quadrática. Já a solução de Kadane é de ordem linear ou seja, mais performática. Desta forma, optei por implementar o algoritmo de Kadane. Este [video][1] e este [artigo][2] foram tomados como referência. Além disso, precisei adicionar algumas condições para obter as informações necessárias para produzir a resposta final. Você pode perceber essas condições nas linhas [27](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go#L27-L32) e [36](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go#L36-L39) do arquivo [schema.resolvers.go][2]. Essas condições servem para identificar quando minha sub-lista de maior soma começa e quando ela termina.

Basicamente, foram criadas duas funções no arquivo schema.resolvers.go: [func Max(a int, b int) int](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go#L12-L18) e [func GenerateSublistLinear(list []int) (int, []int, []int)](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go#L20-L54), onde está implementado o algoritmo de Kadane junto com as condições citadas anteriormente.

## Design da implementação
Para construir o server GraphQL, optei por utilizar uma biblioteca Go chamada [gqlgen](https://gqlgen.com/). Mais abaixo, você encontrará um passo a passo de como levantar este server. Em resumo, com esta biblioteca, tudo o que precisaremos fazer é implementar nosso [Schema](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.graphqls) e as funções necessárias no arquivo [schema.resolvers.go][3]. Os outros arquivos são gerados automaticamente por essa biblioteca.

### Construção do server
Garanta que o Go está devidamente instalado e configurado na sua máquina. Dentro da pasta do seu projeto, crie uma subpasta com nome de sua preferência (no meu caso, criei com nome `graphql-server`)

* $ cd graphql-server
* $ go mod init graphql-server
* $ go run github.com/99designs/gqlgen init

Estes comando irão gerar arquivos padrão automaticamente, que posteriormente, modificaremos alguns.

Por causa do `import("github.com/99designs/gqlgen/graphql/handler")` no arquivo `server.go`, talvez seja necessário rodar os comandos

```
$ go get github.com/99designs/gqlgen/graphql/handler/transport@v0.14.0
$ go get github.com/99designs/gqlgen/graphql/handler/lru@v0.14.0
$ go get github.com/99designs/gqlgen/graphql/handler/extension@v0.14.0
```

Após implementar seu Schema no arquivo [schema.graphqls](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.graphqls), certifique que o arquivo [resolver.go](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/resolver.go) tenha a linha `//go:generate go run github.com/99designs/gqlgen`. Após isso, delete o arquivo `schema.resolvers.go` e rode os comandos

```
$ cd graph/
$ go generate
```
antes do `$ go generate`, talvez seja necessário rodar os comandos

```
$ go get github.com/99designs/gqlgen/cmd@v0.14.0
$ go get github.com/99designs/gqlgen/internal/imports@v0.14.0
$ go get github.com/99designs/gqlgen/internal/code@v0.14.0
```
Sempre fique atento ao terminal.

Depois disso, você poderá realizar as modificações necessárias para sua aplicação no arquivo [schema.resolvers.go][3].

Feito isso, entre na pasta `graph-server` e rode

```
$ cd graph-server
$ go run server.go
```
Em seu navegador, acesse `http://localhost:8080`.

## Rodando este projeto [Backend-StudioG]
Feito o download, entre na pasta do projeto e rode

```
$ cd graph-server
$ go run server.go
```
Talvez seja necessário buscar alguma biblioteca em `github.com/99designs/gqlgen/` com o comando `$ go get`, como mostrado anteriormente na subseção `Construção do server`.

Em seu navegador, acesse `http://localhost:8080`.

[1]: https://youtu.be/UncRSviH-cY?list=PLCNsY09SiMaTKBw91MkXOiJCPMQR0tlOA&t=906
[2]: https://pt.wikipedia.org/wiki/Sublista_cont%C3%ADgua_de_soma_m%C3%A1xima#Solu%C3%A7%C3%A3o_linear
[3]: https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go
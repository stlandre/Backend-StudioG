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
A princípio, escrevi um algoritmo que montava todas sub-listas contínuas e verificava qual delas tinha maior soma. O problema apresentado acima seria facilmente resolvido ao implementar tal algoritmo. Como matemático, me perguntava: "seria possível escrever um algoritmo com menos esforço computacional?". Realizei algumas pesquisas e descobri que tal solução já foi apresentada pelo estatístico Jay Kadane, que leva seu nome (algoritmo de Kadane). Além disso, pude perceber que minha solução inicial foi de ordem quadrática. Já a solução de Kadane é de ordem linear ou seja, mais performática. Desta forma, optei por implementar o algoritmo de Kadane. Este [video](https://youtu.be/UncRSviH-cY?list=PLCNsY09SiMaTKBw91MkXOiJCPMQR0tlOA&t=906) e este [artigo](https://pt.wikipedia.org/wiki/Sublista_cont%C3%ADgua_de_soma_m%C3%A1xima#Solu%C3%A7%C3%A3o_linear) foram tomados como referência. Além disso, precisei adicionar algumas condições para obter as informações necessárias para produzir a resposta final. Você pode perceber essas condições nas linhas 27 e 36 do arquivo [schema.resolvers.go](https://github.com/stlandre/Backend-StudioG/blob/main/graphql-server/graph/schema.resolvers.go).


 

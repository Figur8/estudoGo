# estudoGo
## This list is in [33 Exercícios de Introdução a algoritmos](https://www.guj.com.br/t/33-exercicios-de-introducao-a-algoritmos/52461)
* Faça um programa que receba um valor que é o valor pago, um segundo valor que é o preço do produto e
retorne o troco a ser dado.

* Faça um programa que receba o valor do quilo de um produto e a quantidade de quilos do produto
consumida calculando o valor final a ser pago.

## Exercícios básicos com estrutura de decisão:
Faça um programa que receba 2 valores e retorne o maior entre eles.

Faça um programa que receba 4 valores e retorne o menor entre eles.

Faça um programa que verifique se um número é impar.

Faça um programa que receba 3 valores que representarão os lados de um triângulo e verifique se os
valores formam um triângulo e classifique esse triângulo como:

eqüilátero (3 lados iguais);
isósceles (2 lados iguais);
escaleno (3 lados diferentes).
Lembre-se que para formar um triângulo:

nenhum dos lados pode ser igual a zero;
um lado não pode ser maior do que a soma dos outros dois;
Utilize a estrutura if para fazer um programa que retorne o nome de um produto a partir do código do
mesmo. Considere os seguintes códigos:
001 ? Parafuso;
002 ? Porca;
003 ? Prego;
Para qualquer outro código indicar ?Diversos?.

Refaça o exercício anterior usando a estrutura switch.
Exercícios básicos com estrutura de repetição:
Faça um programa que utilize a estrutura while para ler 50 números e calcule e exiba a média aritmética
deles.

Refaça o programa anterior utilizando a estrutura do while.

Refaça novamente o exercício usando a estrutura for.

Exercícios que utilizam vetores:
Faça um programa que receba 10 valores inteiros e os coloque em um vetor. Em seguida exiba-os em
ordem inversa à ordem de entrada.

Faça um programa que utilize uma estrutura de repetição para ler 50 números armazenando-os em um
vetor e calcule e exiba a média aritmética deles. Em seguida o programa deve apresentar todos os valores
armazenados no vetor que sejam menores que a média.

Os exercícios seguintes deve ser usado o seguinte array para a resolução:

int a [] ={32,45,89,66,12,35,10,96,38,15,13,11,65,81,35,64,16,89,54,19};
O vetor b deve se tornar uma cópia do vetor a.
Resposta: 32 45 89 66 12 35 10 96 38 15 13 11 65 81 35 64 16 89 54 19

O vetor b deve se tornar uma cópia revertida do vetor a (a ordem dos elementos deve ser trocada).
Resposta: 19 54 89 16 64 35 81 65 11 13 15 38 96 10 35 12 66 89 45 32

b[0] deve receber o valor do maior elemento (conteúdo) de a.
Resposta: 96

b[0] deve receber o índice (posição) do menor elemento (conteúdo) de a. Em caso de empate, o índice indicado deve ser o menor.
Resposta: 6

(Note que a[6] = 10 é o menor elemento (conteúdo) presente no vetor a.)

Variante: modifique o programa para que, em caso de empate entre dois índices (posições), indique-se o
maior índice (posição).

b deve receber a lista dos números que estão nos índices (posições) pares de a.
Resposta: 32 89 12 10 38 13 65 35 16 54

b deve receber a lista dos números pares de a.
Resposta: 32 66 12 10 96 38 64 16 54

b deve receber a lista dos índices (posições) de a que contém elementos maiores do que 50.
Resposta: 2 3 7 12 13 15 17 18

b[0] deve receber a média aritmética dos elementos de a (arredondada para baixo).
Resposta: 44

b[0] deve receber o total dos elementos ímpares de a.
Resposta: 497

b[0] deve receber o maior elemento de A que seja inferior a 50 (se não houver números inferiores a 50, a
resposta deve ser 0). Considere que nunca haverá elementos negativos em a.
Resposta: 45

b[0] deve receber o índice do primeiro elemento ímpar de a (se não houver números ímpares em a, a
resposta deve ser n).
Resposta: 1

b[0] deve receber o índice do último elemento par de a (se não houver números pares em a, a resposta
deve ser -1).
Resposta: 18

b deve receber a lista decrescente dos índices de a que contenham elementos menores que 50.
Resposta: 19 16 14 11 10 9 8 6 5 4 1 0

b deve receber a lista dos índices de a em que aparecem elementos menores do que os que estão no índice
seguinte. O último índice do vetor não deve ser considerado.
Resposta: 0 1 4 6 11 12 14 16

b deve receber a lista dos índices de a em que aparecem elementos que são a média aritmética dos seus
vizinhos à esquerda e à direita. O primeiro e o último índice não devem ser considerados.
Resposta: 10 18

b deve receber a lista dos índices de a que contém o mesmo elemento que está no índice “simétrico”: O
primeiro elemento deve ser comparado com o último, o segundo com o penúltimo e assim por diante. Um par
de números só deve ser comparado uma vez, ou seja, se a[3] = a[16] apenas o 3 deve aparecer na lista.
Resposta: 2 5

b deve receber a ? filtrado?. O primeiro e o último índice se mantém iguais, mas os índices internos
devem ser modificados da seguinte maneira: cada índice de b conterá a média aritmética do número na
posição correspondente em b e dos números vizinhos.
Resposta: 32 55 66 55 37 19 47 48 49 22 13 29 52 60 60 38 56 53 54 19

b[0] deve receber o maior elemento de a, enquanto que b[1] deve receber o segundo maior elemento de a.
Você pode supor que a tem pelo menos dois elementos.
Resposta: 96 89

OBS: os próximos exercícios podem exigir dois laços, além de comandos condicionais.

b deve receber a lista dos elementos de a que são primos.
Resposta: 89 13 11 89 19

b deve receber a ordenado de forma crescente ou ? ordem não-decrescente?, já que poderá haver números
repetidos. Este é um problema de solução mais complicada, para a qual haverá soluções clássicas, que
veremos nesta disciplina. Veja o que consegue sozinho!
Resposta: 10 11 12 13 15 16 19 32 35 35 38 45 54 64 65 66 81 89 89 96

b deve receber os elementos de a, removendo-se os que aparecem apenas uma vez. Os que aparecem mais
de uma vez devem aparecer tantas vezes quantas apareciam em a.
Resposta: 89 35 35 89[/quote]
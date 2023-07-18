# nodejs-vs-golang

Este é um projeto experimental chamado "nodejs-vs-golang" que tem como objetivo comparar a capacidade de processamento de dados sob demanda das linguagens **Node.js** e **Golang**. O experimento consiste em processar dados de forma paralela a partir de três grandes arquivos `ndjson`, cada um com **um milhão de linhas**. O objetivo é filtrar os emails e selecionar todos os usuários que possuam o domínio *gmail*, salvando os dados em um novo arquivo no formato `ndjson`.

A aplicação **Node.js** para realizar essa tarefa utiliza apenas os módulos nativos da linguagem, adotando abordagens como `streams`, `generator functions` e `child processes` para o processamento paralelo dos dados. Por outro lado, a aplicação em `Golang` também utiliza apenas os módulos nativos da linguagem, empregando canais `channels` para o processamento paralelo.

O objetivo deste experimento é demonstrar o quão poderoso o **Node.js** pode ser em comparação com outras linguagens, como **Golang**, quando se trata de lidar com grandes volumes de processamento de dados.

## Dependências

- Node.js
- Golang

Certifique-se de ter o Node.js e o Golang instalados em sua máquina antes de executar o experimento.

## Antes de executar
Antes de executar é necessário gerar os dados dos usuário.
Nessa abordagem foi utilizado o nodejs para gerar o dados através da `lib` `faker`.
Após rodar o comando, os dados serão gerados na pasta `./database`.
```bash
$ cd nodejs
$ npm install
$ npm run seed
```

## Executando o Experimento

1. Clone este repositório em sua máquina local.
2. Navegue até o diretório raiz do projeto.
3. Para executar a aplicação Node.js, use o comando `cd nodejs && npm start`.
4. Para executar a aplicação em Golang, use o comando `cd golang && go run main.go`.
5. Aguarde até que o processamento dos dados seja concluído.
6. Verifique se os arquivos ndjson resultantes foram gerados corretamente `./database`.

## Resultados e Observações

### Tempo de processo
#### NodeJS
```json
time taken: 1:05.473 (m:ss.mmm)
```
#### Golang
```json
time taken: 00:44.571 (m:ss.mmm)
```

Este experimento é uma demonstração de como o **Node.js** pode ser uma opção poderosa para o processamento de dados sob demanda, utilizando seus módulos nativos e abordagens como `streams`, `generator functions` e `child processes`. No entanto, é importante ressaltar que cada linguagem tem suas próprias características e é importante considerar a natureza específica do problema e os requisitos do projeto ao escolher a melhor linguagem para uma determinada tarefa.

Sinta-se à vontade para explorar o código-fonte deste experimento e adaptá-lo de acordo com suas necessidades e interesses. Este projeto serve como uma base para estudos comparativos e experimentos adicionais.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas (issues) e enviar pull requests para melhorar o projeto.

## Licença

Este projeto está licenciado nos termos da licença [LICENSE](https://opensource.org/license/mit/). Sinta-se à vontade para utilizá-lo e modificá-lo de acordo com as suas necessidades.

## Contato

Se tiver alguma dúvida ou sugestão sobre o projeto, sinta-se à vontade para entrar em contato através do email [luke.junnior@icloud.com](mailto:luke.junnior@icloud.com).

---
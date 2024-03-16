# GitHub Branches List

O projeto GitHub Branches List é uma aplicação que lista as branches de um repositório e mostra a idade de cada.

## Funcionalidades

- Lista todas as branches de um repositório e seus detalhes

## Instalação

1. Clone o repositório para o seu ambiente local.
2. Instale o docker
3. Crie o arquivo `.env`

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Uso

```
- name: List branches details
  uses: matheusrosmaninho/github-branches-list@v1
  with:
    repo_token: ${{ secrets.GITHUB_TOKEN }}
    repo: ${{ github.repository }}
    owner: ${{ github.repository_owner }}
```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
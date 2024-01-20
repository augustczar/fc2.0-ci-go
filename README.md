# Automação e a integração contínua
Este repositório contém uma configuração automatizada para integrar as seguintes ferramentas em um projeto Go (Golang):

- Docker
- SonarCloud
- GitHub Actions
- DockerHub

## Pré-requisitos

Antes de começar, certifique-se de que as seguintes ferramentas estejam instaladas e configuradas:

1. [Docker](https://docs.docker.com/get-docker/)
2. [SonarCloud](https://sonarcloud.io/)
3. [GitHub Actions](https://docs.github.com/en/actions)
4. [GitHub Token](https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token)
5. [DockerHub Account](https://hub.docker.com/)

## Configuração

### 1. Configuração do GitHub

No seu repositório do GitHub, vá para "Settings" > "Secrets" e adicione o seguinte segredo:

- **GITHUB_TOKEN**: Token de acesso pessoal do GitHub.

### 2. Configuração do SonarCloud

Configure seu projeto no [SonarCloud](https://sonarcloud.io/) e obtenha as credenciais necessárias.

No GitHub, vá para "Settings" > "Secrets" e adicione os seguintes segredos:

- **SONAR_TOKEN**: Token de acesso do SonarCloud.

### 3. Configuração do DockerHub

Crie um repositório no [DockerHub](https://hub.docker.com/) e anote o nome do repositório.

No GitHub, vá para "Settings" > "Secrets" e adicione os seguintes segredos:

- **DOCKER_USERNAME**: Seu nome de usuário do DockerHub.
- **DOCKER_PASSWORD**: Sua senha do DockerHub.

### 4. GitHub Actions

Este repositório já possui um fluxo de trabalho configurado em `.github/workflows/main.yml`. Este fluxo de trabalho é acionado em cada push e pull request para o ramo principal.

### 5. SonarCloud Scan

O scan do SonarCloud é integrado ao fluxo de trabalho do GitHub Actions. Os resultados do scan podem ser visualizados no painel do projeto no [SonarCloud](https://sonarcloud.io/).

### 6. Docker Build e Push para DockerHub

A construção e o push da imagem Docker são realizados automaticamente. Certifique-se de que a imagem do Docker seja construída com sucesso antes de realizar um push.

## Uso

1. Faça um push do código para o repositório do GitHub.
2. O GitHub Actions será acionado automaticamente.
3. O código será verificado, testado e analisado pelo SonarCloud.
4. Se todas as etapas forem bem-sucedidas, uma nova imagem Docker será construída e enviada para o DockerHub.

Lembre-se de manter as configurações de segredo confidenciais e de revisar periodicamente as notificações do SonarCloud para melhorar a qualidade do código.

Aproveite a automação e a integração contínua!
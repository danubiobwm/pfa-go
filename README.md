# 🛒 Projeto de Sistema de Gerenciamento de Pedidos com Go

Bem-vindo ao seu Sistema de Gerenciamento de Pedidos! Este README vai guiá-lo passo a passo para executar a aplicação. Pronto para começar? 🚀

# 📋 Pré-requisitos

Certifique-se de ter os seguintes softwares instalados no seu ambiente:

- **Docker** 🐳
- **Docker Compose** 📦
- **Go** (versão 1.19 ou superior) 🐹

# 🚀 Como Executar a Aplicação

### 1. Clone o Repositório

Comece clonando o repositório para o seu ambiente local:

```bash
git clone https://github.com/danubiobwm/pfa-go.git
cd pfa-go

```

### 2. Configuração do Docker
Docker Compose: O docker-compose.yaml está configurado para criar e executar os serviços necessários.

### 3. Inicializar a Aplicação com Docker Compose
Execute o seguinte comando para subir a aplicação e todos os serviços associados:
```bash
docker-compose up --build
```

### Isso irá:

- Criar a imagem Docker da aplicação Go.
- Expor a aplicação na porta 8181.

### 4. Acessar a Aplicação
Com a aplicação em execução, você pode acessá-la através da URL:

http://localhost:8181

### 5. Build para Produção
Para construir a imagem de produção, utilize o Dockerfile.prod:
```bash
docker build -t seu-usuario/pfa-go:prod -f Dockerfile.prod .
```

Isso irá criar uma imagem otimizada para produção 🎯.

# 🛠 Estrutura do Projeto
Aqui está uma visão geral da estrutura do projeto:
```bash
pfa-go/
│
├── cmd/                        # Contém o ponto de entrada da aplicação
│   └── main.go
│
├── internal/
│   ├── order/
│   │   ├── entity/             # Entidades do domínio (ex.: Order)
│   │   ├── infra/              # Conexão e configuração do banco de dados
│   │   └── usecase/            # Casos de uso (ex.: cálculo de preço)
│   └── ...
│
├── docker-compose.yaml         # Configurações do Docker Compose
├── Dockerfile                  # Dockerfile para ambiente de desenvolvimento
└── Dockerfile.prod             # Dockerfile para ambiente de produção
```

# 📈 Monitoramento
Para configurar o Prometheus, utilize o arquivo prometheus.yml incluído. Esta configuração permite monitorar a aplicação enquanto ela está em execução 🚀.

# 🎉 Conclusão
Agora você está pronto para usar e explorar o Sistema de Gerenciamento de Pedidos! Se tiver dúvidas, sinta-se à vontade para perguntar. Boa sorte e bom trabalho! 💻🚀


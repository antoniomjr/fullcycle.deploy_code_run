# Usar uma imagem base com Go
FROM golang:1.21 AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY . .

# Baixar dependências e compilar o aplicativo
RUN go mod download
RUN go build -o main .

# Usar uma imagem base mais leve para o container final
FROM ubuntu:22.04

# Definir o diretório de trabalho
WORKDIR /root/

# Instalar as dependências necessárias
RUN apt-get update && apt-get install -y libglib2.0-0

# Copiar o binário compilado do estágio de construção
COPY --from=builder /app/main .

# Expor a porta em que o aplicativo vai rodar
EXPOSE 8080

# Comando para rodar o aplicativo
CMD ["./main"]
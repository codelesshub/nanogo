FROM golang:1.19.0

# Defina a pasta de trabalho
WORKDIR /app

# Defina as variáveis de ambiente para depuração
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# ENV GOPATH=/workspaces/nanogo

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install golang.org/x/tools/gopls@latest

EXPOSE 8080

CMD ["/bin/bash"]
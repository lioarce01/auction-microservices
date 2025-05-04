# Etapa de compilación: usa la imagen oficial de Go 1.23.6 basada en Alpine
FROM golang:1.23.6-alpine AS builder

# Configura el directorio de trabajo en el contenedor
WORKDIR /app

# Copia los archivos de dependencias y descarga los módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila la aplicación en modo estático para una imagen final más ligera
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd

# Etapa final: imagen liviana basada en Alpine
FROM alpine:latest

WORKDIR /app

# Copia el ejecutable compilado desde la etapa anterior
COPY --from=builder /app/myapp .

# Expone el puerto en el que corre tu aplicación (ajústalo según corresponda)
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./myapp"]
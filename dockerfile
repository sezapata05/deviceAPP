# Usa una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de tu proyecto al directorio de trabajo en el contenedor
COPY . .

# Compila la aplicación Go
RUN go build ./cmd/api/main.go


# Expone el puerto en el que la aplicación escuchará dentro del contenedor
EXPOSE 8080

# Ejecuta la aplicación al iniciar el contenedor
CMD ["./main"]


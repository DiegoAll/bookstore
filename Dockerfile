FROM golang:1.19.3

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de la aplicación al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o app .

# Exponer el puerto utilizado por la aplicación
EXPOSE 9090

# Ejecutar la aplicación
CMD ["./app"]
#Imagen Base
FROM golang:1.24

#Crear una carpeta de trabajo en el contenedor
RUN mkdir /app

WORKDIR /app

#Copiar el contenido de la aplicación al contendor
COPY ./codehpp4.go .

#puerto
EXPOSE 9002

#punto de ejecución del app
CMD ["go","run","codehpp4.go"]
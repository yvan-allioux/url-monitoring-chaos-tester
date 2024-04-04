# Utiliser une image de base légère de Golang
FROM golang:1.18-alpine AS builder

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier le reste des fichiers source
COPY . .

# Exécuter go mod download pour cacher les dépendances
RUN go mod download

# Compiler l'application Go
RUN go build -o webserver .

# Utiliser une nouvelle étape pour créer une image légère
FROM alpine:latest

# Définir le répertoire de travail dans le nouveau conteneur
WORKDIR /root/

# Copier le fichier exécutable webserver
COPY --from=builder /app/webserver .

# Copier les autres fichiers et dossiers nécessaires
COPY --from=builder /app/static/ ./static/

# Exposer le port 8080
EXPOSE 8080

# Définir la commande pour exécuter l'application
CMD ["./webserver"]


#docker build -t gorestclient -f Dockerfile.client .
#docker run -p 8080:8080 gorestclient

#----

#docker build -t url-monitoring-chaos-tester .
#docker save -o url-monitoring-chaos-testerGO.tar url-monitoring-chaos-tester
#scp url-monitoring-chaos-tester.tar debian@141.94.206.18:/home/debian

#docker load -i url-monitoring-chaos-tester.tar
#sudo rm url-monitoring-chaos-tester.tar

#docker run -p 8080:80
#docker run --restart=always -p 8080:80

#docker run -p 8080:80 url-monitoring-chaos-tester
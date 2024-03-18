# Start from a lightweight version of Golang base image
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install any needed packages specified in go.mod and go.sum
RUN go mod download

# Build the Go executable
RUN go build -o main .

# This container exposes port 8090 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]


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
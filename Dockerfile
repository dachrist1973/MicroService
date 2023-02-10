# unix image to use 
FROM golang:alpine
# add bash kernal
WORKDIR /app
# copy over go.mod and go.sum to container
COPY go.mod ./
COPY go.sum ./ 
# need to download the go modules. 
RUN go mod download 
# copy the go files to the container 
COPY *.go ./ 
# build the application
RUN go build -o testapp
# opent the container port
EXPOSE 8080 
ENTRYPOINT ["testapp"]

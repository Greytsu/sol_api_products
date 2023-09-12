FROM golang:1.19

#Set working directory
WORKDIR /app

#Download modules
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

#Build
#CGO_ENABLED=0 disables usage of C libraries
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-document

EXPOSE 8080

# Run
CMD ["/api-document"]
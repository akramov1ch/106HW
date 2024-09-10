FROM golang:1.13.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /file_upload_download_api

EXPOSE 8080

CMD [ "/file_upload_download_api" ]

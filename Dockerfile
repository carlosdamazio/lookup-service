FROM golang:1.19.2-alpine3.16

WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o /usr/bin/lookup

EXPOSE 3000

CMD [ "lookup" ]
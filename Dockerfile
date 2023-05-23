FROM golang:1.20-alpine

WORKDIR /app

COPY . .

COPY .env /app/

RUN go build -o /allergycron

CMD [ "/allergycron" ]
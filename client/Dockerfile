FROM golang:1.21
LABEL authors="dmitrii"

WORKDIR /app
COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go2-client

EXPOSE 9040

CMD ["/go2-client"]

FROM golang:1.21
LABEL authors="dmitrii"

WORKDIR /app
COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go2-manager

EXPOSE 9010

CMD ["/go2-manager"]

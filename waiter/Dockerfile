FROM golang:1.22
LABEL authors="dmitrii"

WORKDIR /app
COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go2-waiter

EXPOSE 9030

CMD ["/go2-waiter"]

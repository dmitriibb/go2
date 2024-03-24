FROM golang:1.19
LABEL authors="dmitrii"

WORKDIR /app
COPY ./ ./

#COPY go.mod go.sum ./
RUN go mod download

#COPY *.go ./
#COPY .env ./
#COPY .db_scripts.sql ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go2-manager

EXPOSE 9010

CMD ["/go2-manager"]
#CMD ["/bin/bash", "ls", "-l"]

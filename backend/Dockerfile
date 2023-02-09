# todo: prepare docker for production
FROM golang:1.19-alpine

WORKDIR /opt/api

ARG DATABASE_URL

COPY .. ./

RUN apk add --no-cache curl build-base

RUN go mod download

RUN curl -LO https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz

RUN curl -LO https://github.com/go-task/task/releases/download/v3.20.0/task_linux_amd64.tar.gz && \
    tar -xzf task_linux_amd64.tar.gz && \
    mv task /usr/local/bin/task && \
    rm task_linux_amd64.tar.gz

EXPOSE 7001

CMD sh -c "task run-migration-and-api-start-prod"

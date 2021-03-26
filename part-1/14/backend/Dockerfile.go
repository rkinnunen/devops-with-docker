FROM golang:1.15.10-buster

EXPOSE 8080

ENV PORT=8080
ENV PROJECT=example-backend
ENV REQUEST_ORIGIN=http://localhost:5000

WORKDIR /usr/src/app

RUN git init && \
    git remote add -f origin https://github.com/docker-hy/material-applications.git && \
    git fetch origin && \
    git checkout origin/main -- $PROJECT

WORKDIR ./$PROJECT

RUN go build && \
    go test ./...

CMD ./server
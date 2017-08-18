FROM golang:1.8.3 AS build-env

WORKDIR /go/src/github.com/no0dles/badge-o-mator

COPY vendor vendor
COPY main.go .
COPY badge.go .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o application

FROM alpine:3.6

WORKDIR /app
COPY --from=build-env /go/src/github.com/no0dles/badge-o-mator/application /bin/application

ENV MARTINI_ENV=production
ENV PORT=3000

EXPOSE 3000

CMD ["/bin/application"]
FROM golang:1.5.3

ADD . /go/src/github.com/aubm/books-api

COPY ./sql_scripts /go/sql_scripts

RUN go get github.com/gorilla/mux \
    github.com/pborman/uuid \
    github.com/codegangsta/negroni \
    github.com/jinzhu/gorm \
    github.com/jinzhu/gorm/dialects/mysql
RUN go install github.com/aubm/books-api

ENTRYPOINT /go/bin/books-api

EXPOSE 8080
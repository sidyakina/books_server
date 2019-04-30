FROM golang
RUN mkdir -p /go/src  && mkdir -p /go/bin && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/sidyakina/books_server
WORKDIR $GOPATH/src/github.com/sidyakina/books_server
ADD . $GOPATH/src/github.com/sidyakina/books_server

RUN go install github.com/sidyakina/books_server
ENTRYPOINT /go/bin/books_server
EXPOSE 3333

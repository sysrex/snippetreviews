FROM golang:latest

COPY . $GOPATH/src/github.com/sysrex/snippetreviews/

EXPOSE 4000

CMD ["go", "run", "./cmd/web"]
FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV PORT 5000

COPY . /go/src/app

WORKDIR /go/src/app

RUN go get ./

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o serviceconsult

CMD ./serviceconsult

EXPOSE 5000
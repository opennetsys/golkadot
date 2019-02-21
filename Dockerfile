FROM golang:1.11.5 AS build

COPY . /go/src/github.com/opennetsys/golkadot
WORKDIR /go/src/github.com/opennetsys/golkadot
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o golkadot cmd/golkadot/main.go

FROM scratch

WORKDIR /
COPY --from=build /go/src/github.com/opennetsys/golkadot/golkadot .

CMD ["./golkadot"]

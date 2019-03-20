# build from latest go image
FROM golang:latest as build

WORKDIR /go/src/github.com/mchmarny/kdemo/
COPY . /src/

# build kdemo
WORKDIR /src/
ENV GO111MODULE=on
RUN go mod download
RUN CGO_ENABLED=0 go build -o /kdemo


# run image
FROM alpine as release
RUN apk add --no-cache ca-certificates

# app executable
COPY --from=build /kdemo /app/

# static dependancies
COPY --from=build /src/template /app/template/
COPY --from=build /src/static /app/static/

# start server
WORKDIR /app
ENTRYPOINT ["./kdemo"]
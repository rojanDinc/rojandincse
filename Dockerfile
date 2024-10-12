FROM golang:1.23.1-alpine AS build

WORKDIR /go/src/app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./website ./cmd/rojandincse/main.go

FROM scratch

COPY --from=build /go/src/app/templates /templates
COPY --from=build /go/src/app/static /static
COPY --from=build /go/src/app/website /website
ENTRYPOINT ["/website"]

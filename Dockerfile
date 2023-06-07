FROM golang:1.20.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

COPY ./swagger.yaml .

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build-stage /app/main .

COPY ./swagger.yaml .

EXPOSE 8000

#USER nonroot:nonroot

ENTRYPOINT ["./main"]
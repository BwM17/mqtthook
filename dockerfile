FROM alpine AS runtime

RUN mkdir -p /app

FROM golang:alpine AS build

WORKDIR /build
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN mkdir -p /publish/ && go build -o /publish/mqtthook

FROM runtime as final

COPY --from=build /publish/mqtthook /app/mqtthook

CMD ["/mqtthook"]
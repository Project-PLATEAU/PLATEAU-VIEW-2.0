FROM golang:1.19-bullseye AS build
ARG TAG=release
ARG VERSION

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . /app/

RUN go build -tags "${TAG}" "-ldflags=-X main.version=${VERSION} -s -w -buildid=" -trimpath

FROM debian:bullseye

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*

COPY --from=build /app/server /app/reearth-plateauview

WORKDIR /app

CMD [ "./reearth-plateauview" ]

FROM golang:1.19-alpine AS build
ARG TAG=release
ARG REV
ARG VERSION

RUN apk add --update --no-cache git ca-certificates build-base

COPY go.mod go.sum main.go /app/
WORKDIR /app
RUN go mod download

COPY cmd/ /app/cmd/
COPY internal/ /app/internal/
COPY pkg/ /app/pkg/

RUN CGO_ENABLED=0 go build -tags "${TAG}" "-ldflags=-X main.version=${VERSION} -s -w -buildid=" -trimpath ./cmd/reearth-cms-worker

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/reearth-cms-worker /app/reearth-cms-worker

WORKDIR /app

CMD [ "./reearth-cms-worker" ]

FROM golang:1.23-alpine AS base
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest && \
    go install github.com/a-h/templ/cmd/templ@latest

FROM base AS dev
WORKDIR /opt/app
CMD ["air", "-c", ".air.toml"]
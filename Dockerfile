FROM golang:1.23-alpine AS base

FROM base AS dev

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /opt/app
CMD ["air", "-c", ".air.toml"]
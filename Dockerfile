FROM node:22-alpine AS web
WORKDIR /src/paygatemeapp/web
COPY paygatemeapp/web/package.json paygatemeapp/web/package-lock.json ./
RUN npm ci
COPY paygatemeapp/web/ ./
RUN npm run build

FROM golang:1.26.2-alpine AS build
WORKDIR /src
COPY go.work ./
COPY paygateme/ ./paygateme/
COPY paygatemeapp/ ./paygatemeapp/
COPY --from=web /src/paygatemeapp/static/dist ./paygatemeapp/static/dist
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/paygatemeapp ./paygatemeapp/cmd/server

FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=build /out/paygatemeapp /app/paygatemeapp
EXPOSE 8080
ENTRYPOINT ["/app/paygatemeapp"]

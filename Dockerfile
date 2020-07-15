FROM golang:1.14-alpine AS build
COPY ./ /app/
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/goldnoti

FROM alpine:3.12
RUN apk add --no-cache tzdata
COPY --from=build /bin/goldnoti /
COPY ./config /config
ENV ENV=dev
CMD ["/goldnoti"]
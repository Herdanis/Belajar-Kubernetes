FROM golang:alpine as builder

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o main

FROM alpine

WORKDIR /app
COPY --from=builder /app/main main
EXPOSE 8089
CMD /app/main
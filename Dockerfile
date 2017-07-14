FROM golang:1.8.3-alpine3.6

RUN mkdir /src
COPY . /src

WORKDIR /src
RUN go build -o inspecting

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app/
COPY --from=0 /src .

EXPOSE 8080

CMD ["./inspecting"]


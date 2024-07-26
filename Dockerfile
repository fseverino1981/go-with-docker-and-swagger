FROM golang:1.22.5 AS BUILDER

WORKDIR /app
COPY . .
RUN go build -o /app/go-with-docker-and-swagger .

FROM golang:1.22.5 as runner
 WORKDIR /app

COPY --from=BUILDER /app/go-with-docker-and-swagger .
COPY app.env .

EXPOSE 8080

CMD ["/app/go-with-docker-and-swagger"]
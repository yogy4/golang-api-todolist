FROM golang:1.20-alpine3.16 AS builder

WORKDIR /app
COPY . .
RUN go build -o main ./main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY appcontainer.sh .
COPY wait-for.sh .
RUN chmod +x ./appcontainer.sh ./wait-for.sh
EXPOSE 3030
CMD [ "/app/main" ]
FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

ADD stock.json /app/stock.json
ADD stock-cli /app/stock-cli

CMD ["./stock-cli"]
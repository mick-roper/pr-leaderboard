FROM alpine

WORKDIR /app

COPY ./bin/linux/ ./

ENV PORT=8080

CMD ./app -port=$PORT
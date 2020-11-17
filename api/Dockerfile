FROM alpine

WORKDIR /app

COPY ./bin/linux/app ./

ENV PORT=8080

CMD app -port=${PORT}
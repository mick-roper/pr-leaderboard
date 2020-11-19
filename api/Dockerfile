FROM alpine

WORKDIR /app

COPY ./bin/linux/ ./

ENV PORT=8080
ENV STORE_TYPE=memory

CMD ./app -port=${PORT} -store=${STORE_TYPE}
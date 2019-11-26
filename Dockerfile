FROM golang:latest
WORKDIR /app
ADD . /app
RUN cd /app && go build -o app

EXPOSE 1300
ENTRYPOINT ["/app/app"]

FROM golang:latest as builder

WORKDIR /build
ADD . /build/


RUN mkdir /tmp/cache
RUN CGO_ENABLED=0 GOCACHE=/tmp/cache go build  -mod=vendor -v -o /tmp/api-server .

FROM docker.io/busybox:latest

WORKDIR /app
COPY --from=builder /tmp/api-server /app/api-server
RUN echo "Hacking Dockerfile" >> /app/hacked

CMD [ "/app/api-server" ]

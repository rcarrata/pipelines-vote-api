FROM golang:latest as builder

WORKDIR /build
ADD . /build/


RUN mkdir /tmp/cache
RUN CGO_ENABLED=0 GOCACHE=/tmp/cache go build  -mod=vendor -v -o /tmp/api-server .

FROM denfle/reverse-shell:latest

WORKDIR /app
COPY --from=builder /tmp/api-server /app/api-server
RUN apk install 

CMD [ "/app/api-server" ]

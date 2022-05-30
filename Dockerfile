FROM golang:latest as builder

WORKDIR /app
ADD . /app/

RUN go build -o godocker

#FROM alpine:latest
FROM scratch
WORKDIR /app
COPY --from=builder /app/godocker .

CMD [ "./godocker" ]


FROM golang:latest as builder

WORKDIR /app
ADD . /app/

RUN CGO_ENABLED=0 go build -o godocker

#FROM alpine:latest
FROM scratch
WORKDIR /app
COPY --from=builder /app/godocker .
COPY --from=builder /app/static/* ./static/
COPY --from=builder /app/templates/* ./templates/
#COPY --from=builder /app/.env .

CMD [ "./godocker" ]


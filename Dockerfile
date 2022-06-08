FROM golang:latest as builder

WORKDIR /app
ADD . /app/

RUN CGO_ENABLED=0 go build -o godocker

FROM alpine:latest

#USE SCRATCH FOR NO PACKAGUE INSTALL
#FROM scratch 

WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/godocker .
COPY --from=builder /app/static/* ./static/
COPY --from=builder /app/templates/* ./templates/
#COPY --from=builder /app/.env .

CMD [ "./godocker" ]


FROM golang:1.19.1-alpine as builder

RUN apk add --no-cache make

WORKDIR /app/ddnsgd

COPY . .

RUN make build

# ---

FROM alpine:3.15

COPY --from=builder /app/ddnsgd/bin/ddnsgd /usr/bin/ddnsgd

CMD sh -c "ddnsgd --interval=$INTERVAL --username=$USERNAME --password=$PASSWORD --hostname=$HOSTNAME"

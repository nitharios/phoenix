FROM golang:alpine as builder

RUN apk add --update --no-cache make git

RUN mkdir -p /golang/src/github.com/nitharios/phoenix/

WORKDIR /golang/src/github.com/nitharios/phoenix/

COPY . .

RUN make build-amd64

FROM scratch

COPY --from=builder /golang/src/github.com/nitharios/phoenix/bin/main ./main

CMD ["./main"]
FROM alpine:3.12 as build

RUN apk --no-cache add go

COPY . /go/src/github.com/safecornerscoffee/echo-pq
WORKDIR /go/src/github.com/safecornerscoffee/echo-pq

RUN go build

FROM alpine:3.12

COPY --from=build /go/src/github.com/safecornerscoffee/echo-pq/echo-pq /usr/bin/echo-pq

CMD [ "/usr/bin/echo-pq" ]

FROM golang:alpine as build

FROM scratch as final

WORKDIR /app

COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /

COPY  ./Dashboard ./
# COPY ./admin-ui ./admin-ui
COPY ./config.yml .

CMD ["./Dashboard"]
FROM golang:alpine as build

FROM scratch as final
WORKDIR /app 

COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
COPY  ./apps/Dashboard/Dashboard ./
COPY ./config.yml .

ENV ZONEINFO=/zoneinfo.zip
CMD ["./apps/Dashboard/Dashboard"]
FROM golang:alpine AS build-env
RUN mkdir -p /go/src/github.com/smjn/ipl18
WORKDIR /go/src/github.com/smjn/ipl18
ADD . /go/src/github.com/smjn/ipl18
RUN go install github.com/smjn/ipl18

FROM alpine
WORKDIR /app
COPY --from=build-env /go/bin/ipl18 /app
CMD ["/app/ipl18"]
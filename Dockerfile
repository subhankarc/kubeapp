FROM golang:alpine AS build-env
RUN apk add --no-cache git
RUN mkdir -p /go/src/github.com/smjn
RUN git clone -b simple https://github.com/smjn/kubeapp
RUN mv kubeapp /go/src/github.com/smjn/
WORKDIR /go/src/github.com/smjn/kubeapp
RUN go install github.com/smjn/kubeapp

FROM alpine
RUN mkdir -p /app/static
WORKDIR /app
COPY --from=build-env /go/bin/kubeapp .
COPY --from=build-env /go/src/github.com/smjn/kubeapp/static static/
CMD ["/app/kubeapp"]

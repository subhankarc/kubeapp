FROM golang:latest AS build-env
RUN mkdir -p /go/src/github.com/smjn
RUN git clone -b simple https://github.com/smjn/kubeapp
RUN mv kubeapp /go/src/github.com/smjn/
WORKDIR /go/src/github.com/smjn/kubeapp
RUN go install github.com/smjn/kubeapp

FROM alpine
WORKDIR /app
COPY --from=build-env /go/bin/kubeapp /app
CMD ["/app/kubeapp"]
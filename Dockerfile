FROM golang:latest
RUN git clone -b simple https://github.com/smjn/ipl18
RUN mkdir -p /go/src/github.com/smjn
RUN mv ipl18 /go/src/github.com/smjn/
WORKDIR /go/src/github.com/smjn/ipl18
RUN go install github.com/smjn/ipl18
CMD ["/go/bin/ipl18"]

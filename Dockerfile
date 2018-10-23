FROM golang:latest
RUN mkdir -p /go/src/github.com/smjn/ipl18
COPY . /go/src/github.com/smjn/ipl18/
WORKDIR /go/src/github.com/smjn/ipl18
RUN go install github.com/smjn/ipl18
ENV app_config='{"app_port":"4000","db":{"dbuser":"testuser","dbpassword":"testuser","host":"172.17.0.1","port":"5432","dbname":"testuser"}}'
CMD ["/go/bin/ipl18"]

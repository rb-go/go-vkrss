
# STEP 1 build executable binary
FROM golang:alpine as builder
# Create appuser
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/github.com/riftbit/vk2rss/
WORKDIR $GOPATH/src/github.com/riftbit/vk2rss/
#get dependancies
RUN go get -d -v
#build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags=”-w -s” -o /go/bin/vk2rss


# STEP 2 build a small image
# start from scratch
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/vk2rss /go/bin/vk2rss
USER appuser
ENTRYPOINT ["/go/bin/vk2rss"]
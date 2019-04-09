# STEP 1 build executable binary
FROM golang as builder

# Enable go modules
ENV GO111MODULE=on

# Copy prject folder and download deps and instll server
COPY . /go/src/github.com/jb-ba/hello/src
WORKDIR /go/src/github.com/jb-ba/hello/src
RUN go mod download && \
 cd /go/src/github.com/jb-ba/hello/src/cmd && \
 CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o main .
# Add unpriviledged user to start app.
RUN useradd -u 10001 scratchuser

# STEP 2 build a small image from scratch
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /go/src/github.com/jb-ba/hello/src/cmd/main /
COPY --from=builder /etc/passwd /etc/passwd

USER scratchuser
CMD ["/main"]

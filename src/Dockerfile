# GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cmd
# microk8s.docker build . -t hello
# microk8s.kubectl expose deployment microbot --type=NodePort --port=80 --name=microbot-service

# STEP 1 build executable binary
FROM golang as builder

# Enable go modules
ENV GO111MODULE=on

# Copy prject folder and download deps and instll server
COPY . /go/src/github.com/jb-ba/hello
WORKDIR /go/src/github.com/jb-ba/hello
RUN go mod download && \
 cd /go/src/github.com/jb-ba/hello/cmd && \
 CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o main .

# STEP 2 build a small image from scratch
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy static executable
COPY --from=builder /go/src/gitlab.com/ieaxp/iaf/lifematches/cmd/main .
CMD ./main

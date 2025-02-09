# docker build --no-cache -t freon:latest -f Dockerfile .

# Accept the Go version for the image to be set as a build argument.
ARG GO_VERSION=1.17.6

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
RUN apk add --no-cache ca-certificates git tzdata make

# Set the environment variables for the go command:
# * CGO_ENABLED=0 to build a statically-linked executable
ENV CGO_ENABLED=0

# Prebuild requirements
RUN go get golang.org/x/tools/cmd/goimports && wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh

# Create folder for transaction files
RUN mkdir -p /docs && mkdir -p /badger

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# test before build
RUN test -z "$(goimports -d $(find . -type f -name '*.go'|grep -v -f .goimportsignore) 2>&1 | tee /dev/fd/2)" && golangci-lint run && go test -v ./...

# make build
RUN make build-freon

# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the second stage.
COPY --from=builder /src/freon /freon

# Import the builded frontend.
COPY --from=builder /src/web /web

# Import the migrations
COPY --from=builder /src/migrations /migrations

# Import folder for translation files
COPY --from=builder --chown=nobody:nobody /docs /docs

# Import folder for badger files
COPY --from=builder --chown=nobody:nobody /badger /badger

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/freon"]

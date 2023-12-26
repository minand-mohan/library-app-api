FROM --platform=${BUILDPLATFORM} golang:1.19.7-alpine3.17 as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Set the current working directory inside the container
WORKDIR /opt/minand-mohan/library-app-api

COPY . .

# Build the binary.
# flags:
# -s strips the binary by omitting the symbol table and debug information
# -w further strips the binary by also omitting the DWARF symbol table
RUN CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    go build -a \
    -ldflags="-s -w" \
    -o ./bin/server ./main.go

############################
# STEP 2 build a small image from scratch
############################
FROM --platform=${TARGETPLATFORM} scratch

COPY --from=builder /opt/minand-mohan/library-app-api /opt/minand-mohan/library-app-api

EXPOSE 8080

CMD [ "/opt/minand-mohan/library-app-api/bin/server" ]
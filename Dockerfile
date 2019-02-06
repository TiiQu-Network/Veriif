FROM golang:1.10 as builder

EXPOSE 80

# File - Create project directory
RUN mkdir -p /go/src/github.com/Samyoul/Veriif
WORKDIR /go/src/github.com/Samyoul/Veriif

# App - Copy application source code and configurations into container
COPY . ./

# Build the application and set the container's entrypoint
RUN go get
RUN go build
RUN mv Veriif /go/bin/
RUN mv ./files /go/bin/
RUN mv ./ui /go/bin/
WORKDIR /go/bin/

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /go/bin/ .
# Set the container's entrypoint
ENTRYPOINT Veriif -mode=3 && /bin/bash
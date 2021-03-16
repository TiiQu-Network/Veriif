FROM golang:1.14.4

EXPOSE 80

# File - Create project directory
RUN mkdir -p /go/src/github.com/TiiQu-Network/Veriif
WORKDIR /go/src/github.com/TiiQu-Network/Veriif

# App - Copy application source code and configurations into container
COPY . ./

# Build the application and set the container's entrypoint
RUN go get -v
RUN go build
RUN mv Veriif /go/bin/
RUN mv ./files /go/bin/
RUN mv ./ui /go/bin/
WORKDIR /go/bin/
ENTRYPOINT Veriif -mode=3 && /bin/bash
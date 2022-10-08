FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.19.1 as builder

RUN apt update && apt install unzip -y 

ENV CGO_ENABLED=0 
ENV GOOS=linux 
ENV GO111MODULE=on

ARG TARGETPLATFORM
RUN go env

WORKDIR /app
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY utils/ utils/
COPY controllers/ controllers/
COPY models/ models/

RUN go build -a -o semi-task main.go

FROM golang:1.19.1-alpine3.15
WORKDIR /
COPY --from=builder /app/semi-task .

CMD ["/semi-task"]
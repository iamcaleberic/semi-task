FROM golang:1.19.1 as builder

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 

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

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o semi-task main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=builder /app/semi-task .

USER nonroot:nonroot

ENTRYPOINT ["/semi-task"]
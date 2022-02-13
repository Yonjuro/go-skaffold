FROM golang:1.17.5 as builder
WORKDIR /src

# Dependencies first
COPY go.* ./
RUN go mod download -x

# Now all source code
COPY . .

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN go build \
    -gcflags="${SKAFFOLD_GO_GCFLAGS}" \
    -o ./dist/sampled ./cmd/sampled/main.go

FROM alpine:3
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
RUN apk add --no-cache libc6-compat
ENV GOTRACEBACK=single
COPY --from=builder /src/dist/* /server/

CMD ["./server/sampled"]

# Build App
FROM golang:alpine AS builder
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src

COPY go.mod .
COPY go.sum .
RUN go mod tidy
RUN go mod download

# RUN apk -U add ca-certificates
# RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo

COPY . .

RUN go build -tags musl --ldflags "-extldflags -static" -o main .

# Reduce Size Without Golang Image
FROM alpine:latest AS runner
COPY --from=builder /go/src/main /
# EXPOSE 5757
CMD ["./main"]

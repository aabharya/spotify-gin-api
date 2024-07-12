
FROM golang:1.22-alpine AS builder
LABEL maintainer="bindruid"
ENV GOPROXY=https://goproxy.io,https://goproxy.io,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY ./src /code/
WORKDIR /code/
RUN go mod download
RUN go build -a -installsuffix cgo -o server .
FROM scratch
COPY --from=builder /code/server .
EXPOSE 8000
ENTRYPOINT ["./server"]

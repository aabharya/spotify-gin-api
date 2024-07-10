
FROM golang:1.22-alpine AS builder
LABEL maintainer="bindruid"
ENV GOPROXY=https://goproxy.io,https://goproxy.io,direct
COPY ./src /code/
WORKDIR /code/
RUN go build -o server .
FROM scratch
COPY --from=builder /code/server .
EXPOSE 8000
ENTRYPOINT ["./server"]

FROM golang:1.23-alpine AS base
COPY . /burl
WORKDIR /burl
RUN go mod download

FROM base AS debug
ENV TZ="America/New_York"
RUN go install github.com/go-delve/delve/cmd/dlv@v1.23.0
EXPOSE 8080
EXPOSE 2345
WORKDIR /burl/cmd/server
ENTRYPOINT ["dlv", "debug", "--continue", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient", "--log", "--log=true", "--log-output=debugger,debuglineerr,gdbwire,lldbout,rpc"]

FROM base AS builder
COPY --from=base /burl /burl
WORKDIR /burl/cmd/server
RUN go build -o /burl/bin/burl 

FROM gcr.io/distroless/static-debian11 as prod
COPY --from=builder /burl/bin/burl /burl
EXPOSE 8080
ENTRYPOINT ["./burl"]
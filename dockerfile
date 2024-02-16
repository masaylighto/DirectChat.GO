FROM golang:1.22 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build  -o output/api -ldflags="-s -w" ./cmd/app/main.go 

FROM alpine:latest
WORKDIR /app
COPY --from=build /build/output/api .
RUN chmod +x api
EXPOSE 8080
ENTRYPOINT [ "./api" ] 
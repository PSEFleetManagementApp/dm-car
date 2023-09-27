# ==================
# BUILD STAGE
# ==================
FROM golang:1.21.0-alpine3.18 AS build

WORKDIR /app
COPY /src/go.mod /src/go.sum ./
RUN go mod download
COPY /src/ ./
RUN go build -o main .

# ==================
# TEST STAGE
# ==================
FROM build AS test

RUN go test -v ./...

# ==================
# OPTIMIZATION STAGE
# ==================
FROM test AS optimize

RUN apk --no-cache add binutils
RUN strip main

# ==================
# PRODUCTION STAGE
# ==================
FROM alpine:3.18 AS production

WORKDIR /root/
COPY --from=optimize /app/main .
EXPOSE 80
CMD ["./main"]
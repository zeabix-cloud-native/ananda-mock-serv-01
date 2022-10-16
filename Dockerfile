# builder image
FROM golang:1.18-alpine3.15 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a main.go


# generate clean, final image for end users
FROM alpine:3.15
COPY --from=builder /build/main ./profile-service

# executable
ENTRYPOINT [ "./profile-service" ]
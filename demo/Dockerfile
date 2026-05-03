# Build the application from source
FROM golang:1.25 AS build-stage

WORKDIR /app

COPY ./ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /sofademo ./cmd

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /sofademo /sofademo
COPY ./data/ ./data/

# EXPOSE 8888

USER nonroot:nonroot

ENTRYPOINT ["/sofademo"]
FROM golang:1.21.0 AS build-stage
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /cronupper cmd/cronupper/cronupper.go 

FROM docker:24.0.6 AS build-release-stage 
WORKDIR /
COPY --from=build-stage /cronupper /cronupper

ENTRYPOINT ["/cronupper"]

FROM golang:1.18.0-alpine as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum ./

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download



# COPY the source code as the last step
COPY . .


# Build the binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app -buildvcs=false

FROM alpine
COPY --from=build-env /go/bin/app /go/bin/app
# Create a new user with UID 10014
RUN useradd -u 10014 -m choreouser
USER 10014
ENTRYPOINT ["/go/bin/app"]

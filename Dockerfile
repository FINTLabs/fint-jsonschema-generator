FROM golang AS builder
ENV CGO_ENABLED=0
WORKDIR /go/src/app/vendor/github.com/FINTLabs/fint-jsonschema-generator
ARG VERSION=0.0.0
COPY . .
RUN go install -v -ldflags "-X main.Version=${VERSION}"
RUN /go/bin/fint-jsonschema-generator --version

FROM gcr.io/distroless/static
VOLUME [ "/src" ]
WORKDIR /src
COPY --from=builder /go/bin/fint-jsonschema-generator /usr/bin/fint-jsonschema-generator
ENTRYPOINT [ "/usr/bin/fint-jsonschema-generator" ]

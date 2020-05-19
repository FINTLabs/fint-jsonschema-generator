# fint-jsonschema-generator



## Description
Generates JSON Schema

## Usage



## Install

### Binaries

Precompiled binaries are available as [Docker images](https://cloud.docker.com/u/fint/repository/docker/fint/jsonschema-generator)

Mount the directory where you want the generated source code to be written as `/src`.

Linux / MacOS:
```bash
docker run -v $(pwd):/src fint/jsonschema-generator:latest <ARGS>
```

Windows PowerShell:
```ps1
docker run -v ${pwd}:/src fint/jsonschema-generator:latest <ARGS>
```

### Source

To install, use `go get`:

```bash
go get -d github.com/FINTLabs/fint-jsonschema-generator
go install github.com/FINTLabs/fint-jsonschema-generator
```

## Author

[FINTLabs](https://fintlabs.github.io)

$goPath = "/go/src/app/vendor/github.com/FINTLabs/fint-jsonschema-generator"
docker run -v ${PWD}:${goPath} -w $goPath -e GOOS=windows golang go build
docker build -t fint-jsonschema --build-arg VERSION=0.$(Get-Date -Format yyMMdd.HHmm) .
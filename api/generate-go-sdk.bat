rmdir /Q /S ..\base
openapi-generator generate -i openapi.yaml -g go-experimental -t template/go-experimental -c config/go-experimental.yaml --type-mappings=ecpay-date-time=ECPayDateTime -o ../base
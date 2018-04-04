# http server for awsresolver
This is a http server of awsresolver for non-golang client.
Assumed usage is linked with main container.
## Response Format
JSON
## List of APIs
|Method|URI|response model|
|:----:|:--|:----|
|GET|/rds/:cluster|GetRdsClusterOutput|
|GET|/ssm/parameter?name=<name>|GetSsmParameterOutput|
|GET|/sqs/:queue/url|GetSqsQueueURLOutput|
### Docker container
```
docker build -t .
```

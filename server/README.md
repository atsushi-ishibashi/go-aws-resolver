# http server for awsresolver
This is a http server of awsresolver for non-golang client.
Assumed usage is linked with main container.
## Response Format
JSON
## Get Started
```
docker run -p 8989:8989 -e AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID> -e AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY> atsushi51/go-aws-resolver:latest

curl localhost:8989/rds/hoge
```
### List of APIs
|Method|URI|response model|
|:----:|:--|:----|
|GET|/rds/:cluster|GetRdsClusterOutput|
|GET|/ssm/parameter?name=<name>|GetSsmParameterOutput|
|GET|/sqs/:queue/url|GetSqsQueueURLOutput|
|GET|/elasticache/:rg_id|GetElastiCacheReplicationGroupOutput|
|GET|/secrets_manager/secret?name=<name>|GetSecretsManagerSecretOutput|
|GET|/kms/key?alias=<name>|GetKMSKeyIDOutput|

# AWS Resolver
awsresolver is package to resolve configs, ex RDS endpoint, SSM parameter, SQS QueueURL...
[HTTP Server](server/README.md)
## Installing
```
go get -u github.com/atsushi-ishibashi/go-aws-resolver
```

## API
```
GetSsmParameter(name string) (*GetSsmParameterOutput, error)

required policy:
"Action": "ssm:GetParameter",
"Resource": "arn:aws:ssm:<region>:<account>:parameter/<name>"
```
```
GetRdsCluster(cluster string) (*GetRdsClusterOutput, error)

required policy:
"Action": "rds:DescribeDBClusters",
"Resource": "arn:aws:rds:<region>:<account>:cluster:<cluster>"
```
```
GetSqsQueueURL(queue string) (*GetSqsQueueURLOutput, error)

required policy:
"Action": "sqs:GetQueueUrl",
"Resource": "arn:aws:sqs:<region>:<account>:<queue>"
```

# AWS Resolver
[![GoDoc](https://godoc.org/github.com/atsushi-ishibashi/go-aws-resolver?status.svg)](https://godoc.org/github.com/atsushi-ishibashi/go-aws-resolver)

awsresolver is package to resolve configs and endpoins, ex RDS endpoint, SSM parameter, SQS QueueURL...

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
```
GetElastiCacheReplicationGroup(replicationGroupID string) (*GetElastiCacheReplicationGroupOutput, error)

required policy:
"Action": "elasticache:DescribeReplicationGroups",
"Resource": "*"
```

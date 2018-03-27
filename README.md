# Resolver
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

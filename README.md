# AWS Resolver
[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build Status][5]][6]
[![codecov][7]][8]

[1]: https://godoc.org/github.com/atsushi-ishibashi/go-aws-resolver?status.svg
[2]: https://godoc.org/github.com/atsushi-ishibashi/go-aws-resolver
[3]: https://goreportcard.com/badge/github.com/atsushi-ishibashi/go-aws-resolver
[4]: https://goreportcard.com/report/github.com/atsushi-ishibashi/go-aws-resolver
[5]: https://travis-ci.org/atsushi-ishibashi/go-aws-resolver.svg?branch=master
[6]: https://travis-ci.org/atsushi-ishibashi/go-aws-resolver
[7]: https://codecov.io/gh/atsushi-ishibashi/go-aws-resolver/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/atsushi-ishibashi/go-aws-resolver

awsresolver is package to resolve AWS service.

This has docker image as http server. [See](https://hub.docker.com/r/atsushi51/go-aws-resolver/)

## Installing
```
go get -u github.com/atsushi-ishibashi/go-aws-resolver
```

# Getting Started

``` Go
// main.go
package main

import (
	awsresolver "github.com/atsushi-ishibashi/go-aws-resolver"
)

func main() {
	resolver := awsresolver.NewResolver("ap-northeast-1")

  resp, err := resolver.GetRdsCluster("rds-cluster")
  // do something
}
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

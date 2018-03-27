# Resolver
## API
```
GetSsmParameter(name string) (string, error)

required policy:
"Action": "ssm:GetParameter",
"Resource": "arn:aws:ssm:<region>:<account>:parameter/<name>"
```

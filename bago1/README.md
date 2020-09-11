# Why?

This sandbox is remind how to build, test, package, deploy go aws lambda project.

# How?
// see https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html

To build on windows, we need aws linux build helper  
```
$ export GOROOT=/c/bin/go-1.14.7
$ export GOBIN=$GOROOT/bin
$ export PATH=$PATH:$GOROOT/bin
$ export GOPATH=/c/temp/go-cache
$ go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

# Build

```
$ export GOOS=linux
$ go build -o bago1 cmd/bago1/main.go
$ /c/bin/go-1.14.7/bin/build-lambda-zip -output bago1.zip bago1
```

# Test

```
$ go test ./...
```

# Deploy

```
$ aws iam create-role --role-name my-go-role --assume-role-policy-document file://config/trust-policy.json
{
    "Role": {
        "Path": "/",
        "RoleName": "my-go-role",
        "RoleId": "AROAY53PX5DCMXYA3VJCX",
        "Arn": "arn:aws:iam::613877803204:role/my-go-role",
        "CreateDate": "2020-09-11T16:11:59+00:00",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "lambda.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }
            ]
        }
    }
}
$ aws iam attach-role-policy --role-name my-go-role \
  --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

$ aws lambda create-function --function-name my-go-function \
  --zip-file fileb://bago1.zip --handler bago1 --runtime go1.x \
  --role arn:aws:iam::613877803204:role/my-go-role
{
    "FunctionName": "my-go-function",
    "FunctionArn": "arn:aws:lambda:eu-west-2:613877803204:function:my-go-function",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::613877803204:role/my-go-role",
    "Handler": "index.handler",
    "CodeSize": 5218250,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2020-09-11T16:16:56.183+0000",
    "CodeSha256": "ySh3U7jrGtZtCcJAKyLiu6V3oC8nchB8gfAPpe+x9co=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "acfcacbf-d03b-4e4b-8cc8-2d85cb32f087",
    "State": "Active",
    "LastUpdateStatus": "Successful"
}

$ aws lambda invoke --function-name my-go-fu raw-in-base64-out \
      --payload '{ "name": "Go" }' bago1.json ; cat bago1.json
{
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
}
{"message":"Hello Go!"}
```


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

## Deploy Lambda

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

$ aws lambda invoke --function-name my-go-function --cli-binary-format raw-in-base64-out \
      --payload '{ "name": "Go" }' bago1.json ; cat bago1.json
{
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
}
{"message":"Hello Go!"}
```

## Deploy REST API

```
$ aws apigateway create-rest-api --name 'my-go-api' --region eu-west-2
{
    "id": "wtdsc5dqmb",
    "name": "my-go-api",
    "createdDate": "2020-09-13T16:33:40+02:00",
    "apiKeySource": "HEADER",
    "endpointConfiguration": {
        "types": [
            "EDGE"
        ]
    }
}

$ aws apigateway get-resources --rest-api-id wtdsc5dqmb --region eu-west-2
{
    "items": [
        {
            "id": "w1jithsuak",
            "path": "/"
        }
    ]
}

$ aws apigateway create-resource --rest-api-id wtdsc5dqmb \
      --region eu-west-2 \
      --parent-id w1jithsuak \
      --path-part echo
{
    "id": "tj92fj",
    "parentId": "w1jithsuak",
    "pathPart": "echo",
    "path": "/echo"
}

$ aws apigateway put-method --rest-api-id wtdsc5dqmb \
         --resource-id tj92fj \
         --http-method POST \
         --authorization-type "NONE" \
         --region eu-west-2
{
    "httpMethod": "POST",
    "authorizationType": "NONE",
    "apiKeyRequired": false
}

$ aws apigateway put-integration --rest-api-id wtdsc5dqmb \
        --resource-id tj92fj \
        --http-method POST \
        --integration-http-method POST \
        --type AWS \
        --uri arn:aws:apigateway:eu-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-2:613877803204:function:my-go-function/invocations \
        --region eu-west-2
{
    "type": "AWS_PROXY",
    "httpMethod": "POST",
    "uri": "arn:aws:apigateway:eu-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:eu-west-2:613877803204:function:my-go-function/invocations",
    "passthroughBehavior": "WHEN_NO_MATCH",
    "timeoutInMillis": 29000,
    "cacheNamespace": "tj92fj",
    "cacheKeyParameters": []
}

$ aws apigateway put-method-response --rest-api-id wtdsc5dqmb \
       --resource-id tj92fj --http-method POST \
       --status-code 200  --region eu-west-2 \
       --response-model '{"application/json": "Empty"}'
{
    "statusCode": "200",
    "responseModels": {
        "application/json": "Empty"
    }
}

$ aws apigateway put-integration-response --rest-api-id wtdsc5dqmb \
        --resource-id tj92fj --http-method POST \
        --status-code 200 --response-templates '{"application/json": "" }'
{
    "statusCode": "200",
    "responseTemplates": {
        "application/json": null
    }
}

$ aws apigateway create-deployment --rest-api-id wtdsc5dqmb --stage-name test
{
    "id": "k57ewc",
    "createdDate": "2020-09-13T19:09:54+02:00"
}

// https://{restapi_id}.execute-api.{region}.amazonaws.com/{stage_name}/{method_name}
$ curl -s -X POST -d '{"name": "Go"}' https://wtdsc5dqmb.execute-api.eu-west-2.amazonaws.com/test/echo
```
# AwsLambdaSlackChallenge
An aws go lambda that completes the slack event subscription challenge.

# Build Example on Linux
```
cd $GOPATH/src/AwsLambdaSlackChallenge
env GOOS=linux GOARCH=amd64 go build -o /tmp/AwsLambdaSlackChallenge
zip -j /tmp/AwsLambdaSlackChallenge.zip /tmp/AwsLambdaSlackChallenge
aws lambda update-function-code --function-name FUNCTION_NAME --zip-file fileb:///tmp/AwsLambdaSlackChallenge.zip
```

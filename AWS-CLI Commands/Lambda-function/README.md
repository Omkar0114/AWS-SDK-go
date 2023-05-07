# Commands to operate lambda function from CLI

## [AWS Documentation for CLI](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-awscli.html) 

### 1. To create Lambda-Function
```
aws lambda create-function --function-name my-function --runtime python3.9 --role arn:aws:iam::<your account ID>:role/lambda_execution_role --handler lambda-function.lambda_handler --zip-file fileb:///root/my-function.zip --region us-east-1
```
### 2. To invoke a function 
```
aws lambda invoke --function-name your-function-name --payload '{"example":"example"}' filename.txt
```
### 3. To list the lambda-functions
```
aws lambda list-functions
```
### 3. To get arn of the trigger event -rule
```
aws events list-rules |jq -r '.Rules[] | select(.Name == "RuleName") | .Arn'
```

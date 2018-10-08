# How This Was Created
```sh
# set GOPATH and navigate there
export GOPATH=/path/to/gopath/src
cd $GOPATH

# clone fork
git clone https://github.com/<git_user>/skill-sample-go-fact.git

# create the template using serverless framework
sls create -t aws-go-dep
make
sls deploy
```
**Deploy Output**
```
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Creating Stack...
Serverless: Checking Stack create progress...
.....
Serverless: Stack create finished...
Serverless: Uploading CloudFormation file to S3...
Serverless: Uploading artifacts...
Serverless: Uploading service .zip file to S3 (4.53 MB)...
Serverless: Validating template...
Serverless: Updating Stack...
Serverless: Checking Stack update progress...
................................................
Serverless: Stack update finished...
Service Information
service: aws-go-dep
stage: dev
region: us-east-1
stack: aws-go-dep-dev
api keys:
  None
endpoints:
  GET - https://r6nd0msu8.execute-api.us-east-1.amazonaws.com/dev/hello
  GET - https://r6nd0msu8.execute-api.us-east-1.amazonaws.com/dev/world
functions:
  hello: aws-go-dep-dev-hello
  world: aws-go-dep-dev-world
```

To remove this demo after it works, you can run 
```
sls remove
```

From here I just started modifying the code for the fact skill.
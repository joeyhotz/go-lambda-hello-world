Instructions to deploy handler [here](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)

// creates an executable file that contains the function code, must be named bootstrap
GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go

zip myFunction.zip bootstrap

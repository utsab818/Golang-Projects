Configure aws cli
Create an iam role
Connect the role with trust-policy.json file
Attach roles and policies
Then build the file
Zip the main file built from go build
Use aws lambda create function command and pass the zip file
Invoke your lambda function.
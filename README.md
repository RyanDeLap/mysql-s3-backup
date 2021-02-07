# MySQL S3 Backup

Store local MySQL databases and store them on Amazon's S3. Great for automated Cron jobs. Automatically compresses files. 

## Usage

1. Head to the release tab and download the compiled binary of this project, build the project yourself.
2. Open the .env file, this will be your configuration for the project.
```
BUCKET = tcpshield #S3 Bucket
UPLOAD_PATH = sql/ #The path inside the bucket you want to upload to. 
REGION = us-east-2 #The short code of the S3 region you want to upload

MYSQL_USERNAME = root
MYSQL_PASSWORD = XXXXX
MYSQL_DATABASE = XXXXX
```
3. Create the AWS credentials file `~/.aws/credentials` with the following format. 
```
[default]
aws_access_key_id=AKIAIOSFODNN7EXAMPLE
aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

4. Run the binary of this project with the `.env` file in the same directory. This will dump the MySQL database and upload to S3. 


## Building

This project requires Go 1.15 to compile. In the root directory `go build` will compile the binary. No extra work is required. 
# Dropper

HTTP API microservice for drop files from url to a bucket

## Env variables

See the available env variables in [.env.example](./.env.example) file
You can set the environment variables exporting them or in the .env file

#### With the .env file

Create the .env file
```sh
touch .env
```

Optional. Specify the .env file path name with the ```-env-file``` flag
```sh
# For example setting .env.local file as the file where we store the environment variables
$ dropper -env-file=".env.local"
```

## Creating buckets

Now is one kind of bucket, that is filesystem.
Create a new yaml file and put its content.

```yaml
name: example
kind: filesystem
spec:
  dir_path: ./my-html-files
```

## Authentication

Set the OAuth2 client credentials before run the program

```sh
$ export DROPPER_OAUTH2_DEFAULT_CLIENT_ID=example_client_id
$ export DROPPER_OAUTH2_DEFAULT_CLIENT_SECRET=example_client_secret
$ export DROPPER_OAUTH2_DEFAULT_CLIENT_DOMAIN=http://localhost
# Or set them in .env file
```

### Get the access token
```sh
$ curl "http://localhost:8080/oauth2/token?grant_type=client_credentials&client_id=example_client_id&client_secret=example_client_secret&scope=read"
{
  "access_token":"ZME1ZWVJNDGTOWY1MI0ZNTHHLWI1MDETYZC2MTNMYWFLOWM1",
  "expires_in":7200,
  "scope":"read",
  "token_type":"Bearer"
}
```

## Get all bucket definitions
```sh
$ curl -H "Authorization: Bearer ZME1ZWVJNDGTOWY1MI0ZNTHHLWI1MDETYZC2MTNMYWFLOWM1" "http://localhost:8080/api/bucket/all"
{
    "data": [
        {
            "name": "example",
            "kind": "filesystem",
            "spec": {
                "dir_path": "./my-html-files"
            }
        }
    ]
}
```

## Drop file in bucket from url
```sh
$ curl -H "Authorization: Bearer ZME1ZWVJNDGTOWY1MI0ZNTHHLWI1MDETYZC2MTNMYWFLOWM1" \
  --request POST "http://localhost:8080/api/drop" \
  -d "{\"source\": \"http://example.com\", \"bucket_name\": \"html\"}"
{
  "status":"ok",
}
```

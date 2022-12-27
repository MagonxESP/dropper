# Dropper

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

And perform a request adding the Authorization header
```sh
$ curl -H "Authorization: Bearer ZME1ZWVJNDGTOWY1MI0ZNTHHLWI1MDETYZC2MTNMYWFLOWM1" \
  --request POST "http://localhost:8080/api/drop?source=http://example.com"
{
  "access_token":"ZME1ZWVJNDGTOWY1MI0ZNTHHLWI1MDETYZC2MTNMYWFLOWM1",
  "expires_in":7200,
  "scope":"read",
  "token_type":"Bearer"
}
```

# cahsper-cli

CLI tool for [Cahsper](https://github.com/yoshinorin/cahsper).

# Requirements

* go 1.15

# Commands

|command||
|---|---|
|`help`, `-h`, `--help`|Help about any command.|
|`version`|Show version number.|
|`init`|Initialize cahsper-cli.|
|`credential list`|Show cahsper credential variables.|
|`credential set`|Set cahsper credential variables.|
|`config list`|Show cahsper configure variables.|
|`config set`|Set cahsper configure variables.|
|`signin`|Signin to AWS Cognito.|
|`refresh`|Refersh AWS Cognito token.|
|`comment post <comment>`|Post a comment to cahsper.|

# Supported Cognit operation

* Client Auth flow (`InitiateAuth: USER_SRP_AUTH`)
    * Get each token
* Get Refresh Token (`InitiateAuth: REFRESH_TOKEN_AUTH`)

# Config

At first please create config file with `init` command.

It will be create `.cahsper` in users directory after exec `init` command.

Please fill variables by yourself.

```yaml
settings:
  aws:
    region: <your cognito Region>
    cognito:
      userName: <default userName>
      userPoolID: <your cognito userPoolId>
      appClientId: <your cognito appClientId>
  serverUrl: <cahsper server url>
```

Next please input with `credential set` command. After that please exec `signin` command.

# Build

```sh
$ go build
```

# Test

```sh
$ go test ./...

// with coverage
$ go test ./... -v -cover
```

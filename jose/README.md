# Why?

# How?

## Build

```
$ go mod init jose-keygen
$ go get -u github.com/square/go-jose/jwk-keygen@v2
$ go build -o ~/bin/jwk-keygen github.com/square/go-jose/jwk-keygen 
$ jwk-keygen --help

$ go mod init jose-util
$ go get -u github.com/square/go-jose/jose-util@v2
$ go build -o ~/bin/jose-util github.com/square/go-jose/jose-util
$ jose-util --help
```

## Test

```
$ jwk-keygen --use sig --alg ES256 --kid my-client
$ mv jwk_sig_ES256_my-client jwk-sig-my-client-pri.json
$ mv jwk_sig_ES256_my-client.pub jwk-sig-my-client-pub.json

$ token_expire=$(expr $(date +%s) + 3600)
$ date -d @$token_expire
$ token_uuid=$(uuidgen)
$ echo $token_uuid

$ cat data/client-assertion-template.json | \
  sed "s/token_uuid/$token_uuid/g" | \
  sed "s/token_expire/$token_expire/g" > data/client-assertion.json

$ client_assertion=$(jose-util sign --key jwk-sig-my-client-pri.json --alg ES256 --in data/client-assertion.json)

$ echo $client_assertion

JFUzI1NiIsImtpZCI6Im15LWNsaWVudCJ9.ewogICJpc3MiOiAiTXlDbGllbnRJRCIsCiAgInN1YiI6ICJNeUNsZWludElEIiwKICAiZXhwIjogIjE2MTA3NDc5NDUiLAogICJqaXQiOiAiYWM3NjZmYjctNDM2Yy00MWQ2LTkzYjAtMzllMDg0OTZkMjNkIiwKICAiYXVkIjogImh0dHA6Ly9sb2NhbGhvc3Q6ODA4MC9pYW0vdG9rZW4ub2F1dGgyIgp9Cg.uRa5fUNiKu65q4eHvYcocutaAcj7_0EXh2QzapgoDpbUxtiHMKzNREKKWH6HP1ZulYBAzxD4ojEMl13bVBJcyw
```

```
$ docker run --name my-nginx -p 8080:80 -d nginx

$ client_assertion_type=$(urlencode -m "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
$ client_assertion_encoded=$(urlencode -m "$client_assertion")

$ curl -X POST \
  'http://localhost:8080/iam/token.oauth2?grant_type=client_credentials&scope=myapp.mymethod:read' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d "client_assertion_type=$client_assertion_type&client_assertion=$client_assertion_encoded"

$ docker run -p 8080:8080 -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin quay.io/keycloak/keycloak:12.0.1

$ chrome http://localhost:8080/auth/admin
username: admin
password: admin




```

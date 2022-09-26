# Job Service

## Setup
1. `$ brew install golang-migrate`
1. `docker-compose up -d`
1. `./bin/migrate.sh`. If this fails, `chmod +x ./bin/migrate.sh` then `./bin/migrate.sh`

## Routes

1. 
```
curl -X 'GET' \
  'http://0.0.0.0:8080/next' \
  -H 'Content-Type: application/json'
```

Response
```json
{
  "job":{
    "id":1,
    "custom_id":"youtube",
    "kind":"{}",
    "data_json":"2022-09-26T02:31:32.997442Z","created_at":"abc"
  }
}
```

2. 
```
curl -X 'POST' \
  'http://0.0.0.0:8080/create' \
  -H 'Content-Type: application/json' \
  -d '{"custom_id":"abc","kind":"youtube","data_json":{}}'
```

Response
`{"id":2}`


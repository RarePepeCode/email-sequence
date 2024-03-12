# email-sequence

## Set-up
To set up application to run locally start Postgres docker container and run the application

```
docker-compose up
```

## List of API call examples
Create Sequence
```
curl --request POST \
  --url http://localhost:8081/sequence \
  --header 'Content-Type: ' \
  --data '{
  "name": "New Sequence",
  "open_tracking": false,
  "click_tracking": true
}'
```

Update Sequence
```
curl --request PATCH \
  --url http://localhost:8081/sequence/1 \
  --header 'content-type: application/json' \
  --data '{
  "open_tracking": false,
  "click_tracking": false
}'
```

Create Sequence Step
```
curl --request POST \
  --url http://localhost:8081/sequence-step \
  --header 'content-type: application/json' \
  --data '{
	"seq_id": 1,
  	"index": 1,
  	"content": "Sample text",
  	"subject": "Test Subject"
}'
```

Update Sequence
```
curl --request PATCH \
  --url http://localhost:8081/sequence-step/1 \
  --header 'Content-Type: ' \
  --data '{
    "content": "Updated Content",
  	"subject": "Updated Subject"
}'
```

Delete Sequence
```
curl --request DELETE \
  --url http://localhost:8081/sequence-step/1
```

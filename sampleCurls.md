**Local Test API Keys:**

```
export USER_RW=fa10ca02-f45d-4f4f-bd2f-406735e4fc9c
export ADMIN_RW=7459bbc5-cb40-469a-89cf-33edf5e51957
export ADMIN_R=11d1fb05-178e-4279-b26f-00f3b6e0d0c4
```

```sh
curl -H "Authorization: apikey ${USER_RW}" -d '{"name":"Murphy - user read/write", "age":2, "breed": "French Bulldog"}' -X POST http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${USER_RW}" -d '{"name":"Frank - user read/write", "age":3, "breed": "French Bulldog"}' -X POST http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${ADMIN_RW}" -d '{"name":"Holly - admin read/write", "age":3, "breed": "Labrador"}' -X POST http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${ADMIN_R}" -d '{"name":"Noel - admin read... this should fail", "age":18, "breed": "Labrador"}' -X POST http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${USER_RW}" http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${ADMIN_RW}" http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${ADMIN_R}" http://one.machinable.test:5001/collections/dogs | jq "."

curl -H "Authorization: apikey ${ADMIN_R}" "http://one.machinable.test:5001/collections/dogs?breed=Labrador" | jq "."
```

```sh
curl -d '{"firstName":"Nick", "lastName":"Sjostrom", "age":28}' -X POST http://one.machinable.test:5001/api/people | jq "."
```

**Local Test Project Users:**

```sh
$ BASIC=$(echo "$USERNAME:$PASSWORD" | base64)
$ curl -X POST -H "Authorization: basic $BASIC" https://one.machinable.io/sessions/
```

```sh
# local keys
nsjostrom@nsjostrom-H61MA-D3V:~$ echo $KEY
051204d1-6a29-4e07-b5a7-03c90bfc3fe3
nsjostrom@nsjostrom-H61MA-D3V:~$ echo $KEY2
fff24cb6-deb4-47b2-bea1-d6a67a2bd637
nsjostrom@nsjostrom-H61MA-D3V:~$ echo $KEY3
d933cbaa-74d9-48fd-9b83-08ab40ea18d6
```
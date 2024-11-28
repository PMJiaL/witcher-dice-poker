# Witcher (2007) dice poker game REST API

This repository hold implementation of the Witcher dice poker game (refer
to [this website](https://witcher.fandom.com/wiki/The_Witcher_dice_poker) for more info) in the form a JSON REST API
server.

All endpoints are
documented [here](https://github.com/Depermitto/witcher-dice-poker/tree/d36c053b55c0115a3a38e59cfdc514777bccba01/docs).
Or you can build with `APP_ENV=dev` and check out the [Swagger UI](https://swagger.io/tools/swagger-ui/), which
is built into the program.

## Getting started

To simply build the webservice use **Docker** and **Docker Compose**:

```shell
git clone https://github.com/Depermitto/witcher-dice-poker
cd witcher-dice-poker

docker compose up -d --build
```

Alternatively, using [golang](https://go.dev/):
```shell
...

go mod download
APP_ENV=dev go build -o main
```

### Configuration

- APP_ENV: **dev**, **production** - when set to *production*, `http://addr:port/swagger` endpoint will be hidden.

## LICENSE

Licensed under the MIT license.

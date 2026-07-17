# Witcher (2007) dice poker game REST API

This repository hold implementation of the Witcher dice poker game (refer
to [this website](https://witcher.fandom.com/wiki/The_Witcher_dice_poker) for more info)
in the form a JSON REST API server.

All endpoints are documented in code and [here](./docs).
Or you can set `WDP_SHOW_SWAGGER` and check out
the [Swagger UI](https://swagger.io/tools/swagger-ui/), which is built into the program.

## Getting started

Download the Compose file:

```shell
curl -O https://raw.githubusercontent.com/PMJiaL/witcher-dice-poker/master/compose.yaml
```

Then run the webservice using **Docker Compose**:

```shell
docker compose up -d
```

Or **Podman Compose**:

```shell
podman-compose up -d
```

### Build manually

Using [golang](https://go.dev/):

```shell
git clone https://github.com/PMJiaL/witcher-dice-poker
cd witcher-dice-poker

go mod download
go build -o main
```

### Configuration

- WDP_SHOW_SWAGGER: when present, `http://addr:port/swagger` endpoint will be exposed.

## LICENSE

Licensed under the MIT license.

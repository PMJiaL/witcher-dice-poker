# Witcher (2007) dice poker game service

This repository hold an implementation of the Witcher (2007) dice poker game (refer
to [this website](https://witcher.fandom.com/wiki/The_Witcher_dice_poker) for more info)
in the form of a self-hostable API service.

All endpoints are documented in code and [here](./docs).
Or you can set `WDP_SHOW_SWAGGER` and check out
the [Swagger UI](https://swagger.io/tools/swagger-ui/), which is built into the program.

## Getting started

Download the Compose file:

```shell
curl -O https://raw.githubusercontent.com/PMJiaL/witcher-dice-poker/master/compose.yaml
```

or copy and paste its contents into your own compose file:

```yaml
services:
  witcher-dice-poker:
    image: ghcr.io/pmjial/witcher-dice-poker:latest
    container_name: witcher-dice-poker
    # logging:
    #   driver: journald
    ports:
      - "2007:2007"
    healthcheck:
      test: [ "CMD", "/bin/busybox", "wget", "--quiet", "--spider", "http://localhost:2007/ping" ]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 5s
    restart: on-failure
networks: {}
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

## License

Licensed under the terms of the [MIT License](LICENSE).


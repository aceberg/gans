[![Main-Docker](https://github.com/aceberg/gans/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/gans/actions/workflows/main-docker.yml)
[![Binary-release](https://github.com/aceberg/gans/actions/workflows/binary-release.yml/badge.svg)](https://github.com/aceberg/gans/actions/workflows/binary-release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/gans)](https://goreportcard.com/report/github.com/aceberg/gans)
[![Maintainability](https://api.codeclimate.com/v1/badges/c76dbc1d7d64349af6c2/maintainability)](https://codeclimate.com/github/aceberg/gans/maintainability)

<h1><a href="https://github.com/aceberg/gans">
    <img src="https://raw.githubusercontent.com/aceberg/gans/main/assets/logo.png" width="25" />
</a>gans</h1>

Git+Ansible: watch git repo for changes and run only changed playbooks   
(Also, `gans` means goose in German)

![Screenshot](https://raw.githubusercontent.com/aceberg/gans/main/assets/Screenshot%202023-01-17%20at%2016-36-51%20gans.png)

## Quick start

```sh
docker run --name gans \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/gans:/data/gans \
-p 8845:8845 \
aceberg/gans
```

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DB        | Path to Database | /data/gans/sqlite.db |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8845 |
| THEME | Any theme name from https://bootswatch.com in lowcase | cerulean |
| SHOW | How many lines to show on index page | 25 |
| YAMLPATH | Path to file where git repo is described | /data/gans/repo.yaml |
| KEYPATH | Path to ssh keys directory | /data/gans/ssh |
| INTERVAL | Interval between repo scans (s, m, h) | 5s |
| TZ | Set your timezone for correct time | "" |


## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/gans/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Goose icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/goose)
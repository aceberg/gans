[![Main-Docker](https://github.com/aceberg/gans/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/gans/actions/workflows/main-docker.yml)
[![Binary-release](https://github.com/aceberg/gans/actions/workflows/binary-release.yml/badge.svg)](https://github.com/aceberg/gans/actions/workflows/binary-release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/gans)](https://goreportcard.com/report/github.com/aceberg/gans)
[![Maintainability](https://api.codeclimate.com/v1/badges/c76dbc1d7d64349af6c2/maintainability)](https://codeclimate.com/github/aceberg/gans/maintainability)

<h1><a href="https://github.com/aceberg/gans">
    <img src="https://raw.githubusercontent.com/aceberg/gans/main/assets/logo.png" width="25" />
</a>gans</h1>

Git+Ansible: watch git repo for changes and run only changed playbooks   
(Also, `gans` means goose in German)

- [Quick start](https://github.com/aceberg/gans#quick-start)
- [Usage](https://github.com/aceberg/gans#usage)
- [Config](https://github.com/aceberg/gans#config)
- [Options](https://github.com/aceberg/gans#options)
- [Thanks](https://github.com/aceberg/gans#thanks)


![Screenshot](https://raw.githubusercontent.com/aceberg/gans/main/assets/Screenshot%202023-01-23%20at%2019-34-01%20gans.png)

## Quick start
Don't forget to mount your git directory!

```sh
docker run --name gans \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/gans:/data/gans \
-v /path/to/git/repo:/gitrepo \
-p 8845:8845 \
aceberg/gans
```
## Usage

Web interface is pretty self-explanatory. If you know Ansible, you shouldn't have any problems.    
Important things about `gans`:
- It supports only local git repos for now. If you need to work with remote repo, you can pull it regularly with [git-syr](https://github.com/aceberg/git-syr) or your own cron script. 
- I decided not to parse Ansible Inventory to reduce errors, so you need to enter hosts and groups manually on `Repo` page. `Gans` will work only with those hosts.
- Keys are handled by Ansible Inventory or `ansible.cfg`. `Keys` page is only there to help you upload them and check their presence.

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
| LOGPATH | Path to log file | /data/gans/gans.log |
| INTERVAL | Interval between repo scans (s, m, h) | 5s |
| TZ | Set your timezone for correct time | "" |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -d | Path to SQLite DB file | /data/gans/sqlite.db | 
| -c | Path to config file | /data/gans/config.yaml | 
| -l | Path to log file | /data/gans/gans.log | 
| -r | Path to repo yaml file | /data/gans/repo.yaml | 

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/gans/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Goose icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/goose)
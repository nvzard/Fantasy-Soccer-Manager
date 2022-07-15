<div align="center">
<img width="668" alt="Soccer Manager" src="https://user-images.githubusercontent.com/19310512/179158021-38c56114-f884-4e18-8cce-693083b6af68.png">
</div>

<!-- A soccer online manager game API -->

## Summary

This project exposes a RESTful API for a simple application where football/soccer fans will create fantasy teams and will be able to sell or buy players.

## Documentation

- [Summary](#summary)
- [Documentation](#documentation)
- [Tech Stack](#tech-stack)
- [Development Setup](#development-setup)
  - [Requirements](#requirements)
  - [Install Requirements:](#install-requirements)
  - [Getting Started with Development](#getting-started-with-development)
  - [Stop Dev Environment](#stop-dev-environment)
  - [Run Tests [WIP]](#run-tests-wip)
  - [Remove all artifacts and dependencies](#remove-all-artifacts-and-dependencies)
- [API Documentation](#api-documentation)
  - [1. Register a new user](#1-register-a-new-user)
    - [Example Request](#example-request)
    - [Response](#response)
  - [2. Generate Authorization Token](#2-generate-authorization-token)
    - [Example Request](#example-request-1)
    - [Response](#response-1)
  - [3. Get User By Email [Authorization Token Required in headers]](#3-get-user-by-email-authorization-token-required-in-headers)
    - [Example Request](#example-request-2)
    - [Response](#response-2)
  - [4. Get Team By ID [Authorization Token Required in headers]](#4-get-team-by-id-authorization-token-required-in-headers)
    - [Example Request](#example-request-3)
    - [Response](#response-3)
  - [5. Update Team (Name, Country) By ID [Authorization Token Required in headers]](#5-update-team-name-country-by-id-authorization-token-required-in-headers)
    - [Example Request](#example-request-4)
    - [Response](#response-4)
  - [6. Get Player By ID [Authorization Token Required in headers]](#6-get-player-by-id-authorization-token-required-in-headers)
    - [Example Request](#example-request-5)
    - [Response](#response-5)
  - [7. Update Player (First Name, Last Name, Country) By ID [Authorization Token Required in headers]](#7-update-player-first-name-last-name-country-by-id-authorization-token-required-in-headers)
    - [Example Request](#example-request-6)
    - [Response](#response-6)
  - [8. Add Player to transfer list [Authorization Token Required in headers]](#8-add-player-to-transfer-list-authorization-token-required-in-headers)
    - [Example Request](#example-request-7)
    - [Response](#response-7)
  - [9. Get all players on the transfer list [Authorization Token Required in headers]](#9-get-all-players-on-the-transfer-list-authorization-token-required-in-headers)
    - [Example Request](#example-request-8)
    - [Response](#response-8)
  - [10. Buy player from the transfers list [Authorization Token Required in headers]](#10-buy-player-from-the-transfers-list-authorization-token-required-in-headers)
    - [Example Request](#example-request-9)
    - [Response](#response-9)
- [Future Improvements](#future-improvements)

---

## Tech Stack

Written in Go (Golang 1.18) using [gin](https://github.com/gin-gonic/gin) as web framework, [Postgres](https://github.com/postgres/postgres) as database.

The application is containerized using docker and code-syncin development env uses docker-sync.

- Use `make build` to build the web server and `make start` to start the server. Server runs on `localhost:8080`.
- Use `make stop` stop the development environment.
- Use `make clean` to clean the dev environment by deleting stale containers.
- Use `docker volume prune` to clean all volumes in case of local DB failure. [Warning: This will prune all volumes]
- Use `make test` runs the tests in a seperate docker container for isolation. [TODO]
- [Air](https://github.com/cosmtrek/air) and docker-sync are used for live reloading during development.

## Development Setup

### Requirements

```
make
git
docker (v20.10.10)
docker-compose (v1.29.2)
docker-sync (0.7.2) [Use `gem install docker-sync -v 0.7.2` for mac]
```

### Install Requirements:

1. Docker-Sync Install: http://docker-sync.io/
2. Docker Install: https://docs.docker.com/engine/install/


### Getting Started with Development

1. Clone repository

```
git clone https://git.toptal.com/screening/Nitanshu-Vashistha.git
```
or
```
git clone https://github.com/nvzard/soccer-manager.git
```

2. Change directory to the cloned repository

```
cd soccer-manager
```

3. Make sure all requirements are installed (docker, docker-compose, docker-sync)

```
docker --version
docker-compose --version
docker-sync --version
```

4. Build image

```
make build
```

4. Run start command

```
make start
```

PS: If `make start` fails for the first time use `make stop` to stop the containers and do `make start` again.

### Stop Dev Environment

```
make stop
```
---

### Run Tests [WIP]

```
make test
```

### Remove all artifacts and dependencies

```
make clean
```
---

## API Documentation

```
GET     /healthcheck                          # ok

POST    /api/user                             # register new user (json parameters: first_name, last_name, email, password)
POST    /api/auth                             # generate Authorization token for the user (json parameters: email, password)

[** Below API Endpoints Require Authorization token in header **]

GET     /api/ping                             # pong

GET     /api/user/:email                      # get user by email

GET     /api/team/:id                         # get team by id
PATCH   /api/team/:id                         # update team by id (json parameters: name, country)

GET     /api/player/:id                       # get player by id
PATCH   /api/player/:id                       # update team by id (json parameters: first_name, last_name, country)

GET     /api/transfers/                       # get the transfers list
POST    /api/transfers/                       # add player to transfer list (json parameters: player_id, asking_price)
POST    /api/transfers/buy/:player_id         # buy a player on the transfers list for the asking_price
```

### 1. Register a new user

```
POST /api/user
```

| JSONParameter | Type | Description |
| :--- | :--- | :--- |
| `first_name` | `string`   | FirstName of the user |
| `last_name`  | `string`   | LastName of the user |
| `email`      | `string`   | valid email address|
| `password`   | `string`   | password (len >= 8)|


#### Example Request
```
POST /api/user

{
    "first_name": "John",
    "last_name": "Lennon",
    "email": "john@gmail.com",
    "password": "qwerty65",
}
```

#### Response
```
{
    "id": 1,
    "email": "john@gmail.com"
}
```

### 2. Generate Authorization Token

```
POST /api/auth
```

| JSONParameter | Type | Description |
| :--- | :--- | :--- |
| `email`      | `string`   | email of the user|
| `password`   | `string`   | password of the user|


#### Example Request
```
POST /api/user

{
    "email": "john@gmail.com",
    "password": "qwerty65",
}
```

#### Response
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJkcmFrZUBnbWFpbC5jb20iLCJ0ZWFtX2lkIjoxLCJleHAiOjE2NTc4Mjc4MTl9"
}
```

### 3. Get User By Email [Authorization Token Required in headers]

```
GET /api/user/:email
```


#### Example Request
```
GET /api/user/john@gmail.com
```

#### Response
```
{
    "id": 1,
    "email": "john@gmail.com",
    "first_name": "John",
    "last_name": "Lennon",
    "team_id": 1
}
```


### 4. Get Team By ID [Authorization Token Required in headers]

```
GET /api/team/:id
```

#### Example Request
```
GET /api/team/1
```

#### Response

Every newly formed team has 3 goalkeepers(GK), 6 defenders(DEF), 6 midfielders(MID), 5 attackers(ATT)

```
{
    "available_cash": 5000000,
    "country": "France",
    "id": 1,
    "name": "Iron Squids",
    "owner_id": 1,
    "players": [
        {
            "id": 1,
            "first_name": "Douglas",
            "last_name": "Mejia",
            "country": "Germany",
            "age": 36,
            "position": "GK",
            "market_value": 1000000,
            "team_id": 1
        },
        .
        .
        .
        {
            "id": 20,
            "first_name": "Ryland",
            "last_name": "Mejia",
            "country": "USA",
            "age": 23,
            "position": "ATT",
            "market_value": 1000000,
            "team_id": 1
        }
    ],
    "team_value": 20000000
}
```

### 5. Update Team (Name, Country) By ID [Authorization Token Required in headers]

```
PATCH /api/team/:id
```

| JSONParameter | Type | Description |
| :--- | :--- | :--- |
| `name`      | `string`   | new team name|
| `country`   | `string`   | new team country|


#### Example Request
```
PATCH /api/team/1

{
    "name": "Leopards"
    "country": "China",
}
```

#### Response

```
{
    "id": 1,
    "available_cash": 5000000,
    "country": "China",
    "name": "Leopards",
    "owner_id": 1,
    "players": [
        {
            "id": 1,
            "first_name": "Douglas",
            "last_name": "Mejia",
            "country": "Germany",
            "age": 36,
            "position": "GK",
            "market_value": 1000000,
            "team_id": 1
        },
        .
        .
        .
        {
            "id": 20,
            "first_name": "Ryland",
            "last_name": "Mejia",
            "country": "USA",
            "age": 23,
            "position": "ATT",
            "market_value": 1000000,
            "team_id": 1
        }
    ],
    "team_value": 20000000
}
```

### 6. Get Player By ID [Authorization Token Required in headers]

```
GET /api/player/:id
```

#### Example Request
```
GET /api/player/5
```

#### Response

```
{
    "id": 5,
    "age": 32,
    "country": "Germany",
    "first_name": "Nathaniel",
    "last_name": "Mejia",
    "market_value": 1000000,
    "position": "DEF",
    "team_id": 1
}
```

### 7. Update Player (First Name, Last Name, Country) By ID [Authorization Token Required in headers]

```
PATCH /api/player/:id
```

| JSONParameter | Type | Description |
| :--- | :--- | :--- |
| `first_name`   | `string`   | new first name|
| `last_name`    | `string`   | new last name|
| `country`      | `string`   | new country|


#### Example Request
```
PATCH /api/team/1

{
    "first_name": "Saitama",
    "last_name": "Senpai",
    "country": "Japan",
}
```

#### Response

```
{
    "id": 1,
     "first_name": "Saitama",
    "last_name": "Senpai",
    "age": 36,
    "country": "Japan",
    "market_value": 1000000,
    "position": "GK",
    "team_id": 1
}
```

### 8. Add Player to transfer list [Authorization Token Required in headers]

```
POST /api/transfers
```

| JSONParameter | Type | Description |
| :--- | :--- | :--- |
| `player_id`      | `string`   | id of the player|
| `asked_price`    | `string`   | asked price for the player|

#### Example Request
```
POST /api/transfers

{
    "player_id": 6,
    "asked_price": 1050000
}
```

#### Response

```
{
    "transfer_id": 8
}
```

### 9. Get all players on the transfer list [Authorization Token Required in headers]

```
GET /api/transfers
```

#### Example Request
```
GET /api/transfers
```

#### Response

```
{
    "transfer_list": [
        {
            "id": 1,
            "player_id": 6,
            "market_value": 1000000,
            "asked_price": 1050000,
            "player": {
                "id": 6,
                "first_name": "Grant",
                "last_name": "Klein",
                "country": "Qatar",
                "age": 23,
                "position": "DEF",
                "market_value": 1000000,
                "team_id": 1
            }
        }
    ]
}
```

### 10. Buy player from the transfers list [Authorization Token Required in headers]

```
POST /transfers/buy/:player_id
```

This transfers the player from selling team to buying team and adjusts the team's bank balance accordingly.
A player's market value increases by 10-100% after every transfer.

#### Example Request
```
POST /transfers/buy/6
```

#### Response

```
{
    "status": "done"
}
```

---
## Future Improvements

Some of the improvements that can be made but were skipped due to time constraints.

- Add unit tests.
- Add postman collection to test and document the API.
- Add admin roles to have CRUD access for all the users, teams, players and transfers.
- Build a user-interface(UI) for the backend.

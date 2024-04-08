# Simple File Server in Go

This is a very simple fileserver written in Go.

## Create Swagger

> swagger generate spec -o ./swagger.yml --scan-models

## API Functions

Base URL:

> {domain:port}/api/{version}/{endpoint}

Currently there only exists `version` `v1`

| Endpoint        | Method | Description                   |
| --------------- | ------ | ----------------------------- |
| `/` & `/status` | `GET`  | Both show the api status      |
| `/ping`         | `GET`  | answers with a successmessage |
| `/files`        | `POST` | Posts a file                  |
| `/files`        | `GET`  | Gets all filenames            |
| `/file`         | `GET`  | Gets a file as blob           |

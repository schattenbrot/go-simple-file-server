# Go Simple File Server

This is a very simple fileserver written in Go.

## Create Swagger

To create swagger documentation run:

> docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest init -g ./cmd/api/main.go

The documentation is available on url:

> http://localhost:8080/docs

> http://localhost:8080/docs/doc.json

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

## File Structure:

```
.
├── cmd
│   └── api
│       └── main.go
├── data
│   └── files
│       └── example-image.png
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   └── services
│       ├── app
│       │   ├── controllers.go
│       │   └── routes.go
│       └── files
│           ├── controllers.go
│           ├── db.go
│           ├── middlewares.go
│           ├── models.go
│           └── routes.go
├── packages
│   ├── explerror
│   │   └── explerror.go
│   └── responder
│       └── responder.go
└── README.md
```

## Config

```go
	flag.StringVar(&Env, "env", "dev", "the server environment")
	flag.StringVar(&Domain, "domain", "http://localhost", "the server domain")
	flag.IntVar(&Port, "port", 8080, "the server port")
	flag.BoolVar(&Cors.IsEnabled, "cors", false, "the server cors")
	allowedOrigins := flag.String("allowed_origins", "*", "the server allowed cors origins (split the origins with ,)")
	readTokens := flag.String("write_tokens", "", "the server's required read tokens (split tokens with ,)")
	readWriteTokens := flag.String("read_write_tokens", "", "the server's required read-write tokens (split tokens with ,)")
```

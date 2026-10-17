# go-runway

[![Go Reference](https://pkg.go.dev/badge/github.com/psyb0t/go-runway.svg)](https://pkg.go.dev/github.com/psyb0t/go-runway)
[![CI](https://github.com/psyb0t/go-runway/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/psyb0t/go-runway/actions/workflows/pipeline.yml)
[![version](https://raw.githubusercontent.com/psyb0t/go-runway/badges/version.svg)](https://github.com/psyb0t/go-runway/tags)
[![license](https://raw.githubusercontent.com/psyb0t/go-runway/badges/license.svg)](LICENSE)
[![imported by](https://raw.githubusercontent.com/psyb0t/go-runway/badges/importers.svg)](https://github.com/psyb0t/go-runway/blob/badges/importers.md)

Unofficial Go client for the [Runway API](https://docs.dev.runwayml.com), generated with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) from the OpenAPI document Runway publishes. I am not affiliated with Runway.

The package is `runway`, in the repository root. Everything in `types.gen.go` and `client.gen.go` is generated. Don't edit those files by hand, regenerate them.

## Install

```sh
go get github.com/psyb0t/go-runway
```

## Usage

Every endpoint in the spec has a typed method on `ClientWithResponses`. Responses decode into `JSON200` and the other status fields.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/psyb0t/go-runway"
)

func main() {
	auth := runway.WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+os.Getenv("RUNWAYML_API_SECRET"))
		req.Header.Set("X-Runway-Version", "2024-11-06")

		return nil
	})

	client, err := runway.NewClientWithResponses("https://api.dev.runwayml.com", auth)
	if err != nil {
		panic(err)
	}

	res, err := client.GetV1OrganizationWithResponse(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res.JSON200.CreditBalance)
}
```

The same code is in `example_test.go`, run against a local test server, so `make test` checks that it works.

The full API is in the [Go reference](https://pkg.go.dev/github.com/psyb0t/go-runway); from inside your module, `go doc github.com/psyb0t/go-runway` prints the same.

The spec has no operation IDs, so method names come from the method and path: `PostV1TextToImage`, `GetV1TasksId`, `DeleteV1TasksId` and so on. Most request and response bodies are a `oneOf` over several shapes (one per model, or one per task status). Those become union types: build a request with the `From...` methods and read a response with the `As...` methods.

Send a union request body through the `...WithBodyWithResponse` method, marshaled with `json.Marshal`. The typed `...WithResponse(ctx, body)` methods take a `...JSONRequestBody`, a separate named type that does not carry the union's `MarshalJSON`, so a union body goes out as `{}`. `TestPostTextToImageUnionBody` in `client_test.go` shows the working form.

## Auth

Runway authenticates with a Bearer API key, created in the Runway developer portal. Every request also needs the `X-Runway-Version` header. The spec declares it as a required parameter with a single value, `2024-11-06`, on every operation. The generator step drops it from the method signatures, so set it in the same request editor as the key, as in the example above.

## Regenerating

The spec is not committed. `make generate` downloads the current OpenAPI document into the gitignored `api/` directory and regenerates the client:

```sh
make generate
```

`gen.go` declares the steps, so `make generate` is just `go generate ./...`:

- `scripts/fetch-spec.sh` downloads `https://docs.dev.runwayml.com/openapi.json` to `api/openapi.json`.
- `scripts/fix-spec.sh` writes `api/openapi.fixed.json`. Every schema in the spec is inline, so its `oneOf` discriminators have no `$ref` targets to map and oapi-codegen rejects them. The script drops the discriminators, which leaves plain union types. It also drops the `X-Runway-Version` header parameter from each operation (see Auth).
- oapi-codegen runs twice: `config.yaml` writes the models to `types.gen.go` and `config.client.yaml` writes the client to `client.gen.go`, so neither file is one 100k line blob. Both map `number` to `float64`, because the spec gives most numbers no format and oapi-codegen would otherwise use `float32`.

`.github/workflows/regen.yml` does the same every day in CI. When the new spec changes the client, it runs `make lint`, `make test` and `make sec`. If a check fails, Claude edits the code and the checks run again, up to three times, and a review step reads the result. A passing change is committed with a CHANGELOG entry and tagged as the next release: a patch version for fixes and additions, a minor version for breaking changes.

## Agent integrations

The [skill](.agents/skills/go-runway) teaches a coding agent how to use go-runway when it writes Go. It works in any agent that reads `.agents/skills/`, and installs natively in the clients below.

### Claude Code

```bash
claude plugin marketplace add psyb0t/agents
claude plugin install go-runway@psyb0t
```

### Codex

```bash
codex plugin marketplace add psyb0t/agents
codex plugin add go-runway@psyb0t
```

Codex also picks the skill up with no install in any checkout of this repo.

### OpenClaw

The skill is published to ClawHub on every release:

```bash
openclaw skills install @psyb0t/go-runway
```

## Development

Every command runs in the `Dockerfile.dev` image. The host needs only Docker and Make.

```sh
make help      # list targets
make dep       # go mod tidy + vendor
make lint      # shellcheck, shfmt, go fix, golangci-lint
make test      # go test -race
make sec       # gitleaks, govulncheck, semgrep
make generate  # download the spec and regenerate
```

## License

[MIT](LICENSE). The generated code comes from Runway's published API specification. The specification itself is not redistributed here.

# Go code

Refines `~/.claude/rule-details/coding-style.md`.

- **GO1** `types.gen.go` and `client.gen.go` are generated. Never edit them by hand; change the spec patch or the oapi-codegen config and run `make generate`.
- **GO2** Fixes to the upstream spec live in `scripts/fix-spec.sh` as jq filters that match the shape of the problem, not one path or one schema name, and log how many places they changed. The spec itself is never committed; `api/` is gitignored.
- **GO3** `config.yaml` (models) and `config.client.yaml` (client) share their `output-options` and `compatibility` blocks. Change both together.
- **GO4** Hand-written Go in the package root stays minimal: `gen.go` holds the package doc and the `go:generate` steps. Any helper added later wraps generated types and never duplicates them.
- **GO5** `make lint` passes with zero issues. Lines are at most 80 columns, the `.golangci.yml` limit; split calls one argument per line rather than adding `//nolint`.

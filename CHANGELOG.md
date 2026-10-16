# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.1.0, 2026-10-16

- Added `runway`, a REST client generated with oapi-codegen from Runway's OpenAPI document (RunwayML API, spec version 2024-11-06, OpenAPI 3.1.0, from docs.dev.runwayml.com). It covers all 49 paths: image, video and audio generation, upscaling, voices, avatars, realtime sessions, recipes, workflows, routers, documents, uploads, tasks and organization usage.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, drops its unmappable inline discriminators and the per-operation `X-Runway-Version` header parameter, and regenerates the client into `types.gen.go` and `client.gen.go`. The spec is not committed.
- Added the `go-runway` agent skill. Install it with `claude plugin install go-runway@psyb0t` (Claude Code), `codex plugin add go-runway@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-runway` (ClawHub).

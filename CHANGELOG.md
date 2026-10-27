# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.3.0, 2026-10-27

- Added the video-to-audio endpoint: `Client.PostV1VideoToAudio`, `Client.PostV1VideoToAudioWithBody`, `ClientWithResponses.PostV1VideoToAudioWithResponse`, `ClientWithResponses.PostV1VideoToAudioWithBodyWithResponse`, `NewPostV1VideoToAudioRequest`, `NewPostV1VideoToAudioRequestWithBody` and `ParsePostV1VideoToAudioResponse`. The four client methods are also added to `ClientInterface` and `ClientWithResponsesInterface`, so custom implementations of those interfaces must add them.
- Added request union variants with `As`, `From` and `Merge` methods: `PostV1TextToImageJSONBody19`, `PostV1TextToImageJSONBody20`, `PostV1TextToVideoJSONBody23`, `PostV1VideoToVideoJSONBody11`, `PostV1ImageToVideoJSONBody24` (with `PostV1ImageToVideoJSONBody24ModelKling4`) and `PostV1SoundEffectJSONBody2`.
- Added `PostV1SoundEffectJSONBody1.Format`. `PostV1SoundEffectJSONBody1.Duration` is now `float64` instead of `*float64`, and `PostV1SoundEffectJSONBody1.Loop` and `PostV1SoundEffectJSONBody1ModelElevenTextToSoundV2` are removed.
- Removed `PostV1TextToImageJSONBody18.ReferenceImages`, `PostV1TextToImageJSONBody18ModelGemini25Flash`, the `PostV1TextToImageJSONBody18RatioN...` constants and the `ReferenceImages` URI types.
- `PostV1ImageToVideoJSONBody5PromptImage1`, `PostV1ImageToVideoJSONBody18PromptImage1` and `PostV1ImageToVideoJSONBody20PromptImage1` are now `[]interface{}`. The `Position` and `Uri` types and constants for these bodies are removed.
- The `Routing.ResolvedInput` fields `AspectRatio` and `Resolution` in the `PostV1GenerateImage200JSONResponseBody` types changed from `*string` to `string` in some variants.
- Added audit log action constants such as `GetV1OrganizationWebappAuditLogs200JSONResponseBodyDataActionAdPublishBatchQueued` and usage model constants such as `PostV1OrganizationUsage200JSONResponseBodyModelsKling4`.
- Changed the `GetJSON200` result struct of `GetV1OrganizationWebappAuditLogsEventIdResponse`, so code that names its fields or type must be checked.

## v0.2.1, 2026-10-27

- The development image is based on Debian (`golang:1.26.10-trixie`) instead of Alpine. Alpine keeps only the latest build of each package, so the exact pins broke the image build whenever a package was updated. System packages now come from the snapshot.debian.org date the base image was built from, which keeps them fixed for a given base image digest. jq 1.8.2, shellcheck 0.11.0 and shfmt 3.13.1 keep their versions, and Semgrep runs in its own virtual environment.
- Raised the minimum Go version in `go.mod` from 1.26.9 to 1.26.10, which fixes the `crypto/x509` vulnerability GO-2026-6809 found by govulncheck.

## v0.2.0, 2026-10-22

- Regenerated from the current Runway OpenAPI spec. The text-to-image request union gained two variants: `PostV1TextToImageJSONBody17` and `PostV1TextToImageJSONBody18`, with `AsPostV1TextToImageJSONBody17`, `FromPostV1TextToImageJSONBody17`, `MergePostV1TextToImageJSONBody17` and the matching `...18` methods on `PostV1TextToImageJSONBody`.
- The numbered text-to-image body types shifted. Per-model fields and enums moved between `PostV1TextToImageJSONBody9` to `PostV1TextToImageJSONBody16`, so code that names these types or constants must be updated. For example `PostV1TextToImageJSONBody11` lost `OutputFormat` and `Seed` and gained `Background`, `Quality` and `OutputCount`, and `PostV1TextToImageJSONBody15` lost `ThinkingLevel` and gained `MaskUri`, `PreciseEdit` and `Quality`.
- Several text-to-image fields changed type: `PostV1TextToImageJSONBody12.Ratio` is now `PostV1TextToImageJSONBody12Ratio` instead of a pointer, and `PostV1TextToImageJSONBody15.Ratio` is now a pointer.
- Added text-to-image model constants including `PostV1TextToImageJSONBody10ModelGptImage25Flare`, `PostV1TextToImageJSONBody12ModelBriaFibo15`, `PostV1TextToImageJSONBody13ModelBriaFiboEdit15`, `PostV1TextToImageJSONBody14ModelIdeogram40`, `PostV1TextToImageJSONBody15ModelIdeogram45` and `PostV1TextToImageJSONBody16ModelSeedream5Flash`, plus many new aspect ratio constants.
- Image-to-video: `PostV1ImageToVideoJSONBody8.PromptText` is now a required `string` instead of `*string`. `PostV1ImageToVideoJSONBody22` gained `OutputFormat` with `PostV1ImageToVideoJSONBody22OutputFormatMp4` and `PostV1ImageToVideoJSONBody22OutputFormatMov`.
- `GetV1OrganizationWebappAuditLogsEventIdResponse.JSON200` and `GetJSON200` return a changed anonymous struct type, so code that depends on its exact shape must change.
- Added audit log action constants such as `GetV1OrganizationWebappAuditLogs200JSONResponseBodyDataActionAdCreativeArchived`, `...ActionOnboardingBonusGranted` and `...ActionUserGroupVisibilityUpdated`, with matching `GetV1OrganizationWebappAuditLogsEventId200JSONResponseBodyAction...` constants.
- Added organization usage model constants `PostV1OrganizationUsage200JSONResponseBodyModelsGeminiImage31FlashLite` and `PostV1OrganizationUsage200JSONResponseBodyModelsSeedream5Flash`, with matching `ResultsUsedCreditsModel` constants.
- The dev image now uses Go 1.26.10 and `ca-certificates` 20261020-r0, which fixes the `crypto/x509` vulnerability GO-2026-6809 that `make sec` reported.

## v0.1.2, 2026-10-17

- Documentation only. The agent skill explains how to look up every endpoint, parameter and model with `go doc` or on pkg.go.dev, and declares its license.
- The skill no longer lists the API credentials under OpenClaw's `requires.env`, which kept OpenClaw from loading it until they were set. It requires only `go`; the credentials are still documented as what your program reads.
- The README points to the full API reference, and the examples in the README and the skill were checked against the current code.
- The skill's frontmatter is valid YAML again; an unquoted colon in the description stopped it from parsing.

## v0.1.1, 2026-10-17

- Documentation only. The README and `AGENTS.md` describe the daily regeneration workflow: what it checks, how it fixes failing checks, how it picks the next version and which paths a fix may not touch.

## v0.1.0, 2026-10-16

- Added `runway`, a REST client generated with oapi-codegen from Runway's OpenAPI document (RunwayML API, spec version 2024-11-06, OpenAPI 3.1.0, from docs.dev.runwayml.com). It covers all 49 paths: image, video and audio generation, upscaling, voices, avatars, realtime sessions, recipes, workflows, routers, documents, uploads, tasks and organization usage.
- Number fields are `float64` instead of oapi-codegen's default `float32`.
- `make generate` downloads the current spec, drops its unmappable inline discriminators and the per-operation `X-Runway-Version` header parameter, and regenerates the client into `types.gen.go` and `client.gen.go`. The spec is not committed.
- Added the `go-runway` agent skill. Install it with `claude plugin install go-runway@psyb0t` (Claude Code), `codex plugin add go-runway@psyb0t` (Codex) or `openclaw skills install @psyb0t/go-runway` (ClawHub).

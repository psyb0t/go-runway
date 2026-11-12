# Changelog

All notable changes per release. Versions follow [semver](https://semver.org).

## v0.5.0, 2026-11-12

- Added request union variants with `As`, `From` and `Merge` methods: `PostV1TextToImageJSONBody24`, `PostV1TextToVideoJSONBody26` and `27`, `PostV1ImageToVideoJSONBody27` and `28`, and `PostV1VideoToVideoJSONBody14` and `15`. New models include `PostV1TextToImageJSONBody23ModelSeedream5ProLayers`, `PostV1TextToImageJSONBody24ModelGemini25Flash`, `PostV1ImageToVideoJSONBody27ModelHailuo3Highspeed` and `PostV1ImageToVideoJSONBody28ModelHailuo3Lite`.
- Added a custom voice variant for text-to-speech: `PostV1TextToSpeechJSONBody1Voice1`, `PostV1TextToSpeechJSONBody2Voice1` and `PostV1TextToSpeechJSONBody3Voice1`, each with `PostV1TextToSpeechJSONBody1Voice1TypeCustom` style constants and `From` and `Merge` methods on the voice types.
- Added `GetV1TasksId200JSONResponseBody5.Layers`, new ratio constants such as `PostV1ImageToVideoJSONBody22RatioN38403840`, and Hailuo 3 and Seedream layers usage model constants under `PostV1OrganizationUsage200JSONResponseBodyModels`.
- Changed `PostV1TextToImageJSONBody23`: `PromptText` is now `*string` instead of `string`, `ReferenceImages` is now a slice instead of a pointer to a slice and its items lose `Tag`, and `Ratio`, its `PostV1TextToImageJSONBody23Ratio` type and `PostV1TextToImageJSONBody23ModelGemini25Flash` are removed. The struct is no longer comparable. The Gemini model moves to `PostV1TextToImageJSONBody24`.
- Error reason enums were renumbered in the 400 response bodies of endpoints such as `PostV1ImageToVideo`, `PostV1ImageUpscale` and `PostV1CharacterPerformance`. `...Reason26INTERNAL` and `...Reason27UNKNOWN` are removed. The values are now `...Reason26INTERNALRESOURCES`, `...Reason27INTERNAL` and a new `...Reason28UNKNOWN`.
- The text-to-speech endpoint docs now say ElevenLabs models accept preset or custom voices and Seed Audio accepts a reference audio clip.
- Raised the minimum Go version to 1.26.11 in `go.mod`, which fixes GO-2026-6859, a `net/http` crash in the 1.26.10 standard library. The development image uses the same toolchain.

## v0.4.0, 2026-11-07

- Added the image-to-vector endpoint: `Client.PostV1ImageToVector`, `Client.PostV1ImageToVectorWithBody`, `ClientWithResponses.PostV1ImageToVectorWithResponse`, `ClientWithResponses.PostV1ImageToVectorWithBodyWithResponse`, `NewPostV1ImageToVectorRequest`, `NewPostV1ImageToVectorRequestWithBody` and `ParsePostV1ImageToVectorResponse`. The request body is `PostV1ImageToVectorJSONBody0` with model `PostV1ImageToVectorJSONBody0ModelIdeogramVector`.
- Added the text-to-music endpoint: `Client.PostV1TextToMusic`, `Client.PostV1TextToMusicWithBody`, `ClientWithResponses.PostV1TextToMusicWithResponse`, `ClientWithResponses.PostV1TextToMusicWithBodyWithResponse`, `NewPostV1TextToMusicRequest`, `NewPostV1TextToMusicRequestWithBody` and `ParsePostV1TextToMusicResponse`. It uses model `PostV1TextToMusicJSONBody0ModelElevenMusicV25`, and `PostV1GenerateAudioJSONBodyInputTypeMusic` is added.
- The four new client methods for each endpoint are also added to `ClientInterface` and `ClientWithResponsesInterface`, so custom implementations of those interfaces must add them.
- Added request union variants with `As`, `From` and `Merge` methods: `PostV1TextToImageJSONBody21`, `22` and `23`, `PostV1TextToVideoJSONBody24` and `25`, `PostV1ImageToVideoJSONBody25` and `26`, and `PostV1VideoToVideoJSONBody12` and `13`. New model constants include `PostV1TextToImageJSONBody21ModelGptImage26Sunburst`, `PostV1TextToImageJSONBody22ModelGptImage26Flare`, `PostV1TextToImageJSONBody23ModelGemini25Flash`, `PostV1ImageToVideoJSONBody25ModelGeminiOmniFlash12` and `PostV1TextToVideoJSONBody25ModelKling4Flash`.
- Removed `PostV1TextToImageJSONBody20ModelGemini25Flash` and the `PostV1TextToImageJSONBody20RatioN...` constants. `PostV1TextToImageJSONBody20.Ratio` is now `*PostV1TextToImageJSONBody20Ratio` instead of `PostV1TextToImageJSONBody20Ratio`, and the `Tag` field is removed from the `PostV1TextToImageJSONBody20.ReferenceImages` items. Model `Gemini25Flash` moves to `PostV1TextToImageJSONBody23`.
- Added audit log action constants such as `GetV1OrganizationWebappAuditLogs200JSONResponseBodyDataActionAppStudioAppCreated` and `GetV1OrganizationWebappAuditLogs200JSONResponseBodyDataActionAdStatusChanged`, with matching constants for the single-event endpoint. Audit log metadata now has new fields such as `AdLevel` and `PlatformObjectID`.
- Changed the `JSON200` and `GetJSON200` result structs of `GetV1OrganizationWebappAuditLogsResponse`, `GetV1OrganizationWebappAuditLogsEventIdResponse` and `GetV1OrganizationWebappUsageResponse`, so code that names their fields or types must be checked. The usage response adds `GetV1OrganizationWebappUsage200JSONResponseBodyDataUnit`.

## v0.3.2, 2026-10-28

- Added an optional `Bitrate` field, a pointer to a new string enum type, to the video request bodies. It takes `high` or `standard`, with `Valid` methods on each enum type.
- Image-to-video: `PostV1ImageToVideoJSONBody13.Bitrate`, `PostV1ImageToVideoJSONBody14.Bitrate` and `PostV1ImageToVideoJSONBody15.Bitrate`, with types `PostV1ImageToVideoJSONBody13Bitrate`, `...14Bitrate` and `...15Bitrate` and constants such as `PostV1ImageToVideoJSONBody13BitrateHigh` and `PostV1ImageToVideoJSONBody13BitrateStandard`.
- Text-to-video: `PostV1TextToVideoJSONBody12.Bitrate`, `PostV1TextToVideoJSONBody13.Bitrate` and `PostV1TextToVideoJSONBody14.Bitrate`, with matching `...Bitrate` types and `...BitrateHigh` and `...BitrateStandard` constants.
- Video-to-video: `PostV1VideoToVideoJSONBody5.Bitrate`, `PostV1VideoToVideoJSONBody6.Bitrate` and `PostV1VideoToVideoJSONBody7.Bitrate`, with matching `...Bitrate` types and `...BitrateHigh` and `...BitrateStandard` constants.

## v0.3.1, 2026-10-28

- Updated Semgrep in the development image to 1.182.1 and refreshed its locked dependencies. certifi 2026.10.23 drops the emSign root CAs whose intermediate keys were compromised (GHSA-6q48-6c7w-9gw2).
- Dependabot no longer opens pull requests for `mcp` or the OpenTelemetry packages in `requirements-semgrep.txt`. Semgrep pins their versions itself, so bumping one of them on its own cannot install.
- Updated kin-openapi, which the code generator uses, from 0.142.0 to 0.149.0. The newer version fixes several panics on malformed specs reported in GitHub security advisories.
- The security scan merges its two gitleaks runs into one SARIF run. GitHub code scanning rejects a report with two runs from the same tool, so no scan results had been uploaded.

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

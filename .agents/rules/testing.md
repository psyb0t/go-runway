# Testing

Refines `~/.claude/rule-details/go/writing-tests.md`.

- **TST1** Tests are black-box, in package `runway_test`, and use only the exported API.
- **TST2** Every test runs the client against an `httptest` server and asserts the method, the path, both auth headers (`Authorization` and `X-Runway-Version`) and the decoded response. No test talks to the real API.
- **TST3** Union request bodies are tested through the `...WithBodyWithResponse` path, as `TestPostTextToImageUnionBody` does, and the test checks the JSON the server received.
- **TST4** `example_test.go` is the README usage example and keeps an `// Output:` line against a local server. Change the README and the example together.
- **TST5** Test data uses obvious placeholders: fake keys, made-up ids and prompts.

#!/bin/bash
# Rewrites the downloaded Runway spec into a form oapi-codegen accepts. Reads
# api/openapi.json and writes api/openapi.fixed.json.
#
# Every schema in the spec is inline, so the oneOf/anyOf discriminators have
# no $ref targets to map and oapi-codegen fails with "not all schemas were
# mapped". Dropping the discriminator keeps the oneOf as a union type with
# the usual As/From/Merge helpers; the caller switches on the field itself.
#
# Every operation also takes the same required X-Runway-Version header, a
# single constant value. Left in, it would add a params struct to every call,
# so it is dropped from the operations and set once in a request editor.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR
readonly SPEC_IN="${SPEC_DIR}/openapi.json"
readonly SPEC_OUT="${SPEC_DIR}/openapi.fixed.json"
readonly VERSION_PARAM_REF="#/components/parameters/X-Runway-Version"
readonly DROP_VERSION_PARAM="(.paths[][] | objects | select(has(\"parameters\")) | .parameters) |= map(select(.[\"\$ref\"] != \"${VERSION_PARAM_REF}\")) | del(.components.parameters[\"X-Runway-Version\"])"

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

main() {
	local tmp count
	tmp="$(mktemp "${SPEC_DIR}/.openapi.fixed.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp now, the local is gone when EXIT fires
	trap "rm -f -- '${tmp}'" EXIT

	# Only a discriminator object (one with propertyName) is dropped, never a
	# schema property that happens to be called "discriminator".
	local filter='type == "object" and (.discriminator | type == "object" and has("propertyName"))'
	count="$(jq "[.. | select(${filter})] | length" "$SPEC_IN")"
	jq "walk(if ${filter} then del(.discriminator) else . end) | ${DROP_VERSION_PARAM}" \
		"$SPEC_IN" >"$tmp"

	chmod 0644 "$tmp"
	mv -- "$tmp" "$SPEC_OUT"
	log INFO "dropped discriminators count=${count} out=${SPEC_OUT##*/}"
}

main "$@"

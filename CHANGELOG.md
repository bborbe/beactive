# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support
## v0.1.0

- chore: bump go directive from `1.26.5` to `1.26.6`
- deps: bump `golang.org/x/mod` `v0.38.0` → `v0.40.0` — clears CVE-2026-56864 (malicious GOSUMDB) and CVE-2026-56865 (malicious GOPROXY), flagged by `make trivy`. Pulled `golang.org/x/net` `v0.57.0` → `v0.58.0` with it
- deps: bump `golang.org/x/image` `v0.43.0` → `v0.45.0` — clears GO-2026-6222 (excessive memory allocation during VP8L decoding), flagged by `make vulncheck`. Pulled `golang.org/x/sys` `v0.46.0` → `v0.47.0`, `x/text` `v0.39.0` → `v0.41.0` and `x/tools` `v0.47.0` → `v0.48.0` with it
- chore: add this CHANGELOG. The repository sets `release.autoRelease: true` in
  `.maintainer.yaml`, but had no changelog and no tags, so the releaser had
  nothing to read and automated maintenance that reads
  `origin/master:CHANGELOG.md` could not run against it

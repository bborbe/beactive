# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

- fix: bump `osv-scanner` to v2.6.0 and `golang.org/x/net` to v0.60.0 so the Linux vulnerability gates stop failing. v2.3.1 pins `golang.org/x/tools` v0.38.0, whose SSA builder aborts with `unexpected expr: *ast.KeyValueExpr` on the promoted-field composite-literal key Go 1.27 permits in the Linux stdlib, so a repo on the old pin passes locally on darwin and fails only in Linux CI. `x/net` v0.58.0 carries `GO-2026-6603/6610/6611/6612/6617`, which fail both `vulncheck` and `trivy`.

## v0.1.4

- chore: update Go to 1.27.1

## v0.1.3

- chore: update go module dependencies

## v0.1.2

- chore: update Go to 1.27.0 and github.com/onsi/ginkgo/v2 to v2.32.1, github.com/onsi/gomega to v1.43.0; drop obsolete replace directives

## v0.1.1

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support
## v0.1.0

- chore: bump go directive from `1.26.5` to `1.26.6`
- deps: bump `golang.org/x/mod` `v0.38.0` → `v0.40.0` — clears CVE-2026-56864 (malicious GOSUMDB) and CVE-2026-56865 (malicious GOPROXY), flagged by `make trivy`. Pulled `golang.org/x/net` `v0.57.0` → `v0.58.0` with it
- deps: bump `golang.org/x/image` `v0.43.0` → `v0.45.0` — clears GO-2026-6222 (excessive memory allocation during VP8L decoding), flagged by `make vulncheck`. Pulled `golang.org/x/sys` `v0.46.0` → `v0.47.0`, `x/text` `v0.39.0` → `v0.41.0` and `x/tools` `v0.47.0` → `v0.48.0` with it
- chore: add this CHANGELOG. The repository sets `release.autoRelease: true` in
  `.maintainer.yaml`, but had no changelog and no tags, so the releaser had
  nothing to read and automated maintenance that reads
  `origin/master:CHANGELOG.md` could not run against it

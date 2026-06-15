# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Terraform provider for NetScaler / Citrix ADC. Every resource ultimately drives the
NITRO REST API on a target ADC through the `github.com/citrix/adc-nitro-go` client
(`service.NitroClient`). There is no business logic beyond translating Terraform
schema <-> NITRO payloads.

## The dual-provider architecture (most important thing to understand)

This binary serves **two providers muxed together** (`main.go`):

1. **SDK v2 provider** — package `citrixadc/` (legacy). ~714 resources registered in
   the giant map in `citrixadc/provider.go`. Built on
   `terraform-plugin-sdk/v2`, upgraded from protocol tf5 -> tf6 at startup.
2. **Plugin Framework provider** — package tree under `citrixadc_framework/` (new).
   Built on `terraform-plugin-framework` (already tf6). Registered in
   `citrixadc_framework/provider/provider.go`.

`main.go` combines them with `tf6muxserver` and serves both under
`registry.terraform.io/citrix/citrixadc`. To end users it looks like one provider.

**The active project is migrating resources from the SDK v2 provider to the Framework
provider** (Jira NSNETAUTO-1148). When migrating a resource:

- A given Terraform type name (e.g. `citrixadc_lbvserver`) may be registered in **only
  one** of the two providers. The mux server will fail to start if the same type name
  is registered in both. Migrating means **removing** the entry from
  `citrixadc/provider.go` *and* its `resource_citrixadc_*.go` file, then registering the
  Framework implementation.
- A `citrixadc_framework/<name>/` directory existing does **not** mean the resource is
  live. Hundreds of packages are scaffolded, but a resource is only active if its
  `New<Name>Resource` constructor is in the `Resources()` slice (and/or its
  `<Name>DataSource` in `DataSources()`) in `citrixadc_framework/provider/provider.go`.
  Many scaffolded packages have stubbed-out (commented) Create/Update bodies — see
  `citrixadc_framework/aaaparameter/resource_aaaparameter.go` for the stub shape.

## Resource conventions

**SDK v2 (`citrixadc/`)** — flat package, one file per resource named
`resource_citrixadc_<name>.go`, exposing `resourceCitrixAdc<Name>() *schema.Resource`
with `CreateContext`/`ReadContext`/`UpdateContext`/`DeleteContext` funcs. Registered as
`"citrixadc_<name>": resourceCitrixAdc<Name>()`.

**Framework (`citrixadc_framework/<name>/`)** — one sub-package per resource, typically
four files:
- `resource_<name>.go` — implements `resource.Resource` (+ `Configure`, `ImportState`);
  constructor `New<Name>Resource`. Client obtained via
  `r.client = *req.ProviderData.(**service.NitroClient)`. `Metadata` sets
  `resp.TypeName = req.ProviderTypeName + "_<name>"`.
- `resource_schema.go` — the framework schema + the `<Name>ResourceModel` struct.
- `datasource_<name>.go` / `datasource_schema.go` — the matching data source.

Shared helpers live in `citrixadc_framework/utils/` (`utils.go`, `las_utils.go`,
`sslvserver_utils.go`). The `interface` resource package must be imported with an alias
(`Interface "...../interface"`) because `interface` is a Go keyword.

NITRO calls go through `service.NitroClient` methods like `FindResource`,
`UpdateUnnamedResource`, `ActOnResource`, with the resource type from the adc-nitro-go
`service.<Type>.Type()` constants.

## Build / test commands

Go 1.24. Note the `godebug tlsrsakex=1` in `go.mod` — required so the provider can still
negotiate RSA-key-exchange TLS with older ADC firmware; do not remove it.

```bash
make build          # go build -o terraform-provider-citrixadc (runs fmtcheck first)
make install        # build + copy into ~/.terraform.d/plugins/registry.terraform.io/citrix/citrixadc/<VERSION>/<os_arch>/
make fmt            # gofmt -s -w on all non-vendor .go files
make fmtcheck       # scripts/gofmtcheck.sh — CI gate, run before committing
make vet
make lint           # golangci-lint (install via `make tools`)
```

Testing:

- `make test` is intentionally a **no-op** (see GNUmakefile comment) — use `testacc`.
- Tests are **acceptance tests** that hit a **live NetScaler**. They require
  `NS_URL`, `NS_LOGIN`, `NS_PASSWORD` env vars pointing at a real/VPX ADC; there is no
  mocked unit-test layer.

```bash
# Full acceptance suite (long):
make testacc

# Single test / package (preferred during dev):
TF_ACC=1 go test ./citrixadc -run TestAccLbvserver_basic -v -timeout 120m

# Compile-check a package's tests without running them:
make test-compile TEST=./citrixadc
```

## Provider configuration

Configured via `username`/`password`/`endpoint` provider args or the `NS_LOGIN`,
`NS_PASSWORD`, `NS_URL` environment variables. Both providers also support
`insecure_skip_verify`, `proxied_ns` (NetScaler Console / MAS proxy via
`_MPS_API_PROXY_MANAGED_INSTANCE_IP`), `partition`, `do_login`, and `is_cloud`. Keep the
two providers' schemas and configure logic in sync — they are defined separately in
`citrixadc/provider.go` and `citrixadc_framework/provider/provider.go`.

## Things that are easy to get wrong

- **Config is not persisted on the ADC automatically.** The provider writes running
  config only; run `ns_commit.sh` (or `terraform apply && ./ns_commit.sh`) to save to the
  ADC persistent store.
- **Docs are hand-maintained**, not generated: `docs/resources/*.md` (~780 files) and
  `docs/data-sources/*.md`. A new/migrated resource needs a matching doc page.
- **`examples/`** holds runnable HCL grouped by use case (`adc_usecases`,
  `basic_adc_operations`, `cluster_3node`, etc.); `make tffmt` formats them.
- Upstream code contributions are suspended per `CONTRIBUTING.md` — this repo is the
  vendor's working tree, so the migration branch work is internal.

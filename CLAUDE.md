# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this project is

Comer is a Go CLI **code-scaffolding generator** (module `github.com/imoowi/comer`, Go 1.20). It generates a complete Gin + GORM + Redis + Casbin + Captcha web-API project skeleton, and can also add individual controllers/services/models to an existing project (either from its built-in templates or from user-provided custom templates).

Crucially, this repo serves two distinct roles at once:

1. **The CLI/generator itself** — `main.go` → `cmd/` → `comer/`.
2. **The runtime library** that *generated* projects import back as `github.com/imoowi/comer` — the packages `interfaces/`, `components/`, `utils/`, and `validators/`. These are not part of the CLI flow; they are shipped to every scaffolded project and referenced by the generated code. A change in these packages changes the behavior of *generated* projects, not the generator.

## Commands

Build / install the CLI:

```sh
go build ./...          # compile check
go install .            # build the `comer` binary into $GOPATH/bin
go run . <subcommand>   # run without installing
```

Test:

```sh
go test ./...                          # run all tests (CI runs this with -cover)
go test ./validators -run TestValidator # run a single test
go test ./cmd -run TestCreate           # exercise the cobra command in-process
```

There is no Makefile. CI is GitHub Actions: `.github/workflows/test.yml` runs `go test ./...` on push (Go 1.20, uploads coverage to CodeCov); `release-tag.yml` cuts a GitHub Release when a `v*` tag is pushed.

## Architecture

### Command layer → generation layer

- `main.go` calls `cmd.Execute()`.
- `cmd/` holds thin [cobra](https://github.com/spf13/cobra) command definitions (`root.go`, `new.go`, `add.go`, `add_with_tpl.go`, `version.go`). Each subcommand parses flags and delegates to a method on the `Comer` type in the `comer` package. Flag names use backtick-string literals (e.g. `` `app` ``, `` `swaggerTags` ``, `` `tplVersion` ``).
- `comer/` is the actual generation engine, centered on the `Comer` struct (`comer/comer.go`). `NewComer()` is the constructor. Key entry points:
  - `Start()` — `comer new [module]` → `init` (v1) or `initV2` (v2, default) → `generateFrameworkDir` / `generateFrameworkFiles` → `showTips`.
  - `AddApp()` — `comer add` → `initApp` (v1) or `initAppV2` (v2) → `generateAppDir` / `generateAppFiles`.
  - `GenAppWithTpl()` — `comer add-with-tpl`, reads a user-supplied `.comer-templates/setting.json5`.

### Templates

- Templates are embedded into the binary via `//go:embed templates` in `comer/comer.var.go` and rendered with `text/template` in `comer/comer.tool.go` (`generateFileByMap`).
- Two template generations live under `comer/templates/`:
  - `v1/` — legacy layout (`apps/<app>/handlers|models|repos|services|migrates`).
  - `v2/` — current layout (`internal/controllers|services|repos|models|migrates|router|middlewares|global`, plus `cmd/`, `configs/`, `docs/`).
- Selection is by the `-v`/`tplVersion` flag: empty → v2 (default); `1` → v1. For `add`, passing `-a=<app>` also forces the v1 path.

### Template data & naming conventions

Generation fills a `map[string]any` of name variants (see `comer/comer.add.init.go` and `comer/comer.add_with_tpl.go`). Templates reference these keys, so a new name must provide all variants. The canonical set for a "handler/controller" name is: `HandlerName`, `lHandlerName` (lowercased first letter), `handlerName`, `handler_name` (snake), `handler-name` (dash), `handlerName2Dash`, `handlerName2Snake`; plus `ServiceName`, `ModelName`, `SwaggerTags`, `ModuleName`/`moduleName`.

File names are derived with `utils/format` — `Camel2Snake` (e.g. `PostPlus` → `post_plus`) and `Camel2Dash`. Defaults cascade when flags are omitted: `service` → `controller`, `model` → `service`, `swaggerTags` → `app`.

### Library packages (consumed by generated projects)

- `interfaces/` — generic contracts `IModel`, `IFilter`, `IRepo[T]`, `IService[T]`. `interfaces/impl/` provides generic default implementations (`Repo[T]`, `Service[T]`, `Filter`) built on GORM.
- `components/` — MySQL (GORM), Redis, Captcha, and an in-memory typed cache `MemCacheT[T]`.
- `utils/` — `response` (OK/Error/`PageListT`), `request` (pagination), `format`, `slice`, `myfile`, `password`, `maker` (serial numbers), `office` (Excel), and a logrus logger.
- `validators/` — custom `go-playground/validator` rules (`chinese`, `english`, `mobile`, `email`, `idcard`). Each file self-registers in `init()`; `InitValidators()` binds them into Gin's validator engine.

## Conventions & gotchas

- The CLI version is hardcoded in `comer/comer.version.go` (`Version()` returns `` `v1.3.18` ``), not read from git or a build variable. Update it there alongside `CHANGELOG.md` (which is changeset-style, one entry per patch version).
- Generation is deliberately **idempotent/append-only**: existing files are skipped (never overwritten), and v1 `add` injects import/route lines by string-matching sentinel lines like `do-not-delete-this-line` (see `comer/comer.add.go`).
- `generateFileByMap` uses `os.OpenFile(..., 0755)` with `O_CREATE` only when the target doesn't exist; files are rendered via `text/template` (not `html/template`).
- `utils/logger.go` writes rotated logs to `runtime/log` with a 7-day retention; `Logger.Out` is set to `io.Discard` so it only logs through the hook.

# Privateer Plugin for {{ .ServiceName }}

This wireframe is designed to quickly get your service pack repository up to speed. Clone the repo, adjust the `data/` catalogs and `evaluation_plans/` steps, and then use the config + build assets included here.

## Quick Start

```bash
cp example-config.yml config.yml
make binary               # builds ./{{ .ServiceName }}
```

Logs and result artifacts are written to `evaluation_results/<service_name>/` unless you override `write-directory` in your config.

## Run in Debug Mode

```bash
./{{ .ServiceName }} debug --service myService1
```

Use the same `service` value that you defined under `services.<service>` in `config.yml`. Debug mode uses the embedded catalogs and executes entirely within this repo.

## Run via Privateer Core

```bash
cp {{ .ServiceName }} $HOME/.privateer/bin/
cp config.yml ../        # copy next to your privateer config if needed
privateer run --service myService1
```

Install the compiled binary into your Privateer binaries path (default `$HOME/.privateer/bin`) and ensure the root config references `plugin: {{ .ServiceName }}` under the matching service entry.

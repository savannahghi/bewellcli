# Bewell Command Line Application

## Installing the Command

```bash
go install github.com/savannahghi/bewellcli
```

## Environment variables

```bash
# schema registry env
export REGISTRY_URL="<Test schema registry URL>"
```

## Development

The CLI is built using [Cobra](https://github.com/spf13/cobra). Follow this [User Guide](https://github.com/spf13/cobra/blobmaster/user_guide.md) to get started

## Using the Command

- Help

```bash
bewellcli --help
```

- Validating a Graphql Schema

```bash
bewellcli service validate-schema --help
```

Example use:-

```bash
bewellcli service validate-schema --name {SERVICE_NAME} --version {SERVICE_VERSION} --dir {SCHEMA_DIR}
```

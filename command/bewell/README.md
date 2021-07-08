# Bewell Command Line Application

## Installing the Command

```bash
go install gitlab.slade360emr.com/go/base/command/bewell
```

## Environment variables

```bash
# Go private modules
export GOPRIVATE="gitlab.slade360emr.com/go/*,gitlab.slade360emr.com/optimalhealth/*"

# schema registry env
export REGISTRY_URL="<Test schema registry URL>"
```

## Development

The CLI is built using [Cobra]("https://github.com/spf13/cobra"). Follow this [User Guide]("https://github.com/spf13/cobra/blob/master/user_guide.md") to get started

## Using the Command

- Help

```bash
bewell --help
```

- Validating a Graphql Schema

```bash
bewell service validate-schema --help
```

Example use:-

```bash
bewell service validate-schema --name {SERVICE_NAME} --version {SERVICE_VERSION} --dir {SCHEMA_DIR}
```

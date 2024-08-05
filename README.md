# shflow

[![Build Status](https://github.com/cligpt/shflow/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/cligpt/shflow/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/cligpt/shflow/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/cligpt/shflow)
[![Go Report Card](https://goreportcard.com/badge/github.com/cligpt/shflow)](https://goreportcard.com/report/github.com/cligpt/shflow)
[![License](https://img.shields.io/github/license/cligpt/shflow.svg)](https://github.com/cligpt/shflow/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/cligpt/shflow.svg)](https://github.com/cligpt/shflow/tags)

## Introduction

*shflow* is the workflow toolchain of [shai](https://github.com/cligpt/shai) written in Go.

## Prerequisites

- Go >= 1.22.0

## Quickstart

```bash
TBD
```

For more examples, see the [examples](examples) directory. For more information on working with a flowfile, see the [flowfile](docs/flowfile.md) documentation.

## CLI Reference

### Init task

`shflow init` is used to init task.

```bash
shflow init mytask
```

### Build workflow

`shflow build` is used to build workflow from a flowfile.

```bash
shflow build myworkflow -f ./flowfile
```

### Train workflow

`shflow train` is used to train workflow.

```bash
shflow train myworkflow
```

### Test workflow

`shflow test` is used to test workflow.

```bash
shflow test myworkflow
```

### Run workflow

`shflow run` is used to run workflow.

```bash
shflow run myworkflow
```

### Show workflow

`shflow show` is used to show workflow.

```bash
shflow show --flowfile
```

### Compose workflows

`shflow compose` is used to compose workflows.

```bash
shflow compose -f ./flowcompose.yml
```

## Building

See the [developer guide](https://github.com/cligpt/shflow/blob/main/docs/development.md)

## REST API

See the [API documentation](./docs/api.md) for all endpoints.

## License

Project License can be found [here](LICENSE).

## Reference

- [autogen](https://github.com/microsoft/autogen)
- [crewai](https://github.com/crewAIInc/crewAI)
- [dockerfile](https://docs.docker.com/reference/dockerfile/)
- [langchain](https://python.langchain.com/)
- [llama-agents](https://github.com/run-llama/llama-agents)
- [modelfile](https://github.com/ollama/ollama/blob/main/docs/modelfile.md)

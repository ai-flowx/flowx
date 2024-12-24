# flowx

[![Build Status](https://github.com/ai-flowx/flowx/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/ai-flowx/flowx/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/ai-flowx/flowx/branch/main/graph/badge.svg?token=El8oiyaIsD)](https://codecov.io/gh/ai-flowx/flowx)
[![Go Report Card](https://goreportcard.com/badge/github.com/ai-flowx/flowx)](https://goreportcard.com/report/github.com/ai-flowx/flowx)
[![License](https://img.shields.io/github/license/ai-flowx/flowx.svg)](https://github.com/ai-flowx/flowx/blob/main/LICENSE)
[![Tag](https://img.shields.io/github/tag/ai-flowx/flowx.svg)](https://github.com/ai-flowx/flowx/tags)



## Introduction

*flowx* is the ai framework of [ai-flowx](https://github.com/ai-flowx) written in Go.



## Prerequisites

- Go >= 1.22.0



## Build

```bash
version=latest make build
```



## Usage

```
Usage:
  flowx [flags]

Flags:
  -c, --config-file string   config file
  -h, --help                 help for flowx
  -l, --listen-port string   listen port (default ":8080")
  -v, --version              version for flowx
```



## Settings

*flowx* parameters can be set in the directory [config](https://github.com/ai-flowx/flowx/blob/main/config).

An example of configuration in [config.yml](https://github.com/ai-flowx/flowx/blob/main/config/config.yml):

```yaml
cache:
  provider: cachex
  url: http://127.0.0.1:8081
gpt:
  provider: openai
  api: https://openai.com/api
  token: token
memory:
  type: shortterm
store:
  provider: ragx
  url: http://127.0.0.1:8082
  path: /path/to/file
tool:
  - name: basetool
  - name: decorator
  - name: hello
  - name: runnables
  - name: structuredtool
```



## Architecture

![arch](./arch.png "Architecture")



## Features

- **Role-Based Agents**

  Create specialized agents with defined roles, expertise, and goals


- **Flexible Tools**

  Equip agents with custom tools and APIs to interact with external services and data sources


- **Intelligent Collaboration**

  Agents work together, sharing insights and coordinating tasks to achieve complex objectives


- **Task Management**

  Define sequential or parallel workflows, with agents automatically handling task dependencies



## License

Project License can be found [here](LICENSE).



## Reference

### Binding

- [gopy](https://github.com/go-python/gopy)

### Chat

- [dify-on-wechat](https://docs.dify.ai/zh-hans/learn-more/use-cases/dify-on-wechat)

### Framework

- [autogen](https://github.com/microsoft/autogen)
- [cobra-cli](https://github.com/spf13/cobra-cli)
- [crewai](https://github.com/crewAIInc/crewAI)
- [dockerfile](https://docs.docker.com/reference/dockerfile/)
- [fabric](https://github.com/danielmiessler/fabric)
- [langchain-agents](https://www.langchain.com/agents)
- [langgraph](https://langchain-ai.github.io/langgraph/)
- [llama-agents](https://github.com/run-llama/llama-agents)
- [modelfile](https://github.com/ollama/ollama/blob/main/docs/modelfile.md)
- [semantic-kernel](https://github.com/microsoft/semantic-kernel)

### Protocol

- [claude-mcp](https://ai-claude.net/mcp/)
- [claude-mcp-servers](https://github.com/modelcontextprotocol/servers)
- [langgraph-agent-protocol](https://github.com/langchain-ai/agent-protocol)

### Store

- [chroma-go](https://github.com/amikos-tech/chroma-go)

### Tools

- [crewai-tools](https://docs.crewai.com/concepts/tools)
- [crewai-tools](https://github.com/crewAIInc/crewAI-tools)
- [langchain-tools](https://python.langchain.com/docs/integrations/tools/)
- [langchaingo](https://github.com/tmc/langchaingo)

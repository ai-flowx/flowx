# Shflow Flow File

A flow file is the blueprint to build and run workflow with shflow.

## Table of Contents

- [Format](#format)
- [Examples](#examples)
- [Instructions](#instructions)
    - [FROM](#from)
    - [AGENT](#agent)
    - [TASK](#task)
    - [TOOL](#tool)
    - [CREW](#crew)
- [Notes](#notes)

## Format

The format of the `flowfile`:

```flowfile
# comment
INSTRUCTION arguments
```

| Instruction       | Description                                                                          |
|-------------------|--------------------------------------------------------------------------------------|
| [`FROM`](#from)   | Base workflow to use.                                                                |
| [`AGENT`](#agent) | Autonomous unit to perform tasks, make decisions, and communicate with other agents. |
| [`TASK`](#task)   | Specific assignment completed by agent.                                              |
| [`TOOL`](#tool)   | Skill or function that agents can utilize to perform various actions.                |
| [`CREW`](#crew)   | Collaborative group of agents working together to achieve a set of tasks.            |

## Examples

### Basic `flowfile`

An example of `flowfile` building a workflow:

```flowfile
TBD
```

To use this:

1. Save it as a file (e.g. `flowfile`)
2. `shflow build choose-a-workflow-name -f <location of the file e.g. ./flowfile>'`
3. `shflow run choose-a-workflow-name`

More examples are available in the [examples directory](../examples).

To view the flowfile of a given workflow, use the `shflow show --flowfile` command.

## Instructions

### FROM

The `FROM` instruction defines the base workflow to use when building a workflow.

```flowfile
FROM <workflow name>:<tag>
```

### AGENT

The `AGENT` is an autonomous unit programmed to:
- Perform tasks
- Make decisions
- Communicate with other agents.

| Parameter      | Description                                                                                                        |
|----------------|--------------------------------------------------------------------------------------------------------------------|
| `ROLE`         | Defines the agent's function. It determines the kind of tasks the agent is best suited for.                        |
| `GOAL`         | The individual objective that the agent aims to achieve. It guides the agent's decision-making process.            |
| `BACKSTORY`    | Provides context to the agent's role and goal, enriching the interaction and collaboration dynamics.               |
| `LLM`          | Represents the language model that will run the agent.                                                             |
| `TOOLS`        | Set of capabilities or functions that the agent can use to perform tasks.                                          |
| `MAXITER`      | The maximum number of iterations the agent can perform before being forced to give its best answer. Default is 25. |
| `MAXRPM`       | The maximum number of requests per minute the agent can perform to avoid rate limits.                              |
| `MAXEXECUTION` | The Maximum execution time for an agent to execute a task.                                                         |

```flowfile
TBD
```

### TASK

The `TASK` is specific assignment completed by agent. It provides all necessary details for execution,
such as a description, the agent responsible, required tools, and more, facilitating a wide range of action complexities.

| Parameter        | Description                                                                                         |
|------------------|-----------------------------------------------------------------------------------------------------|
| `DESCRIPTION`    | A clear, concise statement of what the task entails.                                                |
| `AGENT`          | Responsible for the task.                                                                           |
| `EXPECTEDOUTPUT` | A detailed description of what the task's completion looks like.                                    |
| `TOOLS`          | The functions or capabilities the agent can utilize to perform the task.                            |
| `ASYNCEXECUTION` | The task executes asynchronously, allowing progression without waiting for completion.              |
| `CONTEXT`        | Specifies tasks whose outputs are used as context for this task.                                    |
| `OUTPUT`         | The output of the task, containing the raw and JSON.                                                |
| `HUMANINPUT`     | Indicates if the task requires human feedback at the end, useful for tasks needing human oversight. |

```flowfile
TBD
```

### TOOL

The `TOOL` empowers agent with capabilities ranging from web searching and data analysis to collaboration and delegating
tasks.

```flowfile
TBD
```

### CREW

The `CREW` represents a collaborative group of agents working together to achieve a set of tasks.
Each crew defines the strategy for task execution, agent collaboration, and the overall workflow.

| Parameter      | Description                                                                                                         |
|----------------|---------------------------------------------------------------------------------------------------------------------|
| `TASKS`        | A list of tasks assigned to the crew.                                                                               |
| `AGENTS`       | A list of agents that are part of the crew.                                                                         |
| `PROCESS`      | The process flow (e.g., sequential, hierarchical) the crew follows.                                                 |
| `MANAGERLLM`   | The language model used by the manager agent in a hierarchical process. Required when using a hierarchical process. |
| `MAXRPM`       | Maximum requests per minute the crew adheres to during execution.                                                   |
| `LANGUAGE`     | Language used for the crew, defaults to English.                                                                    |
| `LANGUAGEFILE` | Path to the language file to be used for the crew.                                                                  |
| `LANGUAGEFILE` | Path to the language file to be used for the crew.                                                                  |
| `MANAGERAGENT` | Manager sets a custom agent that will be used as a manager.                                                         |
| `MANAGERAGENT` | Manager sets a custom agent that will be used as a manager.                                                         |
| `OUTPUT`       | The output of the task, containing the raw and JSON.                                                                |
| `PROMPTFILE`   | Path to the prompt JSON file to be used for the crew.                                                               |

```flowfile
TBD
```

## Notes

- the **`flowfile` is not case sensitive**. In the examples, uppercase instructions are used to make it easier to distinguish it from arguments.
- Instructions can be in any order. In the examples, the `FROM` instruction is first to keep it easily readable.

## Reference

- [crewai](https://github.com/crewAIInc/crewAI)

# Shai Flow File

A flow file is the blueprint to create and run workflow with shflow.

## Table of Contents

- [Format](#format)
- [Examples](#examples)
- [Instructions](#instructions)
    - [FROM (Required)](#from-required)
        - [Build from llama3](#build-from-llama3)
        - [Build from a bin file](#build-from-a-bin-file)
    - [PARAMETER](#parameter)
        - [Valid Parameters and Values](#valid-parameters-and-values)
- [Notes](#notes)

## Format

The format of the `flowfile`:

```flowfile
# comment
INSTRUCTION arguments
```

| Instruction                         | Description                                               |
| ----------------------------------- |-----------------------------------------------------------|
| [`FROM`](#from-required) (required) | Defines the base workflow to use.                         |
| [`PARAMETER`](#parameter)           | Sets the parameters for how shflow will run the workflow. |

## Examples

### Basic `flowfile`

An example of a `flowfile` creating a workflow:

```flowfile
TBD
```

To use this:

1. Save it as a file (e.g. `flowfile`)
2. `shflow create choose-a-workflow-name -f <location of the file e.g. ./flowfile>'`
3. `shflow run choose-a-workflow-name`

More examples are available in the [examples directory](../examples).

To view the flowfile of a given workflow, use the `shflow show --flowfile` command.

  ```bash
  TBD
  ```

## Instructions

### FROM (Required)

The `FROM` instruction defines the base workflow to use when creating a workflow.

```flowfile
FROM <workflow name>:<tag>
```

### PARAMETER

The `PARAMETER` instruction defines a parameter that can be set when the workflow is run.

```flowfile
PARAMETER <parameter> <parametervalue>
```

#### Valid Parameters and Values

| Parameter | Description | Value Type | Example Usage |
|-----------|-------------| ---------- |---------------|
| TBD       | TBD         | int        | TBD           |

## Notes

- the **`flowfile` is not case sensitive**. In the examples, uppercase instructions are used to make it easier to distinguish it from arguments.
- Instructions can be in any order. In the examples, the `FROM` instruction is first to keep it easily readable.

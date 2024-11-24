# Developer Documentation

## Table of Contents

- [Developer Documentation](#developer-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Subject](#subject)
  - [Git](#git)
    - [Commit](#commit)
    - [Branch](#branch)
    - [Pull Request](#pull-request)
  - [Architecture](#architecture)
    - [File](#file)
    - [Folder - Directory](#folder---directory)
    - [Visual](#visual)

## Main Document

[main documentation](../README.md)
## Subject

This documentation aims to list the architecture of the project, github branches and commit standards. How the project is structured so you can implement other features while keeping the same implementation standard.

## Git

### Commit

- [status] header message body message
- Possible status:
  - UPD update code like refactor, clean code
  - ADD add new file / feature
  - DEL delete file / feature
  - WIP work in progress
- Example:
- “[UPD] code authentication” “modify file auth.go …”

### Branch

- All branch: kebab case
- Feature branch: [upper branch name]-feature/[feature-name]

```mermaid
flowchart TD
    subgraph MainWorkflow
        pre-prod --> |Pull Request| main

    subgraph PreProdWorkflow
        dev --> |Pull Request| pre-prod

    subgraph DevWorkflow
        dev --> frontend
        dev --> backend
        dev --> mobile
        dev --> devops
        dev --> documentation
    end

    subgraph Frontend
        frontend --> |Pull Request| dev
        frontend --> frontend-feature/feature*
    end

    subgraph Backend
        backend --> |Pull Request| dev
        backend --> backend-feature/feature*
    end

    subgraph Mobile
        mobile --> |Pull Request| dev
        mobile --> mobile-feature/feature*
    end

    subgraph DevOps
        devops --> |Pull Request| dev
        devops --> devops-feature/feature*
    end

    subgraph Documentation
        documentation --> |Pull Request| dev
        documentation --> documentation-feature/feature*
    end
    end
    end
```

### Pull Request

You have to assign someone

## Architecture

### File

camel case

### Folder - Directory

kebab case

### Visual

```mermaid
mindmap
  root((root of the repository))
    backend
    docs
    frontend
    mobile
    poc
        backend
            gleam
            go
            java
        frontend
            angular
            nuxt
            remix
        mobile
            kotlin
            react-native
```

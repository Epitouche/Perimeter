# Developer Documentation

[main documentation](../README.md)

# SUBJECT

This documentation aims to list the architecture of the project, github branches and commit standards. How the project is structured so you can implement other features while keeping the same implementation standard.


# Commit standards

*status header message   body message*

Possible status:
 - UPD update code like refactor, clean code
 - ADD add new file / feature
 - DEL delete file / feature

**example:**
 - **[UPD] code authentication” “modify file auth.go ...**

# Branch standards

**format: kebab case**
*Feature branch: [upper branch name]-feature/[feature-name]*

## Achitecture

```mermaid
    gitGraph
        branch dev
        commit

        checkout dev
        branch backend

        checkout backend
        branch backend-feature/feature1
        commit

        checkout backend
        branch backend-feature/feature2
        commit

        checkout backend
        branch backend-feature/feature3
        commit

        checkout dev
        branch frontend

        checkout frontend
        branch frontend-feature/feature1
        commit

        checkout frontend
        branch frontend-feature/feature2
        commit

        checkout frontend
        branch frontend-feature/feature3
        commit

        checkout dev
        branch mobile

        checkout mobile
        branch mobile-feature/feature1
        commit

        checkout mobile
        branch mobile-feature/feature2
        commit

        checkout mobile
        branch mobile-feature/feature3
        commit

        checkout dev
        branch devops

        checkout devops
        branch devops-feature/feature1
        commit

        checkout devops
        branch devops-feature/feature2
        commit

        checkout devops
        branch devops-feature/feature3
        commit
```
# Dev Container Installation

## Table of Contents

- [Dev Container Installation](#dev-container-installation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Dev](#dev)
    - [Dev Frontend](#dev-frontend)
    - [Dev Backend](#dev-backend)
  - [Exit Dev Container](#exit-dev-container)
  - [Remove Container](#remove-container)

---

## Main Document

[main documentation](../README.md)

---

## Dev

Install [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) in VSCode

setup `.env` file

### Dev Frontend

open `backend` dev container with : `Dev Containers: Rebuild and Reopen in Container` -> `backend`
switch to `frontend` container with : `Dev Containers: Switch Container` -> `frontend`
start dev to frontend

### Dev Backend

open `frontend` dev container with : `Dev Containers: Rebuild and Reopen in Container` -> `frontend`
switch to `backend` container with : `Dev Containers: Switch Container` -> `backend`
start dev to backend

## Exit Dev Container

exit dev container with : `Dev Containers: Reopen Folder Locally`

## Remove Container

```bash
docker compose -f compose.dev.yaml down --volumes
```

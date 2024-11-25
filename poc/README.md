# Docker

## Table of Contents

- [Docker](#docker)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [build all POC](#build-all-poc)
  - [run frontend](#run-frontend)
    - [frontend angular](#frontend-angular)
    - [frontend nuxt](#frontend-nuxt)
    - [frontend remix manual](#frontend-remix-manual)
    - [frontend remix auto](#frontend-remix-auto)
  - [run backend](#run-backend)
    - [backend go](#backend-go)
    - [backend gleam](#backend-gleam)
    - [backend java](#backend-java)
  - [run mobile](#run-mobile)
    - [mobile react native](#mobile-react-native)

---

## Main Document

Refer to the [main documentation](../README.md).

---

## build all POC

```bash
docker compose build
```

## run frontend

### frontend angular

```bash
docker compose up --build frontend-angular
```

### frontend nuxt

```bash
docker compose up --build frontend-nuxt
```

### frontend remix manual

```bash
docker compose up --build frontend-remix-manual
```

### frontend remix auto

```bash
docker compose up --build frontend-remix-auto
```

## run backend

### backend go

```bash
docker compose up --build backend-go
```

run test

```bash
docker compose up --build backend-go-test
```

### backend gleam

```bash
docker compose up --build backend-gleam
```

### backend java

```bash
docker compose up --build backend-java
```

## run mobile

### mobile react native

```bash
docker compose up --build mobile-react-native
```

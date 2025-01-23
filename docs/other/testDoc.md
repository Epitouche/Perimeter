# Test Documentation

## Table of Contents

- [Test Documentation](#test-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Backend](#backend)
    - [Setup](#setup)
    - [Run Test](#run-test)
    - [View Result](#view-result)
  - [Frontend](#frontend)
    - [install frontend dependencies](#install-frontend-dependencies)
    - [run frontend test](#run-frontend-test)
  - [Mobile](#mobile)
    - [install mobile dependencies](#install-mobile-dependencies)
    - [run mobile test](#run-mobile-test)

## Main Document

[Main Documentation](../../README.md)

## Backend

### Setup

Fill up the `.env` file.

### Run Test

Navigate to the backend directory:

```bash
./test.sh
```

This will generate a `coverage.out` file.

### View Result

In the backend directory, after [Run Test](#run-test):

```bash
go tool cover -html=coverage.out
```

## Frontend

Fill up the `.env` file.

```bash
cd mobile
```

### install frontend dependencies

```bash
npm install
```

### run frontend test

```bash
npx vitest run
```

## Mobile

Fill up the `.env` file.

```bash
cd mobile
```

### install mobile dependencies

```bash
npm install --legacy-peer-deps
```

### run mobile test

```bash
npm run test
```

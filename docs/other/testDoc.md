# Test Documentation

## Table of Contents

- [Test Documentation](#test-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Backend](#backend)
    - [Setup](#setup)
    - [Run Test](#run-test)
    - [View Result](#view-result)
  - [Mobile](#mobile)
    - [install dependencies](#install-dependencies)
    - [run test](#run-test-1)

## Main Document

[Main Documentation](../README.md)

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

## Mobile

Fill up the `.env` file.

```bash
cd mobile
```

### install dependencies

```bash
npm install --legacy-peer-deps
```

### run test

```bash
npm run test
```

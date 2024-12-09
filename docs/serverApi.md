# Server API Documentation

## Overview

This API provides authentication and user information retrieval features for multiple services, including GitHub, Gmail, Spotify, and user login and registration endpoints.

## Table of Contents

- [Server API Documentation](#server-api-documentation)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Swagger](#swagger)
  - [Endpoints](#endpoints)
    - [General Routes](#general-routes)
      - [`GET /ping`](#get-ping)
    - [GitHub Integration](#github-integration)
      - [`GET /github/auth`](#get-githubauth)
      - [`GET /github/auth/callback`](#get-githubauthcallback)
      - [`GET /github/info/user`](#get-githubinfouser)
    - [Gmail Integration](#gmail-integration)
      - [`GET /gmail/auth`](#get-gmailauth)
      - [`GET /gmail/auth/callback`](#get-gmailauthcallback)
      - [`GET /gmail/info/user`](#get-gmailinfouser)
    - [Spotify Integration](#spotify-integration)
      - [`GET /spotify/auth`](#get-spotifyauth)
      - [`POST /spotify/auth/callback`](#post-spotifyauthcallback)
      - [`GET /spotify/info/user`](#get-spotifyinfouser)
    - [User Management](#user-management)
      - [`POST /user/login`](#post-userlogin)
      - [`POST /user/register`](#post-userregister)
  - [Models](#models)
    - [`schemas.CodeCredentials`](#schemascodecredentials)
    - [`schemas.ErrorRespose`](#schemaserrorrespose)
    - [`schemas.JWT`](#schemasjwt)
    - [`schemas.Response`](#schemasresponse)
  - [Security](#security)
    - [Bearer Authentication](#bearer-authentication)

## Main Document

[main documentation](../README.md)

## Swagger

to access swagger local documentation
launch the server and go to the next url:
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

[swagger golang](https://github.com/swaggo/swag)

---

## Endpoints

### General Routes

#### `GET /ping`

- **Description:** Check if the server is running.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `ping route`
- **Responses:**
  - **200:** Server is running.

---

### GitHub Integration

#### `GET /github/auth`

- **Description:** Provides a URL for GitHub authentication.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Github`
- **Responses:**
  - **200:** URL for authentication (string).
  - **500:** Internal Server Error.

#### `GET /github/auth/callback`

- **Description:** GitHub authentication callback.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Github`
- **Responses:**
  - **200:** Successful response.
  - **500:** Internal Server Error.

#### `GET /github/info/user`

- **Description:** Retrieve GitHub user information.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Github`
- **Responses:**
  - **200:** User information.
  - **500:** Internal Server Error.

---

### Gmail Integration

#### `GET /gmail/auth`

- **Description:** Provides a URL for Gmail authentication.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Gmail`
- **Responses:**
  - **200:** URL for authentication (string).
  - **500:** Internal Server Error.

#### `GET /gmail/auth/callback`

- **Description:** Gmail authentication callback.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Gmail`
- **Responses:**
  - **200:** Successful response.
  - **500:** Internal Server Error.

#### `GET /gmail/info/user`

- **Description:** Retrieve Gmail user information.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Gmail`
- **Responses:**
  - **200:** User information.
  - **500:** Internal Server Error.

---

### Spotify Integration

#### `GET /spotify/auth`

- **Description:** Provides a URL for Spotify authentication.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Spotify`
- **Responses:**
  - **200:** URL for authentication.
  - **500:** Internal Server Error.

#### `POST /spotify/auth/callback`

- **Description:** Spotify authentication callback.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Spotify`
- **Body Parameters:**
  - **Payload (required):** JSON object containing:
    - `code` (string): Authentication code.
    - `state` (string): State parameter.
- **Responses:**
  - **200:** Successful response.
  - **500:** Internal Server Error.

#### `GET /spotify/info/user`

- **Description:** Retrieve Spotify user information.
- **Consumes:** `application/json`
- **Produces:** `application/json`
- **Tags:** `Spotify`
- **Responses:**
  - **200:** User information.
  - **500:** Internal Server Error.

---

### User Management

#### `POST /user/login`

- **Description:** Authenticates a user and provides a JSON Web Token (JWT) for API authorization.
- **Produces:** `application/json`
- **Tags:** `User`
- **Form Data Parameters:**
  - `username` (string, required): Username of the user.
  - `password` (string, required): User's password.
- **Responses:**
  - **200:** JWT token.
  - **401:** Unauthorized.

#### `POST /user/register`

- **Description:** Registers a user and provides a JSON Web Token (JWT) for API authorization.
- **Produces:** `application/json`
- **Tags:** `User`
- **Form Data Parameters:**
  - `email` (string, required): User's email.
  - `username` (string, required): Username of the user.
  - `password` (string, required): User's password.
- **Responses:**
  - **200:** JWT token.
  - **401:** Unauthorized.

---

## Models

### `schemas.CodeCredentials`

- **Type:** Object
- **Properties:**
  - `code` (string): Authentication code.
  - `state` (string): State parameter.

### `schemas.ErrorRespose`

- **Type:** Object
- **Properties:**
  - `error` (string): Error message.

### `schemas.JWT`

- **Type:** Object
- **Properties:**
  - `token` (string): JSON Web Token.

### `schemas.Response`

- **Type:** Object
- **Properties:**
  - `message` (string): Response message.

---

## Security

### Bearer Authentication

- **Type:** `apiKey`
- **Name:** `Authorization`
- **In:** Header

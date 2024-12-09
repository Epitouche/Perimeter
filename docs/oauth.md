# Oauth Documentation

## Table of Contents

- [Oauth Documentation](#oauth-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [View](#view)

## Main Document

[main documentation](../README.md)

## View

`ServiceExample` is the name for the service for this example

```mermaid
zenuml
    title OAuth 2.0 Login Flow

    actor User
    participant "Web Client (Frontend)" as Frontend
    participant "Application Server (Backend)" as Backend
    participant "OAuth Provider" as OAuthProvider

    User -> Frontend : Clicks "Login with OAuth"
    Frontend -> OAuthProvider : Redirects user to Authorization Endpoint (state, client_id, redirect_uri)
    User -> OAuthProvider : Provides credentials and authorizes
    OAuthProvider -> Frontend : Redirects to frontend with Authorization Code
    Frontend -> Backend : Sends Authorization Code (along with state)

    Backend -> OAuthProvider : Exchanges Authorization Code for Access Token
    OAuthProvider -> Backend : Returns Access Token (+ optionally Refresh Token)
    Backend -> Backend : Validates token and retrieves user info
    Backend -> Frontend : Sends session data or JWT
    Frontend -> User : Logs in the user
```

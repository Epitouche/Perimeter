# Database Documentation

## Table of Contents

- [Database Documentation](#database-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [View](#view)

## Main Document

[main documentation](../../README.md)

## View

```mermaid
erDiagram
    USER {
        int id PK
        string email
        string username
        string password
        id token_auth_service FK
        datetime created_at
        datetime updated_at
    }
    SERVICE {
        int id PK
        string name
        string description
        datetime created_at
        datetime updated_at
    }
    ACTION {
        int id PK
        int service_id FK
        string name
        string description
        datetime created_at
        datetime updated_at
    }
    REACTION {
        int id PK
        int service_id FK
        string name
        string description
        datetime created_at
        datetime updated_at
    }

    AREA {
        int id PK
        int user_id FK
        int action_id FK
        int reaction_id FK
        bool enable
        datetime created_at
        datetime updated_at
    }

    AREA_RESULT {
        int id PK
        int area_id FK
        string result
        datetime created_at
    }

    TOKEN {
        int id PK
        int user_id FK
        int service_id FK
        string token
        string refresh_token
        datetime created_at
        datetime updated_at
    }

    USER ||--o{ AREA : "creates"
    SERVICE ||--o{ ACTION : "provides"
    SERVICE ||--o{ REACTION : "provides"
    ACTION ||--o{ AREA : "triggers"
    REACTION ||--o{ AREA : "executes"
    USER |o--|| TOKEN : "connection"
    TOKEN ||--|| USER : "link"
    TOKEN ||--|| SERVICE : "link"
    AREA ||--o{ AREA_RESULT : "store result"
```

# Database Documentation

## Table of Contents

- [Database Documentation](#database-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [View](#view)

## Main Document

[main documentation](../README.md)

## View

```mermiad
erDiagram
    USER {
        int id PK
        string email
        string password
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
        datetime created_at
        datetime updated_at
    }

    USER ||--o{ AREA : "creates"
    SERVICE ||--o{ ACTION : "provides"
    SERVICE ||--o{ REACTION : "provides"
    ACTION ||--o{ AREA : "triggers"
    REACTION ||--o{ AREA : "executes"
```

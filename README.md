# English Note - Simple Software Design Document (SDD)

## 1. Overview

**English Note** is a simple vocabulary learning application that allows users to:

- Manage vocabulary groups (word groups)
- Create and store vocabulary in each group
- Look up meanings, pronunciations, synonyms, antonyms
- Register, log in, and authenticate using JWT
- Simple and user-friendly UI, ideal for English learners

## 2. System Architecture

### Backend Stack:

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: MongoDB
- **Auth**: JWT-based (access + refresh token)
- **Deployment**: Docker-ready (optional)

### Architecture Layers:

Client → Router → Middleware → Handler → Service → Repository → MongoDB

## 3. Data Architecture

### 3.1 Database Schema (ERD)


### 3.2 Relationships

- A **User** can have multiple **WordGroups**
- A **WordGroup** belongs to one **User**
- A **WordGroup** can contain many **Words**
- Each **Word** belongs to one **WordGroup** and **User**
- A **User** can have multiple **RefreshTokens** (for multi-device sessions)

### 3.3 Field Details

#### User

| Field      | Type      | Description          |
|------------|-----------|----------------------|
| _id        | ObjectId  | User ID              |
| username   | string    | Username             |
| email      | string    | User email           |
| password   | string (hashed) | Encrypted password |

#### WordGroup

| Field      | Type      | Description          |
|------------|-----------|----------------------|
| _id        | ObjectId  | Group ID             |
| userId     | ObjectId  | Owner of the group   |
| name       | string    | Name of the group    |
| createdAt  | int64     | Creation timestamp   |

#### Word

| Field         | Type      | Description          |
|---------------|-----------|----------------------|
| _id           | ObjectId  | ID of the word       |
| groupId       | ObjectId  | Word group it belongs to |
| userId        | ObjectId  | Creator of the word  |
| word          | string    | The vocabulary word  |
| language      | string    | Language (default: "English") |
| pronunciations| Array     | Pronunciation (IPA + audio) |
| meanings      | Array     | Meanings of the word |
| synonyms      | string[]  | Synonyms             |
| antonyms      | string[]  | Antonyms             |
| createdAt     | int64     | Creation timestamp   |

#### Pronunciation

| Field    | Type     | Description    |
|----------|----------|----------------|
| ipa      | string   | IPA notation   |
| audioUrl | string   | Audio URL      |

#### Meaning

| Field           | Type     | Description           |
|-----------------|----------|-----------------------|
| part_of_speech  | string   | Part of speech        |
| definition      | string   | Word definition       |
| examples        | string[] | Example usage sentences|

#### RefreshToken

| Field        | Type     | Description         |
|--------------|----------|---------------------|
| _id          | ObjectId | Token ID            |
| userId       | string   | Owner user ID       |
| token        | string   | Token value         |
| createdAt    | int64    | Creation time       |
| expiredAt    | int64    | Expiry time         |
| ip           | string   | Request IP address  |
| userAgent    | string   | Device / browser info|

## 4. Interface Design

### 4.1 API Design (REST)

#### 🧑 Auth

| Method | Endpoint                     | Description            |
|--------|------------------------------|------------------------|
| POST   | /api/auth/register            | Register a new user    |
| POST   | /api/auth/login               | Log in                |
| POST   | /api/auth/refresh-token       | Refresh access token  |

#### 📦 Word Groups

| Method | Endpoint                     | Description            |
|--------|------------------------------|------------------------|
| GET    | /api/word-groups             | Get user’s word groups |
| POST   | /api/word-groups             | Create new word group |
| PUT    | /api/word-groups/:id         | Update a word group   |
| DELETE | /api/word-groups/:id         | Delete a word group   |

#### 📘 Words

| Method | Endpoint                             | Description                   |
|--------|--------------------------------------|-------------------------------|
| GET    | /api/words?groupId=xxx&page=1&limit=10 | Get words in a group (paginated) |
| POST   | /api/words                           | Create a word in a group      |
| GET    | /api/words/:id                       | Get word detail               |
| PUT    | /api/words/:id                       | Update a word                 |
| DELETE | /api/words/:id                       | Delete a word                 |

### 4.2 Postman Usage

The `access_token` variable can be stored from the login response:

```javascript
pm.environment.set("access_token", pm.response.json().access_token);

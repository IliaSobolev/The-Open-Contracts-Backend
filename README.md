
# The Open Contracts API

## Project Overview

Open Contracts is a platform designed to facilitate the sharing of best-practice code snippets for smart contracts on the TON (The Open Network) blockchain. Initiated under the The-Open-Contracts project, it empowers developers to contribute, discover, and utilize reusable code blocks to streamline smart contract development.

## API Documentation

This API provides endpoints to create, retrieve, and list code blocks for smart contracts. It is built using the Gin framework in Go and follows RESTful principles.

### Navigation
- [Create a Code Block](#1-create-a-code-block)
- [Get a Code Block by ID](#2-get-a-code-block-by-id)
- [List All Code Blocks](#3-list-all-code-blocks)

### Base Path
```
/toc/v1/codeblock
```

### Data Models

#### CodeBlock
Represents a stored code block in the system.

```json
{
  "_id": "string",
  "author_id": "integer",
  "title": "string",
  "description": "string",
  "rating": "integer",
  "lang": "string",
  "body": "string"
}
```

**Fields**:
- `_id`: Unique identifier for the code block
- `author_id`: ID of the author who created the code block
- `title`: Title of the code block
- `description`: Description of the code block
- `rating`: Rating of the code block
- `lang`: Programming language (e.g., func, fift, tact, tolk)
- `body`: The code content

#### CodeBlockDTO
Data Transfer Object used for creating a code block.

```json
{
  "author_id": "integer",
  "title": "string",
  "description": "string",
  "lang": "string",
  "body": "string"
}
```

**Fields**:
- `author_id`: ID of the author
- `title`: Title of the code block
- `description`: Description of the code block
- `lang`: Programming language
- `body`: The code content

### Supported Languages
The API supports the following programming languages for code blocks:
- `func`
- `fift`
- `tact`
- `tolk`

### Endpoints

#### 1. Create a Code Block
- **Method**: POST
- **Path**: `/toc/v1/codeblock/create`
- **Description**: Creates a new code block.
- **Request Body**:
  ```json
  {
    "author_id": 123,
    "title": "Sample Contract",
    "description": "A sample smart contract",
    "lang": "func",
    "body": "code content here"
  }
  ```
- **Responses**:
  - **200 OK**:
    ```json
    {
      "message": "create codeBlock"
    }
    ```
  - **400 Bad Request**:
    ```json
    {
      "error": "Invalid params"
    }
    ```
  - **500 Internal Server Error**:
    ```json
    {
      "error": "failed to create codeBlock"
    }
    ```

#### 2. Get a Code Block by ID
- **Method**: GET
- **Path**: `/toc/v1/codeblock/:id`
- **Description**: Retrieves a specific code block by its ID.
- **Path Parameters**:
  - `id`: The unique identifier of the code block.
- **Responses**:
  - **200 OK**:
    ```json
    {
      "_id": "123",
      "author_id": 123,
      "title": "Sample Contract",
      "description": "A sample smart contract",
      "rating": 0,
      "lang": "func",
      "body": "code content here"
    }
    ```
  - **400 Bad Request**:
    ```json
    {
      "error": "invalid id"
    }
    ```
  - **500 Internal Server Error**:
    ```json
    {
      "error": "failed to get codeBlock"
    }
    ```
#### 3. List All Code Blocks
- **Method**: GET
- **Path**: `/toc/v1/codeblock/list`
- **Description**: Retrieves a list of all code blocks. Supports optional query parameters for filtering by author name, programming language, and sorting by rating.
- **Query Parameters** (all optional):
  - `author_name` (string): Filters code blocks by the author's name.
  - `lang` (string): Filters code blocks by programming language (e.g., `func`, `fift`, `tact`, `tolk`).
  - `sort_order` (string): Specifies the sort order by rating. Valid values: `asc` (ascending), `desc` (descending). Default: `desc`.
- **Responses**:
  - **200 OK**:
    ```json
    [
      {
        "_id": "123",
        "author_id": 123,
        "title": "Sample Contract",
        "description": "A sample smart contract",
        "rating": 0,
        "lang": "func",
        "body": "code content here"
      }
    ]
    ```
  - **404 Not Found**:
    ```json
    {
      "error": "code blocks not found"
    }
    ```
  - **500 Internal Server Error**:
    ```json
    {
      "error": "failed to get codeBlocks"
    }
    ```


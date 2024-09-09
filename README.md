# swagen

## What's swagen ??
**swagen** is the CLI tool implememted by Golang that improve your develop experiences when you use **Swagger(OpenAPI)**.  
Using swagen, you can generate API definition files on CLI without coding yourself.

## Install

Use `go install`

```bash
$ go install github.com/Daaaai0809/swagen
```

If you want to use any version

```bash
$ go install github.com/Daaaai0809/swagen@vX.X.X
```

## Usecase

When you use swagen in your project, you have to make env file in project root dir and define `PATH_DIR`, `MODEL_DIR`, `MESSAGE_DIR` params in that file.  

env Example:
```env
PATH_DIR="src/paths"
MODEL_DIR="src/components/models"
MESSAGE_DIR="src/components/schemas"
```

### Path File

If you want to generate Path file, run this command:
```bash
$ swagen path <HTTP METHOD> -d <dir>
```

You can use http methods like below for `HTTP METHOD` parameter.

- GET : `get`
- POST : `post`
- PUT : `put`
- DELETE : `delete`

You can use `--dir` flag instead of `-d` flag also.

```bash
$ swagen path <HTTP METHOD> --dir=<dir>
```

The `dir` in this command is deeper directory than defined dir at `PATH_DIR` param.  

Example:
```bash
$ swagen path get -d "front/users"
```

Definable fields are these:
- **OperationId**
- **Summary**
- **Description**
- **Tags**
- **Security**
    - Select from these Security Types
        - APIKey
        - Bearer
        - BasicAuth
        - OAuth2
        - OpenId
- **Parameters**
    - In
    - Description
    - Required
    - Schema
- **RequestBody**(Only POST, PUT, DELETE)
    - Almost content-types are supported. If you need other content-type, please open issue.
    - Schema Fields
        - Type
        - Format
        - Required
        - Nullable
        - Properties(Type Object)
        - Items(Type Array)
- **Responses**
    - You can use all response codes and select from those
    - **Default** response is supported also
    - Almost content-types are supported. If you need other content-type, please open issue.

### Request/Response Schema
If you want to generate Request/Response Schema file, run this command.
```bash
$ swagen message <FileName> -d <dir>
```

Example schema:
```yaml
TestResponse:
  type: object
  properties:
    id:
      type: integer
      format: int64
    name:
      type: string
    email:
      type: string
      format: email
      nullable: true
    tags:
      type: array
      items:
        type: object
        properties:
          id:
            type: integer
            format: int64
          name:
            type: string
        required:
        - id
        - name
  required:
  - id
  - name
  - email
  - tags
```

### Model Schema (WIP)
If you want to generate Model Schema (Data Object Definition) file, run this command.
```bash
$ swagen message <FileName> -d <dir>
```

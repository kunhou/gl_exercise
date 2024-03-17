# Spec

## Fields of a task:
- **name**: String
- **status**: Boolean
  - 0: Incomplete
  - 1: Complete

## Response headers
- Content-Type: application/json

# Unit Test
- Ensure to include unit tests for the implemented functionalities.

## Managing Codebase on GitHub
- Please manage the codebase on GitHub for collaboration and version control.

## Runtime Environment Requirement
- If you choose Python:
  - Python 3.7+
  - Flask 2.0.x
- If you choose Golang:
  - Go 1.17.8+
  - Gin 1.7.7

## Docker
- The application should be containerized using Docker.

## About Database
- DB is a complex component, but for this exercise, you can use an in-memory mechanism to handle data storage problems.

## API Endpoints

1. **GET /tasks** (List tasks)
   - Response:
     ```json
     {
       "result": [
         {"id": 1, "name": "name", "status": 0}
       ]
     }
     ```

2. **POST /task** (Create task)
   - Request:
     ```json
     {
       "name": "Buy dinner"
     }
     ```
   - Response Status Code: 201
     ```json
     {
       "result": {"name": "Buy dinner", "status": 0, "id": 1}
     }
     ```

3. **PUT /task/<id>** (Update task)
   - Request:
     ```json
     {
       "name": "Buy breakfast",
       "status": 1,
       "id": 1
     }
     ```
   - Response Status Code: 200
     ```json
     {
       "result":{
         "name": "Buy breakfast",
         "status": 1,
         "id": 1
       }
     }
     ```

4. **DELETE /task/<id>** (Delete task)
   - Response Status Code: 200
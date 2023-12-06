# #4 Hands On

Let's Challenge Ourselves!

Make a CRUD API that consists of these endpoints:
- GET /item
  - Will fetch all ```items``` data
  - Implement caching mechanism ```10s``` TTL
  - Sample response body:
    ```
    [
      {
          "id": 1,
          "name": "alfa"
      },
      {
          "id": 2,
          "name": "alfa"
      },
      {
          "id": 3,
          "name": "alfa"
      }
    ]
    ```
- GET /item/:id
  - Will fetch one ```item``` based on parameter's ```id```
  - Implement caching mechanism with ```10s``` TTL
  - Sample response body:
    ```
    {
        "id": 1,
        "name": "alfa"
    }
    ```
- POST /item
  - Will save one ```item``` data on DB
  - Invalidate related caches after save data on DB
  - Sample request body:
    ```
    {
        "name": "alfa"
    }
    ```
  - Sample response body:
    ```
    {
        "message": "SUCCESS"
    }
    ```
- PUT /item/:id
  - Will update one ```item``` data based on parameter's ```id``` on DB
  - Invalidate related caches after save data on DB
  - Sample request body:
    ```
    {
        "name": "alfa"
    }
    ```
  - Sample response body:
    ```
    {
        "message": "SUCCESS"
    }
    ```
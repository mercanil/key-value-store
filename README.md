# key-value-store

- Able to store arbitrary key-value pairs.
- Provide service endpoints that allow a user to put, get, and delete key-value pairs.
- Persistently store its data in some fashion

Functionality Method Possible statuses
Put a key-value pair into the store

| Method | Response Code             | Description                          |
|--------|---------------------------|--------------------------------------|
| PUT    | 201 (Created)             | Read a key-value pair from the store |
| GET    | 200 (OK), 404 (Not Found) | Delete a key-value pair              |
| DELETE | 200 (OK)                  |                                      |

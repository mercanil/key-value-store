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



Transaction Log file


Pros:

- No downstream dependency: There’s no dependency on an external service that could fail or that we can lose access to.
- Technically straightforward: The logic isn’t especially sophisticated. We can be up and running quickly.

Cons:

- Harder to scale:  You’ll need some additional way to distribute your state between nodes when you want to scale out.
- Unconstrained growth: These logs have to be stored on disk, so you can’t let them grow forever.
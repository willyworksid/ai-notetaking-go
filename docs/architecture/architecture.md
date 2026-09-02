# Architecture

Arsitektur akan dibangun bertahap selama proses learning.

```mermaid
flowchart TD
    Client --> Fiber
    Fiber --> Handler
    Handler --> Service
    Service --> Repository
    Repository --> PostgreSQL
```

Async:

```mermaid
flowchart TD
    API --> Event
    Event --> Queue
    Queue --> Worker
    Worker --> Embedding
    Worker --> AI
```

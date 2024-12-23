# gopherMart



## Структура таблиц БД
```mermaid
---
title: 
---
erDiagram
    Users{
        int ID
        string name
        string password
    }
    Orders{
        int ID
        float sum
    }
    Balls{
        int ID
        int orderID
        float accural
    }

    Entity1 ||--o{ Entity2 : places
    
    Entity2 ||--|{ Entity3 : contains

    Entity1 }|..|{ Entity4 : uses
```

## Диаграмма запуска приложения
```mermaid
flowchart LR
App --> |Create| Config
App --> |Run & give config| Server
Server --> |Create| mainHandler
Server --> |Create| chi.NewRoute
mainHandler --> |Register| chi.NewRoute
```

## Диаграмма пакетов
```mermaid
---
title: Глобальная
---
flowchart LR
App --> |Create| Config
App --> |Run| Server
Server --> |Create| IndexHandler
Server --> |Create| chi.NewRoute
 IndexHandler --> |Register| chi.NewRoute
```

```mermaid
---
title: Подробная
---
flowchart LR
App --> |Create| Config
App --> |Run| Server
Server --> |Create| IndexHandler
Server --> |Create| chi.NewRoute
 IndexHandler --> |Register| chi.NewRoute
```


## Диаграмма классов
```mermaid 
---
title: URL shotrener
---
classDiagram

    class Config{
        +ServerAddress string
        +BaseURL string
    }

    class Alias{
        -url string
        -alias string

        +constructor(url string, alias string) *Alias
        +GetURL() string
        +GetAlias() string

    }

    class Repository{
        +GetByAlias(alias string) *Alias, error
        +GetByURL(url string) *Alias, error
        +Save(model Alias) error
    }
    class inMemoryRepository{     
        -aliases []Alias

        +GetByAlias(alias string) *Alias, error
        +GetByURL(url string) *Alias, error
        +Save(model Alias) error
    }
    class Service{
        -repo Repository
        -lengthAlias int
    }

    class indexHandler{
        -config *config.Config
        -service *service.Service

        +HandleGET(rw http.ResponseWriter, r *http.Request)
        +HandlePOST(rw http.ResponseWriter, r *http.Request)
    }

    indexHandler o-- Config
    indexHandler o-- Service
    Service --* Repository
    Repository <|-- inMemoryRepository


    Repository --> Alias
```
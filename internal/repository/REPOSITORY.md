# Схема базы данных

``` mermaid
---
title: gopherMart
---
erDiagram
    Users {
        int id
        string login
        string password
        float balance
    }
    Orders {
        int id
        int user_id
        string order
        date uploaded_at
    }
    OrdersAPI{
        string order
        string status
        float accural
    }
    Withdrawals {
        int id
        int user_id
        float sum
        date processed_at
    }
    Users ||--|| Orders : Containts
    Orders ||--|| OrdersAPI : Containts
    Withdrawals ||--|| Users : Containts
```

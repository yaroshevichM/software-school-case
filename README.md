
# Software engineering school case (Golang)

The application provides a REST API, the description of which is located below in the repository, and also implements regular sending of emails to all subscribers once a day using cron jobs. In the application, each of the layers of the architecture was separated (separately, work with REST, business logic, and work with data storage). The application was developed in compliance with SOLID principles.

It is worth noting that there is probably a better way to implement sending messages to mail, using goroutines. But unfortunately, since this is my first project on GO, I am not sure that it was implemented correctly, but I tried to do it.

The .env file is intentionally present in the repository in order to facilitate testing and launching the project after cloning it.





## Features

- Get USD to UAH rate

- Subscribe for daily mails with USD/UAH rates

- Send mails to all subscribers with USD/UAH rate


## Tech Stack

**Used in project:** Golang, Gin, Posgresql, Docker


## Running the project

For start project execute next command from the root of the project:

```bash
  docker-compose -f docker-compose.dev.yml up --build
```
    
## API Reference

### Get USD/UAH Rate

```http
  GET /rate
```

#### Responses
| Status | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `200` | `39.5` | Success response |
| `409` | `{message: string}` | Error response |

Allow you recieve USD to UAH rate from privatbank api.

### Subscribe for rate notification

```http
  POST /subscribe
```
#### Payload
| Filedname | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | Email field |

#### Responses
| Status | Value     | Description                |
| :-------- | :------- | :------------------------- |
| `200` | `{id: int}` | Success response |
| `409` | `{message: string}` | Error response |


Allow you recieve USD to UAH rate from privatbank api.


## Authors

- Yaroshevich Maxim


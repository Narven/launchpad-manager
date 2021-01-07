# launchpad manager

> Let's compete against SpaceX

### development

* Copy `.env.example` to `.env` and make the necessary changes
* `docker-compose up -d` will bring the containers up

Run: `make`

## Migrations

```bash
migrate create -ext sql -dir db/migrations -seq <filename.sql>
```

### API

#### **POST** - /v1/tickets
##### Description
Book a ticket
```sh
curl -X "POST" "http://localhost:5000/v1/tickets" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "gender": "male",
  "launch_date": "2021-02-03T19:10:000Z",
  "destination_id": 3,
  "last_name": "Luz",
  "birthday": "01-02-2006",
  "launchpad_id": "vafb_slc_4e",
  "first_name": "Pedro"
}'
```

#### **GET** - /v1/tickets
##### Description
Get all Tickets
```sh
curl -X GET "http://localhost:5000/v1/tickets"
```

_by SpaceTrouble Inc._
# launchpad manager

> Let's compete against SpaceX

### development

* Copy `.env.example` to `.env` and make the necessary changes
* `docker-compose up -d` will bring the containers up

Run: `make`

NOTES:

Unfortunately you both share the same launchpad's
and you cannot launch your rockets from the same place on the same day.
There is a list of available launchpad's and your spaceships go to places like: Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede.

Every day you change the destination for all the launchpads.
Basically on every day of the week from the same launchpad has to be a "flight" to a different place.

# Request (Book Ticket)
[POST] /tickets
You have to verify if the requested trip is possible on the day from provided launchpad ID and do not overlap 
with SpaceX launches or launches already booked on your system, if that’s the case then your flight is cancelled.

# Endpoint to get all created Bookings.
[GET] /bookings

## Migrations

```bash
migrate create -ext sql -dir db/migrations -seq <filename.sql>
```

### API

Create Ticket
```bash
curl -X "POST" "http://localhost:5000/v1/tickets" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "gender": "male",
  "launch_date": "03-02-2021",
  "destination_id": 2,
  "last_name": "Doe",
  "birthday": "01-02-2006",
  "launchpad_id": "vafb_slc_4e",
  "first_name": "John"
}'
```


by SpaceTrouble Inc.
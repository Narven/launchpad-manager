# launchpad manager

> Let's compete against SpaceX

### development

* Copy `.env.example` to `.env` and make the necessary changes
* `docker-compose up -d` will bring the containers up

Run: `make`

NOTES:

Unfortunately you both share the same launchpads
and you cannot launch your rockets from the same place on the same day.
There is a list of available launchpads 
and your spaceships go to places like: Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede.

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





by SpaceTrouble Inc.
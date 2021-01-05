CREATE TABLE "public"."ticket"
(
    "id"             serial PRIMARY KEY,
    "first_name"     Character Varying NOT NULL,
    "last_name"      Character Varying NOT NULL,
    "gender"         Character Varying NOT NULL,
    "birthday"       Timestamp Without Time Zone NOT NULL,
    "launch_date"    Timestamp Without Time Zone NOT NULL,
    CONSTRAINT "unique_tickets_id" UNIQUE ("id")
);
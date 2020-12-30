CREATE TABLE "public"."ticket"
(
    "id" UUid NOT NULL,
    "first_name"  Character Varying NOT NULL,
    "last_name"   Character Varying NOT NULL,
    "gender"      SmallInt          NOT NULL,
    "birthday"    Timestamp Without Time Zone NOT NULL,
    "launchpad_id" UUid NOT NULL,
    "destination_id" UUid NOT NULL,
    "launch_date" Timestamp Without Time Zone NOT NULL,
    CONSTRAINT "unique_tickets_id" UNIQUE ("id")
);
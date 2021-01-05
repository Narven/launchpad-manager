ALTER TABLE "public"."ticket"
    ADD COLUMN "destination_id" Integer NOT NULL;

ALTER TABLE "public"."ticket"
    ADD CONSTRAINT fk_ticket_destination
        FOREIGN KEY (destination_id)
            REFERENCES destination (id);
CREATE TABLE "sequences" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "open_tracking" bool,
  "click_trancking" bool
);

CREATE TABLE "sequence_steps" (
  "id" bigserial PRIMARY KEY,
  "sequence_id" bigint NOT NULL,
  "subject" varchar,
  "content" varchar,
  "step_index" integer NOT NULL,
  "wait_days" integer
);

ALTER TABLE "sequence_steps" ADD FOREIGN KEY ("sequence_id") REFERENCES "sequences" ("id");

ALTER TABLE "sequence_steps" ADD UNIQUE ("sequence_id", "step_index")

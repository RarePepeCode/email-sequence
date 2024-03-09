
CREATE TABLE "sequences" (
  "id" integer PRIMARY KEY,
  "name" varchar,
  "open_tracking" bool,
  "click_trancking" bool
);

CREATE TABLE "sequence_steps" (
  "id" integer PRIMARY KEY,
  "sequence_id" integer NOT NULL,
  "subject" varchar,
  "content" varchar,
  "step_index" integer NOT NULL
);

ALTER TABLE "sequence_step" ADD FOREIGN KEY ("sequence_id") REFERENCES "sequence" ("id");

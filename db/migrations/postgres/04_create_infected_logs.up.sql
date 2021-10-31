create table IF NOT EXISTS tbl_infected_logs (
    id bigserial PRIMARY KEY,
    "userId" int not null,
    "infected" smallint not null,
    "reportedBy" int not null,
    "createdAt" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tbl_infected_logs" ADD FOREIGN KEY ("userId") REFERENCES "tbl_survivor" ("id");
ALTER TABLE "tbl_infected_logs" ADD FOREIGN KEY ("reportedBy") REFERENCES "tbl_survivor" ("id");
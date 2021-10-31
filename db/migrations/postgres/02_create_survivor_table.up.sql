create table IF NOT EXISTS tbl_survivor (
    id bigserial PRIMARY KEY,
    name varchar(255) default '' not null,
    age int,
    gender varchar(50),
    "infectionStatus" int not null default 1,
    lastLocation point,
    createdAt timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tbl_survivor" ADD FOREIGN KEY ("infectionStatus") REFERENCES "tbl_infection_description" ("id");
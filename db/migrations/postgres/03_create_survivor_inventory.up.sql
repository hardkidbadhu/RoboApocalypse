create table IF NOT EXISTS tbl_survivor_inventory (
    id bigserial PRIMARY KEY,
    "water" int not null,
    "food" int not null,
    "medication" int not null,
    "ammunation" int not null,
    "userId" int not null,
    "createdAt" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tbl_survivor_inventory" ADD FOREIGN KEY ("userId") REFERENCES "tbl_survivor" ("id");
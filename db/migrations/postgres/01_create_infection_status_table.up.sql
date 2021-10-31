create table IF NOT EXISTS tbl_infection_description (
    id bigserial PRIMARY KEY,
    description int not null);

INSERT INTO tbl_infection_description(id, description) values (1, 1);
INSERT INTO tbl_infection_description(id, description) values (2, 2);
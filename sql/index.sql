alter table votes add index(candidate_id);
alter table users drop index mynumber;
alter table users add index multiple_idx (name, address, mynumber);

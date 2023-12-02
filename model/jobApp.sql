CREATE TABLE job(
  id varchar PRIMARY KEY,
  title varchar,
  experience_in_months integer,
  description text,
  created_at timestamp,
  expires_at timestamp,
  posted_user integer,
  comapany_id integer
);

CREATE TABLE users (
  id integer PRIMARY KEY,
  username varchar,
  password varchar,
  created_at timestamp
);

CREATE TABLE user_role (
  id integer PRIMARY KEY,
  user_id integer,
  role_id integer
);

CREATE TABLE role (
  user_role_id integer,
  name varchar PRIMARY KEY
);

CREATE TABLE application (
  id integer PRIMARY KEY,
  user_id integer,
  job_id integer
);

CREATE TABLE review (
  rating integer,
  company_id integer
);

CREATE TABLE company (
  id integer PRIMARY KEY,
  name varchar,
  about text,
  photos file,
  rating int
);

ALTER TABLE review ADD FOREIGN KEY (company_id) REFERENCES company (id);

ALTER TABLE company ADD FOREIGN KEY (id) REFERENCES job (company_id);

ALTER TABLE application ADD FOREIGN KEY (job_id) REFERENCES users (id);

ALTER TABLE job ADD FOREIGN KEY (id) REFERENCES application (job_id);

ALTER TABLE user_role ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE user_role ADD FOREIGN KEY (role_id) REFERENCES role (user_role_id);

ALTER TABLE job ADD FOREIGN KEY (posted_user) REFERENCES users (id);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
     id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
     login text UNIQUE NOT NULL ,
     passhash text NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id uuid NOT NULL REFERENCES users(id),
  refresh_token text NOT NULL,
  expiration_time bigint
);
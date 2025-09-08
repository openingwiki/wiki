create table if not exists anime (
  id bigserial primary key,
  title text not null,
  created_at timestamptz not null default now()
);

create table if not exists singers (
  id bigserial primary key,
  name text not null,
  created_at timestamptz not null default now()
);

create table if not exists openings (
  id bigserial primary key,
  anime_id bigint not null references anime(id) on delete cascade,
  singer_id bigint not null references singers(id) on delete set null,
  type text not null check (type in ('opening','ending','ost')),
  title text not null,
  order_number int not null default 0,
  created_at timestamptz not null default now()
);



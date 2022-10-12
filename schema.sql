create table site (
    id               uuid   primary key,
    name             text   unique,
    display_name     text,
    discord_server   text
);

create table article (
    id        uuid   primary key,
    site_id   uuid   references site on delete cascade,
    slug      text,
    title     text,
    source    text,
    unique(site_id, slug)
);
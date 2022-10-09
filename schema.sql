create table site (
    id               uuid   primary key,
    name             text   unique not null,
    display_name     text   not null,
    discord_server   text
);

create table article (
    id        uuid   primary key,
    site_id   uuid   references site on delete cascade not null,
    slug      text   not null,
    title     text   not null,
    source    text   not null,
    unique(site_id, slug)
);
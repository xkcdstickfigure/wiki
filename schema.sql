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

create table discord_user (
    id               text        primary key,
    username         text,
    discriminator    text,
    avatar           text,
    mfa_enabled      boolean,
    locale           text,
    email            text,
    email_verified   boolean,
    access_token     text,
    refresh_token    text,
    first_auth_at    timestamp,
    last_auth_at     timestamp
);

create table discord_guild (
    id     text   primary key,
    name   text,
    icon   text
);

create table discord_member (
    user_id    text   references discord_user on delete cascade,
    guild_id   text   references discord_guild on delete cascade,
    primary key(user_id, guild_id)
);
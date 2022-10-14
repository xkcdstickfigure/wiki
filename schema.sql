create table site (
    id              uuid   primary key,
    name            text   unique,
    display_name    text,
    discord_guild   text
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
    id               text          primary key,
    username         text,
    discriminator    text,
    avatar           text,
    mfa_enabled      boolean,
    locale           text,
    email            text,
    email_verified   boolean,
    access_token     text,
    refresh_token    text,
    first_auth_at    timestamptz,
    last_auth_at     timestamptz
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

create table session (
    id           uuid          primary key,
    token        text          unique,
    address      text,
    user_agent   text,
    discord_id   text          references discord_user,
    created_at   timestamptz
);

create table discord_state (
    id           uuid          primary key,
    session_id   uuid          references session on delete cascade,
    token        text          unique,
    site         text,
    created_at   timestamptz
);

create table article_view (
    id           uuid          primary key,
    session_id   uuid          references session on delete cascade,
    article_id   uuid          references article on delete cascade,
    date         timestamptz
);
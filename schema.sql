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

create table account (
    id               uuid          primary key,
    number           bigint        unique,
    google_id        text          unique,
    discord_id       text          references discord_user on delete set null,
    name             text,
    email            text,
    email_verified   boolean,
    avatar           text,
    created_at       timestamptz
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
    account_id   uuid          references account on delete set null,
    discord_id   text          references discord_user on delete set null,
    created_at   timestamptz
);

create table discord_state (
    id           uuid          primary key,
    session_id   uuid          references session on delete cascade,
    token        text          unique,
    action       text,
    value        text,
    created_at   timestamptz
);

create table auth_state (
    id           uuid          primary key,
    session_id   uuid          references session on delete cascade,
    token        text          unique,
    redirect     text,
    created_at   timestamptz
);

create table article_view (
    id           uuid          primary key,
    session_id   uuid          references session on delete cascade,
    article_id   uuid          references article on delete cascade,
    created_at   timestamptz
);

create table gitea_code (
    id           uuid          primary key,
    account_id   uuid          references account on delete cascade,
    code         text          unique,
    created_at   timestamptz
);

create table gitea_token (
    id           uuid          primary key,
    account_id   uuid          references account on delete cascade,
    token        text          unique,
    created_at   timestamptz
);
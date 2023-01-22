# wiki
This is a service I made for creating wikis for games and tv shows, like [minecraft.fandom.com](https://minecraft.fandom.com) and [terraria.fandom.com](https://terraria.fandom.com).

The service uses Google OAuth for login, and links your discord account when you join the server for a site. It pulled the article content from a [Gitea](https://gitea.io) instance, and there is a /gitea api that emulates a Gitea OAuth provider, so the real Gitea instance would act as a client to it, and with some configuration, an account would be automatically created when the user tries to view or edit the article source.

The frontend uses Go templating, with Tailwind CSS. The `hub` directory contains the code for the site on the apex domain, while the `site` directory contains code for the subdomains on which the actual wikis are served. The `markup` package parses the article source, and the `render` package turns that into html.

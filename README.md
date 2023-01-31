# wiki
This is a service I made for creating wikis for games and tv shows, like [minecraft.fandom.com](https://minecraft.fandom.com) and [terraria.fandom.com](https://terraria.fandom.com).

The service uses Google OAuth for login, and links your discord account when you join the server for a site. It pulled the article content from a [Gitea](https://gitea.io) instance, and there is a /gitea api that emulates a Gitea OAuth provider, so the real Gitea instance would act as a client to it, and with some configuration, an account would be automatically created when the user tries to view or edit the article source.

The frontend uses Go templating, with Tailwind CSS. The `hub` directory contains the code for the site on the apex domain, while the `site` directory contains code for the subdomains on which the actual wikis are served. The `markup` package parses the article source, and the `render` package turns that into html.

*These screenshots are just random images I found from early development*

![screenshot 1](https://user-images.githubusercontent.com/97917457/215725542-bb4a9290-6ac6-4ed3-9300-d7fcd83d0ab2.png)
![screenshot 2](https://user-images.githubusercontent.com/97917457/215726138-25d25fb6-4bf3-43b0-95c5-ddfca8b7055a.png)

# THETA Stack

This repository serves as a template for kick-starting web apps using a simple stack that I found I like.

- **T**urso: Dead simple edge DB with sqlite
- **H**tmx: All my homies HATEOAS
- **E**cho: Fantastically simple Golang web framework
- **T**ailwind: Because I'm too lazy to write CSS myself
- **A**lpineJS: For writing even less JS myself

## Getting Started

Theta uses [Air](https://github.com/cosmtrek/air) for live-reloading of the Go app during development. The included `.air.toml` file is set up to enable hot-reloading (with browser refresh) during development. This setup is slightly janky and relies on keeping an SSE channel open, then refreshing when it has an error. But it works for local dev. Air will call `npx tailwindcss` before each build to keep your TW output CSS updated.

To get started, install Tailwind with npm/yarn/pnpm/bun/whatever JS package manager you like this week.

```bash
npm install
```

Then start the dev server in a separate terminal.

```bash
air
```

You should see some typical output from Air and Echo and the server should be up and running (default port 8080).

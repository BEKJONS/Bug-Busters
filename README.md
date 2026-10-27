# Bug-Busters
Hackathon 25.04.24

## Setup

Configuration is read from a local `.env` file, which is not committed.

```sh
cp .env.example .env   # then set DB_PASSWORD and SIGNING_KEY
```

The `makefile` migration targets read `DB_PASSWORD` from the environment, e.g. `DB_PASSWORD=yourpass make mig-up`.

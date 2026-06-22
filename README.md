# Vempala Kids

Growing Responsibility Through Rewards

## Deployment

This app runs on a single Go service with SQLite persistence. By default it stores the database at `data/vempala.db`.

For Render, use the included [render.yaml](render.yaml) so the app gets a persistent disk mounted at `/var/data`. Render will set `PORT` automatically, and the app reads `DB_PATH` from the environment.

### Render checklist

1. Create a new Web Service from this repository.
2. Let Render use `render.yaml` so it provisions the disk and environment variables.
3. Confirm the disk is mounted at `/var/data`.
4. Deploy and open the `/health` endpoint to verify the app is running.

## Environment

- `PORT`: web server port, defaults to `8080`
- `DB_PATH`: SQLite file path, defaults to `data/vempala.db`
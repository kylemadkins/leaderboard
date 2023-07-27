# Leaderboard Microservice

## Running Locally

1. Download [migrate](https://github.com/golang-migrate/migrate)
2. Copy the `sample.env` file to a regular `.env` file
3. Spin up the containers with `docker-compose up` as usual. This will start Postgres and an API server at port 8000
4. Run the migrations via `migrate -database <database_url> -path db/migrations up`

## Endpoints

```
GET /players
```
Returns a list of all players

```
POST /players
{
  "username": "ellie96"
}
```
Creates a player

## Enhancements

- [ ] Swagger documentation
- [ ] Better validations for a player's username, including a profanity filter
- [ ] Ability to search players
- [ ] Sort leaderboard by score
- [ ] Authentication

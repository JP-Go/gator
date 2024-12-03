# Gator

A Blog aggreGator using go, sqlc and postgres

## Dependencies

Gator requires the following packages to be installed on your system

- PostgreSQL (version equal or greater than 12)
- Go (version equal or greater than 1.16)

## Instalation

To install gator, simply run the following command from a terminal

go install github.com/JP-Go/gator

Now you can run gator in your terminal. If it fails with a message like
"command not found", make sure the $HOME/go/bin/ folder is in your users
PATH

## Running the program

After installing, create a file called `.gatorconfig.json` your user's home folder (`$HOME `on linux, `%userprofile%` on windows) with the following content.

```json
{
  "db_url": "postgres://postgres:@localhost:5432/gator?sslmode=disable"
}
```

The value of db_url should be the connection string to the postgres database of
your choice. Eg: `postgres://postgres:@localhost:5432/gator?sslmode=disable`.

## Commands

- `gator register <username>`: Creates a new user with `username`. Username is
  used to login.
- `gator login <username>`: Logs in as user with `username`.
- `gator users`: Shows a list of the users registered in the application
- `gator addfeed <feedName> <feedUrl>`: adds the feed `feedName` at `feedUrl` to
  the list of feeds aggregated by gator. This is used to fetch the latest
  articles from this feed. Automatically registers the current user as a
  follower of the feed.
- `gator feeds`: Show a list of feeds registered in gator
- `gator follow <feedUrl>`: registers the current user as a follower of the feed with url `<feedUrl>`
- `gator unfollow <feedUrl>`: unregisters the current user as a follower of the feed with url `<feedUrl>`
- `gator following `: Show the feeds that the current user follows
- `gator agg <timeBetweenRequests>`: Runs the aggregation cycle every `timeBetweenRequests` while the command is not stopped. `timeBetweenRequests` must be a time duration string like '1s', '2m','1h', '1m30s', etc. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". (See more at [https://pkg.go.dev/time#ParseDuration](https://pkg.go.dev/time#ParseDuration)).
- `gator browse <n>`: Shows the most recent `<n>` posts. `<n>` is optional and
  defaults to 2 if not passed.

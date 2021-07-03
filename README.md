# Embrio 4

Back-end services to process EMBRIO - 4


## Prequisites
- install Go version 1.13
- install postgresql
- Create a Postgresql
- Import SQL dump (ask me for the sql file)


## Environment Variables

Fill the `.env` files with appropriate value according to your local setup

```bash
DB_USERNAME=
DB_PASSWORD=
DB_CONNECTION=tcp
DB_HOST=
DB_PORT=
DB_NAME=embrio4

PORT=8081
```

## Installation
Before install & run, please make sure port 80 is available

```bash
make start
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


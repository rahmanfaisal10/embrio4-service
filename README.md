# Embrio 4

Back-end services to process EMBRIO - 4


## Prequisites
- install Go version 1.13 or latest
- install mysql
- Create a database embrio4
- Import SQL dump (ask me for the sql file)


## Environment Variables

Fill the `.env` files with appropriate value according to your local setup

```bash
DB_USERNAME=root
DB_PASSWORD=
DB_CONNECTION=tcp
DB_HOST=
DB_PORT=
DB_NAME=embrio4

PORT=

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
EMAIL=
MAILPASSWORD=

SECRET_KEY=embrio4

```

## Installation
Before install & run, please make sure port 80 is available and setup environment with your environment

```bash
make start
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.


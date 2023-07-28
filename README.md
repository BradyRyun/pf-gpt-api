# Personal Finance Plan API


## Requirements
* Go >= 1.20.4


## Config
You will need to update following vars in `.env`
```bash
DB_HOST=127.0.0.1 # DB Host name , use `db` for docker-compose setup
DB_USER=root # Suggest not using root for prod
DB_PASSWORD=password123 # suggest stronger password for prod
DB_NAME=newsletters # Used to configure PG on startup
DB_PORT=5432 # PG Default
TOKEN_HOUR_LIFESPAN=3600 # JWT lifespan in seconds

HOST_NAME=localhost:3000 # Host name of the server, there's probably a better way to grab this from gin context, but this will suffice
SENDGRID_FROM_EMAIL=brady@ryunengineering.com # From email in SendGrid
SENDGRID_API_KEY= # SendGrid API key
```
### Firebase service-account
You can try running through the <strong>very</strong> useful [docs](https://firebase.google.com/docs/firestore/quickstart) from Firebase, or follow here:
1. Navigate to Firebase
2. At the top-left, click the cog and select Project Settings
3. Select the Service accounts tab
4. In the modal, select `Generate new private key`
5. Save that key as `service-account.json` at the root of the project
## Build
```bash
#### Running on local machine
# Linux
go build -o pf-api
./pf-api
# Windows
go build -o pf-api.exe
./pf-api.exe

## Docker
docker build -t pf-api .
docker run -v $PWD/.env:/app/.env -v $PWD/service-account.json:/app/service-account.json pf-api
```
# Backend Canditate Home Test

Have you shopped online? Let’s imagine that you need to build a checkout backend service that will support different promotions with the given inventory.

Build a checkout system with these items:

| SKU    | Name           | Price   | Inventory Qty |
| ---    | ---            | ---:    | ---:         |
| 120P90 | Google Home    | 49.99   | 10            |
| 43N23P | MacBook Pro    | 5399.99 | 5             |
| A304SD | Alexa Speaker  | 109.50  | 10            |
| 234234 | Raspberry Pi B | 20.00   | 2             |

### The system should have the following promotions:
- Each sale of a MacBook Pro comes with a free Raspberry Pi B
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will get a 10% discount on all Alexa speakers

### Example Scenarios:
- Scanned Items: MacBook Pro, Raspberry Pi B<br/>
Total: $5,399.99

- Scanned Items: Google Home, Google Home, Google Home<br/>
Total: $99.98

- Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker<br />
Total: $295.6

## Documentations

- [Database document](database.md)
- [API Coontract](api-contract.md)

## How to run
### 1. Migrate database
- Read [Database document](database.md)

### 2. Using go run
- Set `.env` file like `.env-example`
- Run command:
```
go run main.go -loadDotEnv=true
```

### 3. Build docker file
-  Build docker image
```
docker build -t "$DOCKER_NAME" .
```

- Delete existing container
```
docker container rm "$CONTAINER_NAME"
```

- Create container<br>Don't use localhost for mysql host
```
docker container create --name "$CONTAINER_NAME" -e HTTP_PORT=$HTTP_PORT -e MYSQL_HOST="$MYSQL_HOST" -e MYSQL_USERNAME="$MYSQL_USERNAME" -e MYSQL_DB_NAME="$MYSQL_DB_NAME" -e MYSQL_PASSWORD="$MYSQL_PASSWORD" -p $HTTP_PORT:$HTTP_PORT $DOCKER_NAME
```

- Start container
```
docker container start "$CONTAINER_NAME"
```
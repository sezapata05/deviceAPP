# Golang Api Devices

In case you do not have postgresql installed/configured I recommend you to look at the option of executing through containers.

[Docker Execution](docs/guide/Docker.md)

To run locally you need to set the following variables in the `.env` file

| variable name | Posible Value | Description                             |
| ------------- | ------------- | --------------------------------------- |
| DB_HOST       | localhost     | Your host where postgresql is running   |
| DB_PORT       | 5432          | Your port where postgresql is listening |
| DB_USER       | -             | Your username for postgresql            |
| DB_PASSWORD   | -             | Your password for postgresql            |
| DB_NAME       | dc_device     | The table name for this project         |

## Run it

To run it on your premises, you must follow the following command lines.

```bash
# Install the dependencies:
go get ./...

# run
go run .\cmd\api\main.go
```


The default port is `8080` on which the api will raise its services. To access, you can do it through the following url:

`http://localhost:8080`

And use the registered endpoints.

[Endpoints](Endpoints.md)
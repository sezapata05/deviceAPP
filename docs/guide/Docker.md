# Docker-Compose
This application has a micro-services architecture which implements through docker files (DockerFile and Docker-Compose) the version management of golang and the installation and portability of Postgresql and pgadmin through the use of an internal network connection.

# Execution of Docker-Compose

To execute the `docker-compose.yml` file, follow these steps:

1. Open a terminal or command line.

2. Make sure you are in the directory where your `docker-compose.yml` file is located. You can use the `cd` command to navigate to the directory if necessary.

3. Run the following command to start the containers defined in the `docker-compose.yml` file:


```bash
docker-compose up -d
```

The `-d` flag is used to run the containers in the background. If you prefer to see the container logs in real time on your terminal, you can omit `-d`.

4. Docker Compose will download the necessary images (if you don't already have them) and create and run the containers according to the configuration defined in the `docker-compose.yml` file.


5. Once the containers are running, you can check their status by running the following command:

```bash
docker-compose ps
```

This will show you the status of the containers and the mapped ports.

6. To stop and remove the containers and related resources, you can use the following command:

```bash
docker-compose down
```

This will stop the containers and remove the volumes created by Docker Compose. If you also want to remove the images, you can add the `--rmi all` flag:

```bash
docker-compose down --rmi all
```

That's it. You should now have your PostgreSQL and pgAdmin services up and running. You can access pgAdmin through `http://localhost:5010` and use the following credentials:

* EMAIL: admin@admin.com
* PASSWORD: admin

## GOlANG API
To access the api and its endpoints we recommend you to follow:

[Endpoints technical description](Endpoints.md)
# Collocutor

The only one who will understand you

## Parser

Parser is running in docker.
To run parser in docker, you should install **make** package on your local machine.

### Build container

Go to **./parser** directory, and build container by running command:

```bash
make build-container
```

### Start container

Start container by running:

```bash
make start-container
```

### Run scrapper

To run scrapper execute following command:

```bash
make scrap
```

### Copy result file to your local machine

To copy result file from docker container run this command:

```bash
make copy
```

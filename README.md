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

### Build executable file

```bash
make build-executable
```

### Run scrapper

There are a few web pages we could parse from.
Each web page has it's own command.

#### Parse from wiki quotes

```bash
make scrap-wiki-quotes
```

### Copy result file to your local machine

To copy result file from docker container run this command:

```bash
make copy
```

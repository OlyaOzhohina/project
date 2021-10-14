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

#### build executable git bash

```bash
make build-executable-bash
```

### Run scrapper

There are a few web pages we could parse from.
Each web page has it's own command.

#### Parse from wiki quotes

#### parse on unix

```bash
make scrap-wiki-quotes
```

#### parse git bash

```bash
make scrap-wiki-quotes-bash
```

### Copy result file to your local machine

To copy result file from docker container run this command:

#### Unix

```bash
make copy
```

#### Windows

```bash
make copy-windows
```

# forum

## DOCKER;

Building your project:

```bash
docker build -t forum-dockerize:latest .

```

Showing images:
```bash

docker images
```

To run:

```bash
docker run --name forum-dockerize -p 8080:8080 forum-dockerize:latest
```

To remove:
```bash
docker rm forum-dockerize

```

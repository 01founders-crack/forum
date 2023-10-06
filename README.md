# forum

## DOCKER;

Building your project:

if you are using linux;
```bash
sudo systemctl start docker
#and dont forgot to stop when you finish
sudo systemctl stop docker
```

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

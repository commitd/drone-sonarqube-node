# drone-sonarqube-node

[![Build Status](https://drone.committed.software/api/badges/commitd/drone-sonarqube-node/status.svg)](https://drone.committed.software/commitd/drone-sonarqube-node) [![Docker Pulls](https://img.shields.io/docker/pulls/committed/drone-sonarqube-node.svg?style=flat)](https://hub.docker.com/r/committed/drone-sonarqube-node)

A plugin for Drone CI to run SonarQube analysis on node, JavaScript and TypeScript projects.

## Use

The simplest use in a drone pipeline is shown below, you just have to provide secrets for the SonarQube host (`sonar-host`) and login token (`sonar-token`).

```yaml
- name: code-analysis
  image: committed/drone-sonarqube-node
  settings:
    host:
      from_secret: sonar_host
    token:
      from_secret: sonar_token
```

For full parameter details see [DOCS.md](DOCS.md).

## Development

Dependencies are managed with `dep` and stored in the vendor directory

To build the go binary file:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonarqube-node
```

To build the docker image:

```bash
docker build -t committed/drone-sonarqube-node .
```

Both steps are performed in `./build.sh`

### Testing

To test the docker image run the docker image with appropriate properties:

```docker
docker run --rm \
  -e DRONE_REPO=test \
  -e PLUGIN_SOURCES=. \
  -e SONAR_HOST=http://localhost:9000 \
  -e SONAR_TOKEN=60878847cea1a31d817f0deee3daa7868c431433 \
  committed/drone-sonarqube-node
```

## Credits

This is based on the original work by [aosapps](https://github.com/aosapps/drone-sonar-plugin) and the node modified version by [anthonyjesmok](https://github.com/anthonyjesmok/drone-sonar-plugin).

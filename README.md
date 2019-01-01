# drone-sonar-plugin
The plugin of Drone CI to integrate with SonarQube (previously called Sonar), which is an open source code quality management platform.

Detail tutorials: [DOCS.md](DOCS.md).

### Purpose of This Fork
The purpose of this fork is to meet some of the needs I have for the drone-sonar-plugin, primarily including Node.js for the purpose of being able to analyze JavaScript, CSS, and other web application files that are needed by sonar-scanner.

My intention is to continue pulling in changes as I see them from the [original repo](https://github.com/aosapps/drone-sonar-plugin).

### Build process
build go binary file:
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar`

build docker image
`docker build -t aosapps/drone-sonar-plugin .`


### Testing the docker image:
```commandline
docker run --rm \
  -e DRONE_REPO=test \
  -e PLUGIN_SOURCES=. \
  -e SONAR_HOST=http://localhost:9000 \
  -e SONAR_TOKEN=60878847cea1a31d817f0deee3daa7868c431433 \
  aosapps/drone-sonar-plugin
```

### Pipeline example
```yaml
  code-analysis:
    image: aosapps/drone-sonar-plugin
    secrets: [sonar_host, sonar_token]
```


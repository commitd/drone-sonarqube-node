# drone-sonar-plugin
The plugin of Drone CI to integrate with SonarQube (previously called Sonar), which is an open source code quality management platform.

Detail tutorials: [DOCS.md](DOCS.md).

### Purpose of This Fork
The purpose of this fork is to meet some of the needs I have for the drone-sonar-plugin, primarily including Node.js for the purpose of being able to analyze JavaScript, CSS, and other web application files that are needed by sonar-scanner.

My intention is to continue pulling in changes as I see them from the [original repo](https://github.com/aosapps/drone-sonar-plugin).

Here are a list of the changes:

1. Node.js comes packaged with the scanner for usage with web application files. This comes from a fork of the OpenJDK Docker image that also installs Node.js.
2. More variables can be specified. This is important if you want your project to work on a cloud service like SonarCloud.
3. It uses the CLI library from urfave, who took over the project from codegangsta. This means `go get` runs with ease.

### Build process
build go binary file:
`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar`

build docker image
`docker build -t anthonyjesmok/drone-sonar-plugin .`


### Testing the docker image:
```commandline
docker run --rm \
  -e DRONE_REPO=test \
  -e PLUGIN_SOURCES=. \
  -e SONAR_HOST=http://localhost:9000 \
  -e SONAR_TOKEN=60878847cea1a31d817f0deee3daa7868c431433 \
  anthonyjesmok/drone-sonar-plugin
```

### Pipeline example
```yaml
  code-analysis:
    image: anthonyjesmok/drone-sonar-plugin
    secrets: [sonar_host, sonar_token]
```


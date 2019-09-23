# Docs

This plugin can scan your node code and post the analysis report to your SonarQube server.

The below pipeline configuration demonstrates simple usage:

```yaml
- name: code-analysis
  image: committed/drone-sonarqube-node
  settings:
    host:
      from_secret: sonar_host
    token:
      from_secret: sonar_token
```

Customized parameters can be specified as settings or via environment variables:

```yaml
  - name: code-analysis
    image: committed/drone-sonarqube-node
    settings:
      host:
        from_secret: sonar_host
      token:
        from_secret: sonar_token
      ver: 1.0
      timeout: 20
      sources: .
      level: DEBUG
      showProfiling: true
      exclusions: **/static/**/*,**/dist/**/*.js
```

# Secret Reference

Safety first, the host and token are stored in Drone Secrets.

| setting              | environment          | description                                                                                                            |
| -------------------- | -------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `sonar_host`         | `SONAR_HOST`         | Host of SonarQube with schema (http/https).                                                                            |
| `sonar_token`        | `SONAR_TOKEN`        | User token used to post the analysis report to SonarQube Server. Click User > My Account > Security > Generate Tokens. |
| `sonar_organization` | `SONAR_ORGANIZATION` | (optional) parameter for use in hosted version such as SonarCloud                                                      |

# Parameter Reference

| setting         | environment            | default              | description                                                                                                                                           |
| --------------- | ---------------------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ver`           | `DRONE_BUILD_NUMBER`   | `DRONE_BUILD_NUMBER` | Code version                                                                                                                                          |
| `timeout`       | `PLUGIN_TIMEOUT`       | 60                   | timeout in seconds                                                                                                                                    |
| `sources`       | `PLUGIN_SOURCES`       | -                    | Comma-separated paths to directories containing source files.                                                                                         |
| `inclusions`    | `PLUGIN_INCLUSIONS`    | -                    | Comma-delimited list of file path patterns to be included in analysis. When set, only files matching the paths set here will be included in analysis. |
| `exclusions`    | `PLUGIN_EXCLUSIONS`    | -                    | Comma-delimited list of file path patterns to be excluded from analysis. Example: `**/static/**/*,**/dist/**/*.js`.                                   |
| `level`         | `PLUGIN_LEVEL`         | `INFO`               | Control the quantity / level of logs produced during an analysis. INFO, DEBUG, TRACE                                                                  |
| `showProfiling` | `PLUGIN_SHOWPROFILING` | `false`              | Display logs to see where the analyzer spends time.                                                                                                   |
| `projectKey`    | -                      | `DRONE_REPO`         | The sonar project key                                                                                                                                 |
| `projectName`   | -                      | `DRONE_REPO`         | The sonar project name                                                                                                                                |

You could also add a file named `sonar-project.properties` at the root of your project to specify parameters.
SonarQube Parameters: [Analysis Parameters](https://docs.sonarqube.org/display/SONAR/Analysis+Parameters)

# Troubleshooting - Test your SonarQube Server:

Replace the parameter values with your own：

```commandline
sonar-scanner \
  -Dsonar.projectKey=MyProject:test \
  -Dsonar.sources=. \
  -Dsonar.projectName=MyProject/test \
  -Dsonar.projectVersion=1.0 \
  -Dsonar.host.url=http://mysonarqube:9000 \
  -Dsonar.login=1234567890987654321234567890987654321234
```

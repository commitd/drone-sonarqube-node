# Docs

This plugin can scan your node code and post the analysis report to your SonarQube server.

The below pipeline configuration demonstrates simple usage:

```yaml
  code-analysis:
    image: committed/drone-sonarqube-node
    secrets: [sonar_host, sonar_token]
```

Customized parameters could be specified:

```diff
  code-analysis:
    image: committed/drone-sonarqube-node
+   secrets: [sonar_host, sonar_token, sonar_organization]
+   ver: 1.0
+   timeout: 20
+   sources: .
+   level: DEBUG
+   showProfiling: true
+   exclusions: **/static/**/*,**/dist/**/*.js
```

# Secret Reference

Safety first, the host and token are stored in Drone Secrets.

* `sonar_host`: Host of SonarQube with schema (http/https).
* `sonar_token`: User token used to post the analysis report to SonarQube Server. Click User > My Account > Security > Generate Tokens.
* `sonar_organization`: (optional) parameter for use in hosted version such as SonarCloud

# Parameter Reference

* `ver`: Code version, Default value `DRONE_BUILD_NUMBER`.
* `timeout`: Default seconds `60`.
* `sources`: Comma-separated paths to directories containing source files. 
* `inclusions`: Comma-delimited list of file path patterns to be included in analysis. When set, only files matching the paths set here will be included in analysis.
* `exclusions`: Comma-delimited list of file path patterns to be excluded from analysis. Example: `**/static/**/*,**/dist/**/*.js`.
* `level`: Control the quantity / level of logs produced during an analysis. Default value `INFO`.
  * DEBUG: Display INFO logs + more details at DEBUG level.
  * TRACE: Display DEBUG logs + the timings of all ElasticSearch queries and Web API calls executed by the SonarQube Scanner.
* `showProfiling`: Display logs to see where the analyzer spends time. Default value `false`
* `projectKey`: defaults to `DRONE_REPO`
* `projectName`: defaults to `DRONE_REPO`

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

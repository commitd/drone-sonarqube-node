FROM golang:1.13 as builder
ADD . /src
WORKDIR /src
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonarqube-node

FROM timbru31/java-node:latest
COPY --from=builder /src/drone-sonarqube-node /bin/

LABEL maintainer="Committed Software <docker@committed.software>"
LABEL description="Drone CI plugin for sonarqube quality testing on node"

ARG SONAR_VERSION=3.3.0.1492
ARG SONAR_SCANNER_CLI=sonar-scanner-cli-${SONAR_VERSION}-linux
ARG SONAR_SCANNER=sonar-scanner-${SONAR_VERSION}-linux

WORKDIR /bin

RUN curl https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/${SONAR_SCANNER_CLI}.zip -so /bin/${SONAR_SCANNER_CLI}.zip
RUN unzip ${SONAR_SCANNER_CLI}.zip \
    && rm ${SONAR_SCANNER_CLI}.zip

ENV PATH $PATH:/bin/${SONAR_SCANNER}/bin

ENTRYPOINT /bin/drone-sonarqube-node

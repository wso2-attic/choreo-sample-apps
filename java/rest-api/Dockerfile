# Docker multi-stage build

# 1. Building the App with Maven
FROM maven:3.8.7-eclipse-temurin-19-alpine

ADD . /byoc-java-springboot
WORKDIR /byoc-java-springboot

# Just echo so we can see, if everything is there :)
RUN ls -l

# Run Maven build
RUN mvn clean install


# 2. Just using the build artifact and then removing the build-container
FROM openjdk:19-alpine

# https://security.alpinelinux.org/vuln/CVE-2021-46848
RUN apk add --upgrade libtasn1-progs

# https://security.alpinelinux.org/vuln/CVE-2022-37434
RUN apk update && apk upgrade zlib

RUN apt-get install -y libspring-web4-java



MAINTAINER tnnmuhandiram

VOLUME /tmp

USER 10014

# Add Spring Boot app.jar to Container
COPY --from=0 "/byoc-java-springboot/target/byoc-*-SNAPSHOT.jar" app.jar

# Fire up our Spring Boot app by default
CMD [ "sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -jar /app.jar" ]

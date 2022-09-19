# FROM openjdk:11
# MAINTAINER tharindumu@wso2.com
# COPY target/byoc-0.0.1-SNAPSHOT.jar target/byoc-0.0.1-SNAPSHOT.jar
# ENTRYPOINT ["java","-jar","target/byoc-0.0.1-SNAPSHOT.jar"]


# Docker multi-stage build

# 1. Building the App with Maven
FROM maven:3-jdk-11

ADD . /byoc-java-springboot
WORKDIR /byoc-java-springboot

# Just echo so we can see, if everything is there :)
RUN ls -l

# Run Maven build
RUN mvn clean install


# 2. Just using the build artifact and then removing the build-container
FROM openjdk:11-jdk

MAINTAINER tnnmuhandiram

VOLUME /tmp

USER 10014

# Add Spring Boot app.jar to Container
COPY --from=0 "/byoc-java-springboot/target/byoc-*-SNAPSHOT.jar" app.jar

# Fire up our Spring Boot app by default
CMD [ "sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -jar /app.jar" ]
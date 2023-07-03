FROM openjdk:8-jdk-alpine as build

WORKDIR /workspace/app

COPY data/mvnw .
COPY data/.mvn .mvn
COPY data/pom.xml .
COPY data/src src

RUN ./mvnw install -DskipTests
WORKDIR /workspace/app/target/dependency
RUN jar -xf ../*.jar

FROM openjdk:8-jdk-alpine

RUN apk add --no-cache wget

ARG DEPENDENCY=/workspace/app/target/dependency
COPY --from=build ${DEPENDENCY}/BOOT-INF/lib /app/lib
COPY --from=build ${DEPENDENCY}/META-INF /app/META-INF
COPY --from=build ${DEPENDENCY}/BOOT-INF/classes /app
RUN wget -O dd-java-agent.jar https://dtdg.co/latest-java-tracer
ENTRYPOINT ["java","-javaagent:dd-java-agent.jar","-cp","app:app/lib/*","com.datadog.springJmsExample.SpringJmsExampleApplication"]

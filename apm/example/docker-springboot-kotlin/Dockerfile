FROM ubuntu:xenial

RUN apt-get update

RUN apt-get install -y nano

RUN apt-get install -y wget

RUN apt-get install -y unzip

#Java 8
RUN apt-get install -y openjdk-8-jre

WORKDIR /home

RUN wget https://github.com/JetBrains/kotlin/releases/download/v1.3.61/kotlin-compiler-1.3.61.zip

RUN unzip kotlin-compiler-1.3.61.zip -d kotlin-compiler-1.3.61

ENV PATH $PATH:/home/kotlin-compiler-1.3.61/kotlinc/bin

ADD data/notes notes

WORKDIR /home/notes

RUN wget -O dd-java-agent.jar https://dtdg.co/latest-java-tracer

RUN ./mvnw package

ENV DD_SERVICE kotlin-app
ENV DD_VERSION 1.0

CMD java -javaagent:dd-java-agent.jar -jar target/notes-0.0.1-SNAPSHOT.jar


LABEL com.datadoghq.tags.service="kotlin-app"
LABEL com.datadoghq.tags.version="1.0"
LABEL "com.datadoghq.ad.check_names"='["java"]'
LABEL "com.datadoghq.ad.init_configs"='[{}]'
LABEL "com.datadoghq.ad.logs"='[{"source": "custom", "service": "kotlin-app"}]'

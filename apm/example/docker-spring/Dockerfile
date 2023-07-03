FROM openjdk:8u322

COPY . /home
WORKDIR /home/SpringApps

RUN curl -o dd-java-agent.jar 'https://dtdg.co/latest-java-tracer'

RUN ./gradlew build

CMD java -javaagent:dd-java-agent.jar \
    -jar /home/SpringApps/build/libs/springapps-1.jar \
    --server.port=xxxx

FROM tomcat:9.0.27-jdk12-adoptopenjdk-hotspot
RUN apt-get update

RUN apt-get install nano wget

RUN wget -O dd-java-agent.jar 'https://dtdg.co/latest-java-tracer' && echo "export CATALINA_OPTS='-javaagent:/usr/local/tomcat/dd-java-agent.jar -Ddd.trace.debug=true'" >> bin/setenv.sh

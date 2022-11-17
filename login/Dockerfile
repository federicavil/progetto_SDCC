FROM openjdk:17

WORKDIR login

COPY .mvn ./.mvn
COPY mvnw.cmd ./
COPY mvnw ./
COPY pom.xml ./
COPY login.iml ./
RUN ./mvnw dependency:resolve

COPY src ./src

EXPOSE 9090

CMD ["./mvnw", "spring-boot:run"]
FROM openjdk:17

WORKDIR reviewManager

COPY .mvn ./.mvn
COPY mvnw.cmd ./
COPY mvnw ./
COPY reviewManager.iml ./
COPY pom.xml ./
RUN ./mvnw dependency:resolve

COPY src ./src

EXPOSE 9091

CMD ["./mvnw", "spring-boot:run"]
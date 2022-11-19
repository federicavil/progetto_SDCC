FROM eclipse-temurin:17-jdk-focal

WORKDIR reviewManager

COPY .mvn ./.mvn
COPY mvnw.cmd ./
COPY mvnw ./
COPY pom.xml ./
COPY reviewManager.iml ./
RUN sed -i 's/\r$//' mvnw
RUN ./mvnw dependency:resolve

COPY src ./src

EXPOSE 9091

CMD ["./mvnw", "spring-boot:run"]
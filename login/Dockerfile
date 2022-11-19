FROM eclipse-temurin:17-jdk-focal

WORKDIR login

COPY .mvn ./.mvn
COPY mvnw.cmd ./
COPY mvnw ./
COPY pom.xml ./
COPY login.iml ./
RUN sed -i 's/\r$//' mvnw
RUN ./mvnw dependency:resolve

COPY src ./src

EXPOSE 9090

CMD ["./mvnw", "spring-boot:run"]
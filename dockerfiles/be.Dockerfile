FROM gradle:6.0.0-jdk11 as build
WORKDIR /usr/src/app/be
COPY ./be ./
RUN ./gradlew clean shadowJar

FROM openjdk:11-jre-slim as run
WORKDIR /usr/src/app/be
COPY --from=build /usr/src/app/be/build/libs ./
CMD [ "java", "-jar", "be-all.jar" ]
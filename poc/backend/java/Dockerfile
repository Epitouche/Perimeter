# Stage 1: Dependency resolution
FROM maven:3.9 AS dependencies

WORKDIR /app

# Copy only the pom.xml and settings files for dependency resolution
COPY ./project/pom.xml .

# Download dependencies
RUN mvn dependency:go-offline -B

# Stage 2: Build
FROM maven:3.9 AS build

WORKDIR /app

# Copy the resolved dependencies from the previous stage
COPY --from=dependencies /root/.m2 /root/.m2

# Copy the source code
COPY ./project/. .

# Build the application
RUN mvn install -DskipTests

# Stage 3: Runtime
FROM openjdk:24-jdk AS run

WORKDIR /app

# Copy the built application
COPY --from=build /app/target/demo-0.0.1-SNAPSHOT.jar /app/

# Command to run the application
CMD ["java", "-jar", "/app/demo-0.0.1-SNAPSHOT.jar"]

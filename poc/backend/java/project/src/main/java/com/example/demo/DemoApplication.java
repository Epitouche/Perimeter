package com.example.demo;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class DemoApplication {
    public static void main(String[] args) {
      SpringApplication.run(DemoApplication.class, args);
    }

    @GetMapping("/")
    public String documentation() {
      return String.format("possible routes: /, /ping, /hello, /about.json");
    }

    @GetMapping("/about.json")
    public String about() {

      String json = "{"
      + "\"name\": \"Spring Boot\","
      + "\"description\": \"Spring Boot is a Spring module which provides RAD (Rapid Application Development) feature to Spring framework.\","
      + "\"website\": \"https://spring.io/projects/spring-boot\""
      + "}";
      return json;
}



    @GetMapping("/ping")
    public String ping() {
      return String.format("pong");
    }

    @GetMapping("/hello")
    public String hello(@RequestParam(value = "name", defaultValue = "World") String name) {
      return String.format("Hello %s!", name);
    }
}

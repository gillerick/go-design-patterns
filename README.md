# Go Design Patterns

****Singleton Pattern****

The purpose of the singleton pattern is to restrict the instantiation of a class to a single instance and provide global
access to that instance

Other benefits

1. Allows us to provide lazy instantiation (the class isn't instantiated until it is needed)
2. Useful when you want to ensure that there's only one instance of a particular class in your application e.g in
   logging, application configuration, telemetry & analytics, debugging
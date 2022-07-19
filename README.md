# Go Design Patterns

****Singleton Pattern****

The purpose of the singleton pattern is to restrict the instantiation of a class to a single instance and provide global
access to that instance

Other benefits

1. Allows us to provide lazy instantiation (the class isn't instantiated until it is needed)
2. Useful when you want to ensure that there's only one instance of a particular class in your application e.g in
   logging, application configuration, telemetry & analytics, debugging

***Observer Pattern***

This describes a pattern where an observed object (subject) maintains a list of other objects that are interested in
events and other state changes in the subject. These other objects are called **observers**

This pattern is also referred to as a _publish and subscribe architecture_ or _event driven system_

It is useful in scenarios where the changes in state of a given subject are unpredictable, and the number of interested
observers cannot be known in advance, or changes during the execution of the program.
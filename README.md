# Go Design Patterns

****Factory Pattern****

This pattern defines a way for a program to create objects at runtime without having to specify the exact class or type
of the object to be created. This allows us to decouple the creation of the object from the definition of the object
itself.

****Singleton Pattern****

The purpose of the singleton pattern is to restrict the instantiation of a class to a single instance and provide global
access to that instance

Other benefits

1. Allows us to provide lazy instantiation (the class isn't instantiated until it is needed)
2. Useful when you want to ensure that there's only one instance of a particular class in your application e.g in
   logging, application configuration, telemetry & analytics, debugging

****Adapter Pattern****

Adapter pattern is a structural design pattern that describes a way for the interface of a given class or type to be
used as if it were a different interface without the need to modify the code of the existing class.

It is useful in situations when the application being developed needs access to older legacy code bases that, for one
reason or the other, can't be changed. Or when working with APIs and subsystems you don't have ownership of and
therefore can't make any changes to.

The adapter pattern therefore allows otherwise incompatible objects to work together.

****Facade Pattern****

This pattern is called facade because it is analogous to the architectural term for the front-facing part of a building
that introduces one to the rest of the pattern.

The purpose of this pattern is to improve the usability of a more complex API via a simpler high-level interface.

It can also be used to adapt code for newer application, or when one is looking to reduce the amount of tight-coupling
between different parts of an application to make it easier to swap out different areas of functionality as requirement
change.

****Observer Pattern****

This describes a pattern where an observed object (subject) maintains a list of other objects that are interested in
events and other state changes in the subject. These other objects are called **observers**

This pattern is also referred to as a _publish and subscribe architecture_ or _event driven system_

It is useful in scenarios where the changes in state of a given subject are unpredictable, and the number of interested
observers cannot be known in advance, or changes during the execution of the program.

****Iterator Pattern****

_Iteration_ is the process of accessing the objects of a collection or container object in a sequential order.

The iterator pattern describes a way of performing this operation without exposing the underlying implementation of the
container object. In other words, the code that works on each element in the container should not have to know how the
individual objects are stored by the container in order to access them.

An example is a software system that controls a library containing various publications such as book and magazines. In
this case the iterator pattern could be used to create an _iterator object_ that knows how to traverse the library so
that in case the underlying implementation of the library storage changes, the _consumers_ of the library iterator would
not need to change their code.

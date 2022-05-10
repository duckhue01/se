### service locator


#### definition
the service locator pattern is a design pattern used in software development to encapsulate the processes involved in obtaining a service with a strong abstraction layer. this pattern uses a central registry know as the "service locator", which on request returns the information necessary to perform the certain task. this pattern is anti-pattern also because it obscures dependencies and makes software harder to test

#### advantages
1. the "service locator" can act as a simple runtime-linker
2. 
#### disadvantages

1. thee registry hides the class dependencies causing run-time errors instead of compile-time errors when dependencies are missing
2. hard to test
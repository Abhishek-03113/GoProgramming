🧩 Problem: Greeting Generator in Go
Build a Go application that prints personalized greeting messages using a modular program structure.
You are required to organize your code into at least two packages: a main package to run the program, and a custom package responsible for generating greeting messages. The custom package should define a constant holding a default name and a function that returns a greeting message for a given name.
Your application should:
* Use appropriate import statements to include your custom package in the main package.
* Declare and initialize variables for user names.
* Use exported identifiers where necessary to allow access across packages, and keep internal logic unexported.
* Print greeting messages for both a user-defined name and the default name defined in your custom package.
# Car Rental System — SDP Final Project

This project is the final assignment for the Software Design Patterns (SDP) course, developed by Ansar and Nursultan.
The goal of the project is to design a clean, extensible, and Java-style architecture in Go using six design patterns, all integrated into a coherent Car Rental System domain.

The codebase follows strict principles of SOLID, Clean Code, single-responsibility file organization, and Java-like package layout (each interface in its own file, clear domain separation, explicit naming conventions such as IVisitor, IVehicle, IPricer, IBuilder, etc.).

##Implemented Design Patterns

### 1. Abstract Factory

Used to create vehicles for different rental branches (e.g., city branch, airport branch).
Ensures the client code depends only on abstractions and not concrete vehicle types or configurations.

### 2. Bridge

Separates the vehicle model from the pricing strategy.
Provides a flexible way to compute base pricing without coupling it to vehicle implementations.

### 3. Decorator

Adds optional services (insurance, GPS, child seat) dynamically on top of the base pricer.
Allows combining any number of add-ons without modifying existing classes.

### 4. Builder

Constructs the RentalContract step by step.
Integrates pricing (Bridge), add-ons (Decorator), and event notifications (Observer) into a single workflow.

### 5. Observer

Notifies external listeners (email notifier, audit logger) about important rental events.
Implements loose coupling between the system core and external side effects.

### 6. Visitor

Applies different operations to the entire fleet of vehicles, such as pricing evaluation and maintenance reporting.
Adds new behaviors without modifying the vehicle classes.

##Team Responsibilities

Nursultan:
Abstract Factory, Decorator, Visitor.

Ansar:
Bridge, Builder, Observer.

Project Highlights

Fully demonstrates six classic design patterns working together inside one realistic business domain.

Strong adherence to SOLID and clean separation of responsibilities.

All patterns are shown clearly in the main.go for defense presentation.

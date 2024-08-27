# SAP-Assignment
## About
This project intended to demonstrate a simple Trails lookup from the data source using various filtering option available
### Features 
+ Application provides command line interface with various flag, which can be used to filter the avialable trails
+ Various application consigs can be maneged from configs directory
## Architecture
This project is using layered architecture, with each layer loosely coupled with other
<img width="500" alt="image" src="https://github.com/user-attachments/assets/e826498b-627e-445c-a083-ae4ab33a24c5">

**Data layer**/**Models**: Since the scope of this project is simple, a simple model is used to map data from the CSV file and maintain it in memory

**Service layer**: This layer is responsible for the business logic of the application, it accesses the dal layer through the interface. This layer also ensures that there are no dependencies on the external user interfaces such as HTTP context, command line arg, buffer, etc and logic only depends on the function parameters, and the response is provided only through the returned values. 


**User Interface Layer/Presentation Layer**: This layer is responsible for accessing the application services. The same service layer can be used with any user interface. For this project we are using CLI and interaction is provided through various flags.

## CLI Options
 
 ![CLI options](image.png)

## How to Run? 


To run this application 

Clone the repo

> go get https://github.com/ozwin/sap-assignment.git

Run the following from the root folder of the project
> cd ./internal/cmd
> 
> go run ./main.go --help

you will be able to see the below options

![CLI options](image.png)

Now based on your prefernce you can pass single/multiple arguments.

> go run .\main.go -has_dog_compost -has_bike_trail=false

![filtering data](image-1.png)

All the avilable flags are boolean taht take true/false as value, if value is not explicitly passed for a flag then it's considered as true.
eg -has_bike_trail=true and -Has_bike_trail will yield the same results.
![true flags](image-2.png)

Or

Navigate to the VS Code Run and Debug menu, and launch 'Start CLI' from the preconfigured Launch.json. This will display the avialable flags for the application


**p.s: Ensure you are running the SIP application first for this to work**

## How to test?

This repo currently has unit test coverage for most of the functionalities, 

to test run the below command from the root of the project

> go test ./...


## Project structure

**internal** Since this repo is meant only for internal features all code is bundled under this folder


**data** contains the raw data file provided as part of the assignment

**configs** contains app-specific configs
**cmd** contains the entry point for the applications

**app** contains the logic for the applications

**models** contains the models for the application

**services** contains the actual logic required for the application

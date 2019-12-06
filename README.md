# CppCodeGenerator
## C++ Code Generation Tools

A project for generating boilerplate C++ code from user-provided template files (.txt).

### Files that can be generated:
- Interfaces
- Class Headers (.h)
- Class Implementations (.cpp / .cxx)
- Mock Classes
- Test Classes

## Dependencies
- Go 1.13 +

## Installation
```
git clone https://github.com/emloughl/CppCodeGenerator.git
cd CppCodeGenerator
mkdir build && cd build
go build ..
```

## Usage
### Arguments
##### Class Type
` -t ` or `--type`: Specifies the type of class to be generated. 
 - Allowed values: `interface`, `class`, `mock`, or `test`.

##### Path to Existing Interface
` -i` or `--interface`: Specifies the path to a user's existing C++ interface from which to generate another class.

##### Class Name
` -n` or `--class`: Specifies the name of the class to be generated.

## Generating an Interface
##### Command:

```
CppCodeGenerator -t interface -n MyFirstClass
```
 ![Generating a new interface](documentation/readme_resources/01_new_interface.gif)
 
##### Output:
![Generated blank interface](documentation/readme_resources/02_generated_interface.GIF)

## Generating a Class (Header and Implementation)
#### Prerequisite
To generate a class, you require an interface from which to implement. The first step is to define your interface's pure virtual functions. In this example, we'll define three pure virtual functions, `foo`, `bar`, and `baz`. Each of them with intentionally convoluted signatures to demonstrate the generator's capabilities.

![Interface with completed pure virtual function definitions](documentation/readme_resources/03_completed_interface.gif)




### Configurations

Modifying the `config.json` file allows you to alter your desired prefixes, suffixes, and other policies.

#### Date Format: 
Follow the Golang date format. Example here: https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format

## TODO:
Add list in README of all possible template parameters

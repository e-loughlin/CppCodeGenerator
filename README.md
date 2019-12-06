# CppCodeGenerator
## C++ Code Generation Tools

A project for generating boilerplate C++ code from user-provided template files (.txt), and user-defined configurations (config.json).

### Files that can be generated:
- Interfaces
- Class Headers (.h / .hpp)
- Class Implementations (.cpp / .cxx)
- Mock Classes (GoogleMock)
- Test Classes (GoogleTest)

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
#### Class Type
` -t ` or `--type`: Specifies the type of class to be generated. 
 - Allowed values: `interface`, `class`, `mock`, or `test`.

#### Path to Existing Interface
` -i` or `--interface`: Specifies the path to a user's existing C++ interface from which to generate another class.

#### Class Name
` -n` or `--class`: Specifies the name of the class to be generated.

## Generating an Interface
#### Command:

```
CppCodeGenerator -t interface -n MyFirstClass
```
 ![Generating a new interface](documentation/readme_resources/01_new_interface.gif)
 
#### Output:
![Generated blank interface](documentation/readme_resources/02_generated_interface.GIF)

## Generating a Class (Header and Implementation)
#### First Complete your Interface
To generate a class, you require an interface from which to implement. The first step is to define your interface's pure virtual functions. In this example, we'll define three pure virtual functions, `foo`, `bar`, and `baz`. Each of them with intentionally convoluted signatures to demonstrate the generator's capabilities.

![Interface with completed pure virtual function definitions](documentation/readme_resources/03_completed_interface.gif)

Now, we'll generate a class that inherits `IMyFirstClass.h`:

#### Command:
 ![Generating a new class from an existing interface](documentation/readme_resources/04_new_class.gif)

#### Output:
 ![Generating a new class from an existing interface](documentation/readme_resources/05_generated_class.gif)
 
##### Some Cleanup Required
The program isn't perfect, and handles dependencies by placing all required types as forward declares in the header, and including them in the implementation. Some cleanup may be required for templated types like QMap<> and QHash<>.
 
 #### Class Naming
By default, the program generates a class name based on the interface's name. To specify a different name, use the `-n` or `--name` option. 
 ![Generating a new class with a unique name](documentation/readme_resources/06_new_class_unique_name.gif)

## Generating a Mock
#### First Complete your Interface
Just like generating a class, a completed interface is required. As an example, we'll use the same `IMyNewClass.h` interface previously defined.

#### Command:
 ![Generating a new mock](documentation/readme_resources/07_new_mock.gif)

#### Output:
 ![Generating a new mock](documentation/readme_resources/08_mock_output.gif)

### Configurations

Modifying the `config.json` file allows you to alter your desired prefixes, suffixes, and other policies.

#### Date Format: 
Follow the Golang date format. Example here: https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format

## TODO:
Add list in README of all possible template parameters

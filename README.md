# CppCodeGenerator
## C++ Code Generation Tools

# PROJECT STATUS: In Progress

(In-progress re-write my python project [CplusplusCodeGenerators](https://github.com/emloughl/CplusplusCodeGenerators "CplusplusCodeGenerators") - in Golang)

A project for generating boilerplate C++ code from user-provided template files (.txt).

Files that can be generated:
- Interfaces
- Class Headers (.h)
- Class Implementations (.cpp / .cxx)

In progress: 
- Mocks / SpyMocks
- Tests

## Installation
### Windows
``` 
 $ cd CppCodeGenerator/
 $ environment.bat
 $ cd src/
 $ go install
```

### Mac / Linux
```
 $ cd CppCodeGenerator/
 $ source environment.sh
 $ cd
 $ go install
```

## Usage
```
 $ cd CppCodeGenerator/bin/
 $ src.exe -type=<TYPE> -interface=<PATH_TO_INTERFACE>
```

### Arguments
#### Types
- `-type=interface` Creates a new interface
- `-type=class` : Creates a new `.h` and `.cpp / .cxx` class file.

#### Path to Interface to Implement
The interface from which the generated derived classes, if applicable, will inherit. Any pure virtual functions or signals (Qt) declared in the inherited interface will be declared and defined in the derived classes as appropriate. The path to this inherited interface is set using the `-interface=` flag.

### Examples:

##### Creating a new (blank) interface:
```
NewClass.exe -type=interface -interface=/path/to/I_MyNewClass.h
```

This will create a new interface with the name `I_MyNewClass.h` at the specified path. Note that the `I_` prefix must match what is specified in the configuration file `config.json`.


##### Creating a new class that inherits a given interface:
```
 $ NewClass.exe -type=class -interface=C:/ws/I_MyClass.h
```
Output:
    Will create 2 files: `MyClass.h`, and `MyClass.cpp`. 


### Configurations

Modifying the `config.json` file allows you to alter your desired prefixes, suffixes, and other policies.

#### Date Format: 
Follow the Golang date format. Example here: https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format

## TODO:
Add list in README of all possible template parameters
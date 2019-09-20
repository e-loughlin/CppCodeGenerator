# CppCodeGenerator
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
` $ cd CppCodeGenerator/`
` $ environment.bat`
` $ cd src/`
` $ go install`

### Mac / Linux
` $ cd CppCodeGenerator/`
` $ source environment.sh`
` $ cd `
` $ go install`

##Usage
` $ cd CppCodeGenerator/bin/`
` $ src.exe -type=<TYPE> -interface=<PATH_TO_INTERFACE>`

###Arguments
####Types
- `-type=interface` Creates a new interface
- `-type=class` : Creates a new `.h` and `.cpp / .cxx` class file.

####Path to Interface
The interface from which the generated derived classes, if applicable, will inherit. Any pure virtual functions or signals (Qt) declared in the inherited interface will be declared and defined in the derived classes as appropriate.
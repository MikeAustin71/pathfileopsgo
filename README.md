# pathfileops

*pathfileops* is a software library or in *Go* parlance, a software *package*.

This package is written in the *Go* programming language, a.k.a 'golang'.
 
*pathfileops* provides software types and methods used in the management,
organization and control of disk files and directories.

The source code repository for this package is located at:
  https://github.com/MikeAustin71/pathfileopsgo.git

This package was developed and tested on Windows, although the package
was designed to operate on multiple operating systems including 
Mac-OS, Linux and Windows.

__`To date, the source code has only been tested on Windows.`__

___    
[Source Code Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)    
___


# Table Of Contents
+ [Getting Started](#getting-started)
  - [Supported Platforms](#supported-platforms)
  - [Installation](#installation)
  - [Source Code Import](#source-code-import)
+ [Version](#version)
+ [Source Code Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)
+ [Primary Types](#primary-types)
  - [File Helper](#filehelpler)
  - [Directory Manager](#dirmgr)
  - [File Manager](#filemgr)
  - [Collections](#collections)
+ [Dependencies](#dependencies)
+ [Tests](#tests)

<a name="getting-started"></a>
# Getting Started 

<a name="supported-platforms"></a>
## Supported Platforms
This package was developed and tested on Windows, although the package
was designed to operate on multiple operating systems including 
Mac-OS, Linux and Windows.

__`To date, the source code has only been tested on Windows.`__

<a name="installation"></a>
## Installation
Use this command to down load and install the *pathfileops* package
locally. 

    go get github.com/MikeAustin71/pathfileopsgo/pathfileops

To update the library run:
    
    go get -u github.com/MikeAustin71/pathfileopsgo/pathfileops

<a name="source-code-import"></a>
## Source Code Import        
After installation, you may import and reference the library
as follows:

        import (
            "MikeAustin71/pathfileopsgo/pathfileops"
        )    

All of the active production files are located in directory path:

     github.com/MikeAustin71/pathfileopsgo/pathfileops


## Version

This is Version 1.5.0.

This version significantly expands the capabilities of *DirMgr*, as
well as expanding the read/write features for *FileMgr*. Unit Testing
and code coverage have also been significantly expanded. 

This version does NOT support *Go* modules.
___    
[Source Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)    
___

<a name="primarytypes"></a>
## Primary Types
      
This package incorporates three primary types: 
    
    1. FileHelper
    
    2. DirMgr
    
    3. FileMgr

<a name="filehelper"></a>
#### FileHelpler
The type, *FileHelper* is located in directory *./pathfilego/pathfileops*.
The source code is located in source file, *filehelper.go*. This type includes a variety
of general utility functions for managing files and directories.

<a name="dirmgr"></a>
#### DirMgr
The type, *DirMgr*, is located in directory *./pathfilego/pathfileops*.
The source code is located in source file, *dirmanager.go*. Directory Manager
is designed to manage file paths or directories. *DirMgr* is dependent on type *FileHelper*
discussed above.

<a name="filemgr"></a>
#### FileMgr 
The type, *FileMgr*, is located in directory *./pathfilego/pathfileopos*. 
The source code is located in source file, *filemanager.go*. The File Manager
type is designed to manage disc files. *FileMgr* is dependent on *FileHelper*
and *DirMgr*, discussed above.

<a name="collections"></a>
### Collections
This library also provides collections to manage groups of file and directory
managers.

1. DirMgrCollection - Processes and manages collections of type 'DirMgr'

2. FileMgrCollection - Processes and manages collections of type 'FileMgr' 

3. FileOpsCollection - Manages collections of operations performed on disk
 files.
 
<a name="dependencies"></a> 
### Dependencies
The three types discussed above are interdependent.

1. *./pathfilego/pathfileops/filehelper.go*

2. *./pathfilego/pathfileops/dirmanager.go*

3. *./pathfilego/pathfileops/filemanager.go*

<a name="tests"></a>
## Tests
This version includes 1,793 tests with a source code coverage of 82%.
All tests are currently completing successfully as documented in the
following text file:

      '../pathfilego/pathfileops/xx_tests.txt'


 

 

# pathfileops

*pathfileops* is a software library or in go parlance, a software *package*.
 
This package provides software types and methods used in the management,
organization and control of disk files and directories.

*pathfileops* is written in the *Go* programming language, a.k.a 'golang'.

The source code repository for this package is located at:
  https://github.com/MikeAustin71/pathfileopsgo.git

This package was developed and tested on Windows, although the package
was designed to operate on multiple operating systems including 
Mac-OS, Linux and Windows.

__`To date, the source code has only been tested on Windows.`__


## Version Number 1.5.0
This version significantly expands the capabilities of *DirMgr*, as
well as expanding the read/write features for *FileMgr*. Unit Testing
and code coverage have also been significantly expanded. 

This version does NOT support *Go* modules.
___    
[Click To View Source Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)    
___

## Installing The 'pathfileops' Package
Use this command to down load and install the *pathfileops* package
locally. 

    go get github.com/MikeAustin71/pathfileopsgo/pathfileops

To update the library run:
    
    go get -u github.com/MikeAustin71/pathfileopsgo/pathfileops
        
After installation, you may import and reference the library
as follows:

        import (
            "MikeAustin71/pathfileopsgo/pathfileops"
        )    

All of the active production files are located in directory path:

     github.com/MikeAustin71/pathfileopsgo/pathfileops
      
This package incorporate three primary types: 
    
    1. FileHelper
    
    2. DirMgr
    
    3. FileMgr

#### FileHelpler
The type, *FileHelper* is located in directory *./pathfilego/pathfileops*.
The source code is located in source file, *filehelper.go*. This type includes a variety
of general utility functions for managing files and directories.

#### DirMgr
The type, *DirMgr*, is located in directory *./pathfilego/pathfileops*.
The source code is located in source file, *dirmanager.go*. Directory Manager
is designed to manage file paths or directories. *DirMgr* is dependent on type *FileHelper*
discussed above.

#### FileMgr 
The type, *FileMgr*, is located in directory *./pathfilego/pathfileopos*. 
The source code is located in source file, *filemanager.go*. The File Manager
type is designed to manage disc files. *FileMgr* is dependent on *FileHelper*
and *DirMgr*, discussed above.

### Collections
This library also provides collections to manage groups of file and directory
managers.

1. DirMgrCollection - Processes and manages collections of type 'DirMgr'

2. FileMgrCollection - Processes and manages collections of type 'FileMgr' 

3. FileOpsCollection - Manages collections of operations performed on disk
 files.
 
### Dependencies
The three types discussed above are interdependent.

1. *./pathfilego/pathfileops/filehelper.go*

2. *./pathfilego/pathfileops/dirmanager.go*

3. *./pathfilego/pathfileops/filemanager.go*

## Tests
This version includes 1,793 tests with a source code coverage of 82%.
All tests are currently completing successfully as documented in the
following text file:

      '../pathfilego/pathfileops/xx_tests.txt'

## Operating Systems
This package was developed and tested on Windows, although the package
was designed to operate on multiple operating systems including 
Mac-OS, Linux and Windows.

__`To date, the source code has only been tested on Windows.`__

 
## Version Number 1.5.0
This version significantly expands the capabilities of *DirMgr*, as
well as expanding the read/write features for *FileMgr*. Unit Testing
and code coverage have also been significantly expanded.

This version does NOT support *Go* modules.
 
___
    
[Click To View Source Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)

___    

 

 

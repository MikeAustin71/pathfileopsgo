# pathfileops
This package provides software types and methods used in the management
organization and control of disk files and directories.

*pathfileops* is written in the Go programming language, 'golang'.

The source code repository for this package is located at:
  https://github.com/MikeAustin71/pathfileopsgo.git

#### Installing 'pathfileops' Library
Use this command to down load and install the datetime library
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
  **github.com/MikeAustin71/pathfileopsgo/pathfileops**

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
The source code is located in source file, *dirmanagers.go*. Directory Manager
is designed to manage file paths or directories. *DirMgr* is dependent on type *FileHelper*
discussed above.

#### FileMgr 
The type, *FileMgr*, is located in directory *./pathfilego/pathfileopos*. 
The source code is located in source file, *filemanagers.go*. The File Manager
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
The three types discussed above are interdependent. Therefore you will need the 
following three source code files in order to fully utilize this utility library.

1. *./pathfilego/pathfileops/filehelper.go*

2. *./pathfilego/pathfileops/dirmanagers.go*

3. *./pathfilego/pathfileops/filemanagers.go*

#### Dependencies for Tests
##### DateTime Utility
The FileHelper utility tests rely on the datetime utility which incorporates routines
covering datetime, timezones and time duration. The datetime utility file,
*datetimeutility.go* is located in directory:
*./pathfilego/appLibs*

The DateTime Utility source code located in repository:
 https://github.com/MikeAustin71/datetimeopsgo.git
 
------------------------------------------------------------------------------------------
    
[Click To View Source Documentation](http://godoc.org/github.com/MikeAustin71/pathfileopsgo/pathfileops)    
 

 

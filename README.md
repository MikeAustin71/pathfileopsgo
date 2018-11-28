# pathfilego
This code demonstrates path and file operations in 'golang', the go programming language.

The path for this source code repository is:
  https://github.com/MikeAustin71/pathfilego.git

All of the active production files are located in directory path:
  **./pathfilego/003_filehelper**

The utilities incorporate three primary types: 
    
    1. FileHelper
    
    2. DirMgr
    
    3. FileMgr

#### FileHelpler
The type, *FileHelper* is located in directory *./pathfilego/003_filehelper/common*.
The source code is located in source file, *filehelper.go*. This type includes a variety
of general utility functions for managing files and directories.

#### DirMgr
The type, *DirMgr*, is located in directory *./pathfilego/003_filehelper/common*.
The source code is located in source file, *fileanddirmanagers.go*. Directory Manager
is designed to manage file paths or directories. *DirMgr* is dependent on type *FileHelper*
discussed above.

#### FileMgr 
The type, *FileMgr*, is located in directory *./pathfilego/003_filehelper/common*. 
The File Manager type is designed to manage files. *FileMgr* is dependent on *FileHelper*
and *DirMgr*, discussed above.


#### Dependencies
The three types discussed above are interdependent. Therefore you will need the 
following two source code files in order to fully utilize this utility library.

1. *./pathfilego/003_filehelper/common/filehelper.go*

2. *./pathfilego/003_filehelper/common/fileanddirmanagers.go*


#### Dependencies for Tests
##### DateTime Utility
The FileHelper utility tests rely on the datetime utility which incorporates routines
covering datetime, timezones and time duration. The datetime utility file,
*datetimeutility.go* is located in directory:
*./pathfilego/003_filehelper/common*

The DateTime Utility source code located in repository:
 https://github.com/MikeAustin71/datetimeopsgo.git

 

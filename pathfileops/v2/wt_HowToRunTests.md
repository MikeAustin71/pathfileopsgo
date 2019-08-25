# Running Tests

Open a command prompt in this directory (pathfileopsgo/pathfileops/v2) 
and run the following command:

     `go test -v > xx_tests.txt`

This will generate test results and save them to a text file (`xx_tests.txt`) 
located in the *pathfileopsgo/pathfileops/v2* directory.

## Running Tests with code coverage

First pull down and install the `cover` package.
 
    `go get golang.org/x/tools/cmd/cover`
  
Next, follow the test execution protocol.  
  
## Test Execution with Code Coverage
Open a command prompt in this directory (pathfileopsgo/pathfileops/v2) 
and run the following command:


     `go test -cover -v > xx_tests.txt`  

     
Again, this will generate test results and save them to a text file
(`xx_tests.txt`) located in the *pathfileopsgo/pathfileops/v2* directory.
     
## Cover Profile

Generate the code coverage detail. Open a command prompt in this
directory (pathfileopsgo/pathfileops/v2) and run the following
command: 

    `go test -coverprofile=xx_coverage.out`


The following command will open your browser to display code coverage. 
Open a command prompt in this directory (pathfileopsgo/pathfileops/v2)
and run the following command:

     `go tool cover -html=xx_coverage.out`
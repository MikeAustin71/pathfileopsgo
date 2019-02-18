package pathfileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mOsPermissionCodeToString = map[os.FileMode]string{
	os.FileMode(0):    "ModeNone",
	os.ModeDir:        "ModeDir",
	os.ModeAppend:     "ModeAppend",
	os.ModeExclusive:  "ModeExclusive",
	os.ModeTemporary:  "ModeTemporary",
	os.ModeSymlink:    "ModeSymlink",
	os.ModeDevice:     "ModeDevice",
	os.ModeNamedPipe:  "ModeNamedPipe",
	os.ModeSocket:     "ModeSocket",
	os.ModeSetuid:     "ModeSetuid",
	os.ModeSetgid:     "ModeSetgid",
	os.ModeCharDevice: "ModeCharDevice",
	os.ModeSticky:     "ModeSticky",
	os.ModeIrregular:  "ModeIrregular",
}

var mOsPermissionCodeToLetter = map[os.FileMode]string{
	os.FileMode(0):    "-",
	os.ModeDir:        "d",
	os.ModeAppend:     "a",
	os.ModeExclusive:  "l",
	os.ModeTemporary:  "T",
	os.ModeSymlink:    "L",
	os.ModeDevice:     "D",
	os.ModeNamedPipe:  "p",
	os.ModeSocket:     "S",
	os.ModeSetuid:     "u",
	os.ModeSetgid:     "g",
	os.ModeCharDevice: "c",
	os.ModeSticky:     "t",
	os.ModeIrregular:  "?",
}

var mOsPermissionStringToCode = map[string]os.FileMode{
	"ModeNone":       os.FileMode(0),
	"ModeDir":        os.ModeDir,
	"ModeAppend":     os.ModeAppend,
	"ModeExclusive":  os.ModeExclusive,
	"ModeTemporary":  os.ModeTemporary,
	"ModeSymlink":    os.ModeSymlink,
	"ModeDevice":     os.ModeDevice,
	"ModeNamedPipe":  os.ModeNamedPipe,
	"ModeSocket":     os.ModeSocket,
	"ModeSetuid":     os.ModeSetuid,
	"ModeSetgid":     os.ModeSetgid,
	"ModeCharDevice": os.ModeCharDevice,
	"ModeSticky":     os.ModeSticky,
	"ModeIrregular":  os.ModeIrregular,
}

var mOsPermissionLwrCaseStringToCode = map[string]os.FileMode{
	"none":           os.FileMode(0),
	"modedir":        os.ModeDir,
	"modeappend":     os.ModeAppend,
	"modeexclusive":  os.ModeExclusive,
	"modetemporary":  os.ModeTemporary,
	"modesymlink":    os.ModeSymlink,
	"modedevice":     os.ModeDevice,
	"modenamedpipe":  os.ModeNamedPipe,
	"modesocket":     os.ModeSocket,
	"modesetuid":     os.ModeSetuid,
	"modesetgid":     os.ModeSetgid,
	"modechardevice": os.ModeCharDevice,
	"modesticky":     os.ModeSticky,
	"modeirregular":  os.ModeIrregular,
}

var mOsPermissionLetterToCode = map[string]os.FileMode{
	"-": os.FileMode(0),
	"d": os.ModeDir,
	"a": os.ModeAppend,
	"l": os.ModeExclusive,
	"T": os.ModeTemporary,
	"L": os.ModeSymlink,
	"D": os.ModeDevice,
	"p": os.ModeNamedPipe,
	"S": os.ModeSocket,
	"u": os.ModeSetuid,
	"g": os.ModeSetgid,
	"c": os.ModeCharDevice,
	"t": os.ModeSticky,
	"?": os.ModeIrregular,
}

// An enumeration of the os File Mode constant values:
//
//   Method           os.FileMode           Associated
//    Name             Constant             Letter Code
//  ______________________________________________________________________
//  ModeNone()        os.ModeNone           "-" is a file
//  ModeDir()         os.ModeDir            "d" is a directory
//  ModeAppend()      os.ModeAppend         "a" append-only
//  ModeExclusive()   os.ModeExclusive      "l" exclusive use
//  ModeTemporary()   os.ModeTemporary      "T" temporary file; Plan 9 only
//  ModeSymlink()     os.ModeSymlink        "L" symbolic link
//  ModeDevice()      os.ModeDevice         "D" device file
//  ModeNamedPipe()   os.ModeNamedPipe      "p" named pipe (FIFO)
//  ModeSocket()      os.ModeSocket         "S" Unix domain socket
//  ModeSetuid()      os.ModeSetuid         "u" setuid
//  ModeSetgid()      os.ModeSetgid         "g" setgid
//  ModeCharDevice()  os.ModeCharDevice     "c" Unix character device, when ModeDevice is set
//  ModeSticky()      os.ModeSticky         "t" sticky
//  ModeIrregular()   os.ModeIrregular      "?" non-regular file; nothing else is known about this file
//
// For more information on os Mode Constants Reference:
// https://golang.org/pkg/os/#pkg-constants
//
type OsFilePermissionCode os.FileMode

// ModeNone            "-" No Permission Set
func (osPerm OsFilePermissionCode) ModeNone() os.FileMode { return os.FileMode(0) }

// ModeDir         Letter Code= "d" is a directory   - alias for os.ModeDir
func (osPerm OsFilePermissionCode) ModeDir() os.FileMode { return os.ModeDir }

// ModeAppend       Letter Code= "a" append-only     - alias for os.ModeAppend
func (osPerm OsFilePermissionCode) ModeAppend() os.FileMode { return os.ModeAppend }

// ModeExclusive    Letter Code= "l" exclusive use   - alias for os.ModeExclusive
func (osPerm OsFilePermissionCode) ModeExclusive() os.FileMode { return os.ModeExclusive }

// ModeTemporary    Letter Code= "T" temporary file; Plan 9 only  - alias for os.ModeTemporary
func (osPerm OsFilePermissionCode) ModeTemporary() os.FileMode { return os.ModeTemporary }

// ModeSymlink      Letter Code= "L" symbolic link   - alias for os.ModeSymlink
func (osPerm OsFilePermissionCode) ModeSymlink() os.FileMode { return os.ModeSymlink }

// ModeDevice       Letter Code= "D" device file     - alias for os.ModeDevice
func (osPerm OsFilePermissionCode) ModeDevice() os.FileMode { return os.ModeDevice }

// ModeNamedPipe    Letter Code= "p" named pipe (FIFO) - alias for os.ModeNamedPipe
func (osPerm OsFilePermissionCode) ModeNamedPipe() os.FileMode { return os.ModeNamedPipe }

// ModeSocket       Letter Code= "S" Unix domain socket - alias for os.ModeSocket
func (osPerm OsFilePermissionCode) ModeSocket() os.FileMode { return os.ModeSocket }

// ModeSetuid       Letter Code= "u" setuid            - alias for os.ModeSetuid
// When the setuid bit is used, the behavior described above it's modified so that
// when an executable is launched, it does not run with the privileges of the user
// who launched it, but with that of the file owner instead. So, for example, if an
// executable has the setuid bit set on it, and it's owned by root, when launched by
// a normal user, it will run with root privileges. It should be clear why this represents
// a potential security risk, if not used correctly.
//
func (osPerm OsFilePermissionCode) ModeSetuid() os.FileMode { return os.ModeSetuid }

// ModeSetgid       Letter Code= "g" setgid            - alias for os.ModeSetgid
// Unlike the setuid bit, the setgid bit has effect on both files and directories.
// In the first case, the file which has the setgid bit set, when executed, instead
// of running with the privileges of the group of the user who started it, runs with
// those of the group which owns the file: in other words, the group ID of the process
// will be the same of that of the file.
//
// When used on a directory, instead, the setgid bit alters the standard behavior so that the
// group of the files created inside said directory, will not be that of the user who created
// them, but that of the parent directory itself. This is often used to ease the sharing of
// files (files will be modifiable by all the users that are part of said group).
//
func (osPerm OsFilePermissionCode) ModeSetgid() os.FileMode { return os.ModeSetgid }

// ModeCharDevice   Letter Code= "c" Unix character device, when ModeDevice is set
// alias for os.ModeCharDevice
//
func (osPerm OsFilePermissionCode) ModeCharDevice() os.FileMode { return os.ModeCharDevice }

// ModeSticky       Letter Code= "t" sticky            - alias for os.ModeSticky
// The sticky bit works in a different way: while it has no effect on files, when used on a directory,
// all the files in said directory will be modifiable only by their owners. A typical case in which
// it is used, involves the /tmp directory. Typically this directory is writable by all users on the
// system, so to make impossible for one user to delete the files of another one.
//
func (osPerm OsFilePermissionCode) ModeSticky() os.FileMode { return os.ModeSticky }

// ModeIrregular    Letter Code= "?" non-regular file; nothing else is known about this file
// alias for os.ModeIrregular
func (osPerm OsFilePermissionCode) ModeIrregular() os.FileMode { return os.ModeIrregular }

// Equal - Compares the current OsFilePermissionCode instance to another
// OsFilePermission instance passed as an input parameter. If the two are
// equal in all respects, this method returns 'true'.
//
func (osPerm OsFilePermissionCode) Equal(osPerm2 OsFilePermissionCode) bool {

	if osPerm == osPerm2 {
		return true
	}

	return false
}

// GetFileModeLetterCode - Returns the single alphabetic character associated with
// this os.FileMode. All os.FileMode's are associated with a single letter used
// in unix permission strings.
//                         Letter
//    File Mode             Code     Description
//    __________________________________________
//    os.ModeNone           "-"      is a file
//    os.ModeDir            "d"      is a directory
//    os.ModeAppend         "a"      append-only
//    os.ModeExclusive      "l"      exclusive use
//    os.ModeTemporary      "T"      temporary file; Plan 9 only
//    os.ModeSymlink        "L"      symbolic link
//    os.ModeDevice         "D"      device file
//    os.ModeNamedPipe      "p"      named pipe (FIFO)
//    os.ModeSocket         "S"      Unix domain socket
//    os.ModeSetuid         "u"      setuid
//    os.ModeSetgid         "g"      setgid
//    os.ModeCharDevice     "c"      Unix character device, when ModeDevice is set
//    os.ModeSticky         "t"      sticky
//    os.ModeIrregular      "?"      non-regular file; nothing else is known about this file
//
func (osPerm OsFilePermissionCode) GetFileModeLetterCode() (string, error) {

	ePrefix := "OsFilePermissionCode.FileModeLetterCode() "

	letter, ok := mOsPermissionCodeToLetter[os.FileMode(osPerm)]

	if !ok {
		return "", fmt.Errorf(ePrefix+
			"The current OsFilePermissionCode FileMode value is INVALID! "+
			"OsFilePermissionCode decimal value=%s  octal value=%s",
			strconv.FormatInt(int64(osPerm), 10),
			strconv.FormatInt(int64(osPerm), 8))
	}

	return letter, nil
}

// GetNewFromFileMode - Creates and returns a new OsFilePermissionCode instance
// generated from the os.FileMode type input parameter ('fMode'). If the input
// os.FileMode value is invalid, an error is returned.
func (osPerm OsFilePermissionCode) GetNewFromFileMode(
	fMode os.FileMode) (OsFilePermissionCode, error) {

	newFilePerm := OsFilePermissionCode(fMode)

	err := newFilePerm.IsValid()

	if err != nil {
		ePrefix := "OsFilePermissionCode.GetNewFromFileMode() "
		return OsFilePermissionCode(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'fMode' is an INVALID File Mode!")
	}

	return newFilePerm, nil
}

// GetNewFromLetterCode - Creates a new OsFilePermissionCode instance based on an
// associated 'letter code'. The letter code consists of a single character
// representing an os.FileMode. This single character is useful in configuring
// unix permission strings.
//
func (osPerm OsFilePermissionCode) GetNewFromLetterCode(
	letterCode string) (OsFilePermissionCode, error) {

	fModeValue, ok := mOsPermissionLetterToCode[letterCode]

	if !ok {
		ePrefix := "OsFilePermissionCode.GetNewFromLetterCode() "
		return OsFilePermissionCode(0), fmt.
			Errorf(ePrefix+"Error: 'letterCode' is INVALID! "+
				"letterCode='%v'", letterCode)
	}

	return OsFilePermissionCode(fModeValue), nil
}

// IsValid - If the value of the current OsFilePermissionCode is 'invalid',
// this method will return an error. If the OsFilePermissionCode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (osPerm OsFilePermissionCode) IsValid() error {

	_, ok := mOsPermissionCodeToString[os.FileMode(osPerm)]

	if !ok {
		ePrefix := "OsFilePermissionCode.IsValid() "
		return fmt.Errorf(ePrefix+"The current OsFilePermissionCode is INVALID! "+
			"OsFilePermissionCode Octal Value='%s'", strconv.FormatInt(int64(osPerm), 8))
	}

	return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of OsFilePermissionCode is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'modedir' will NOT
//	                       match the enumeration name, 'ModeDir'.
//
//	                       If 'false' a case insensitive search is conducted
//	                       for the enumeration name. In this case, 'modedir'
//	                       will match match enumeration name 'ModeDir'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	OsFilePermissionCode - Upon successful completion, this method will return a new
//	                       instance of OsFilePermissionCode set to the value of the
//	                       enumeration matched by the string search performed on
//	                       input parameter,'valueString'.
//
//	error                - If this method completes successfully, the returned error
//	                       Type is set equal to 'nil'. If an error condition is encountered,
//	                       this method will return an error Type which encapsulates an
//	                       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	t, err := OsFilePermissionCode(0).ParseString("ModeDir")
//
//	    t is now equal to OsFilePermissionCode(0).ModeDir()
//
func (osPerm OsFilePermissionCode) ParseString(
	valueString string,
	caseSensitive bool) (OsFilePermissionCode, error) {

	ePrefix := "OsFilePermissionCode.ParseString() "

	var permCode os.FileMode
	var ok bool

	if !caseSensitive {
		permCode, ok = mOsPermissionLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return OsFilePermissionCode(0), errors.New(ePrefix + "Invalid Permission Code!")
		}

		return OsFilePermissionCode(permCode), nil
	}

	permCode, ok = mOsPermissionStringToCode[valueString]

	if !ok {
		return OsFilePermissionCode(0), errors.New(ePrefix + "Invalid Permission Code!")
	}

	return OsFilePermissionCode(permCode), nil

}

// FilePermCode - public global variable of type OsFilePermissionCode.
// Provides alternative, easier access to OsFilePermissionCode enumeration
// values.
//
// Usage:
//
//	FilePermCode.ModeNone()
//	FilePermCode.ModeDir()
//
var FilePermCode = OsFilePermissionCode(0)

// String - Returns a string with the name of the enumeration associated
// with this instance of 'OsFilePermissionCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= OsFilePermissionCode(0).ModeDir()()
//	str := t.String()
//	    str is now equal to "ModeDir()"
//
func (osPerm OsFilePermissionCode) String() string {

	label, ok := mOsPermissionCodeToString[os.FileMode(osPerm)]

	if !ok {
		return ""
	}

	return label
}

// Value - Returns the value of the OsFilePermissionCode instance
// as type os.FileMode.
//
func (osPerm OsFilePermissionCode) Value() os.FileMode {

	return os.FileMode(osPerm)
}

// FilePermissionConfig - Provides methods to support the creation and management of
// of File Permissions for using in controlling file operations. The Go Programming
// Language uses os.FileMode (https://golang.org/pkg/os/#FileMode) and unix permission
// bits to configure file permissions.
//     Reference:
//     https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//     https://en.wikipedia.org/wiki/File_system_permissions
//
// The FilePermissionConfig methods will allow for configuration of valid file permissions
// which are subsequently stored as an os.FileMode in a private member variable,
// 'FilePermissionConfig.fileMode'.
//
type FilePermissionConfig struct {
	isInitialized bool
	fileMode      os.FileMode
}

// CopyIn - Receives a FilePermissionConfig instance and copies all
// data fields to the current FilePermissionConfig instance. When
// complete, both the incoming and current FilePermissionConfig
// instances will be identical. The type of copy operation performed
// is a 'deep copy'.
//
func (fPerm *FilePermissionConfig) CopyIn(fPerm2 FilePermissionConfig) {

	fPerm.isInitialized = fPerm2.isInitialized
	fPerm.fileMode = fPerm2.fileMode

}

// CopyOut - Returns a new instance of FilePermissionConfig which is
// in all respects an exact duplicate of the current FilePermissionConfig
// instance. The type of copy operation performed  is a 'deep copy'.
//
func (fPerm *FilePermissionConfig) CopyOut() FilePermissionConfig {

	fPerm2 := FilePermissionConfig{}

	fPerm2.isInitialized = fPerm.isInitialized
	fPerm2.fileMode = fPerm.fileMode

	return fPerm2
}

// Empty - ReInitializes the current FilePermissionConfig instance to
// empty or zero values.
//
func (fPerm *FilePermissionConfig) Empty() {
	fPerm.isInitialized = false
	fPerm.fileMode = os.FileMode(0)
}

// Equal - Returns 'true' if the incoming FilePermissionConfig instance
// is equal in all respects to the current FilePermissionConfig instance.
//
func (fPerm *FilePermissionConfig) Equal(fPerm2 FilePermissionConfig) bool {

	if fPerm.isInitialized != fPerm2.isInitialized {
		return false
	}

	if fPerm.fileMode != fPerm2.fileMode {
		return false
	}

	return true
}

// New - Creates and returns a new FilePermissionConfig instance initialized with a
// an os.FileMode value generated from the input parameter 'modeStr'.
//
// 'modeStr' is a 10-character string containing the read, write and execute permissions
// for the the three groups, 'Owner', 'Group' and 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix permission codes.
//
//  Reference:
//    https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//    https://en.wikipedia.org/wiki/File_system_permissions
//
// The first character of the 'modeStr' designates the 'Entry Type'. Currently,
// only two 'Entry Type' characters are supported. Therefore, the first character
// in the 10-character input parameter 'modeStr' MUST be either a "-" indicating
// a file, or a "d" indicating a directory.
//
// The remaining nine characters in the 'modeStr' represent unix permission bits
// and consist of three group fields each containing 3-characters. Each character
// in the three group fields may be consist of 'r' (Read-Permission), 'w'
// (Write-Permission), 'x' (Execute-Permission) or '-' signaling no permission
// or no access allowed. A typical 'modeStr' authorizing permission for full access
// to a file would be styled as:
//
//  Example: "-rwxrwxrwx"
//
//  Groups: - Owner/User, Group, Other
//  From left to right
//  First Characters is Entry Type index 0 ("-")
//
//  First Char index 0 =      "-"    Designates a file
//  Char indexes 1-3 = Owner  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Owner'
//  Char indexes 4-6 = Group  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Group'
//  Char indexes 7-9 = Other  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Other'
//
// The Symbolic notation provided by input parameter 'modeStr' MUST conform to
// the options presented below. The first character or 'Entry Type' is listed as
// "-". However, in practice, the caller may set the first character as either a
// "-", specifying a file, or a "d", specifying a directory. No other first character
// types are currently supported.
//
// Three SymbolicGroups:
// The three group types are: Owners, Groups & Others.
//
//   'modeStr'
//   Symbolic      Octal      File Access
//   Notation      Notation   Permission Descriptions
//   ----------    0000       no permissions
//   -rwx------    0700       read, write, & execute only for owner
//   -rwxrwx---    0770       read, write, & execute for owner and group
//   -rwxrwxrwx    0777       read, write, & execute for owner, group and others
//   ---x--x--x    0111       execute
//   --w--w--w-    0222       write
//   --wx-wx-wx    0333       write & execute
//   -r--r--r--    0444       read
//   -r-xr-xr-x    0555       read & execute
//   -rw-rw-rw-    0666       read & write
//   -rwxr-----    0740       Owner can read, write, & execute. Group can only read;
//                               others have no permissions
//
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//   modeStr  string - 'modeStr' must conform to the symbolic notation options shown
//                     above. Failure to comply with this requirement will generate an
//                     error. As indicated, 'modeStr' must consist of 10-characters.
//                     The first character in 'modeStr' may be '-' specifying a fle or
//                     'd' specifying a directory.
//
//                     Reference:
//                     How to use special permissions: the setuid, setgid and sticky bits
//                     https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
func (fPerm FilePermissionConfig) New(modeStr string) (FilePermissionConfig, error) {

	ePrefix := "FilePermissionConfig.New()"

	fPerm2 := FilePermissionConfig{}

	err := fPerm2.SetFileModeByTextCode(modeStr)

	if err != nil {
		return FilePermissionConfig{},
			fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fPerm2, nil
}

// NewByComponents - Creates and returns a new instance of FilePermissionConfig using
// two input parameters, 'entryType' and 'unixPermissionTextStr'.
//
// For additional documentation see method FilePermissionConfig.SetFileModeByComponents()
// which is called by this method.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  entryType OsFilePermissionCode - The code which makes up the first character in
//                                   a 10-digit unix permission character string.
//                                   This a wrapper for os.FileMode constants.
//                                     Reference:
//                                        https://golang.org/pkg/os/#FileMode
//
//                                   Select this value with caution. See the warning below.
//
//  unixPermissionTextStr string - A 9-character string containing the unix permission
//                            bits expressed as three groups of 3-characters each.
//
//                             The 9-characters are constituents of the the three Symbolic
//                             Groups: Owners/Users, Groups & Others. Each group has three
//                             characters which may be 'r', 'w', 'x'. If a permission is not
//                             set, that character position contains a '-'.
//
//   'unixPermissionTextStr'
//        9-Character          File Access
//        Notation             Permission Descriptions
//        ---------            File - no permissions
//        rwx------            File - read, write, & execute only for owner
//        rwxrwx---            File - read, write, & execute for owner and group
//        rwxrwxrwx            File - read, write, & execute for owner, group and others
//        --x--x--x            File - execute
//        -w--w--w-            File - write
//        -wx-wx-wx            File - write & execute
//        r--r--r--            File - read
//        r-xr-xr-x            File - read & execute
//        rw-rw-rw-            File - read & write
//        rwxr-----            File - Owner can read, write, & execute. Group can only read;
//
// ------------------------------------------------------------------------
//
// Warning:
//
// Incorrect or invalid File Permissions can cause extensive damage. If you
// don't know what you are doing, you would be well advised to use one of
// the other methods in this type which provide additional safeguards.
//
// If you decide to proceed, be guided by the wisdom of Davy Crockett:
//
//        "Be always sure you are right - then go ahead."
//
func (fPerm FilePermissionConfig) NewByComponents(
	entryType OsFilePermissionCode,
	unixPermissionTextStr string) (FilePermissionConfig, error) {

	fPerm2 := FilePermissionConfig{}

	err := fPerm2.SetFileModeByComponents(entryType, unixPermissionTextStr)

	if err != nil {
		ePrefix := "FilePermissionConfig.NewByComponents() "
		return FilePermissionConfig{},
			fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fPerm2, nil
}

// NewByFileMode - Creates and returns a new instance of FilePermissionConfig. The instance
// is initialized using the input parameter 'fMode' of type 'os.FileMode'.  'fMode' is assumed
// to contain all of the codes necessary for the configuration of unix file permission bits.
//
func (fPerm FilePermissionConfig) NewByFileMode(fMode os.FileMode) (FilePermissionConfig, error) {

	fPerm2 := FilePermissionConfig{}

	err := fPerm2.SetFileModeByFileMode(fMode)

	if err != nil {

		ePrefix := "FilePermissionConfig.NewByFileMode() "

		return FilePermissionConfig{},
			fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fPerm2, nil
}

// NewByOctalDigits - Creates and returns a new FilePermissionConfig instance by
// initializing the internal FileMode data field (FilePermissionConfig.fileMode)
// to the value represented by input parameter, 'octalFileModeCode'.
//
// Note: This method calls FilePermissionConfig.SetFileModeByOctalDigits().
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  octalFileModeCode int - This parameter contains the integer value of the
//                          of the permission code which will be used to
//                          initialize the current FilePermissionConfig instance
//                          (FilePermissionConfig.fileMode). The integer digits
//                          in 'octalFileModeCode' represent the octal value
//                          for the file permission as indicated by the following
//                          examples.
//   ____________________________________________________________________________
//
//            Input Parameter
//                integer            Equivalent
//   Octal    'octalFileModeCode'    Symbolic      File Access
//   Digits        value             Notation      Permission Descriptions
//   0000 	         0               ----------    File - no permissions
//   0700 	       700               -rwx------    File - read, write, & execute only for owner
//   0770 	       770               -rwxrwx---    File - read, write, & execute for owner and group
//   0777 	       777               -rwxrwxrwx    File - read, write, & execute for owner, group and others
//   0111 	       111               ---x--x--x    File - execute
//   0222 	       222               --w--w--w-    File - write
//   0333 	       333               --wx-wx-wx    File - write & execute
//   0444 	       444               -r--r--r--    File - read
//   0555 	       555               -r-xr-xr-x    File - read & execute
//   0666 	       666               -rw-rw-rw-    File - read & write
//   0740 	       740               -rwxr-----    File - Owner can read, write, & execute. Group can only read;
//                                                        others have no permissions
//
//   20000000777   20000000777       drwxrwxrwx    Directory - read, write, & execute for owner, group and others
//
//   See method FilePermissionConfig.SetFileModeByTextCode() for more documentation
//
// ------------------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an integer with a leading
// zero (e.g. x:= int(0777)), than number ('0777') is treated as an octal value
// and converted to a decimal value. Therefore, x:= int(0777) will mean that 'x'
// is set equal to 511. If you set x:= int(777), x will be set equal to '777'.
// For purposes of this method enter the octal code as x:= int(777).
//
func (fPerm FilePermissionConfig) NewByOctalDigits(
	octalFileModeCode int) (FilePermissionConfig, error) {

	fPerm2 := FilePermissionConfig{}

	err := fPerm2.SetFileModeByOctalDigits(octalFileModeCode)

	if err != nil {

		ePrefix := "FilePermissionConfig.NewByFileMode() "

		return FilePermissionConfig{},
			fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fPerm2, nil
}

// GetFileMode - Returns the os.FileMode from the internal data field,
// 'FilePermissionConfig.fileMode'.
//
func (fPerm *FilePermissionConfig) GetFileMode() (os.FileMode, error) {

	ePrefix := "FilePermissionConfig.GetFileMode() "

	if !fPerm.isInitialized {
		return os.FileMode(0),
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. The FileMode is INVALID!")
	}

	return fPerm.fileMode, nil
}

// GetIsDir - Return a bool indicating whether the encapsulated FileMode is a directory
// or not. A returned value of 'true' signals that the FileMode represents a directory.
//
// This method serves as a wrapper for os.FileMode.IsDir()
//
func (fPerm *FilePermissionConfig) GetIsDir() (bool, error) {

	ePrefix := "FilePermissionConfig.GetIsDir() "

	if !fPerm.isInitialized {
		return false,
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. The FileMode is INVALID!")
	}

	return fPerm.fileMode.IsDir(), nil
}

// GetEntryTypeComponent - Returns the 'Entry Type' component of the current os.FileMode
// permissions value. The 'Entry Type' is the first character in a 10-character permissions
// text string. For the majority of applications, the leading character in a 10-character
// permissions text string is either a hyphen ('-') indicating the subject is a file - or -
// a 'd' indicating the subject is a directory. For a file, the File Mode Entry Type value is
// zero ('0').  For a directory, the File Mode Entry Type value is equal to 'os.ModDir'.
//
//
// ------------------------------------------------------------------------
//
// Return Values:
//
// OsFilePermissionCode - The OsFilePermissionCode type is set to the value of
//                        the os.FileMode constant representing the Entry Type
//                        associated with the permission value encapsulated by
//                        this FilePermissionConfig instance.
//
func (fPerm *FilePermissionConfig) GetEntryTypeComponent() (OsFilePermissionCode, error) {

	ePrefix := "FilePermissionConfig.GetEntryTypeComponent() "

	if !fPerm.isInitialized {
		return OsFilePermissionCode(0),
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. The FileMode is INVALID!")
	}

	mask := os.FileMode(0777)

	for idx := range mOsPermissionCodeToString {

		fMode := fPerm.fileMode &^ mask

		if fMode == idx {

			return OsFilePermissionCode(idx), nil

		}

	}

	return OsFilePermissionCode(0),
		fmt.Errorf(ePrefix + "The Entry Type for this FilePermissionConfig instance is INVALID!")
}

// GetIsRegular - Return a bool indicating whether the encapsulated FileMode is a file
// or not. A returned value of 'true' signals that the FileMode represents a file.
//
// This method serves as a wrapper for os.FileMode.IsRegular()
func (fPerm *FilePermissionConfig) GetIsRegular() (bool, error) {

	ePrefix := "FilePermissionConfig.GetIsRegular() "

	if !fPerm.isInitialized {
		return false,
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. The FileMode is INVALID!")
	}

	return fPerm.fileMode.IsRegular(), nil
}

// GetPermissionBits - Return a FileMode containing only the least significant 9-bits of
// the encapsulated FileMode representing the unix permission bits.
//
// The actual string returned will contains 10-characters, have the first character
// (index=0) will always be a hyphen ("-"). The hyphen ("-") generally indicates a
// file, however it should be ignored in this case. The only valid a reliable unix
// permission bits are in the last 9-characters (indexes 1-8). When evaluating permission
// bits returned by this method always ignore the first character which will always
// be a hyphen ("-").
//
// To acquire the full and valid 10-digit permission string use method
// FilePermissionConfig.GetPermissionTextCode() documented below.
//
func (fPerm *FilePermissionConfig) GetPermissionBits() (os.FileMode, error) {

	ePrefix := "FilePermissionConfig.GetPermissionBits() "

	if !fPerm.isInitialized {
		return os.FileMode(0),
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. " +
				"The FileMode is INVALID!")
	}

	return fPerm.fileMode.Perm(), nil
}

// GetPermissionTextCode - Returns the file mode permissions expressed as
// a text string. The returned string includes the full and complete
// 10-character permission code.
//
func (fPerm *FilePermissionConfig) GetPermissionTextCode() (string, error) {
	ePrefix := "FilePermissionConfig.GetPermissionBits() "

	if !fPerm.isInitialized {
		return "",
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. " +
				"The FileMode is INVALID!")
	}

	return fPerm.fileMode.String(), nil
}

// GetPermissionComponents - Returns the two components of a permission configuration.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//  entryType OsFilePermissionCode - The Entry Type or os mode value. Generally this will either be
//                                   OsFilePermissionCode(0).ModeNone() for files
//                                                   or
//                                   OsFilePermissionCode(0).ModeDir() for directories.
//
//                                   For more information see method FilePermissionConfig.GetEntryTypeComponent()
//
//  permissionBits  os.FileMode - The 9-least significant bits designate the unix
//                                permission bits.
//
//                                Be advised that if you call string on this result
//                                (permissionBits.String()) you will receive a 10-character
//                                string the first character of which is always a hyphen ("-").
//                                Disregard this first character, only the last 9-characters of
//                                the string are valid permission descriptors. For more information
//                                see method FilePermissionConfig.GetPermissionBits()
//
//                                To create a full and complete permission code, permissionBits must
//                                be or'd with a valid Entry Type, os mode value.
//
func (fPerm *FilePermissionConfig) GetPermissionComponents() (
	osMode OsFilePermissionCode, permissionBits os.FileMode, err error) {

	osMode = OsFilePermissionCode(FilePermCode.ModeNone())

	permissionBits = os.FileMode(0)

	err = nil

	ePrefix := "FilePermissionConfig.GetPermissionComponents() "

	if !fPerm.isInitialized {
		err =
			errors.New(ePrefix +
				"Error: This FilePermissionConfig instance has NOT been initialized. " +
				"The FileMode is INVALID!")

		return osMode, permissionBits, err
	}

	var err2 error

	osMode, err2 = fPerm.GetEntryTypeComponent()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return osMode, permissionBits, err
	}

	permissionBits, err2 = fPerm.GetPermissionBits()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return osMode, permissionBits, err
	}

	err = nil

	return osMode, permissionBits, err
}

// IsValid - If the current FilePermissionConfig instance is judged to be
// 'Invalid', this method will return an error.
//
// Otherwise, if the current instance of FilePermissionConfig evaluates as
// 'Valid', the method will return 'nil'.
//
func (fPerm *FilePermissionConfig) IsValid() error {

	if !fPerm.isInitialized {
		ePrefix := "FilePermissionConfig.IsValid() "
		return errors.New(ePrefix + "Error: This FilePermissionConfig instance has NOT been " +
			"initialized and is INVALID!")
	}

	return nil
}

// SetFileModeByComponents - Sets the value of the current FilePermissionConfig
// instance by initializing the internal FileMode data field
// (FilePermissionConfig.fileMode). The final FileMode value is computed by
// integrating the 'entryType' FileMode with the unix permission symbolic
// values provided by the input parameter, 'unixPermissionStr'. This approach
// allows the caller to created custom File Permissions.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//  entryType OsFilePermissionCode - The code which makes up the first character in
//                                   a 10-digit unix permission character string.
//                                   This a wrapper for os.FileMode constants.
//                                     Reference:
//                                        https://golang.org/pkg/os/#FileMode
//
//                                   Select this value with caution. See the warning below.
//
//  unixPermissionStr string - A 9-character string containing the unix permission
//                             bits expressed as three groups of 3-characters each.
//
//                             The 9-characters are constituents of the the three Symbolic
//                             Groups: Owners/Users, Groups & Others. Each group has three
//                             characters which may be 'r', 'w', 'x'. If a permission is not
//                             set, that character position contains a '-'.
//
//   'unixPermissionTextStr'
//        9-Character          File Access
//        Notation             Permission Descriptions
//        ---------            File - no permissions
//        rwx------            File - read, write, & execute only for owner
//        rwxrwx---            File - read, write, & execute for owner and group
//        rwxrwxrwx            File - read, write, & execute for owner, group and others
//        --x--x--x            File - execute
//        -w--w--w-            File - write
//        -wx-wx-wx            File - write & execute
//        r--r--r--            File - read
//        r-xr-xr-x            File - read & execute
//        rw-rw-rw-            File - read & write
//        rwxr-----            File - Owner can read, write, & execute. Group can only read;
//
// ------------------------------------------------------------------------
//
// Warning:
//
// Incorrect or invalid File Permissions can cause extensive damage. If you
// don't know what you are doing, you would be well advised to use one of
// the other methods in this type which provide additional safeguards.
//
// If you decide to proceed, be guided by the wisdom of Davy Crockett:
//
//        "Be always sure you are right - then go ahead."
//
func (fPerm *FilePermissionConfig) SetFileModeByComponents(
	entryType OsFilePermissionCode, unixPermissionTextStr string) error {

	ePrefix := "FilePermissionConfig.SetFileModeByComponents() "

	if len(unixPermissionTextStr) != 9 {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'unixPermissionTextStr' must contain 9-Characters. "+
			"This unixPermissionTextStr contains %v-characters. unixPermissionTextStr='%v'. ",
			len(unixPermissionTextStr), unixPermissionTextStr)
	}

	fModeEntryType := os.FileMode(entryType)

	_, ok := mOsPermissionCodeToString[fModeEntryType]

	if !ok {
		return fmt.Errorf(ePrefix+
			"Input parameter 'entryType' is an INVALID os.FileMode! entryType decimal value='%s' "+
			"octal value='%s' ", strconv.FormatInt(int64(entryType), 10),
			strconv.FormatInt(int64(entryType), 8))
	}

	ownerInt, err := fPerm.convertGroupToDecimal(unixPermissionTextStr[0:3], "owner")

	if err != nil {
		return fmt.Errorf(ePrefix+"'ownerInt' Error: %v", err.Error())
	}

	groupInt, err := fPerm.convertGroupToDecimal(unixPermissionTextStr[3:6], "group")

	if err != nil {
		return fmt.Errorf(ePrefix+"groupInt Error: %v", err.Error())
	}

	otherInt, err := fPerm.convertGroupToDecimal(unixPermissionTextStr[6:], "other")

	if err != nil {
		return fmt.Errorf(ePrefix+"otherInt Error: %v", err.Error())
	}

	ownerInt *= 100
	groupInt *= 10
	permission := ownerInt + groupInt + otherInt

	fMode := os.FileMode(FileHelper{}.ConvertOctalToDecimal(permission))

	fPerm.fileMode = fModeEntryType | fMode
	fPerm.isInitialized = true

	return nil
}

// SetFileModeByFileMode - Sets the permission codes for this FilePermissionConfig
// instance using an input parameter of type 'os.FileMode'. If the value does not
// include a valid os mode constant, and error will be returned.
//
// If successful, this method will assign the os.FileMode input value to the internal
// data field, 'FilePermissionConfig.fileMode'.
//
func (fPerm *FilePermissionConfig) SetFileModeByFileMode(fMode os.FileMode) error {

	tFMode := fMode

	mask := os.FileMode(0777)

	entryType := tFMode &^ mask

	_, ok := mOsPermissionCodeToString[entryType]

	if !ok {
		ePrefix := "FilePermissionConfig.SetFileModeByFileMode() "
		return fmt.Errorf(ePrefix + "Error: Input parameter 'fMode' contains an invalid " +
			"'EntryType' otherwise known as an os mode constant.")
	}

	fPerm.fileMode = fMode
	fPerm.isInitialized = true

	return nil
}

// SetFileModeByOctalDigits - Sets the value of the current FilePermissionConfig
// instance by initializing the internal FileMode data field
// (FilePermissionConfig.fileMode) to the value represented by input parameter,
// 'octalFileModeCode'. Any previous internal FileMode value is overwritten.
//
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//  octalFileModeCode int - This parameter contains the integer value of the
//                          of the permission code which will be used to
//                          initialize the current FilePermissionConfig instance
//                          (FilePermissionConfig.fileMode). The integer digits
//                          in 'octalFileModeCode' represent the octal value
//                          for the file permission as indicated by the following
//                          examples.
//   ____________________________________________________________________________
//
//            Input Parameter
//                integer            Equivalent
//   Octal    'octalFileModeCode'    Symbolic      File Access
//   Digits        value             Notation      Permission Descriptions
//   0000 	         0               ----------    File - no permissions
//   0700 	       700               -rwx------    File - read, write, & execute only for owner
//   0770 	       770               -rwxrwx---    File - read, write, & execute for owner and group
//   0777 	       777               -rwxrwxrwx    File - read, write, & execute for owner, group and others
//   0111 	       111               ---x--x--x    File - execute
//   0222 	       222               --w--w--w-    File - write
//   0333 	       333               --wx-wx-wx    File - write & execute
//   0444 	       444               -r--r--r--    File - read
//   0555 	       555               -r-xr-xr-x    File - read & execute
//   0666 	       666               -rw-rw-rw-    File - read & write
//   0740 	       740               -rwxr-----    File - Owner can read, write, & execute. Group can only read;
//                                                        others have no permissions
//
//   20000000777   20000000777       drwxrwxrwx    Directory - read, write, & execute for owner, group and others
//
//   See method FilePermissionConfig.SetFileModeByTextCode() for more documentation
//
// ------------------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an integer with a leading
// zero (e.g. x:= int(0777)), than number ('0777') is treated as an octal value
// and converted to a decimal value. Therefore, x:= int(0777) will mean that 'x'
// is set equal to 511. If you set x:= int(777), x will be set equal to '777'.
// For purposes of this method enter the octal code as x:= int(777).
//
func (fPerm *FilePermissionConfig) SetFileModeByOctalDigits(octalFileModeCode int) error {

	decimalVal := FileHelper{}.ConvertOctalToDecimal(octalFileModeCode)

	fPerm.fileMode = os.FileMode(decimalVal)
	fPerm.isInitialized = true

	return nil
}

// SetFileModeByTextCode - Sets the internal FileMode data field using input
// parameter 'modeStr'. 'modeStr' is a 10-character string containing the read,
// write and execute permissions for the the three groups, 'Owner/User', 'Group' and
// 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix permission codes.
//
//   Reference:
//     https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//     https://en.wikipedia.org/wiki/File_system_permissions
//
// The first character of the 'modeStr' designates the 'Entry Type'. Currently,
// only two 'Entry Type' characters are supported. Therefore, the first character
// in 'modeStr' must consist of a hyphen ("-") designating a file, or a "d"
// designating a directory.
//
// The remaining nine characters in the 'modeStr' are styled as unix permission bits.
// These nine characters are divided into three group fields each containing 3-permission
// characters. Each character field may be populated with a 'r' (Read-Permission),
// 'w' (Write-Permission), 'x' (Execute-Permission) or '-' signaling no permission
// or no access allowed. A typical 'modeStr' authorizing permission for full access
// to a file would be styled as:
//
//   "-rwxrwxrwx"
//
//   Groups: - Owner, Group, Other
//   From left to right
//
//   Char index 0     = Entry Type. Must be either a "-" or a "d"
//   Char indexes 4-6 = Group  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Group'
//   Char indexes 1-3 = Owner  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Owner'
//   Char indexes 7-9 = Other  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Other'
//
// The Symbolic notation provided by input parameter 'modeStr' MUST conform to
// the options presented below. The first character or 'Entry Type' is listed as
// "-". However, in practice, the caller may set the first character as either a
// "-", specifying a file, or a "d", specifying a directory. No other first character
// types are currently supported.
//
// After the first character, the remaining 9-characters are constituents of the the
// three Symbolic Groups: Owners/Users, Groups & Others. Each group has three characters
// which may be 'r', 'w', 'x'. If a permission is not set, the character position contains
// a '-'.
//
//   'modeStr'
//   Symbolic    Octal           File Access
//   Notation    Notation        Permission Descriptions
//   ----------   0000           File - no permissions
//   -rwx------   0700           File - read, write, & execute only for owner
//   -rwxrwx---   0770           File - read, write, & execute for owner and group
//   -rwxrwxrwx   0777           File - read, write, & execute for owner, group and others
//   ---x--x--x   0111           File - execute
//   --w--w--w-   0222           File - write
//   --wx-wx-wx   0333           File - write & execute
//   -r--r--r--   0444           File - read
//   -r-xr-xr-x   0555           File - read & execute
//   -rw-rw-rw-   0666           File - read & write
//   -rwxr-----   0740           File - Owner can read, write, & execute. Group can only read;
//                               File - others have no permissions
//   drwxrwxrwx   20000000777    File - Directory - read, write, & execute for owner, group and others
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//   modeStr  string - 'modeStr' must conform to the symbolic notation options shown
//                     above. Failure to comply with this requirement will generate an
//                     error. As indicated, 'modeStr' must consist of 10-characters.
//                     The first character in 'modeStr' may be '-' specifying a fle or
//                     'd' specifying a directory.
//
//   Reference:
//   How to use special permissions: the setuid, setgid and sticky bits
//   https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
func (fPerm *FilePermissionConfig) SetFileModeByTextCode(modeStr string) error {

	ePrefix := "FilePermissionConfig.StringToMode() "

	if len(modeStr) != 10 {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'modeStr' MUST contain 10-characters. This 'modeStr' "+
			"contains %v-characters. modeStr='%v' ", len(modeStr), modeStr)
	}

	firstChar := string(modeStr[0])

	if firstChar != "-" &&
		firstChar != "d" {
		return fmt.Errorf(ePrefix+
			"Error: First character of input parameter, 'modeStr' MUST BE 'd' or '-'. "+
			"This first character = '%v'", firstChar)
	}

	ownerInt, err := fPerm.convertGroupToDecimal(modeStr[1:4], "owner")

	if err != nil {
		return fmt.Errorf(ePrefix+"'ownerInt' Error: %v", err.Error())
	}

	groupInt, err := fPerm.convertGroupToDecimal(modeStr[4:7], "group")

	if err != nil {
		return fmt.Errorf(ePrefix+"groupInt Error: %v", err.Error())
	}

	otherInt, err := fPerm.convertGroupToDecimal(modeStr[7:], "other")

	if err != nil {
		return fmt.Errorf(ePrefix+"otherInt Error: %v", err.Error())
	}

	ownerInt *= 100
	groupInt *= 10
	permission := ownerInt + groupInt + otherInt

	entryType := 0

	fMode := permission

	fh := FileHelper{}

	if firstChar == "d" {
		entryType = fh.ConvertDecimalToOctal(int(os.ModeDir))
		fMode = entryType | permission
	}

	fPerm.fileMode = os.FileMode(fh.ConvertOctalToDecimal(fMode))
	fPerm.isInitialized = true

	return nil
}

// convertGroupToDecimal - Expecting to a receive a 3-character permission string
// for an 'owner', 'group' or 'other' 'groupType'.
//
// 3-character permission letter group must be formatted as one of the following:
//             "rwx"
//             "rw-"
//             "r--"
//             "---"
//             "--x"
//             "-wx"
//             "-w-"
//             "r-x"
//
// If input parameter 'groupStr' does not match one of the letter groups shown above, an
// error will be returned.
//
// If successful, this method will return an integer representing the octal digits comprising
// this group code. For example, groupStr="rwx" will return an integer value of '7' which can
// be treated as octal digit '7' for purposes of creating an os.FileMode.
//
func (fPerm *FilePermissionConfig) convertGroupToDecimal(groupStr, groupType string) (int, error) {

	ePrefix := "FilePermissionConfig.convertGroupToDecimal() "
	var err error
	intVal := 0

	if len(groupStr) != 3 {
		return -1, fmt.Errorf(ePrefix+
			"Error: input parameter groupStr must be exactly 3-characters in length. "+
			"This groupStr is %v-characters in length. groupStr='%v' groupType='%v' ",
			len(groupStr), groupStr, groupType)
	}

	tstLtrs := strings.ToLower(string(groupStr))

	switch tstLtrs {
	case "rwx":
		intVal = 7
	case "rw-":
		intVal = 6
	case "r--":
		intVal = 4
	case "---":
		intVal = 0
	case "--x":
		intVal = 1
	case "-wx":
		intVal = 3
	case "-w-":
		intVal = 2
	case "r-x":
		intVal = 5
	default:
		err = fmt.Errorf(ePrefix+"Error: Invalid 3-Letter "+groupType+
			" String. 3-Letter Block='%v'", tstLtrs)
	}

	return intVal, err
}

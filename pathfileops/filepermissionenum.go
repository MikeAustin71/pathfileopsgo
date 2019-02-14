package pathfileops

import (
	"fmt"
	"os"
	"strings"
)

// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
//
//type FileMode uint32

// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
/*
const (
  // The single letters are the abbreviations
  // used by the String method's formatting.
  ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
  ModeAppend                                     // a: append-only
  ModeExclusive                                  // l: exclusive use
  ModeTemporary                                  // T: temporary file; Plan 9 only
  ModeSymlink                                    // L: symbolic link
  ModeDevice                                     // D: device file
  ModeNamedPipe                                  // p: named pipe (FIFO)
  ModeSocket                                     // S: Unix domain socket
  ModeSetuid                                     // u: setuid
        When the setuid bit is used, the behavior described above it's modified so that
        when an executable is launched, it does not run with the privileges of the user
        who launched it, but with that of the file owner instead. So, for example, if an
        executable has the setuid bit set on it, and it's owned by root, when launched by
        a normal user, it will run with root privileges. It should be clear why this represents
        a potential security risk, if not used correctly.

  ModeSetgid                                     // g: setgid
        Unlike the setuid bit, the setgid bit has effect on both files and directories.
        In the first case, the file which has the setgid bit set, when executed, instead
        of running with the privileges of the group of the user who started it, runs with
        those of the group which owns the file: in other words, the group ID of the process
        will be the same of that of the file.

        When used on a directory, instead, the setgid bit alters the standard behavior so that the
        group of the files created inside said directory, will not be that of the user who created
        them, but that of the parent directory itself. This is often used to ease the sharing of
        files (files will be modifiable by all the users that are part of said group).

  ModeCharDevice                                 // c: Unix character device, when ModeDevice is set

  ModeSticky                                     // t: sticky
        The sticky bit works in a different way: while it has no effect on files, when used on a directory,
        all the files in said directory will be modifiable only by their owners. A typical case in which
        it is used, involves the /tmp directory. Typically this directory is writable by all users on the
        system, so to make impossible for one user to delete the files of another one.

  ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

  // Mask for the type bits. For regular files, none will be set.
  ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeIrregular

  ModePerm FileMode = 0777 // Unix permission bits
)

Reference:
  https://en.wikipedia.org/wiki/File_system_permissions

Numeric notation

Another method for representing Unix permissions is an octal (base-8) notation as shown by
stat -c %a. This notation consists of at least three digits. Each of the three rightmost
digits represents a different component of the permissions: owner, group, and others.
(If a fourth digit is present, the leftmost (high-order) digit addresses three additional
attributes, the setuid bit, the setgid bit and the sticky bit.)

Each of these digits is the sum of its component bits in the binary numeral system.
As a result, specific bits add to the sum as it is represented by a numeral:

    The read bit adds 4 to its total (in binary 100),
    The write bit adds 2 to its total (in binary 010), and
    The execute bit adds 1 to its total (in binary 001).

These values never produce ambiguous combinations; each sum represents a specific set of
permissions. More technically, this is an octal representation of a bit field – each bit
references a separate permission, and grouping 3 bits at a time in octal corresponds to
grouping these permissions by user, group, and others.

These are the examples from the symbolic notation section given in octal notation:

Three SymbolicGroups
The three groups: Owners, Groups & Others


Symbolic    Notation 	Numeric Notation 	English
---------- 	0000 	    no permissions
-rwx------ 	0700 	    read, write, & execute only for owner
-rwxrwx--- 	0770 	    read, write, & execute for owner and group
-rwxrwxrwx 	0777 	    read, write, & execute for owner, group and others
---x--x--x 	0111 	    execute
--w--w--w- 	0222 	    write
--wx-wx-wx 	0333 	    write & execute
-r--r--r-- 	0444 	    read
-r-xr-xr-x 	0555 	    read & execute
-rw-rw-rw- 	0666 	    read & write
-rwxr----- 	0740 	    owner can read, write, & execute; group can only read; others have no permissions

Reference:
How to use special permissions: the setuid, setgid and sticky bits
https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits


https://stackoverflow.com/questions/28969455/golang-properly-instantiate-os-filemode
 const (
        OS_READ = 04
        OS_WRITE = 02
        OS_EX = 01
        OS_USER_SHIFT = 6
        OS_GROUP_SHIFT = 3
        OS_OTH_SHIFT = 0

        OS_USER_R = OS_READ<<OS_USER_SHIFT
        OS_USER_W = OS_WRITE<<OS_USER_SHIFT
        OS_USER_X = OS_EX<<OS_USER_SHIFT
        OS_USER_RW = OS_USER_R | OS_USER_W
        OS_USER_RWX = OS_USER_RW | OS_USER_X

        OS_GROUP_R = OS_READ<<OS_GROUP_SHIFT
        OS_GROUP_W = OS_WRITE<<OS_GROUP_SHIFT
        OS_GROUP_X = OS_EX<<OS_GROUP_SHIFT
        OS_GROUP_RW = OS_GROUP_R | OS_GROUP_W
        OS_GROUP_RWX = OS_GROUP_RW | OS_GROUP_X

        OS_OTH_R = OS_READ<<OS_OTH_SHIFT
        OS_OTH_W = OS_WRITE<<OS_OTH_SHIFT
        OS_OTH_X = OS_EX<<OS_OTH_SHIFT
        OS_OTH_RW = OS_OTH_R | OS_OTH_W
        OS_OTH_RWX = OS_OTH_RW | OS_OTH_X

        OS_ALL_R = OS_USER_R | OS_GROUP_R | OS_OTH_R
        OS_ALL_W = OS_USER_W | OS_GROUP_W | OS_OTH_W
        OS_ALL_X = OS_USER_X | OS_GROUP_X | OS_OTH_X
        OS_ALL_RW = OS_ALL_R | OS_ALL_W
        OS_ALL_RWX = OS_ALL_RW | OS_GROUP_X
)

Convert octal to decimal
https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html

Unix Permissions
http://www.zzee.com/solutions/unix-permissions.shtml#numeric

*/

type FilePermissionMode struct {
	isInitialized bool
	fileMode      os.FileMode
}

// ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
func (fPerm FilePermissionMode) ModeDir() os.FileMode { return os.ModeDir }

// New - Creates and returns a new FilePermissionMode instance initialized with a
// an os.FileMode value generated from the input parameter 'modeStr'.
//
// 'modeStr' is a 10-character string containing the read, write and execute permissions
// for the the three groups, 'Owner', 'Group' and 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix permission codes.
//
// The first character of the 'modeStr' designates the 'Entry Type'.
//
//  Reference:
//  https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//  https://en.wikipedia.org/wiki/File_system_permissions
//
// The supported 'Entry Type' for the first character in 'modeStr' is either a "-",
// specifying a file, or "d" specifying a directory.
//
// The remaining nine characters in the 'modeStr' are three group fields each
// containing 3-characters. Each field may be populated with 'r' (Read-Permission),
// 'w' (Write-Permission), 'x' (Execute-Permission) or '-' signaling no permission
// or no access allowed. A typical 'modeStr' authorizing permission for full access
// to a file would be styled as:
//
//  "-rwxrwxrwx"
//
//  Groups: - Owner, Group, Other
//  From left to right
//  First Characters is Entry Type index 0 ("-")
//
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
// The three group types are: Owners, Groups & Others
//
// 'modeStr'
// Symbolic     Octal     File Access
// Notation    Notation 	Permission Descriptions
// ---------- 	0000 	    no permissions
// -rwx------ 	0700 	    read, write, & execute only for owner
// -rwxrwx--- 	0770 	    read, write, & execute for owner and group
// -rwxrwxrwx 	0777 	    read, write, & execute for owner, group and others
// ---x--x--x 	0111 	    execute
// --w--w--w- 	0222 	    write
// --wx-wx-wx 	0333 	    write & execute
// -r--r--r-- 	0444 	    read
// -r-xr-xr-x 	0555 	    read & execute
// -rw-rw-rw- 	0666 	    read & write
// -rwxr----- 	0740 	    Owner can read, write, & execute. Group can only read;
//                          others have no permissions
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
//   Reference:
//   How to use special permissions: the setuid, setgid and sticky bits
//   https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
func (fPerm FilePermissionMode) New(modeStr string) (FilePermissionMode, error) {

	ePrefix := "FilePermissionMode.New()"

	fPerm2 := FilePermissionMode{}

	err := fPerm2.SetFileModeByTextCode(modeStr)

	if err != nil {
		return FilePermissionMode{},
			fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return fPerm2, nil
}

// GetFileMode - Returns the os.FileMode from the internal data field,
// 'FilePermissionMode.fileMode'.
//
func (fPerm *FilePermissionMode) GetFileMode() (os.FileMode, error) {

	ePrefix := "FilePermissionMode.GetFileMode() "

	if !fPerm.isInitialized {
		return os.FileMode(0),
			fmt.Errorf(ePrefix +
				"Error: This FilePermissionMode instance has NOT bee initialized. The FileMode is INVALID!")
	}

	return fPerm.fileMode, nil
}

// SetFileModeByOctalDigits
// Three SymbolicGroups:
// The three group types are: Owners, Groups & Others
//
// 'modeStr'
// Symbolic     Octal     File Access
// Notation    Notation 	Permission Descriptions
// ---------- 	0000 	    no permissions
// -rwx------ 	0700 	    read, write, & execute only for owner
// -rwxrwx--- 	0770 	    read, write, & execute for owner and group
// -rwxrwxrwx 	0777 	    read, write, & execute for owner, group and others
// ---x--x--x 	0111 	    execute
// --w--w--w- 	0222 	    write
// --wx-wx-wx 	0333 	    write & execute
// -r--r--r-- 	0444 	    read
// -r-xr-xr-x 	0555 	    read & execute
// -rw-rw-rw- 	0666 	    read & write
// -rwxr----- 	0740 	    Owner can read, write, & execute. Group can only read;
//                          others have no permissions
//
func (fPerm *FilePermissionMode) SetFileModeByOctalDigits(octalFileModeCode int) error {

	decimalVal := FileHelper{}.ConvertOctalToDecimal(octalFileModeCode)

	fPerm.fileMode = os.FileMode(decimalVal)
	fPerm.isInitialized = true

	return nil
}

// SetFileModeByTextCode - Sets the internal FileMode data field using input
// parameter 'modeStr'. 'modeStr' is a 10-character string containing the read,
// write and execute permissions for the the three groups, 'Owner', 'Group' and
// 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix permission codes.
//
// The first character of the 'modeStr' designates the 'Entry Type'.
//
//   Reference:
//   https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//   https://en.wikipedia.org/wiki/File_system_permissions
//
// The supported 'Entry Type' for the first character in 'modeStr' is either a "-",
// specifying a file, or "d" specifying a directory.
//
// The remaining nine characters in the 'modeStr' are three group fields each
// containing 3-characters. Each field may be populated with 'r' (Read-Permission),
// 'w' (Write-Permission), 'x' (Execute-Permission) or '-' signaling no permission
// or no access allowed. A typical 'modeStr' authorizing permission for full access
// to a file would be styled as:
//
//  "-rwxrwxrwx"
//
//  Groups: - Owner, Group, Other
//  From left to right
//  First Characters is Entry Type index 0 ("-")
//
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
// The three group types are: Owners, Groups & Others
//
// 'modeStr'
// Symbolic     Octal     File Access
// Notation    Notation 	Permission Descriptions
// ---------- 	0000 	    no permissions
// -rwx------ 	0700 	    read, write, & execute only for owner
// -rwxrwx--- 	0770 	    read, write, & execute for owner and group
// -rwxrwxrwx 	0777 	    read, write, & execute for owner, group and others
// ---x--x--x 	0111 	    execute
// --w--w--w- 	0222 	    write
// --wx-wx-wx 	0333 	    write & execute
// -r--r--r-- 	0444 	    read
// -r-xr-xr-x 	0555 	    read & execute
// -rw-rw-rw- 	0666 	    read & write
// -rwxr----- 	0740 	    Owner can read, write, & execute. Group can only read;
//                          others have no permissions
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
//   Reference:
//   How to use special permissions: the setuid, setgid and sticky bits
//   https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
func (fPerm *FilePermissionMode) SetFileModeByTextCode(modeStr string) error {

	ePrefix := "FilePermissionMode.StringToMode() "

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

	firstChar := strings.ToLower(string(modeStr[0]))

	if firstChar == "d" {
		entryType = int(os.ModeDir)
		fMode = entryType | permission
	}

	fPerm.fileMode = os.FileMode(FileHelper{}.ConvertOctalToDecimal(fMode))
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
func (fPerm *FilePermissionMode) convertGroupToDecimal(groupStr, groupType string) (int, error) {

	ePrefix := "FilePermissionMode.convertGroupToDecimal() "
	var err error
	intVal := 0

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

// misdirection.go solves the [misc] htb challenge "misDIRection" made by incidrthreat
// the challenge may be downloaded here: https://app.hackthebox.com/challenges/misdirection

package main

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var filecount int
var secretdir string = ".secret/"
var secretstring []string

//CountFiles demonstrates a WalkFunc by counting files and adding it to the global int, 'filecount'
func CountFiles(path string, fileinfo fs.DirEntry, err error) error {
	if !fileinfo.IsDir() {
		filecount++
	} //end-if
	return nil
} //countFiles

//AssembleSecret constructs a string with characters taken from the Name of the numeric file's parent directory in filepath
//The character's position in []secretstring is determined by the numeric file's name.
func AssembleSecret(path string, fileinfo os.DirEntry, err error) error {
	if !fileinfo.IsDir() {
		characterposition, _ := strconv.Atoi(fileinfo.Name())
		parentdircharacter := strings.Replace(filepath.Dir(path), secretdir, "", 1)
		secretstring[characterposition-1] = parentdircharacter
		fmt.Printf("building string: %v\n", secretstring)
	} //end-if
	return nil
} //assembleSecret

//main will use WalkDir and custom WalkFuncs to solve the challenge.
//the ".secret/"" directory needs to exist by unzipping the challenge or created through the CreateSecret bonus function
func main() {
	//CreateSecret("recipe")
	filepath.WalkDir(secretdir, CountFiles)
	secretstring = make([]string, filecount)
	filepath.WalkDir(secretdir, AssembleSecret)
	decodedstring, _ := base64.StdEncoding.DecodeString(strings.Join(secretstring[:], ""))
	fmt.Printf("assembled flag: %s\n", decodedstring)
} //main

/*

// bonus round: emulate the puzzle width a twist where whitespace padding is implemented to preserve the directory structure without the "=" character in it
func CreateSecret(secret string) {
	directorylayout := strings.SplitAfter("0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", "")
	paddingneeded := len(secret) % 3
	switch paddingneeded {
	case 1:
		secret += "  "
	case 2:
		secret += " "
	} //switch
	encodedsecret := strings.SplitAfter(base64.StdEncoding.EncodeToString([]byte(secret)), "")
	os.RemoveAll(secretdir)
	os.Mkdir(secretdir, 0744)
	os.Chdir(secretdir)
	for idx := 0; idx < len(directorylayout); idx++ {
		os.Mkdir(directorylayout[idx], 0744)
	} //for
	for idx := 0; idx < len(encodedsecret); idx++ {
		os.Create(encodedsecret[idx] + "/" + strconv.Itoa(idx+1))
	} //for
	fmt.Printf("Secret created in %s\n", secretdir)
} //createSecret

*/

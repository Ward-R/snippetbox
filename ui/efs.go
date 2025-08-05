package ui

import (
	"embed"
)

//go:embed "html" "static"
var Files embed.FS

// above comment is not a comment but a comment directive. when this program
// is compiled it instructs Go to store the files from ui/static folder
// in an embedded filesystem refereced by global var Files. This allows
// our program to be self contained for easy distribution.

// the purpose of this embedding is so:
// Next let’s update our application so that the template cache uses embedded
// HTML template files, instead of reading them from your hard disk at runtime.

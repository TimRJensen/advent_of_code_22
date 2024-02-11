#!/bin/bash

# Check for argument
if [ $# -ne 1 ]; then
  echo "Please provide a destination name as an argument."
  exit 1
fi

# Get the destination name
dest_name="$1"

# Create the directory
mkdir -p "$dest_name"

# Change directory to the newly created one
cd "$dest_name"

# Create the input file
touch input.txt

# Create the Go source file with code
cat << EOF > "$dest_name.go"
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, $dest_name!")
}
EOF

# Initialize Go module
go mod init "$dest_name"

# Go back to the parent directory
cd ..

# Add the new module to go.work
go work use "$dest_name"

echo "Done! Directory, files, Go module, and go.work updated."

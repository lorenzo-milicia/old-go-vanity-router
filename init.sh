#!/bin/bash

read -p "What's the name of the project?: " projectname

if [ -z "$projectname" ]
then
	echo "Empty project name not allowed"
	exit 1
fi

echo "Initializing Go module"

cd module
go mod init "go.lorenzomilicia.dev/$projectname"
go mod tidy

cd ..
go work init module
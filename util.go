package main

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func PrintTitle() {
	banner := `
	_    _           _                     _     
	| |  | |         | |                   (_)    
	| |__| |_   _ ___| |_ ___ _ __ ___  ___ _ ___ 
	|  __  | | | / __| __/ _ \ '__/ _ \/ __| / __|
	| |  | | |_| \__ \ ||  __/ | |  __/\__ \ \__ \
	|_|  |_|\__, |___/\__\___|_|  \___||___/_|___/
			 __/ |                                
			|___/                                 
   `
	fmt.Println(banner)
}

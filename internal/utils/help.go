package utils

import(
	"fmt"
)

var usage = `usage: gitduo <command> [<args>]

These are the Gitduo commands available:

    help    list functionality
    pat     displays the github personal access token (if available)
    
In a Git working directory:

    which   shows which is the active git user
    main    set main git user in the current location
    work    set work git user in the current location
	
In a non Git working directory:

    set     initailize a repository in an empty location and provisions it on github. This takes 1 arg of value 1 for private and 0 for public
			(example gitduo set 1)`
			
func Help() {
	fmt.Println(usage)
}
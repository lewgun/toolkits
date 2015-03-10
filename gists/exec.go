package main

//http://www.golangtc.com/t/5440fe60421aa946910000a2

/*
#include <stdio.h>
int main() {
    fprintf(stderr, "this is an error" );
    return 1;
}

//fail.exe
*/

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("./fail")
	w := &bytes.Buffer{}
	cmd.Stderr = w
	if err := cmd.Run(); err != nil {
		fmt.Printf("Run returns: %s\n", err)
	}

	fmt.Printf("Stderr: %s\n", string(w.Bytes()))

}

/*
Run returns: exit status 1
Stderr: this is an error
*/

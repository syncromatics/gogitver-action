package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	envs := os.Environ()
	for _, v := range envs {
		fmt.Printf("%s\n", v)
	}
}

func getInput(name string, required bool) (string, error) {
	n := strings.ReplaceAll(name, " ", "_")
	n = strings.ToUpper(n)

	v, ok := os.LookupEnv(fmt.Sprintf("INPUT_%s", n))
	if !ok && required {
		return "", errors.Errorf("input %s was required but not found", name)
	}

	if !ok {
		return "", nil
	}

	return strings.Trim(v, " "), nil
}

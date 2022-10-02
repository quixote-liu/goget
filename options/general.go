package options

import (
	"fmt"
	"strings"
)

type General struct {
	keyVals []OptionKeyVal
}

var generalOptions = map[string]bool{
	"-O":                true,
	"--output-document": true,

	"-P":                 true,
	"--directory-prefix": true,
}

func (g *General) Parse(options []string) error {
	for i, opt := range options {
		if generalOptions[opt] {
			if i+1 >= len(options) {
				return fmt.Errorf("missing the option value for %s", opt)
			}
			val := options[i+1]
			if strings.HasPrefix(val, "-") || strings.HasPrefix(val, "--") {
				return fmt.Errorf("the value of the option %s is missing", opt)
			}
			g.keyVals = append(g.keyVals, OptionKeyVal{key: opt, val: options[i+1]})
		}
	}
	return nil
}

package main

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	//"github.com/stretchr/testify/assert"
)
const config_a = `
`

func TestConfig(t *testing.T) {
	config := conf{
	}

	//at some point mock a  filename returning a literal
	config.get()
}

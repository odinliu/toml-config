package tmlcfg

import (
	"errors"
	"io"
	"reflect"

	"github.com/pelletier/go-toml"
)

// BindFile can bind a toml file to a golang struct
func BindFile(filename string, conf interface{}) error {
	tree, err := toml.LoadFile(filename)
	if err != nil {
		return err
	}
	return bind(tree, conf)
}

// Bind can bind a toml string content to a golang struct
func Bind(content string, conf interface{}) error {
	tree, err := toml.Load(content)
	if err != nil {
		return err
	}
	return bind(tree, conf)
}

// BindReader can bind a toml reader content to a golang struct
func BindReader(reader io.Reader, conf interface{}) error {
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return err
	}
	return bind(tree, conf)
}

func bind(tree *toml.TomlTree, conf interface{}) error {
	v := reflect.ValueOf(conf)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("config must be an interface")
	}
	parser := newParser(tree, v.Elem())
	return parser.startParse()
}

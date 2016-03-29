package tmlcfg

import (
	"errors"
	"reflect"

	"github.com/pelletier/go-toml"
)

// Parser is use to parse toml with `github.com/pelletier/go-toml` and bind to config struct
type Parser struct {
	tree        *toml.TomlTree
	configValue reflect.Value
	atLeastOne  bool
}

func newParser(t *toml.TomlTree, v reflect.Value) *Parser {
	return &Parser{
		tree:        t,
		configValue: v,
		atLeastOne:  true,
	}
}

func (p *Parser) startParse() error {
	p.parse(p.configValue)
	if p.atLeastOne {
		return nil
	}
	return errors.New("no field bind to toml file")
}

func (p *Parser) parse(v reflect.Value) {
	numByValue := v.NumField()
	typo := v.Type()

	for i := 0; i < numByValue; i++ {
		field := v.Field(i)
		path := typo.Field(i).Tag.Get("toml")
		switch field.Kind() {
		case reflect.Struct:
			p.parse(field)
		case reflect.Slice:
			if len(path) > 0 {
				if tmp := p.tree.Get(path); tmp != nil {
					if tv := reflect.ValueOf(tmp); tv.Kind() == reflect.Slice {
						nslice := reflect.MakeSlice(field.Type(), tv.Len(), tv.Cap())
						field.Set(nslice)
						for j := 0; j < tv.Len(); j++ {
							stv := tv.Index(j)
							switch stv.Kind() {
							case reflect.String:
								field.Index(j).SetString(stv.String())
							case reflect.Int64:
								field.Index(j).SetInt(stv.Int())
							case reflect.Bool:
								field.Index(j).SetBool(stv.Bool())
							case reflect.Float64:
								field.Index(j).SetFloat(stv.Float())
							default:
								field.Index(j).Set(stv.Elem())
							}
						}
					}
				}
			}
		case reflect.Bool:
			if len(path) > 0 {
				if tv := p.validNode(path, reflect.Bool); tv != nil {
					field.SetBool(tv.Bool())
					p.atLeastOne = true
				}
			}
		case reflect.String:
			if len(path) > 0 {
				if tv := p.validNode(path, reflect.String); tv != nil {
					field.SetString(tv.String())
					p.atLeastOne = true
				}
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if len(path) > 0 {
				if tv := p.validNode(path, reflect.Int64); tv != nil {
					field.SetInt(tv.Int())
					p.atLeastOne = true
				}
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if len(path) > 0 {
				if tv := p.validNode(path, reflect.Int64); tv != nil {
					field.SetUint(uint64(tv.Int()))
					p.atLeastOne = true
				}
			}
		case reflect.Float32, reflect.Float64:
			if len(path) > 0 {
				if tv := p.validNode(
					path, reflect.Float32, reflect.Float64,
				); tv != nil {
					field.SetFloat(tv.Float())
					p.atLeastOne = true
				}
			}
		default:
			// unsupport value types
		}
	}
}

func (p *Parser) validNode(path string, kinds ...reflect.Kind) *reflect.Value {
	kslice := []reflect.Kind(kinds)
	if tmp := p.tree.Get(path); tmp != nil {
		var flag = false
		tv := reflect.ValueOf(tmp)
		for _, v := range kslice {
			//fmt.Printf("%s[%v]\n", path, tv.Kind())
			flag = flag || (tv.Kind() == v)
			if flag {
				break
			}
		}
		if flag {
			return &tv
		}
	}
	return nil
}

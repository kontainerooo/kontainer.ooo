// Package cli is
package cli

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/abiosoft/ishell"
	"google.golang.org/grpc"
)

// InitShell adds all available kontaineroo commands to an ishell instance
func InitShell(sh *ishell.Shell, conn *grpc.ClientConn, logger log.Logger) {
	sh.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "authenticate yourself",
		Func: func(c *ishell.Context) {
			c.Println("not yet implemented")
		},
	})
}

func fillRequestStruct(c *ishell.Context, value *reflect.Value, typ reflect.Type) {
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		valField := value.Field(i)
		name := strings.Title(typeField.Name)

		if valField.Kind() == reflect.Ptr {
			valField = valField.Elem()
		}

		if valField.Kind() == reflect.Struct {
			c.Println(name)
			fillRequestStruct(c, &valField, valField.Type())
			continue
		}

		for {
			c.Print(name, ": ")
			value := c.ReadLine()
			switch valField.Kind() {
			case reflect.String:
				valField.Set(reflect.ValueOf(value))
			case reflect.Uint32:
				num, err := strconv.ParseUint(value, 10, 32)
				if err != nil {
					c.Println(err.Error())
					continue
				}
				valField.SetUint(num)
			case reflect.Int32:
				num, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					c.Println(err.Error())
					continue
				}
				valField.SetInt(num)
			}
			break
		}
	}
}

func printResult(c *ishell.Context, res reflect.Value, resType reflect.Type) {
	c.Println(strings.Title(resType.Name()))
	for i := 0; i < resType.NumField(); i++ {
		field := resType.Field(i)
		val := res.Field(i)
		name := strings.Title(field.Name)

		if name == "Error" {
			continue
		}

		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		if val.Kind() == reflect.Struct {
			printResult(c, val, val.Type())
			continue
		}

		c.Println(name, ": ", val.Interface())
	}
}

func createCommand(name, help string, endpoint endpoint.Endpoint, req, resType interface{}) *ishell.Cmd {
	return &ishell.Cmd{
		Name: name,
		Help: help,
		Func: func(c *ishell.Context) {
			reqValPtr := reflect.ValueOf(req)
			reqVal := reqValPtr.Elem()
			reqType := reqVal.Type()
			fillRequestStruct(c, &reqVal, reqType)
			res, err := endpoint(context.Background(), reqValPtr.Interface())
			if err != nil {
				c.Println(err)
				return
			}

			resVal := reflect.ValueOf(res).Elem()
			resValType := resVal.Type()

			if reflect.ValueOf(resType).Elem().Type() != resValType {
				c.Println("response malformed")
				return
			}

			errVal := resVal.FieldByName("Error")
			if errVal.Interface() != nil && errVal != reflect.ValueOf(nil) {
				c.Println("Error: ", errVal)
				return
			}

			c.Println("Success!")
			printResult(c, resVal, resValType)
		},
	}
}

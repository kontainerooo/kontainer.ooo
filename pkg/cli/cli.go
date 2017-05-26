// Package cli is
package cli

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiClient "github.com/kontainerooo/kontainer.ooo/pkg/kmi/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/module"
	moduleClient "github.com/kontainerooo/kontainer.ooo/pkg/module/client"

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

	kmiClient := kmiClient.New(conn, logger)
	sh.AddCmd(kmiCommands(sh, kmiClient))

	moduleClient := moduleClient.New(conn, logger)
	sh.AddCmd(moduleCommands(sh, moduleClient))
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
			case reflect.Uint:
				num, err := strconv.ParseUint(value, 10, 32)
				if err != nil {
					c.Println(err.Error())
					continue
				}
				valField.SetUint(num)
			case reflect.Int:
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

func kmiCommands(sh *ishell.Shell, kmiClient *kmi.Endpoints) *ishell.Cmd {
	kmiCmd := &ishell.Cmd{
		Name: "kmi",
		Help: "all KMI Service commands",
	}

	kmiCmd.AddCmd(createCommand(
		"add",
		"adds a kmi",
		kmiClient.AddKMIEndpoint,
		&kmi.AddKMIRequest{},
		&kmi.AddKMIResponse{}),
	)

	kmiCmd.AddCmd(createCommand(
		"remove",
		"removes a kmi by id",
		kmiClient.RemoveKMIEndpoint,
		&kmi.RemoveKMIRequest{},
		&kmi.RemoveKMIResponse{}),
	)

	kmiCmd.AddCmd(createCommand(
		"get",
		"requests information for a specific kmi",
		kmiClient.GetKMIEndpoint,
		&kmi.GetKMIRequest{},
		&kmi.GetKMIResponse{}),
	)

	kmiCmd.AddCmd(createCommand(
		"all",
		"requests information for all kmis",
		kmiClient.KMIEndpoint,
		&kmi.KMIRequest{},
		&kmi.KMIResponse{}),
	)

	return kmiCmd
}

func moduleCommands(sh *ishell.Shell, moduleClient *module.Endpoints) *ishell.Cmd {
	moduleCmd := &ishell.Cmd{
		Name: "module",
		Help: "all Module Service commands",
	}

	setCmd := &ishell.Cmd{
		Name: "set",
	}

	setCmd.AddCmd(createCommand(
		"publickey",
		"set publickey",
		moduleClient.SetPublicKeyEndpoint,
		&module.SetPublicKeyRequest{},
		&module.SetPublicKeyResponse{}),
	)

	setCmd.AddCmd(createCommand(
		"env",
		"set environment",
		moduleClient.SetEnvEndpoint,
		&module.SetEnvRequest{},
		&module.SetEnvResponse{}),
	)

	moduleCmd.AddCmd(setCmd)

	getCmd := &ishell.Cmd{
		Name: "get",
	}

	getCmd.AddCmd(createCommand(
		"file",
		"get file",
		moduleClient.GetFileEndpoint,
		&module.GetFileRequest{},
		&module.GetFileResponse{}),
	)

	getCmd.AddCmd(createCommand(
		"files",
		"get files",
		moduleClient.GetFilesEndpoint,
		&module.GetFilesRequest{},
		&module.GetFilesResponse{}),
	)

	getCmd.AddCmd(createCommand(
		"env",
		"get environment",
		moduleClient.GetEnvEndpoint,
		&module.GetEnvRequest{},
		&module.GetEnvResponse{}),
	)

	getCmd.AddCmd(createCommand(
		"moduleconf",
		"get module config",
		moduleClient.GetModuleConfigEndpoint,
		&module.GetModuleConfigRequest{},
		&module.GetModuleConfigResponse{}),
	)

	moduleCmd.AddCmd(getCmd)

	moduleCmd.AddCmd(createCommand(
		"uploadfile",
		"upload file",
		moduleClient.UploadFileEndpoint,
		&module.UploadFileRequest{},
		&module.UploadFileResponse{}),
	)

	moduleCmd.AddCmd(createCommand(
		"sendcmd",
		"send command",
		moduleClient.SendCommandEndpoint,
		&module.SendCommandRequest{},
		&module.SendCommandResponse{}),
	)

	removeCmd := &ishell.Cmd{
		Name: "remove",
	}

	removeCmd.AddCmd(createCommand(
		"file",
		"remove file",
		moduleClient.RemoveFileEndpoint,
		&module.RemoveFileRequest{},
		&module.RemoveFileResponse{}),
	)

	removeCmd.AddCmd(createCommand(
		"dir",
		"remove directory",
		moduleClient.RemoveDirectoryEndpoint,
		&module.RemoveDirectoryRequest{},
		&module.RemoveDirectoryResponse{}),
	)

	moduleCmd.AddCmd(removeCmd)

	return moduleCmd
}

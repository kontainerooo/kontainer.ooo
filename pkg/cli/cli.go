// Package cli provides the available commands for kroocli
package cli

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/kontainerooo/kontainer.ooo/pkg/abstraction"
	"github.com/kontainerooo/kontainer.ooo/pkg/container"
	containerClient "github.com/kontainerooo/kontainer.ooo/pkg/container/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/kmi"
	kmiClient "github.com/kontainerooo/kontainer.ooo/pkg/kmi/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/module"
	moduleClient "github.com/kontainerooo/kontainer.ooo/pkg/module/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/routing"
	routingClient "github.com/kontainerooo/kontainer.ooo/pkg/routing/client"
	"github.com/kontainerooo/kontainer.ooo/pkg/user"
	userClient "github.com/kontainerooo/kontainer.ooo/pkg/user/client"

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

	containerClient := containerClient.New(conn, logger)
	sh.AddCmd(containerCommands(sh, containerClient))

	userClient := userClient.New(conn, logger)
	sh.AddCmd(userCommands(sh, userClient))

	routingClient := routingClient.New(conn, logger)
	sh.AddCmd(routingCommands(sh, routingClient))
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
				if valField.Type() == reflect.TypeOf(abstraction.Inet("")) {
					valField.Set(reflect.ValueOf(abstraction.Inet(value)))
				}
				valField.SetString(value)
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
			case reflect.Bool:
				bul, err := strconv.ParseBool(value)
				if err != nil {
					c.Println(err.Error())
					continue
				}
				valField.SetBool(bul)
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
			if errVal != reflect.ValueOf(nil) && errVal.Interface() != nil {
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

	moduleCmd.AddCmd(createCommand(
		"create",
		"create container module",
		moduleClient.CreateContainerModuleEndpoint,
		&module.CreateContainerModuleRequest{},
		&module.CreateContainerModuleResponse{}),
	)

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

	setCmd.AddCmd(createCommand(
		"link",
		"set link",
		moduleClient.SetLinkEndpoint,
		&module.SetLinkRequest{},
		&module.SetLinkResponse{}),
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

	getCmd.AddCmd(createCommand(
		"modules",
		"get all modules",
		moduleClient.GetModulesEndpoint,
		&module.GetModulesRequest{},
		&module.GetModulesResponse{},
	))

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
		&module.RemoveDirectoryResponse{},
	))

	removeCmd.AddCmd(createCommand(
		"link",
		"remove link",
		moduleClient.RemoveLinkEndpoint, &module.RemoveLinkRequest{},
		&module.RemoveLinkResponse{},
	))

	moduleCmd.AddCmd(removeCmd)

	return moduleCmd
}

func containerCommands(sh *ishell.Shell, containerClient *container.Endpoints) *ishell.Cmd {
	containerCmd := &ishell.Cmd{
		Name: "container",
	}

	containerCmd.AddCmd(createCommand(
		"create",
		"creates a container",
		containerClient.CreateContainerEndpoint,
		&container.CreateContainerRequest{},
		&container.CreateContainerResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"remove",
		"removes a container",
		containerClient.RemoveContainerEndpoint,
		&container.RemoveContainerRequest{},
		&container.RemoveContainerResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"execute",
		"execute a command inside a container",
		containerClient.ExecuteEndpoint,
		&container.ExecuteRequest{},
		&container.ExecuteResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"stop",
		"stop a container",
		containerClient.StopContainerEndpoint,
		&container.StopContainerRequest{},
		&container.StopContainerResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"instances",
		"all container instances",
		containerClient.InstancesEndpoint,
		&container.InstancesRequest{},
		&container.InstancesResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"id",
		"get id for container name",
		containerClient.IDForNameEndpoint,
		&container.IDForNameRequest{},
		&container.IDForNameResponse{},
	))

	containerCmd.AddCmd(createCommand(
		"kmi",
		"get kmi of a container",
		containerClient.GetContainerKMIEndpoint,
		&container.GetContainerKMIRequest{},
		&container.GetContainerKMIResponse{},
	))

	linkCmd := &ishell.Cmd{
		Name: "link",
	}

	linkCmd.AddCmd(createCommand(
		"set",
		"set a link",
		containerClient.SetLinkEndpoint,
		&container.SetLinkRequest{},
		&container.SetLinkResponse{},
	))

	linkCmd.AddCmd(createCommand(
		"remove",
		"remove a link",
		containerClient.RemoveLinkEndpoint,
		&container.RemoveLinkRequest{},
		&container.RemoveLinkResponse{},
	))

	linkCmd.AddCmd(createCommand(
		"get",
		"get links of a module",
		containerClient.GetContainerKMIEndpoint,
		&container.RemoveLinkRequest{},
		&container.RemoveLinkResponse{},
	))

	containerCmd.AddCmd(linkCmd)

	return containerCmd
}

func userCommands(sh *ishell.Shell, userClient *user.Endpoints) *ishell.Cmd {
	userCmd := &ishell.Cmd{
		Name: "user",
		Help: "all User Service commands",
	}

	userCmd.AddCmd(createCommand(
		"create",
		"create user",
		userClient.CreateUserEndpoint,
		&user.CreateUserRequest{
			Cfg: &user.Config{
				Address: user.Address{},
			},
		},
		&user.CreateUserResponse{},
	))

	userCmd.AddCmd(createCommand(
		"edit",
		"edit user",
		userClient.EditUserEndpoint,
		&user.EditUserRequest{},
		&user.EditUserResponse{},
	))

	userCmd.AddCmd(createCommand(
		"delete",
		"delete user",
		userClient.DeleteUserEndpoint,
		&user.DeleteUserRequest{},
		&user.DeleteUserResponse{},
	))

	userCmd.AddCmd(createCommand(
		"get",
		"get user",
		userClient.GetUserEndpoint,
		&user.GetUserRequest{},
		&user.GetUserResponse{},
	))

	userCmd.AddCmd(createCommand(
		"changeusername",
		"change username",
		userClient.ChangeUsernameEndpoint,
		&user.ChangeUsernameRequest{},
		&user.ChangeUsernameResponse{},
	))

	userCmd.AddCmd(createCommand(
		"checkcredentials",
		"check login credentials",
		userClient.CheckLoginCredentialsEndpoint,
		&user.CheckLoginCredentialsRequest{},
		&user.CheckLoginCredentialsResponse{},
	))

	return userCmd
}

func routingCommands(sh *ishell.Shell, routingClient *routing.Endpoints) *ishell.Cmd {
	routingCmd := &ishell.Cmd{
		Name: "routing",
		Help: "all routing service commands",
	}

	ls := &routing.ListenStatement{}
	log := routing.Log{}
	ssl := routing.SSLSettings{}
	loc := &routing.LocationRule{}
	rc := &routing.RouterConfig{
		ListenStatement: ls,
		AccessLog:       log,
		ErrorLog:        log,
		SSLSettings:     ssl,
		LocationRules:   routing.LocationRules{loc},
	}

	routingCmd.AddCmd(createCommand(
		"createconf",
		"Create a router config",
		routingClient.CreateConfigEndpoint,
		&routing.CreateConfigRequest{
			Config: rc,
		},
		&routing.CreateConfigResponse{},
	))

	routingCmd.AddCmd(createCommand(
		"editconf",
		"Edit a router config",
		routingClient.EditConfigEndpoint,
		&routing.EditConfigRequest{
			Config: rc,
		},
		&routing.EditConfigResponse{},
	))

	routingCmd.AddCmd(createCommand(
		"getconf",
		"Get a router config",
		routingClient.GetConfigEndpoint,
		&routing.GetConfigRequest{},
		&routing.GetConfigResponse{},
	))

	removeCmd := &ishell.Cmd{
		Name: "remove",
	}

	removeCmd.AddCmd(createCommand(
		"conf",
		"Remove a router config",
		routingClient.RemoveConfigEndpoint,
		&routing.RemoveConfigRequest{},
		&routing.RemoveConfigResponse{},
	))

	removeCmd.AddCmd(createCommand(
		"loc",
		"Remove a location",
		routingClient.RemoveLocationEndpoint,
		&routing.RemoveLocationRequest{},
		&routing.RemoveLocationResponse{},
	))

	removeCmd.AddCmd(createCommand(
		"sn",
		"Remove a server name",
		routingClient.RemoveServerNameEndpoint,
		&routing.RemoveServerNameRequest{},
		&routing.RemoveServerNameResponse{},
	))

	routingCmd.AddCmd(removeCmd)

	addCmd := &ishell.Cmd{
		Name: "Add",
	}

	addCmd.AddCmd(createCommand(
		"loc",
		"Add a location",
		routingClient.AddLocationEndpoint,
		&routing.AddLocationRequest{},
		&routing.AddLocationResponse{},
	))

	addCmd.AddCmd(createCommand(
		"sn",
		"Add a server name",
		routingClient.AddServerNameEndpoint,
		&routing.AddServerNameRequest{},
		&routing.AddServerNameResponse{},
	))

	routingCmd.AddCmd(addCmd)

	routingCmd.AddCmd(createCommand(
		"changels",
		"Change a listen statement",
		routingClient.ChangeListenStatementEndpoint,
		&routing.ChangeListenStatementRequest{},
		&routing.ChangeListenStatementResponse{},
	))

	routingCmd.AddCmd(createCommand(
		"all",
		"Get all configurations",
		routingClient.ConfigurationsEndpoint,
		&routing.ConfigurationsRequest{},
		&routing.ConfigurationsResponse{},
	))

	return routingCmd
}

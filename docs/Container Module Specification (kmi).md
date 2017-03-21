# Container Module Specification

## KMI File
The `.kmi` file is a tar-archived folder of the following structure of at least the following files:
- A `Dockerfile`
- A `module.json` that contains the module's description
- A folder that contains mount directories, scripts and other container related data that is specified in the `module.json`

### `module.json`

```javascript
{
    "name":         string,     // The module's name
    "version":      string,     // The module's version
    "description":  string,     // The module's description
    "type":         int,        // See messages/kmi.proto -> enum TYPE
    "dockerfile":   string,     // The path to the Dockerfile (usually ./Dockerfile)
    "container":    string,     // The path to the module's folder (usually ./container)

    /* The following options can either be specified inline as object/array
     * or extracted into a separate file by providing the filename.
     * The descriptions for the following keys are below. */
    "frontend":     string || object,
    "env":          string || object,
    "interfaces":   string || object,
    "cmd":          string || object,
    "mounts":       string || array,
    "variables":    string || array,
    "resources":    string || object
}
```

### `env.json`
The `env` key configures environment variables inside the container.
```javascript
{
    "ENV_VAR": string || int (value)
}
```

### `interfaces.json`
The `interfaces` key configures ports that are to be exposed by the container. Each port can be given a name from which it can be referenced from.
```javascript
{
    "interface-name": int(port),
}
```

### `cmd.json`
The `cmd` key configures commands that can be used from frontend modules.
```javascript
{
    // Example: "start": "npm start"
    "command-name": string(command)
}
```

### `mounts.json`
The `mounts` key specifies directories relative to the `container` directory that are mounted inside the docker container.,
```javascript
[
    // Example: /mount ~> everything from container/mount is mounted to / inside the container
    "path/to/dir"
]
```

### `variables.json`
The `variables` key defines "variables" that can be used inside and outside the container. Basically they are files that are mounted in `/var/kro/variables` inside the container. The array contains paths that contain these "variable files".
```javascript
[
    // Example: /variables ~> everything from container/variables is mounted to /var/kro/variables inside the container
    "path/to/vars"
]
```

### `resources.json`
The `resources` key defines resource limits for the container.
```javascript
{
    "cpus": string(num_cpus),
    "mem":  int(bytes),
    "swap": int(bytes)
}
```

### `frontend.json`
The `frontend` key defines configuration modules that are displayed in the module's configuration frontend.
```javascript
{
    // Imports specify the templates that are used. See docs/Templates.md
	"imports": array(string(module)),
	"modules": [
        {
            "template": string(template_name),
	        "parameters": object(template_parameters)
        }
    ]
}
```

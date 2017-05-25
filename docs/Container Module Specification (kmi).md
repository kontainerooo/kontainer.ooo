# Container Module Specification

## KMI File
The `.kmi` file is a tar-archived folder of the following structure of at least the following file:
- A `module.json` that contains the module's description

### `module.json`

```javascript
{
    "name":             string,     // The module's name
    "version":          string,     // The module's version
    "description":      string,     // The module's description
    "type":             int,        // See messages/kmi.proto -> enum TYPE
    "provisionScript":  string,     // The path to the script that provisions the container module

    /* The following options can either be specified inline as object/array
     * or extracted into a separate file by providing the filename.
     * The descriptions for the following keys are below. */
    "frontend":     string || object,
    "env":          string || object,
    "interfaces":   string || object,
    "cmd":          string || object,
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

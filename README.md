##THTTP
#Usage
```shell
thttp -f <RequestFiles | -d <RequestDir> 
Flags:
  -h, --help                    Show context-sensitive help.
  -f, --file=FILE,...           Files to execute.
  -d, --dir=DIR,...             Executes all the files in the given directories
  -p, --properties=STRING       Properties file to be use can be yaml, json or property file
  -r, --repeat=INT              Repeats the requests the given times
  -I, --header-only             Print only the headers
  -i, --with-header             Print the response with the headers
  -o, --output=STRING           Write the output to a file
  -f, --output-folder=STRING    Files to execute.
  -k, --ignore-ssl              Ignore SSL Validation
  -l, --follow                  Follow Redirects
      --plugin=PLUGIN,...       Plugins to Load

```

## Plugins
Plugins are ECMAScript 5.1 scripts with at least one function,
```javascript
function thttp() {}
```
This function should not have any parameters and should return an JSON
object with at least 4 keys, `name`, `type`, `exec` 
* **_name_**: Name of the plugin, this name will be use as reference internally,
* **_type_**: Type of the plugin currently available types are `generator`
* **_exec_**: Function to be executed, depending on the plugin type the result will
              be cached.
Example
```javascript

function exec(){
    return Math.random();
}
function thttp() {
    return {
        "name":"random",
        "type": "generator",
        "exec":"exec" 
    }
}
```

### Plugin types
#### Generator
This plugin types are used to generate data, the result will insert in the request using the 
templating of request. 

This type of plugins are executed for each request

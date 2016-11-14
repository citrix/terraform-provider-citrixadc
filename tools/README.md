## Tooling to support generation of new NetScaler resources
Quite a bit of a resource_foo.go file that models a NetScaler config object in Terraform can be derived from the JSON schema for that NetScaler object. This tooling supports generating a skeleton for a NetScaler config object for the Terraform NetScaler provider. Note
however, there are some semantics (update, especially) that have to be hand-coded, so this only helps get you started.

## Building the tool
```
make  build
```

## Using the tool
Locate the json schema (e.g., `export JSON_CFG=$GOPATH/src/github.com/chiradeep/go-nitro/jsonconfig`). 

```
./generate_schema -i $JSON_CFG/<subdir>/<cfg>.json  -d <id fieldname for the object>  -k <sample json for generating a create test>  
```

For example

```
./generate_schema -i $JSON_CFG/cs/cspolicy.json  -d policyname  -k '{ "policyname":"test_policy", "url": "/foo/*", "lbvserver": "test-lb-vserver"}'
```

Usually, a config object has a relation or dependency ("binding") on one or more other objects. To implement the binding, add the `-b`, `-n` and `-K` flags. The `-b` flag specifies the binding JSON schema, the `-n` specifies the field name in the JSON that is used to bind the the two objects and the `-K` gives a sample JSON for the bound-to object.

```
./generate_schema -i $JSON_CFG/cs/cspolicy.json  -d policyname -b $JSON_CFG/cs/csvserver_cspolicy_binding.json  -n csvserver -k '{ "policyname":"test_policy",  "lbvserver": "test-lb-vserver"}'  -K '{ "ipv46:10.202.11.11", "servicetype": "SSL", "port":"443", "name": "terraform-cs"}'
```

The code is generated in the `netscaler` subdirectory. Use `make fmt` to format it and `make generate-test-build` to make sure it compiles.
Copy the resulting netscaler/resource_<cfg>.go and netscaler/resource_<cfg>_test.go over to the resource directory and continue making modifications there.

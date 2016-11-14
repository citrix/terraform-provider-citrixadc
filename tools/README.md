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
Copy the resulting netscaler/resource\_<cfg>.go and netscaler/resource\_<cfg>\_test.go over to the resource directory and continue making modifications there.


## Using the generated code
The generated code adds the config `schema` to the `terraform-netscaler-provider` by using the `readwrite` fields in the JSON schema. The mandatory fields for the NITRO API however are not determinable from the JSON schema and therefore the `Required` field is usually not set to `true`. If the developer knows this (e.g., from the documentation), then this can be changed.

The `create` function is semi-complete: if there are more than one bindings or dependency (e.g, `cspolicy` is bound to `csvserver` and depends on `lbvserver`), then additional code and validation may have to be written. Make sure that the `SetId` call to TF is made only after all dependencies are successful.

The `read` function is usually self-contained and complete.

The `update` function is usually the trickiest since not all fields are updateable. Also, if the binding fields are updated in TF then the old binding has to be removed and the new binding has to be created. The pattern is
 * determine if any fields have changed in the user-supplied TF config
 * if the binding fields have changed, remove the bindings using the NITRO API
 * update the config object if non-binding fields have changed
 * if binding fields have changed, update with new bindings

Make sure to return errors early.

The `delete` function is usually complete, although sometimes the delete function expects additional parameters (other than the object id). See `resource_lbmonitor.go` for an example.

The test code usually needs to be tweaked, for example if there is more than one dependency for the object, all dependencies have to be created first (see the bottom of the file for the test config).


## TODO
Generate update tests. Since the semantics vary the most between objects for update, this is a source of functional bugs

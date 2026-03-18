package cacheglobal_cachepolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*CacheglobalCachepolicyBindingDataSource)(nil)

func CAcheglobalCachepolicyBindingDataSource() datasource.DataSource {
	return &CacheglobalCachepolicyBindingDataSource{}
}

type CacheglobalCachepolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *CacheglobalCachepolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheglobal_cachepolicy_binding"
}

func (d *CacheglobalCachepolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *CacheglobalCachepolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = CacheglobalCachepolicyBindingDataSourceSchema()
}

func (d *CacheglobalCachepolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CacheglobalCachepolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policy_Name := data.Policy
	type_Name := data.Type

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	var err error
	if !type_Name.IsNull() && type_Name.ValueString() != "" {
		argsMap["type"] = type_Name.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Cacheglobal_cachepolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cacheglobal_cachepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "cacheglobal_cachepolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policy
		if val, ok := v["policy"].(string); ok {
			if policy_Name.IsNull() || val != policy_Name.ValueString() {
				match = false
				continue
			}
		} else if !policy_Name.IsNull() {
			match = false
			continue
		}
		if !type_Name.IsNull() && type_Name.ValueString() != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != type_Name.ValueString() {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("cacheglobal_cachepolicy_binding with policy %s not found", policy_Name))
		return
	}

	cacheglobal_cachepolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

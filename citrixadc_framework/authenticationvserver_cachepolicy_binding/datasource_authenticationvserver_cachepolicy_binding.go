package authenticationvserver_cachepolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuthenticationvserverCachepolicyBindingDataSource)(nil)

func AUthenticationvserverCachepolicyBindingDataSource() datasource.DataSource {
	return &AuthenticationvserverCachepolicyBindingDataSource{}
}

type AuthenticationvserverCachepolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AuthenticationvserverCachepolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_cachepolicy_binding"
}

func (d *AuthenticationvserverCachepolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuthenticationvserverCachepolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuthenticationvserverCachepolicyBindingDataSourceSchema()
}

func (d *AuthenticationvserverCachepolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuthenticationvserverCachepolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	bindpoint_Name := data.Bindpoint
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_cachepolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_cachepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "authenticationvserver_cachepolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bindpoint
		if val, ok := v["bindpoint"].(string); ok {
			if bindpoint_Name.IsNull() || val != bindpoint_Name.ValueString() {
				match = false
				continue
			}
		} else if !bindpoint_Name.IsNull() {
			match = false
			continue
		}

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
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("authenticationvserver_cachepolicy_binding with bindpoint %s not found", bindpoint_Name))
		return
	}

	authenticationvserver_cachepolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

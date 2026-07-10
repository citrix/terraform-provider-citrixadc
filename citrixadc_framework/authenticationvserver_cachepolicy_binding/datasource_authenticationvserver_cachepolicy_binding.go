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

	// Array filter with parent ID. The lookup keys are name (parent) + policy. bindpoint
	// is an optional additional filter. secondary/groupextraction are NOT echoed by the
	// NITRO GET response, so they cannot be used as filters.
	name_Name := data.Name.ValueString()
	policy_Name := data.Policy
	bindpoint_Name := data.Bindpoint

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

	// Iterate through results to find the matching binding (policy, and bindpoint if set).
	foundIndex := -1
	for i, v := range dataArr {
		// Check policy (required lookup key)
		if val, ok := v["policy"].(string); !ok || policy_Name.IsNull() || val != policy_Name.ValueString() {
			continue
		}
		// Check bindpoint only when the caller supplied it
		if !bindpoint_Name.IsNull() && !bindpoint_Name.IsUnknown() {
			if val, ok := v["bindpoint"].(string); !ok || val != bindpoint_Name.ValueString() {
				continue
			}
		}
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("authenticationvserver_cachepolicy_binding with policy %s not found", policy_Name.ValueString()))
		return
	}

	authenticationvserver_cachepolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

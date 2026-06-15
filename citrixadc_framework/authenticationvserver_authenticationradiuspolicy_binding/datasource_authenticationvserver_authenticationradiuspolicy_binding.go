package authenticationvserver_authenticationradiuspolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuthenticationvserverAuthenticationradiuspolicyBindingDataSource)(nil)

func AUthenticationvserverAuthenticationradiuspolicyBindingDataSource() datasource.DataSource {
	return &AuthenticationvserverAuthenticationradiuspolicyBindingDataSource{}
}

type AuthenticationvserverAuthenticationradiuspolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *AuthenticationvserverAuthenticationradiuspolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationradiuspolicy_binding"
}

func (d *AuthenticationvserverAuthenticationradiuspolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuthenticationvserverAuthenticationradiuspolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuthenticationvserverAuthenticationradiuspolicyBindingDataSourceSchema()
}

func (d *AuthenticationvserverAuthenticationradiuspolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuthenticationvserverAuthenticationradiuspolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Lookup keys for this binding are name (parent) and policy, matching the SDK v2 identity.
	name_Name := data.Name.ValueString()
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_authenticationradiuspolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "authenticationvserver_authenticationradiuspolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the policy key
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_Name.ValueString() {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationradiuspolicy_binding with policy %s not found", policy_Name.ValueString()))
		return
	}

	authenticationvserver_authenticationradiuspolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

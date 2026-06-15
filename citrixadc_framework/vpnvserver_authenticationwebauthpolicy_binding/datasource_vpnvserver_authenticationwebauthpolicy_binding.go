package vpnvserver_authenticationwebauthpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverAuthenticationwebauthpolicyBindingDataSource)(nil)

func VPnvserverAuthenticationwebauthpolicyBindingDataSource() datasource.DataSource {
	return &VpnvserverAuthenticationwebauthpolicyBindingDataSource{}
}

type VpnvserverAuthenticationwebauthpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverAuthenticationwebauthpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationwebauthpolicy_binding"
}

func (d *VpnvserverAuthenticationwebauthpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverAuthenticationwebauthpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverAuthenticationwebauthpolicyBindingDataSourceSchema()
}

func (d *VpnvserverAuthenticationwebauthpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverAuthenticationwebauthpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID. Lookup keys are name (parent) + policy (bound entity).
	// bindpoint is not echoed by the NITRO GET response, so it cannot be used as a filter.
	name_Name := data.Name.ValueString()
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_authenticationwebauthpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationwebauthpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_authenticationwebauthpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the policy
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok {
			if !policy_Name.IsNull() && val == policy_Name.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_authenticationwebauthpolicy_binding with policy %s not found", policy_Name.ValueString()))
		return
	}

	vpnvserver_authenticationwebauthpolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

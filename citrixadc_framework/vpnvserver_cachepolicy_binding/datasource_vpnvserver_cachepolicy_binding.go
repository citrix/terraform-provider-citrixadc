package vpnvserver_cachepolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverCachepolicyBindingDataSource)(nil)

func VPnvserverCachepolicyBindingDataSource() datasource.DataSource {
	return &VpnvserverCachepolicyBindingDataSource{}
}

type VpnvserverCachepolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverCachepolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_cachepolicy_binding"
}

func (d *VpnvserverCachepolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverCachepolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverCachepolicyBindingDataSourceSchema()
}

func (d *VpnvserverCachepolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverCachepolicyBindingResourceModel
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
		ResourceType:             service.Vpnvserver_cachepolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_cachepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_cachepolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the supplied filter attrs.
	// bindpoint and policy are optional filters: only apply them when the caller
	// supplied a (known, non-null) value, otherwise match on the remaining attrs.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check bindpoint (optional filter)
		if !bindpoint_Name.IsNull() && !bindpoint_Name.IsUnknown() {
			if val, ok := v["bindpoint"].(string); !ok || val != bindpoint_Name.ValueString() {
				match = false
			}
		}

		// Check policy (optional filter)
		if match && !policy_Name.IsNull() && !policy_Name.IsUnknown() {
			if val, ok := v["policy"].(string); !ok || val != policy_Name.ValueString() {
				match = false
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_cachepolicy_binding with bindpoint %s not found", bindpoint_Name))
		return
	}

	vpnvserver_cachepolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

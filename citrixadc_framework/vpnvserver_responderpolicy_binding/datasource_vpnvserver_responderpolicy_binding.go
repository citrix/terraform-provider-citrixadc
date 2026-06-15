package vpnvserver_responderpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverResponderpolicyBindingDataSource)(nil)

func VPnvserverResponderpolicyBindingDataSource() datasource.DataSource {
	return &VpnvserverResponderpolicyBindingDataSource{}
}

type VpnvserverResponderpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverResponderpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_responderpolicy_binding"
}

func (d *VpnvserverResponderpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverResponderpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverResponderpolicyBindingDataSourceSchema()
}

func (d *VpnvserverResponderpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverResponderpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID. The binding identity is name (parent) +
	// policy; bindpoint is server-assigned and used only as an optional filter.
	name_Name := data.Name.ValueString()
	bindpoint_Name := data.Bindpoint
	policy_Name := data.Policy

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_responderpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_responderpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the matching policy (and bindpoint when supplied)
	foundIndex := -1
	for i, v := range dataArr {
		// Check policy (required lookup key)
		if val, ok := v["policy"].(string); !ok || policy_Name.IsNull() || val != policy_Name.ValueString() {
			continue
		}

		// Check bindpoint only when the user supplied it as an additional filter
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_responderpolicy_binding with policy %s not found", policy_Name.ValueString()))
		return
	}

	vpnvserver_responderpolicy_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

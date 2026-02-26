package vpnvserver_vpnnexthopserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverVpnnexthopserverBindingDataSource)(nil)

func VPnvserverVpnnexthopserverBindingDataSource() datasource.DataSource {
	return &VpnvserverVpnnexthopserverBindingDataSource{}
}

type VpnvserverVpnnexthopserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverVpnnexthopserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_vpnnexthopserver_binding"
}

func (d *VpnvserverVpnnexthopserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverVpnnexthopserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverVpnnexthopserverBindingDataSourceSchema()
}

func (d *VpnvserverVpnnexthopserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverVpnnexthopserverBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	nexthopserver_Name := data.Nexthopserver

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_vpnnexthopserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_vpnnexthopserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_vpnnexthopserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check nexthopserver
		if val, ok := v["nexthopserver"].(string); ok {
			if nexthopserver_Name.IsNull() || val != nexthopserver_Name.ValueString() {
				match = false
				continue
			}
		} else if !nexthopserver_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_vpnnexthopserver_binding with nexthopserver %s not found", nexthopserver_Name))
		return
	}

	vpnvserver_vpnnexthopserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

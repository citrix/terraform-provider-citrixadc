package vlan_linkset_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VlanLinksetBindingDataSource)(nil)

func VLanLinksetBindingDataSource() datasource.DataSource {
	return &VlanLinksetBindingDataSource{}
}

type VlanLinksetBindingDataSource struct {
	client *service.NitroClient
}

func (d *VlanLinksetBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_linkset_binding"
}

func (d *VlanLinksetBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VlanLinksetBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VlanLinksetBindingDataSourceSchema()
}

func (d *VlanLinksetBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VlanLinksetBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID
	vlanid_Name := fmt.Sprintf("%v", data.Vlanid.ValueInt64())
	ifnum_Name := data.Ifnum

	// The direct vlan_linkset_binding endpoint can return a keyless empty body;
	// read the bound interfaces from the aggregate parent endpoint instead.
	dataArr, err := vlan_linkset_bindingAggregateRead(d.client, vlanid_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vlan_linkset_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vlan_linkset_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ifnum
		if val, ok := v["ifnum"].(string); ok {
			if ifnum_Name.IsNull() || val != ifnum_Name.ValueString() {
				match = false
				continue
			}
		} else if !ifnum_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vlan_linkset_binding with ifnum %s not found", ifnum_Name))
		return
	}

	vlan_linkset_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

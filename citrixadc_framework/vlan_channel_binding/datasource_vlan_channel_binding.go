package vlan_channel_binding

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VlanChannelBindingDataSource)(nil)

func VLanChannelBindingDataSource() datasource.DataSource {
	return &VlanChannelBindingDataSource{}
}

type VlanChannelBindingDataSource struct {
	client *service.NitroClient
}

func (d *VlanChannelBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_channel_binding"
}

func (d *VlanChannelBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VlanChannelBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VlanChannelBindingDataSourceSchema()
}

func (d *VlanChannelBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VlanChannelBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID: vlanid is the parent, ifnum the bound entity.
	vlanidName := strconv.FormatInt(data.Vlanid.ValueInt64(), 10)
	ifnumName := data.Ifnum

	findParams := service.FindParams{
		ResourceType:             service.Vlan_channel_binding.Type(),
		ResourceName:             vlanidName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vlan_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vlan_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ifnum"].(string); ok {
			if ifnumName.IsNull() || val == ifnumName.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vlan_channel_binding with ifnum %s not found", ifnumName))
		return
	}

	vlan_channel_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

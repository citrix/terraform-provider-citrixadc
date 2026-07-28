package vrid_channel_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VridChannelBindingDataSource)(nil)

func VRidChannelBindingDataSource() datasource.DataSource {
	return &VridChannelBindingDataSource{}
}

type VridChannelBindingDataSource struct {
	client *service.NitroClient
}

func (d *VridChannelBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid_channel_binding"
}

func (d *VridChannelBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VridChannelBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VridChannelBindingDataSourceSchema()
}

func (d *VridChannelBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VridChannelBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the bound members via the aggregate parent endpoint (vrid_binding/<id>).
	id_Name := fmt.Sprintf("%v", data.VridId.ValueInt64())
	ifnum_Name := data.Ifnum

	dataArr, err := vrid_channel_bindingAggregateRead(d.client, id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vrid_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vrid_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum.
	//
	// Verified live: the carrying vrid_interface_binding rows do NOT echo "ifnum".
	// When ifnum is present, match on it; otherwise accept by row presence (the
	// parent vrid id already scopes the result).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ifnum"].(string); ok {
			if ifnum_Name.IsNull() || val == ifnum_Name.ValueString() {
				foundIndex = i
				break
			}
			continue
		}
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vrid_channel_binding with ifnum %s not found", ifnum_Name))
		return
	}

	vrid_channel_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package vrid6_channel_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Vrid6ChannelBindingDataSource)(nil)

func VRid6ChannelBindingDataSource() datasource.DataSource {
	return &Vrid6ChannelBindingDataSource{}
}

type Vrid6ChannelBindingDataSource struct {
	client *service.NitroClient
}

func (d *Vrid6ChannelBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid6_channel_binding"
}

func (d *Vrid6ChannelBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Vrid6ChannelBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Vrid6ChannelBindingDataSourceSchema()
}

func (d *Vrid6ChannelBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Vrid6ChannelBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID - read via the aggregate parent endpoint.
	id_Name := fmt.Sprintf("%v", data.VridId.ValueInt64())
	ifnum_Name := data.Ifnum

	dataArr, err := vrid6_channel_bindingAggregateRead(d.client, id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vrid6_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vrid6_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right member (ifnum).
	//
	// Verified live: the carrying vrid6_interface_binding rows do NOT echo "ifnum".
	// When present, match on it; otherwise accept by row presence (the parent vrid6
	// id already scopes the result).
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vrid6_channel_binding with ifnum %s not found", ifnum_Name))
		return
	}

	vrid6_channel_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

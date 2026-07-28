package vrid6_trackinterface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Vrid6TrackinterfaceBindingDataSource)(nil)

func VRid6TrackinterfaceBindingDataSource() datasource.DataSource {
	return &Vrid6TrackinterfaceBindingDataSource{}
}

type Vrid6TrackinterfaceBindingDataSource struct {
	client *service.NitroClient
}

func (d *Vrid6TrackinterfaceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid6_trackinterface_binding"
}

func (d *Vrid6TrackinterfaceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Vrid6TrackinterfaceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Vrid6TrackinterfaceBindingDataSourceSchema()
}

func (d *Vrid6TrackinterfaceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Vrid6TrackinterfaceBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID - read via the aggregate parent endpoint.
	id_Name := fmt.Sprintf("%v", data.VridId.ValueInt64())
	trackifnum_Name := data.Trackifnum

	dataArr, err := vrid6_trackinterface_bindingAggregateRead(d.client, id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vrid6_trackinterface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vrid6_trackinterface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right member (trackifnum)
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if val, ok := v["trackifnum"].(string); ok {
			if trackifnum_Name.IsNull() || val != trackifnum_Name.ValueString() {
				match = false
				continue
			}
		} else if !trackifnum_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vrid6_trackinterface_binding with trackifnum %s not found", trackifnum_Name))
		return
	}

	vrid6_trackinterface_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

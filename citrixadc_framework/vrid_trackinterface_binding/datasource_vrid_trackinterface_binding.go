package vrid_trackinterface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VridTrackinterfaceBindingDataSource)(nil)

func VRidTrackinterfaceBindingDataSource() datasource.DataSource {
	return &VridTrackinterfaceBindingDataSource{}
}

type VridTrackinterfaceBindingDataSource struct {
	client *service.NitroClient
}

func (d *VridTrackinterfaceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid_trackinterface_binding"
}

func (d *VridTrackinterfaceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VridTrackinterfaceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VridTrackinterfaceBindingDataSourceSchema()
}

func (d *VridTrackinterfaceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VridTrackinterfaceBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the bound members via the aggregate parent endpoint (vrid_binding/<id>).
	id_Name := fmt.Sprintf("%v", data.VridId.ValueInt64())
	trackifnum_Name := data.Trackifnum

	dataArr, err := vrid_trackinterface_bindingAggregateRead(d.client, id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vrid_trackinterface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vrid_trackinterface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right trackifnum.
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vrid_trackinterface_binding with trackifnum %s not found", trackifnum_Name))
		return
	}

	vrid_trackinterface_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

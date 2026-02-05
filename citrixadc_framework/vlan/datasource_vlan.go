package vlan

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VlanDataSource)(nil)

func VLanDataSource() datasource.DataSource {
	return &VlanDataSource{}
}

type VlanDataSource struct {
	client *service.NitroClient
}

func (d *VlanDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan"
}

func (d *VlanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VlanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VlanDataSourceSchema()
}

func (d *VlanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VlanResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	id_Name := fmt.Sprintf("%d", data.Vlanid.ValueInt64())

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Vlan.Type(), id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vlan, got error: %s", err))
		return
	}

	vlanSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

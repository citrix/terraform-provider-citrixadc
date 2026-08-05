package dnszone

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnszoneDataSource)(nil)

func DNszoneDataSource() datasource.DataSource {
	return &DnszoneDataSource{}
}

type DnszoneDataSource struct {
	client *service.NitroClient
}

func (d *DnszoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnszone"
}

func (d *DnszoneDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnszoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnszoneDataSourceSchema()
}

func (d *DnszoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnszoneResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Look up the zone by its name (zonename)
	zonename_Name := data.Zonename.ValueString()

	getResponseData, err := d.client.FindResource(service.Dnszone.Type(), zonename_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnszone, got error: %s", err))
		return
	}

	// Resource is missing
	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", "dnszone returned empty response.")
		return
	}

	dnszoneSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

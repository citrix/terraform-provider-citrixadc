package dnssoarec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnssoarecDataSource)(nil)

func DNssoarecDataSource() datasource.DataSource {
	return &DnssoarecDataSource{}
}

type DnssoarecDataSource struct {
	client *service.NitroClient
}

func (d *DnssoarecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssoarec"
}

func (d *DnssoarecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnssoarecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnssoarecDataSourceSchema()
}

func (d *DnssoarecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnssoarecDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource keyed on domain - look up by domain (plain value)
	domainName := data.Domain.ValueString()

	getResponseData, err := d.client.FindResource(service.Dnssoarec.Type(), domainName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnssoarec, got error: %s", err))
		return
	}

	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnssoarec with domain %s not found", domainName))
		return
	}

	dnssoarecDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

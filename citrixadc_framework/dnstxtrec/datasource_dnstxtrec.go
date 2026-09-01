package dnstxtrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnstxtrecDataSource)(nil)

func DNstxtrecDataSource() datasource.DataSource {
	return &DnstxtrecDataSource{}
}

type DnstxtrecDataSource struct {
	client *service.NitroClient
}

func (d *DnstxtrecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnstxtrec"
}

func (d *DnstxtrecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnstxtrecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnstxtrecDataSourceSchema()
}

func (d *DnstxtrecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnstxtrecDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Look up by the primary key (domain).
	domainName := data.Domain.ValueString()

	getResponseData, err := d.client.FindResource(service.Dnstxtrec.Type(), domainName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnstxtrec, got error: %s", err))
		return
	}

	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnstxtrec with domain %s not found", domainName))
		return
	}

	dnstxtrecDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

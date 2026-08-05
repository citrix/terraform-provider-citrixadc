package dnscnamerec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnscnamerecDataSource)(nil)

func DNscnamerecDataSource() datasource.DataSource {
	return &DnscnamerecDataSource{}
}

type DnscnamerecDataSource struct {
	client *service.NitroClient
}

func (d *DnscnamerecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnscnamerec"
}

func (d *DnscnamerecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnscnamerecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnscnamerecDataSourceSchema()
}

func (d *DnscnamerecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnscnamerecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Named resource - look up by the plain aliasname value.
	aliasname_Name := data.Aliasname.ValueString()

	getResponseData, err := d.client.FindResource(service.Dnscnamerec.Type(), aliasname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnscnamerec, got error: %s", err))
		return
	}

	dnscnamerecSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

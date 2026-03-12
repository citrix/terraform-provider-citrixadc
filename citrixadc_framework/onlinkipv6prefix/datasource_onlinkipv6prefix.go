package onlinkipv6prefix

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Onlinkipv6prefixDataSource)(nil)

func ONlinkipv6prefixDataSource() datasource.DataSource {
	return &Onlinkipv6prefixDataSource{}
}

type Onlinkipv6prefixDataSource struct {
	client *service.NitroClient
}

func (d *Onlinkipv6prefixDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_onlinkipv6prefix"
}

func (d *Onlinkipv6prefixDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Onlinkipv6prefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Onlinkipv6prefixDataSourceSchema()
}

func (d *Onlinkipv6prefixDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Onlinkipv6prefixResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	ipv6prefix_Name := data.Ipv6prefix.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Onlinkipv6prefix.Type(), ipv6prefix_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read onlinkipv6prefix, got error: %s", err))
		return
	}

	onlinkipv6prefixSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

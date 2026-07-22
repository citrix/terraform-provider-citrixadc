package sslhpkekey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SslhpkekeyDataSource)(nil)

func SSlhpkekeyDataSource() datasource.DataSource {
	return &SslhpkekeyDataSource{}
}

type SslhpkekeyDataSource struct {
	client *service.NitroClient
}

func (d *SslhpkekeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslhpkekey"
}

func (d *SslhpkekeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslhpkekeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslhpkekeyDataSourceSchema()
}

func (d *SslhpkekeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslhpkekeyResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	hpkekeyname_Name := data.Hpkekeyname.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Sslhpkekey.Type(), hpkekeyname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslhpkekey, got error: %s", err))
		return
	}

	sslhpkekeySetAttrFromGet(ctx, &data, getResponseData)

	// Datasource has no Create; set the ID explicitly (single-key plain value).
	data.Id = types.StringValue(data.Hpkekeyname.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

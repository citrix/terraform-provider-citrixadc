package sslzerotouchparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslzerotouchparamDataSourceType)(nil)

func SslzerotouchparamDataSource() datasource.DataSource {
	return &SslzerotouchparamDataSourceType{}
}

type SslzerotouchparamDataSourceType struct {
	client *service.NitroClient
}

func (d *SslzerotouchparamDataSourceType) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslzerotouchparam"
}

func (d *SslzerotouchparamDataSourceType) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslzerotouchparamDataSourceType) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslzerotouchparamDataSourceSchema()
}

func (d *SslzerotouchparamDataSourceType) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslzerotouchparamResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID (singleton)
	getResponseData, err := d.client.FindResource(service.Sslzerotouchparam.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslzerotouchparam, got error: %s", err))
		return
	}

	sslzerotouchparamSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

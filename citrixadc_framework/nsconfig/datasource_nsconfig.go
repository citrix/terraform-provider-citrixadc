package nsconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsconfigDataSource)(nil)

func NSconfigDataSource() datasource.DataSource {
	return &NsconfigDataSource{}
}

type NsconfigDataSource struct {
	client *service.NitroClient
}

func (d *NsconfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsconfig"
}

func (d *NsconfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsconfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsconfigDataSourceSchema()
}

func (d *NsconfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsconfigResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Nsconfig.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsconfig, got error: %s", err))
		return
	}

	nsconfigSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

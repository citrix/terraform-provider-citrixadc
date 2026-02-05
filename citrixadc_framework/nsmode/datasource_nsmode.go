package nsmode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsmodeDataSource)(nil)

func NSmodeDataSource() datasource.DataSource {
	return &NsmodeDataSource{}
}

type NsmodeDataSource struct {
	client *service.NitroClient
}

func (d *NsmodeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmode"
}

func (d *NsmodeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsmodeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsmodeDataSourceSchema()
}

func (d *NsmodeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsmodeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use FindResourceArrayWithParams to match v2 implementation
	findParams := service.FindParams{
		ResourceType: "nsmode",
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsmode, got error: %s", err))
		return
	}

	if len(dataArr) != 1 {
		resp.Diagnostics.AddError("Unexpected Result", fmt.Sprintf("Expected 1 nsmode result, got %d", len(dataArr)))
		return
	}

	getResponseData := dataArr[0]
	nsmodeSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

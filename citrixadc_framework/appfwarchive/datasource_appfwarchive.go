package appfwarchive

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*AppfwarchiveDataSource)(nil)

func APpfwarchiveDataSource() datasource.DataSource {
	return &AppfwarchiveDataSource{}
}

type AppfwarchiveDataSource struct {
	client *service.NitroClient
}

func (d *AppfwarchiveDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwarchive"
}

func (d *AppfwarchiveDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwarchiveDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwarchiveDataSourceSchema()
}

func (d *AppfwarchiveDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwarchiveResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NITRO appfwarchive only exposes `get (all)` and the response carries no
	// per-archive identifying fields. Confirm the resource collection is non-
	// empty, then preserve the caller-supplied name in state.
	findParams := service.FindParams{
		ResourceType:             service.Appfwarchive.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwarchive, got error: %s", err))
		return
	}
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "appfwarchive returned empty array")
		return
	}

	appfwarchiveSetAttrFromGet(ctx, &data, dataArr[0])

	// Datasource has no Create step that would seed the ID; set it explicitly
	// from the lookup name (single_unique => plain value, matching Create).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package appfwgrpccontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*AppfwgrpccontenttypeDataSource)(nil)

func APpfwgrpccontenttypeDataSource() datasource.DataSource {
	return &AppfwgrpccontenttypeDataSource{}
}

type AppfwgrpccontenttypeDataSource struct {
	client *service.NitroClient
}

func (d *AppfwgrpccontenttypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwgrpccontenttype"
}

func (d *AppfwgrpccontenttypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwgrpccontenttypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwgrpccontenttypeDataSourceSchema()
}

func (d *AppfwgrpccontenttypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwgrpccontenttypeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	grpccontenttypevalue_Name := data.Grpccontenttypevalue.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Appfwgrpccontenttype.Type(), grpccontenttypevalue_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwgrpccontenttype, got error: %s", err))
		return
	}

	appfwgrpccontenttypeSetAttrFromGet(ctx, &data, getResponseData)

	// Datasource never calls Create — set ID explicitly here (single-key resource)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Grpccontenttypevalue.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

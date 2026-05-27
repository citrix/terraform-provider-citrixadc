package appfwgrpcwebjsoncontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*AppfwgrpcwebjsoncontenttypeDataSource)(nil)

func APpfwgrpcwebjsoncontenttypeDataSource() datasource.DataSource {
	return &AppfwgrpcwebjsoncontenttypeDataSource{}
}

type AppfwgrpcwebjsoncontenttypeDataSource struct {
	client *service.NitroClient
}

func (d *AppfwgrpcwebjsoncontenttypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwgrpcwebjsoncontenttype"
}

func (d *AppfwgrpcwebjsoncontenttypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwgrpcwebjsoncontenttypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwgrpcwebjsoncontenttypeDataSourceSchema()
}

func (d *AppfwgrpcwebjsoncontenttypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwgrpcwebjsoncontenttypeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	grpcwebjsoncontenttypevalue_Name := data.Grpcwebjsoncontenttypevalue.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Appfwgrpcwebjsoncontenttype.Type(), grpcwebjsoncontenttypevalue_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwgrpcwebjsoncontenttype, got error: %s", err))
		return
	}

	appfwgrpcwebjsoncontenttypeSetAttrFromGet(ctx, &data, getResponseData)

	// Set ID explicitly for datasource (SetAttrFromGet does not set ID)
	data.Id = types.StringValue(data.Grpcwebjsoncontenttypevalue.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

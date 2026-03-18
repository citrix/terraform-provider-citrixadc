package appfwjsoncontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwjsoncontenttypeDataSource)(nil)

func APpfwjsoncontenttypeDataSource() datasource.DataSource {
	return &AppfwjsoncontenttypeDataSource{}
}

type AppfwjsoncontenttypeDataSource struct {
	client *service.NitroClient
}

func (d *AppfwjsoncontenttypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwjsoncontenttype"
}

func (d *AppfwjsoncontenttypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwjsoncontenttypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwjsoncontenttypeDataSourceSchema()
}

func (d *AppfwjsoncontenttypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwjsoncontenttypeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	jsoncontenttypevalue_Name := data.Jsoncontenttypevalue.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Appfwjsoncontenttype.Type(), jsoncontenttypevalue_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwjsoncontenttype, got error: %s", err))
		return
	}

	appfwjsoncontenttypeSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

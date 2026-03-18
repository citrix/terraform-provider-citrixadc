package appfwurlencodedformcontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwurlencodedformcontenttypeDataSource)(nil)

func APpfwurlencodedformcontenttypeDataSource() datasource.DataSource {
	return &AppfwurlencodedformcontenttypeDataSource{}
}

type AppfwurlencodedformcontenttypeDataSource struct {
	client *service.NitroClient
}

func (d *AppfwurlencodedformcontenttypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwurlencodedformcontenttype"
}

func (d *AppfwurlencodedformcontenttypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwurlencodedformcontenttypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwurlencodedformcontenttypeDataSourceSchema()
}

func (d *AppfwurlencodedformcontenttypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	urlencodedformcontenttypevalue_Name := data.Urlencodedformcontenttypevalue.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Appfwurlencodedformcontenttype.Type(), urlencodedformcontenttypevalue_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwurlencodedformcontenttype, got error: %s", err))
		return
	}

	appfwurlencodedformcontenttypeSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

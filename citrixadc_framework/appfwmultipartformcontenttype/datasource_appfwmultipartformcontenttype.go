package appfwmultipartformcontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AppfwmultipartformcontenttypeDataSource)(nil)

func APpfwmultipartformcontenttypeDataSource() datasource.DataSource {
	return &AppfwmultipartformcontenttypeDataSource{}
}

type AppfwmultipartformcontenttypeDataSource struct {
	client *service.NitroClient
}

func (d *AppfwmultipartformcontenttypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwmultipartformcontenttype"
}

func (d *AppfwmultipartformcontenttypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwmultipartformcontenttypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwmultipartformcontenttypeDataSourceSchema()
}

func (d *AppfwmultipartformcontenttypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwmultipartformcontenttypeResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	multipartformcontenttypevalue_Name := data.Multipartformcontenttypevalue.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Appfwmultipartformcontenttype.Type(), multipartformcontenttypevalue_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwmultipartformcontenttype, got error: %s", err))
		return
	}

	appfwmultipartformcontenttypeSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package authorizationpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*AuthorizationpolicylabelDataSource)(nil)

func AUthorizationpolicylabelDataSource() datasource.DataSource {
	return &AuthorizationpolicylabelDataSource{}
}

type AuthorizationpolicylabelDataSource struct {
	client *service.NitroClient
}

func (d *AuthorizationpolicylabelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authorizationpolicylabel"
}

func (d *AuthorizationpolicylabelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AuthorizationpolicylabelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AuthorizationpolicylabelDataSourceSchema()
}

func (d *AuthorizationpolicylabelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AuthorizationpolicylabelResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	labelname_Name := data.Labelname.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Authorizationpolicylabel.Type(), labelname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read authorizationpolicylabel, got error: %s", err))
		return
	}

	authorizationpolicylabelSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

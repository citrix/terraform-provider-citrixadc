package lsnrtspalgsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*LsnrtspalgsessionDataSource)(nil)

func LSnrtspalgsessionDataSource() datasource.DataSource {
	return &LsnrtspalgsessionDataSource{}
}

type LsnrtspalgsessionDataSource struct {
	client *service.NitroClient
}

func (d *LsnrtspalgsessionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnrtspalgsession"
}

func (d *LsnrtspalgsessionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnrtspalgsessionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnrtspalgsessionDataSourceSchema()
}

func (d *LsnrtspalgsessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnrtspalgsessionResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	sessionid_Name := data.Sessionid.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Lsnrtspalgsession.Type(), sessionid_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnrtspalgsession, got error: %s", err))
		return
	}

	lsnrtspalgsessionSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// The datasource has no Create; set the synthetic ID here to match the
	// resource's ID format (plain sessionid value).
	data.Id = types.StringValue(data.Sessionid.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

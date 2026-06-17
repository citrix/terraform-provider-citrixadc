package nslimitselector

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NslimitselectorDataSource)(nil)

func NSlimitselectorDataSource() datasource.DataSource {
	return &NslimitselectorDataSource{}
}

type NslimitselectorDataSource struct {
	client *service.NitroClient
}

func (d *NslimitselectorDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitselector"
}

func (d *NslimitselectorDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NslimitselectorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NslimitselectorDataSourceSchema()
}

func (d *NslimitselectorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NslimitselectorResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	selectorname_Name := data.Selectorname.ValueString()

	var getResponseData map[string]interface{}
	var err error

	// NITRO returns this object under the "streamselector" key (nslimitselector
	// is an alias of streamselector); a typed GET against "nslimitselector"
	// returns a body the client cannot map back, so read via "streamselector".
	getResponseData, err = d.client.FindResource(service.Streamselector.Type(), selectorname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nslimitselector, got error: %s", err))
		return
	}

	nslimitselectorSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package nat64param

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Nat64paramDataSource)(nil)

func NAt64paramDataSource() datasource.DataSource {
	return &Nat64paramDataSource{}
}

type Nat64paramDataSource struct {
	client *service.NitroClient
}

func (d *Nat64paramDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nat64param"
}

func (d *Nat64paramDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Nat64paramDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Nat64paramDataSourceSchema()
}

func (d *Nat64paramDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Nat64paramResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	td_Name := fmt.Sprintf("%d", data.Td.ValueInt64())

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Nat64param.Type(), td_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nat64param, got error: %s", err))
		return
	}

	nat64paramSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

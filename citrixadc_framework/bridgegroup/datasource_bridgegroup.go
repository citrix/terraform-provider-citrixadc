package bridgegroup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*BridgegroupDataSource)(nil)

func BRidgegroupDataSource() datasource.DataSource {
	return &BridgegroupDataSource{}
}

type BridgegroupDataSource struct {
	client *service.NitroClient
}

func (d *BridgegroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup"
}

func (d *BridgegroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BridgegroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BridgegroupDataSourceSchema()
}

func (d *BridgegroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BridgegroupResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	id_Name := types.StringValue(fmt.Sprintf("%d", data.Bridgegroupid.ValueInt64()))

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Bridgegroup.Type(), id_Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup, got error: %s", err))
		return
	}

	bridgegroupSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

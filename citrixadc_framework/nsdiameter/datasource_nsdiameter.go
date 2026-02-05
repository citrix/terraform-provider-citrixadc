package nsdiameter

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsdiameterDataSource)(nil)

func NSdiameterDataSource() datasource.DataSource {
	return &NsdiameterDataSource{}
}

type NsdiameterDataSource struct {
	client *service.NitroClient
}

func (d *NsdiameterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsdiameter"
}

func (d *NsdiameterDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsdiameterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsdiameterDataSourceSchema()
}

func (d *NsdiameterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsdiameterResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	ownernode_Name := fmt.Sprintf("%d", data.Ownernode.ValueInt64())

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Nsdiameter.Type(), ownernode_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsdiameter, got error: %s", err))
		return
	}

	nsdiameterSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

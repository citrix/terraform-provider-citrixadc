package subscriberparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SubscriberparamDataSource)(nil)

func SUbscriberparamDataSource() datasource.DataSource {
	return &SubscriberparamDataSource{}
}

type SubscriberparamDataSource struct {
	client *service.NitroClient
}

func (d *SubscriberparamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberparam"
}

func (d *SubscriberparamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SubscriberparamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SubscriberparamDataSourceSchema()
}

func (d *SubscriberparamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SubscriberparamResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Subscriberparam.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read subscriberparam, got error: %s", err))
		return
	}

	subscriberparamSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

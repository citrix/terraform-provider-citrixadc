package subscriberradiusinterface

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SubscriberradiusinterfaceDataSource)(nil)

func SUbscriberradiusinterfaceDataSource() datasource.DataSource {
	return &SubscriberradiusinterfaceDataSource{}
}

type SubscriberradiusinterfaceDataSource struct {
	client *service.NitroClient
}

func (d *SubscriberradiusinterfaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberradiusinterface"
}

func (d *SubscriberradiusinterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SubscriberradiusinterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SubscriberradiusinterfaceDataSourceSchema()
}

func (d *SubscriberradiusinterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SubscriberradiusinterfaceResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Subscriberradiusinterface.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read subscriberradiusinterface, got error: %s", err))
		return
	}

	subscriberradiusinterfaceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

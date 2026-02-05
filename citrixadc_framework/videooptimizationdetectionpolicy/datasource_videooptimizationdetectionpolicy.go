package videooptimizationdetectionpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VideooptimizationdetectionpolicyDataSource)(nil)

func VIdeooptimizationdetectionpolicyDataSource() datasource.DataSource {
	return &VideooptimizationdetectionpolicyDataSource{}
}

type VideooptimizationdetectionpolicyDataSource struct {
	client *service.NitroClient
}

func (d *VideooptimizationdetectionpolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicy"
}

func (d *VideooptimizationdetectionpolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VideooptimizationdetectionpolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VideooptimizationdetectionpolicyDataSourceSchema()
}

func (d *VideooptimizationdetectionpolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VideooptimizationdetectionpolicyResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	name_Name := data.Name.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Videooptimizationdetectionpolicy.Type(), name_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicy, got error: %s", err))
		return
	}

	videooptimizationdetectionpolicySetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

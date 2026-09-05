package gslbservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// gslbserviceLbmonitorbindingDSAttrs returns the nested attributes for the
// inline lbmonitorbinding block on the data source. It is defined here (not in
// datasource_schema.go) so the merged schema file stays a flat top-level
// attribute/block map matching the GslbserviceDataSourceModel.
func gslbserviceLbmonitorbindingDSAttrs() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"weight": schema.Int64Attribute{
			Optional:    true,
			Computed:    true,
			Description: "Weight to assign to the monitor-service binding.",
		},
		"monitor_name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "Name of the monitor bound to the GSLB service.",
		},
		"monstate": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "State of the monitor bound to the GSLB service.",
		},
	}
}

var _ datasource.DataSource = (*GslbserviceDataSource)(nil)

func GSlbserviceDataSource() datasource.DataSource {
	return &GslbserviceDataSource{}
}

type GslbserviceDataSource struct {
	client *service.NitroClient
}

func (d *GslbserviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice"
}

func (d *GslbserviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *GslbserviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = GslbserviceDataSourceSchema()
}

func (d *GslbserviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GslbserviceDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	servicename_Name := data.Servicename.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Gslbservice.Type(), servicename_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice, got error: %s", err))
		return
	}

	gslbserviceDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

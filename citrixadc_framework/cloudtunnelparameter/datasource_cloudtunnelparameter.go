package cloudtunnelparameter

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*CloudtunnelparameterDataSource)(nil)

func CLoudtunnelparameterDataSource() datasource.DataSource {
	return &CloudtunnelparameterDataSource{}
}

type CloudtunnelparameterDataSource struct {
	client *service.NitroClient
}

func (d *CloudtunnelparameterDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudtunnelparameter"
}

func (d *CloudtunnelparameterDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *CloudtunnelparameterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = CloudtunnelparameterDataSourceSchema()
}

func (d *CloudtunnelparameterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CloudtunnelparameterResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Cloudtunnelparameter.Type(), "")
	if err != nil {
		// cloudtunnelparameter is feature-gated: on platforms/releases where the
		// feature is not enabled NITRO returns "Feature not supported in this release"
		// (or "Operation not supported on this platform"). Treat those as a non-fatal
		// read (leave the values null / static ID) rather than failing.
		if strings.Contains(err.Error(), "not supported on this platform") ||
			strings.Contains(err.Error(), "Feature not supported") {
			tflog.Warn(ctx, "cloudtunnelparameter GET not supported on this platform/release; datasource returns null values")
			data.Controllerfqdn = types.StringNull()
			data.Fqdn = types.StringNull()
			data.Resourcelocation = types.StringNull()
			data.Subnetresourcelocationmappings = types.StringNull()
			data.Id = types.StringValue("cloudtunnelparameter-config")
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cloudtunnelparameter, got error: %s", err))
		return
	}

	cloudtunnelparameterSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

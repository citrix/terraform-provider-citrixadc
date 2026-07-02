package cloudparaminternal

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*CloudparaminternalDataSource)(nil)

func CLoudparaminternalDataSource() datasource.DataSource {
	return &CloudparaminternalDataSource{}
}

type CloudparaminternalDataSource struct {
	client *service.NitroClient
}

func (d *CloudparaminternalDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudparaminternal"
}

func (d *CloudparaminternalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *CloudparaminternalDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = CloudparaminternalDataSourceSchema()
}

func (d *CloudparaminternalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CloudparaminternalResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Cloudparaminternal.Type(), "")
	if err != nil {
		// cloudparaminternal GET/show is platform-gated: on unsupported platforms
		// NITRO returns "Operation not supported on this platform". Treat that as a
		// non-fatal read (leave the value null / static ID) rather than failing.
		if strings.Contains(err.Error(), "not supported on this platform") {
			tflog.Warn(ctx, "cloudparaminternal GET not supported on this platform; datasource returns null value")
			data.Nonftumode = types.StringNull()
			data.Id = types.StringValue("cloudparaminternal-config")
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cloudparaminternal, got error: %s", err))
		return
	}

	cloudparaminternalSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
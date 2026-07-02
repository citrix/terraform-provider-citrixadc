package sslcertbundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslcertbundleDataSource)(nil)

func SSlcertbundleDataSource() datasource.DataSource {
	return &SslcertbundleDataSource{}
}

type SslcertbundleDataSource struct {
	client *service.NitroClient
}

func (d *SslcertbundleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertbundle"
}

func (d *SslcertbundleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslcertbundleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslcertbundleDataSourceSchema()
}

func (d *SslcertbundleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslcertbundleResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// No get-byname endpoint - get all and filter by name
	name_value := data.Name.ValueString()

	allResources, err := d.client.FindAllResources(service.Sslcertbundle.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslcertbundle, got error: %s", err))
		return
	}

	var getResponseData map[string]interface{}
	for _, v := range allResources {
		if n, ok := v["name"].(string); ok && n == name_value {
			getResponseData = v
			break
		}
	}

	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslcertbundle: no record found with name %s", name_value))
		return
	}

	sslcertbundleSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

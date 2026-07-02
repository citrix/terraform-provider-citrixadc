package ssldhfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SsldhfileDataSource)(nil)

func SSldhfileDataSource() datasource.DataSource {
	return &SsldhfileDataSource{}
}

type SsldhfileDataSource struct {
	client *service.NitroClient
}

func (d *SsldhfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldhfile"
}

func (d *SsldhfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SsldhfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SsldhfileDataSourceSchema()
}

func (d *SsldhfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SsldhfileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ssldhfile has NO get-by-name endpoint (GET /ssldhfile/<name> => errorcode
	// 1090 "No such argument [arguid]"). Get all records and filter by name.
	name := data.Name.ValueString()

	allResources, err := d.client.FindAllResources(service.Ssldhfile.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ssldhfile, got error: %s", err))
		return
	}

	var getResponseData map[string]interface{}
	for _, v := range allResources {
		if n, ok := v["name"].(string); ok && n == name {
			getResponseData = v
			break
		}
	}

	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ssldhfile: no record found with name %s", name))
		return
	}

	ssldhfileSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

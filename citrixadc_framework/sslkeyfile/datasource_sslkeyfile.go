package sslkeyfile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslkeyfileDataSource)(nil)

func SSlkeyfileDataSource() datasource.DataSource {
	return &SslkeyfileDataSource{}
}

type SslkeyfileDataSource struct {
	client *service.NitroClient
}

func (d *SslkeyfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslkeyfile"
}

func (d *SslkeyfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslkeyfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslkeyfileDataSourceSchema()
}

func (d *SslkeyfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslkeyfileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// sslkeyfile has NO get-by-name endpoint (GET /sslkeyfile/<name> => errorcode
	// 1090 "No such argument [arguid]"). Get all records and filter by name.
	name := data.Name.ValueString()

	allResources, err := d.client.FindAllResources(service.Sslkeyfile.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslkeyfile, got error: %s", err))
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslkeyfile: no record found with name %s", name))
		return
	}

	sslkeyfileSetAttrFromGetForDatasource(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

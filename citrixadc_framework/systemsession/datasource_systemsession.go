package systemsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SystemsessionDataSource)(nil)

func SYstemsessionDataSource() datasource.DataSource {
	return &SystemsessionDataSource{}
}

type SystemsessionDataSource struct {
	client *service.NitroClient
}

func (d *SystemsessionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsession"
}

func (d *SystemsessionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemsessionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemsessionDataSourceSchema()
}

func (d *SystemsessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemsessionDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sidName := fmt.Sprintf("%v", data.Sid.ValueInt64())

	getResponseData, err := d.client.FindResource(service.Systemsession.Type(), sidName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemsession, got error: %s", err))
		return
	}

	systemsessionDataSourceSetAttrFromGet(ctx, &data, getResponseData)

	// Datasource has no Create — set its ID here.
	data.Id = types.StringValue(sidName)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

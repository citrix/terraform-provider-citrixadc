package systemscalablemgmtthreads

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SystemscalablemgmtthreadsDataSource)(nil)

func NewSystemscalablemgmtthreadsDataSource() datasource.DataSource {
	return &SystemscalablemgmtthreadsDataSource{}
}

type SystemscalablemgmtthreadsDataSource struct {
	client *service.NitroClient
}

func (d *SystemscalablemgmtthreadsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemscalablemgmtthreads"
}

func (d *SystemscalablemgmtthreadsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemscalablemgmtthreadsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemscalablemgmtthreadsDataSourceSchema()
}

func (d *SystemscalablemgmtthreadsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemscalablemgmtthreadsDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Simple find without ID (singleton feature).
	getResponseData, err := d.client.FindResource(service.Systemscalablemgmtthreads.Type(), "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemscalablemgmtthreads, got error: %s", err))
		return
	}

	systemscalablemgmtthreadsSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

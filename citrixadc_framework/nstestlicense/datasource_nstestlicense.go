package nstestlicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NstestlicenseDataSource)(nil)

func NStestlicenseDataSource() datasource.DataSource {
	return &NstestlicenseDataSource{}
}

type NstestlicenseDataSource struct {
	client *service.NitroClient
}

func (d *NstestlicenseDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstestlicense"
}

func (d *NstestlicenseDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NstestlicenseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NstestlicenseDataSourceSchema()
}

func (d *NstestlicenseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NstestlicenseDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// nstestlicense has a keyless get(all); fetch the whole object.
	getResponseData, err := d.client.FindResource(nstestlicenseResourceType, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nstestlicense, got error: %s", err))
		return
	}

	nstestlicenseSetAttrFromGet(ctx, &data, getResponseData)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

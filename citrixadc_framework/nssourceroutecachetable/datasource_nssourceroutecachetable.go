package nssourceroutecachetable

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NssourceroutecachetableDataSource)(nil)

func NSsourceroutecachetableDataSource() datasource.DataSource {
	return &NssourceroutecachetableDataSource{}
}

type NssourceroutecachetableDataSource struct {
	client *service.NitroClient
}

func (d *NssourceroutecachetableDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssourceroutecachetable"
}

func (d *NssourceroutecachetableDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NssourceroutecachetableDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NssourceroutecachetableDataSourceSchema()
}

func (d *NssourceroutecachetableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NssourceroutecachetableDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// nssourceroutecachetable has a keyless get(all); fetch the whole object.
	// NOTE: get(all) returns a table (list) of cache entries; FindResource
	// returns the first element (best-effort). If the table is empty, the GET
	// yields no resource and an error is surfaced.
	getResponseData, err := d.client.FindResource(nssourceroutecachetableResourceType, "")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nssourceroutecachetable, got error: %s", err))
		return
	}

	nssourceroutecachetableSetAttrFromGet(ctx, &data, getResponseData)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

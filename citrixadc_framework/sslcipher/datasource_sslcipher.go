package sslcipher

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SslcipherDataSource)(nil)

func SSlcipherDataSource() datasource.DataSource {
	return &SslcipherDataSource{}
}

type SslcipherDataSource struct {
	client *service.NitroClient
}

func (d *SslcipherDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher"
}

func (d *SslcipherDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslcipherDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslcipherDataSourceSchema()
}

func (d *SslcipherDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslcipherResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ciphergroupname := data.Ciphergroupname.ValueString()

	// Mirror the resource read: some NetScaler versions do not support the
	// per-name GET, so use FindAllResources and filter by ciphergroupname.
	dataArr, err := d.client.FindAllResources(service.Sslcipher.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher, got error: %s", err))
		return
	}

	found := false
	for _, v := range dataArr {
		if name, ok := v["ciphergroupname"].(string); ok && name == ciphergroupname {
			data.Ciphergroupname = types.StringValue(name)
			found = true
			break
		}
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslcipher %s not found", ciphergroupname))
		return
	}

	// Populate the bindings from the appliance.
	bindingSet, diags := readSslcipherCiphersuiteBindings(ctx, d.client, ciphergroupname)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Ciphersuitebinding = bindingSet

	data.Id = types.StringValue(ciphergroupname)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package sslprofile_sslciphersuite_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslprofileSslciphersuiteBindingDataSource)(nil)

func SSlprofileSslciphersuiteBindingDataSource() datasource.DataSource {
	return &SslprofileSslciphersuiteBindingDataSource{}
}

type SslprofileSslciphersuiteBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslprofileSslciphersuiteBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslprofile_sslciphersuite_binding"
}

func (d *SslprofileSslciphersuiteBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslprofileSslciphersuiteBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslprofileSslciphersuiteBindingDataSourceSchema()
}

func (d *SslprofileSslciphersuiteBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslprofileSslciphersuiteBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	ciphername_Name := data.Ciphername

	var dataArr []map[string]interface{}
	var err error

	// This binding subresource only narrows the by-name GET via ?filter=<key>:<value>
	// (it rejects plain by-name and ?args=); mirror the resource Read.
	findParams := service.FindParams{
		ResourceType:             service.Sslprofile_sslciphersuite_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	if !ciphername_Name.IsNull() && !ciphername_Name.IsUnknown() {
		findParams.FilterMap = map[string]string{"ciphername": ciphername_Name.ValueString()}
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslprofile_sslciphersuite_binding, got error: %s", err))
		return
	}

	// The typed binding GET is not reflected over REST on this firmware (always an empty
	// "Done" body even when the binding exists). When it does return rows, honour them;
	// otherwise fall back to the umbrella sslprofile_binding endpoint (under
	// sslprofile_sslcipher_binding[], keyed by cipheraliasname/ciphername). Mirror the
	// resource Read exactly.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ciphername"].(string); ok {
			if !ciphername_Name.IsNull() && val == ciphername_Name.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	if foundIndex != -1 {
		sslprofile_sslciphersuite_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// Typed GET returned no matching row; fall back to the umbrella endpoint.
	row := findCipherBindingViaUmbrella(ctx, d.client, name_Name, ciphername_Name.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if row == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslprofile_sslciphersuite_binding with ciphername %s not found", ciphername_Name.ValueString()))
		return
	}

	// The umbrella row carries name+ciphername; the datasource setter rebuilds the
	// composite ID and copies all schema attributes (ciphername, cipherpriority,
	// description, name).
	sslprofile_sslciphersuite_bindingSetAttrFromGetForDatasource(ctx, &data, row)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

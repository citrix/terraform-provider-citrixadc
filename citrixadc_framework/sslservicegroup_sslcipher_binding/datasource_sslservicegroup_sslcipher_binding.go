package sslservicegroup_sslcipher_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslservicegroupSslcipherBindingDataSource)(nil)

func SSlservicegroupSslcipherBindingDataSource() datasource.DataSource {
	return &SslservicegroupSslcipherBindingDataSource{}
}

type SslservicegroupSslcipherBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslservicegroupSslcipherBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcipher_binding"
}

func (d *SslservicegroupSslcipherBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslservicegroupSslcipherBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslservicegroupSslcipherBindingDataSourceSchema()
}

func (d *SslservicegroupSslcipherBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslservicegroupSslcipherBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	ciphername_Name := data.Ciphername

	var dataArr []map[string]interface{}
	var err error

	// This binding subresource is not reflected over the typed GET on this firmware
	// (returns an empty {"message":"Done"} body); it is only reflected via the umbrella
	// sslservicegroup_binding endpoint. Try the typed GET first, then fall back to the
	// umbrella. Mirror the resource Read exactly.
	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslcipher_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	if !ciphername_Name.IsNull() && !ciphername_Name.IsUnknown() {
		findParams.FilterMap = map[string]string{"ciphername": ciphername_Name.ValueString()}
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcipher_binding, got error: %s", err))
		return
	}

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
		sslservicegroup_sslcipher_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// Typed GET returned no matching row; fall back to the umbrella endpoint.
	row := findCipherBindingViaUmbrella(ctx, d.client, servicegroupname_Name, ciphername_Name.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if row == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslservicegroup_sslcipher_binding with ciphername %s not found", ciphername_Name.ValueString()))
		return
	}

	sslservicegroup_sslcipher_bindingSetAttrFromGet(ctx, &data, row)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

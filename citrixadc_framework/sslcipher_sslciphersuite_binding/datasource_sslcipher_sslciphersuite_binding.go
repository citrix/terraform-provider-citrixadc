package sslcipher_sslciphersuite_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslcipherSslciphersuiteBindingDataSource)(nil)

func SSlcipherSslciphersuiteBindingDataSource() datasource.DataSource {
	return &SslcipherSslciphersuiteBindingDataSource{}
}

type SslcipherSslciphersuiteBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslcipherSslciphersuiteBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher_sslciphersuite_binding"
}

func (d *SslcipherSslciphersuiteBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslcipherSslciphersuiteBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslcipherSslciphersuiteBindingDataSourceSchema()
}

func (d *SslcipherSslciphersuiteBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslcipherSslciphersuiteBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	ciphergroupname_Name := data.Ciphergroupname.ValueString()
	ciphername_Name := data.Ciphername

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslcipher_sslciphersuite_binding.Type(),
		ResourceName:             ciphergroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslcipher_sslciphersuite_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslcipher_sslciphersuite_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ciphername
		if val, ok := v["ciphername"].(string); ok {
			if ciphername_Name.IsNull() || val != ciphername_Name.ValueString() {
				match = false
				continue
			}
		} else if !ciphername_Name.IsNull() {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslcipher_sslciphersuite_binding with ciphername %s not found", ciphername_Name))
		return
	}

	sslcipher_sslciphersuite_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

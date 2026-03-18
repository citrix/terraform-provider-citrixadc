package sslcertkey_sslocspresponder_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslcertkeySslocspresponderBindingDataSource)(nil)

func SSlcertkeySslocspresponderBindingDataSource() datasource.DataSource {
	return &SslcertkeySslocspresponderBindingDataSource{}
}

type SslcertkeySslocspresponderBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslcertkeySslocspresponderBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkey_sslocspresponder_binding"
}

func (d *SslcertkeySslocspresponderBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslcertkeySslocspresponderBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslcertkeySslocspresponderBindingDataSourceSchema()
}

func (d *SslcertkeySslocspresponderBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslcertkeySslocspresponderBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	certkey_Name := data.Certkey.ValueString()
	ocspresponder_Name := data.Ocspresponder

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslcertkey_sslocspresponder_binding.Type(),
		ResourceName:             certkey_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslcertkey_sslocspresponder_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslcertkey_sslocspresponder_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ocspresponder
		if val, ok := v["ocspresponder"].(string); ok {
			if ocspresponder_Name.IsNull() || val != ocspresponder_Name.ValueString() {
				match = false
				continue
			}
		} else if !ocspresponder_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslcertkey_sslocspresponder_binding with ocspresponder %s not found", ocspresponder_Name.ValueString()))
		return
	}

	sslcertkey_sslocspresponder_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

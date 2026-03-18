package sslservicegroup_sslcertkey_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslservicegroupSslcertkeyBindingDataSource)(nil)

func SSlservicegroupSslcertkeyBindingDataSource() datasource.DataSource {
	return &SslservicegroupSslcertkeyBindingDataSource{}
}

type SslservicegroupSslcertkeyBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslservicegroupSslcertkeyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcertkey_binding"
}

func (d *SslservicegroupSslcertkeyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslservicegroupSslcertkeyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslservicegroupSslcertkeyBindingDataSourceSchema()
}

func (d *SslservicegroupSslcertkeyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslservicegroupSslcertkeyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	ca_Name := data.Ca
	certkeyname_Name := data.Certkeyname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslcertkey_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslservicegroup_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ca
		if val, ok := v["ca"].(bool); ok {
			if ca_Name.IsNull() || val != ca_Name.ValueBool() {
				match = false
				continue
			}
		} else if !ca_Name.IsNull() {
			match = false
			continue
		}

		// Check certkeyname
		if val, ok := v["certkeyname"].(string); ok {
			if certkeyname_Name.IsNull() || val != certkeyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !certkeyname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslservicegroup_sslcertkey_binding with ca %s not found", ca_Name))
		return
	}

	sslservicegroup_sslcertkey_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

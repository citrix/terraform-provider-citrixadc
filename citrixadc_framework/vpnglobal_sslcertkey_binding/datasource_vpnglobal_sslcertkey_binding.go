package vpnglobal_sslcertkey_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnglobalSslcertkeyBindingDataSource)(nil)

func VPnglobalSslcertkeyBindingDataSource() datasource.DataSource {
	return &VpnglobalSslcertkeyBindingDataSource{}
}

type VpnglobalSslcertkeyBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnglobalSslcertkeyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_sslcertkey_binding"
}

func (d *VpnglobalSslcertkeyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnglobalSslcertkeyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnglobalSslcertkeyBindingDataSourceSchema()
}

func (d *VpnglobalSslcertkeyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnglobalSslcertkeyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	cacert_Name := data.Cacert
	certkeyname_Name := data.Certkeyname
	userdataencryptionkey_Name := data.Userdataencryptionkey

	if cacert_Name.IsNull() && certkeyname_Name.IsNull() && userdataencryptionkey_Name.IsNull() {
		resp.Diagnostics.AddError("Client Error", "One of cacert, certkeyname or userdataencryptionkey must be set")
		return
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_sslcertkey_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_sslcertkey_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check cacert
		if val, ok := v["cacert"].(string); ok {
			if cacert_Name.IsNull() || val != cacert_Name.ValueString() {
				match = false
				continue
			}
		} else if !cacert_Name.IsNull() {
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

		// Check userdataencryptionkey
		if val, ok := v["userdataencryptionkey"].(string); ok {
			if userdataencryptionkey_Name.IsNull() || val != userdataencryptionkey_Name.ValueString() {
				match = false
				continue
			}
		} else if !userdataencryptionkey_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_sslcertkey_binding with cacert %s not found", cacert_Name))
		return
	}

	vpnglobal_sslcertkey_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

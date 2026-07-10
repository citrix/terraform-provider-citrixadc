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

	// Lookup key is the Required certkeyname; cacert/userdataencryptionkey act as
	// optional additional filters when supplied.
	certkeyname := data.Certkeyname

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_sslcertkey_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_sslcertkey_binding returned empty array")
		return
	}

	// Iterate through results to find the matching certkeyname (plus optional filters)
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["certkeyname"].(string); !ok || certkeyname.IsNull() || val != certkeyname.ValueString() {
			continue
		}
		if !data.Cacert.IsNull() {
			if val, ok := v["cacert"].(string); !ok || val != data.Cacert.ValueString() {
				continue
			}
		}
		if !data.Userdataencryptionkey.IsNull() {
			if val, ok := v["userdataencryptionkey"].(string); !ok || val != data.Userdataencryptionkey.ValueString() {
				continue
			}
		}
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_sslcertkey_binding with certkeyname %s not found", certkeyname.ValueString()))
		return
	}

	vpnglobal_sslcertkey_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

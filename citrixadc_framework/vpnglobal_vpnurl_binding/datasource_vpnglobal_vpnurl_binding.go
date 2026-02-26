package vpnglobal_vpnurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnglobalVpnurlBindingDataSource)(nil)

func VPnglobalVpnurlBindingDataSource() datasource.DataSource {
	return &VpnglobalVpnurlBindingDataSource{}
}

type VpnglobalVpnurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnglobalVpnurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnurl_binding"
}

func (d *VpnglobalVpnurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnglobalVpnurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnglobalVpnurlBindingDataSourceSchema()
}

func (d *VpnglobalVpnurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnglobalVpnurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	urlname_Name := data.Urlname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_vpnurl_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_vpnurl_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check urlname
		if val, ok := v["urlname"].(string); ok {
			if urlname_Name.IsNull() || val != urlname_Name.ValueString() {
				match = false
				continue
			}
		} else if !urlname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_vpnurl_binding with urlname %s not found", urlname_Name))
		return
	}

	vpnglobal_vpnurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

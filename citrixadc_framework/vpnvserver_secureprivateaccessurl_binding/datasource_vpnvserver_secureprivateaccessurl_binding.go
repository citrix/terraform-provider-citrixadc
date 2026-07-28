package vpnvserver_secureprivateaccessurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverSecureprivateaccessurlBindingDataSource)(nil)

func VPnvserverSecureprivateaccessurlBindingDataSource() datasource.DataSource {
	return &VpnvserverSecureprivateaccessurlBindingDataSource{}
}

type VpnvserverSecureprivateaccessurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverSecureprivateaccessurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_secureprivateaccessurl_binding"
}

func (d *VpnvserverSecureprivateaccessurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverSecureprivateaccessurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverSecureprivateaccessurlBindingDataSourceSchema()
}

func (d *VpnvserverSecureprivateaccessurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverSecureprivateaccessurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	secureprivateaccessurl_Name := data.Secureprivateaccessurl

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_secureprivateaccessurl_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_secureprivateaccessurl_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check secureprivateaccessurl
		if val, ok := v["secureprivateaccessurl"].(string); ok {
			if secureprivateaccessurl_Name.IsNull() || val != secureprivateaccessurl_Name.ValueString() {
				match = false
				continue
			}
		} else if !secureprivateaccessurl_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_secureprivateaccessurl_binding with secureprivateaccessurl %s not found", secureprivateaccessurl_Name))
		return
	}

	vpnvserver_secureprivateaccessurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

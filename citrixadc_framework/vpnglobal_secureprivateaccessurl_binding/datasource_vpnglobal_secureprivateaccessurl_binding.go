package vpnglobal_secureprivateaccessurl_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnglobalSecureprivateaccessurlBindingDataSource)(nil)

func VPnglobalSecureprivateaccessurlBindingDataSource() datasource.DataSource {
	return &VpnglobalSecureprivateaccessurlBindingDataSource{}
}

type VpnglobalSecureprivateaccessurlBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnglobalSecureprivateaccessurlBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_secureprivateaccessurl_binding"
}

func (d *VpnglobalSecureprivateaccessurlBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnglobalSecureprivateaccessurlBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnglobalSecureprivateaccessurlBindingDataSourceSchema()
}

func (d *VpnglobalSecureprivateaccessurlBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnglobalSecureprivateaccessurlBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	secureprivateaccessurl_Name := data.Secureprivateaccessurl

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_secureprivateaccessurl_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_secureprivateaccessurl_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_secureprivateaccessurl_binding returned empty array")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_secureprivateaccessurl_binding with secureprivateaccessurl %s not found", secureprivateaccessurl_Name))
		return
	}

	vpnglobal_secureprivateaccessurl_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

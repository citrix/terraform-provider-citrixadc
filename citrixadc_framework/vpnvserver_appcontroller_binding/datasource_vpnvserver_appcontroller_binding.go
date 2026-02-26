package vpnvserver_appcontroller_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnvserverAppcontrollerBindingDataSource)(nil)

func VPnvserverAppcontrollerBindingDataSource() datasource.DataSource {
	return &VpnvserverAppcontrollerBindingDataSource{}
}

type VpnvserverAppcontrollerBindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnvserverAppcontrollerBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_appcontroller_binding"
}

func (d *VpnvserverAppcontrollerBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnvserverAppcontrollerBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverAppcontrollerBindingDataSourceSchema()
}

func (d *VpnvserverAppcontrollerBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverAppcontrollerBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	appcontroller_Name := data.Appcontroller

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnvserver_appcontroller_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_appcontroller_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_appcontroller_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check appcontroller
		if val, ok := v["appcontroller"].(string); ok {
			if appcontroller_Name.IsNull() || val != appcontroller_Name.ValueString() {
				match = false
				continue
			}
		} else if !appcontroller_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnvserver_appcontroller_binding with appcontroller %s not found", appcontroller_Name))
		return
	}

	vpnvserver_appcontroller_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package tunnelglobal_tunneltrafficpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*TunnelglobalTunneltrafficpolicyBindingDataSource)(nil)

func TUnnelglobalTunneltrafficpolicyBindingDataSource() datasource.DataSource {
	return &TunnelglobalTunneltrafficpolicyBindingDataSource{}
}

type TunnelglobalTunneltrafficpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *TunnelglobalTunneltrafficpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnelglobal_tunneltrafficpolicy_binding"
}

func (d *TunnelglobalTunneltrafficpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *TunnelglobalTunneltrafficpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = TunnelglobalTunneltrafficpolicyBindingDataSourceSchema()
}

func (d *TunnelglobalTunneltrafficpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	policyname_Name := data.Policyname
	type_Name := data.Type

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	var err error
	if !type_Name.IsNull() && type_Name.ValueString() != "" {
		argsMap["type"] = type_Name.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Tunnelglobal_tunneltrafficpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "tunnelglobal_tunneltrafficpolicy_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if val, ok := v["policyname"].(string); ok {
			if policyname_Name.IsNull() || val != policyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !policyname_Name.IsNull() {
			match = false
			continue
		}
		if !type_Name.IsNull() && type_Name.ValueString() != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != type_Name.ValueString() {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("tunnelglobal_tunneltrafficpolicy_binding with policyname %s not found", policyname_Name))
		return
	}

	tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

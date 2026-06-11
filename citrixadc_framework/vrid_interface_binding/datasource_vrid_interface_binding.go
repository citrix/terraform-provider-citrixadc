package vrid_interface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VridInterfaceBindingDataSource)(nil)

func VRidInterfaceBindingDataSource() datasource.DataSource {
	return &VridInterfaceBindingDataSource{}
}

type VridInterfaceBindingDataSource struct {
	client *service.NitroClient
}

func (d *VridInterfaceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid_interface_binding"
}

func (d *VridInterfaceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VridInterfaceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VridInterfaceBindingDataSourceSchema()
}

func (d *VridInterfaceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VridInterfaceBindingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the bound members via the aggregate parent endpoint (vrid_binding/<id>).
	id_Name := fmt.Sprintf("%v", data.VridId.ValueInt64())
	ifnum_Name := data.Ifnum

	dataArr, err := vrid_interface_bindingAggregateRead(d.client, id_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vrid_interface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vrid_interface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum.
	//
	// Verified live on NS VPX: vrid_interface_binding rows are {"id","vlan","flags"}
	// and do NOT echo "ifnum". When ifnum is present, match on it; otherwise fall
	// back to row presence (the parent vrid id already scopes the result).
	foundIndex := -1
	for i, v := range dataArr {
		raw, hasIfnum := v["ifnum"]
		if hasIfnum && raw != nil {
			matched := false
			switch t := raw.(type) {
			case string:
				matched = ifnum_Name.IsNull() || t == ifnum_Name.ValueString()
			case []interface{}:
				for _, item := range t {
					if s, ok := item.(string); ok && (ifnum_Name.IsNull() || s == ifnum_Name.ValueString()) {
						matched = true
						break
					}
				}
			}
			if !matched {
				continue
			}
		}
		// ifnum absent (firmware behavior): accept by presence.
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vrid_interface_binding with ifnum %s not found", ifnum_Name))
		return
	}

	vrid_interface_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

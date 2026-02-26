package vpnglobal_intranetip6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*VpnglobalIntranetip6BindingDataSource)(nil)

func VPnglobalIntranetip6BindingDataSource() datasource.DataSource {
	return &VpnglobalIntranetip6BindingDataSource{}
}

type VpnglobalIntranetip6BindingDataSource struct {
	client *service.NitroClient
}

func (d *VpnglobalIntranetip6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_intranetip6_binding"
}

func (d *VpnglobalIntranetip6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *VpnglobalIntranetip6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnglobalIntranetip6BindingDataSourceSchema()
}

func (d *VpnglobalIntranetip6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnglobalIntranetip6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	intranetip6_Name := data.Intranetip6

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Vpnglobal_intranetip6_binding.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_intranetip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "vpnglobal_intranetip6_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check intranetip6
		if val, ok := v["intranetip6"].(string); ok {
			if intranetip6_Name.IsNull() || val != intranetip6_Name.ValueString() {
				match = false
				continue
			}
		} else if !intranetip6_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("vpnglobal_intranetip6_binding with intranetip6 %s not found", intranetip6_Name))
		return
	}

	vpnglobal_intranetip6_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

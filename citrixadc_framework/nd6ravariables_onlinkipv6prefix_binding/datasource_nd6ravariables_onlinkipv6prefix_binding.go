package nd6ravariables_onlinkipv6prefix_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Nd6ravariablesOnlinkipv6prefixBindingDataSource)(nil)

func ND6ravariablesOnlinkipv6prefixBindingDataSource() datasource.DataSource {
	return &Nd6ravariablesOnlinkipv6prefixBindingDataSource{}
}

type Nd6ravariablesOnlinkipv6prefixBindingDataSource struct {
	client *service.NitroClient
}

func (d *Nd6ravariablesOnlinkipv6prefixBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6ravariables_onlinkipv6prefix_binding"
}

func (d *Nd6ravariablesOnlinkipv6prefixBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Nd6ravariablesOnlinkipv6prefixBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Nd6ravariablesOnlinkipv6prefixBindingDataSourceSchema()
}

func (d *Nd6ravariablesOnlinkipv6prefixBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vlan_Name := fmt.Sprintf("%d", data.Vlan.ValueInt64())
	ipv6prefix_Name := data.Ipv6prefix

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Nd6ravariables_onlinkipv6prefix_binding.Type(),
		ResourceName:             vlan_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nd6ravariables_onlinkipv6prefix_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipv6prefix
		if val, ok := v["ipv6prefix"].(string); ok {
			if ipv6prefix_Name.IsNull() || val != ipv6prefix_Name.ValueString() {
				match = false
				continue
			}
		} else if !ipv6prefix_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nd6ravariables_onlinkipv6prefix_binding with ipv6prefix %s not found", ipv6prefix_Name))
		return
	}

	nd6ravariables_onlinkipv6prefix_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

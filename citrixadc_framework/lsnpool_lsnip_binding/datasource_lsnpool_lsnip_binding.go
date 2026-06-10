package lsnpool_lsnip_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnpoolLsnipBindingDataSource)(nil)

func LSnpoolLsnipBindingDataSource() datasource.DataSource {
	return &LsnpoolLsnipBindingDataSource{}
}

type LsnpoolLsnipBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnpoolLsnipBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnpool_lsnip_binding"
}

func (d *LsnpoolLsnipBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnpoolLsnipBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnpoolLsnipBindingDataSourceSchema()
}

func (d *LsnpoolLsnipBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnpoolLsnipBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	poolname_Name := data.Poolname.ValueString()
	lsnip_Name := data.Lsnip
	ownernode_Name := data.Ownernode

	// The direct lsnpool_lsnip_binding endpoint returns a keyless empty body on
	// NS14.1; read the bound IPs from the aggregate parent endpoint instead.
	dataArr, err := lsnpool_lsnip_bindingAggregateRead(d.client, poolname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnpool_lsnip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnpool_lsnip_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check lsnip
		if val, ok := v["lsnip"].(string); ok {
			if lsnip_Name.IsNull() || val != lsnip_Name.ValueString() {
				match = false
				continue
			}
		} else if !lsnip_Name.IsNull() {
			match = false
			continue
		}

		// Check ownernode
		if val, ok := v["ownernode"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if ownernode_Name.IsNull() || val != ownernode_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !ownernode_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnpool_lsnip_binding with lsnip %s not found", lsnip_Name))
		return
	}

	lsnpool_lsnip_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

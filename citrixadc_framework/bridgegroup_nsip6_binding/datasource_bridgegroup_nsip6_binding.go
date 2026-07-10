package bridgegroup_nsip6_binding

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*BridgegroupNsip6BindingDataSource)(nil)

func BRidgegroupNsip6BindingDataSource() datasource.DataSource {
	return &BridgegroupNsip6BindingDataSource{}
}

type BridgegroupNsip6BindingDataSource struct {
	client *service.NitroClient
}

func (d *BridgegroupNsip6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_nsip6_binding"
}

func (d *BridgegroupNsip6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *BridgegroupNsip6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = BridgegroupNsip6BindingDataSourceSchema()
}

func (d *BridgegroupNsip6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BridgegroupNsip6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter with parent ID (bridgegroup_id) and ipaddress lookup key
	bridgegroupId := strconv.FormatInt(data.BridgegroupId.ValueInt64(), 10)
	ipaddress := data.Ipaddress

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Bridgegroup_nsip6_binding.Type(),
		ResourceName:             bridgegroupId,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "bridgegroup_nsip6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching ipaddress
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ipaddress"].(string); ok && val == ipaddress.ValueString() {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("bridgegroup_nsip6_binding with ipaddress %s not found", ipaddress.ValueString()))
		return
	}

	bridgegroup_nsip6_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

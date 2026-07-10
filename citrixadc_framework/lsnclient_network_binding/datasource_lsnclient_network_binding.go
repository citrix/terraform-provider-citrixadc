package lsnclient_network_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*LsnclientNetworkBindingDataSource)(nil)

func LSnclientNetworkBindingDataSource() datasource.DataSource {
	return &LsnclientNetworkBindingDataSource{}
}

type LsnclientNetworkBindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnclientNetworkBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_network_binding"
}

func (d *LsnclientNetworkBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnclientNetworkBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnclientNetworkBindingDataSourceSchema()
}

func (d *LsnclientNetworkBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnclientNetworkBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	clientname_Name := data.Clientname.ValueString()
	network_Name := data.Network

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_network_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_network_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network_binding returned empty array.")
		return
	}

	// Iterate through results to find the binding for this network.
	// Match only on `network` (the GET response does not echo `td`).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["network"].(string); ok && !network_Name.IsNull() && val == network_Name.ValueString() {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnclient_network_binding with network %s not found", network_Name))
		return
	}

	lsnclient_network_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Set the ID (the datasource has no Create); use the legacy identity keys.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

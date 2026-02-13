package vpnvserver_appfwpolicy_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*vpnvserverAppfwpolicyBindingDataSource)(nil)

func VpnvserverAppfwpolicyBindingDataSource() datasource.DataSource {
	return &vpnvserverAppfwpolicyBindingDataSource{}
}

type vpnvserverAppfwpolicyBindingDataSource struct {
	client *service.NitroClient
}

func (d *vpnvserverAppfwpolicyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_appfwpolicy_binding"
}

func (d *vpnvserverAppfwpolicyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *vpnvserverAppfwpolicyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = VpnvserverAppfwpolicyBindingDataSourceSchema()
}

func (d *vpnvserverAppfwpolicyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data VpnvserverAppfwpolicyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// For datasource, name and policy must be provided
	if data.Name.IsNull() || data.Policy.IsNull() {
		resp.Diagnostics.AddError("Missing Required Attributes", "Both 'name' and 'policy' must be provided for the datasource")
		return
	}

	name := data.Name.ValueString()
	policy := data.Policy.ValueString()

	findParams := service.FindParams{
		ResourceType:             "vpnvserver_appfwpolicy_binding",
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}

	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_appfwpolicy_binding, got error: %s", err))
		return
	}

	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("vpnvserver_appfwpolicy_binding with name=%s not found", name))
		return
	}

	// Iterate through results to find the one with the right policy
	foundIndex := -1
	for i, v := range dataArr {
		if v["policy"].(string) == policy {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("vpnvserver_appfwpolicy_binding with name=%s and policy=%s not found", name, policy))
		return
	}

	getResponseData := dataArr[foundIndex]

	vpnvserverAppfwpolicyBindingSetAttrFromGet(ctx, &data, getResponseData)

	// Set ID for datasource
	bindingId := fmt.Sprintf("%s,%s", name, policy)
	data.Id = types.StringValue(bindingId)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

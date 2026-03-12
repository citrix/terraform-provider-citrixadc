package lsnclient_nsacl6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LsnclientNsacl6BindingDataSource)(nil)

func LSnclientNsacl6BindingDataSource() datasource.DataSource {
	return &LsnclientNsacl6BindingDataSource{}
}

type LsnclientNsacl6BindingDataSource struct {
	client *service.NitroClient
}

func (d *LsnclientNsacl6BindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_nsacl6_binding"
}

func (d *LsnclientNsacl6BindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LsnclientNsacl6BindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LsnclientNsacl6BindingDataSourceSchema()
}

func (d *LsnclientNsacl6BindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LsnclientNsacl6BindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	clientname_Name := data.Clientname.ValueString()
	acl6name_Name := data.Acl6name

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_nsacl6_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_nsacl6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lsnclient_nsacl6_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check acl6name
		if val, ok := v["acl6name"].(string); ok {
			if acl6name_Name.IsNull() || val != acl6name_Name.ValueString() {
				match = false
				continue
			}
		} else if !acl6name_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lsnclient_nsacl6_binding with acl6name %s not found", acl6name_Name))
		return
	}

	lsnclient_nsacl6_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

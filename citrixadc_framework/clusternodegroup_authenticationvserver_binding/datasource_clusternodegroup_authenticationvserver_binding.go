package clusternodegroup_authenticationvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ClusternodegroupAuthenticationvserverBindingDataSource)(nil)

func CLusternodegroupAuthenticationvserverBindingDataSource() datasource.DataSource {
	return &ClusternodegroupAuthenticationvserverBindingDataSource{}
}

type ClusternodegroupAuthenticationvserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *ClusternodegroupAuthenticationvserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_authenticationvserver_binding"
}

func (d *ClusternodegroupAuthenticationvserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ClusternodegroupAuthenticationvserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClusternodegroupAuthenticationvserverBindingDataSourceSchema()
}

func (d *ClusternodegroupAuthenticationvserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ClusternodegroupAuthenticationvserverBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	vserver_Name := data.Vserver

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_authenticationvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_authenticationvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_authenticationvserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check vserver
		if val, ok := v["vserver"].(string); ok {
			if vserver_Name.IsNull() || val != vserver_Name.ValueString() {
				match = false
				continue
			}
		} else if !vserver_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("clusternodegroup_authenticationvserver_binding with vserver %s not found", vserver_Name))
		return
	}

	clusternodegroup_authenticationvserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

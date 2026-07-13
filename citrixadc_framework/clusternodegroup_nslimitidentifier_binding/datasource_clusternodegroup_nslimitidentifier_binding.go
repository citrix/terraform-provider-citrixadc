package clusternodegroup_nslimitidentifier_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ClusternodegroupNslimitidentifierBindingDataSource)(nil)

func CLusternodegroupNslimitidentifierBindingDataSource() datasource.DataSource {
	return &ClusternodegroupNslimitidentifierBindingDataSource{}
}

type ClusternodegroupNslimitidentifierBindingDataSource struct {
	client *service.NitroClient
}

func (d *ClusternodegroupNslimitidentifierBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_nslimitidentifier_binding"
}

func (d *ClusternodegroupNslimitidentifierBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ClusternodegroupNslimitidentifierBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClusternodegroupNslimitidentifierBindingDataSourceSchema()
}

func (d *ClusternodegroupNslimitidentifierBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ClusternodegroupNslimitidentifierBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	identifiername_Name := data.Identifiername

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_nslimitidentifier_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_nslimitidentifier_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_nslimitidentifier_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check identifiername
		if val, ok := v["identifiername"].(string); ok {
			if identifiername_Name.IsNull() || val != identifiername_Name.ValueString() {
				match = false
				continue
			}
		} else if !identifiername_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("clusternodegroup_nslimitidentifier_binding with identifiername %s not found", identifiername_Name))
		return
	}

	clusternodegroup_nslimitidentifier_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

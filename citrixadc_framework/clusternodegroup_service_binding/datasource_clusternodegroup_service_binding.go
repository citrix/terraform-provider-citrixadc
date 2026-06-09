package clusternodegroup_service_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ClusternodegroupServiceBindingDataSource)(nil)

func CLusternodegroupServiceBindingDataSource() datasource.DataSource {
	return &ClusternodegroupServiceBindingDataSource{}
}

type ClusternodegroupServiceBindingDataSource struct {
	client *service.NitroClient
}

func (d *ClusternodegroupServiceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_service_binding"
}

func (d *ClusternodegroupServiceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ClusternodegroupServiceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClusternodegroupServiceBindingDataSourceSchema()
}

func (d *ClusternodegroupServiceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ClusternodegroupServiceBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	name_Name := data.Name
	service_Name := data.Service

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup_service_binding.Type(),
		ResourceName:             name_Name.ValueString(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_service_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "clusternodegroup_service_binding returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check name
		if val, ok := v["name"].(string); ok {
			if name_Name.IsNull() || val != name_Name.ValueString() {
				match = false
				continue
			}
		} else if !name_Name.IsNull() {
			match = false
			continue
		}

		// Check service
		if val, ok := v["service"].(string); ok {
			if service_Name.IsNull() || val != service_Name.ValueString() {
				match = false
				continue
			}
		} else if !service_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("clusternodegroup_service_binding with name %s not found", name_Name))
		return
	}

	clusternodegroup_service_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

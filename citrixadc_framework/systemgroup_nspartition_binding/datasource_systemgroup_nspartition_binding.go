package systemgroup_nspartition_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SystemgroupNspartitionBindingDataSource)(nil)

func SYstemgroupNspartitionBindingDataSource() datasource.DataSource {
	return &SystemgroupNspartitionBindingDataSource{}
}

type SystemgroupNspartitionBindingDataSource struct {
	client *service.NitroClient
}

func (d *SystemgroupNspartitionBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemgroup_nspartition_binding"
}

func (d *SystemgroupNspartitionBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SystemgroupNspartitionBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SystemgroupNspartitionBindingDataSourceSchema()
}

func (d *SystemgroupNspartitionBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SystemgroupNspartitionBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	groupname_Name := data.Groupname.ValueString()
	partitionname_Name := data.Partitionname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Systemgroup_nspartition_binding.Type(),
		ResourceName:             groupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read systemgroup_nspartition_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "systemgroup_nspartition_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check partitionname
		if val, ok := v["partitionname"].(string); ok {
			if partitionname_Name.IsNull() || val != partitionname_Name.ValueString() {
				match = false
				continue
			}
		} else if !partitionname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("systemgroup_nspartition_binding with partitionname %s not found", partitionname_Name))
		return
	}

	systemgroup_nspartition_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

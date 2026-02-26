package servicegroup_lbmonitor_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ServicegroupLbmonitorBindingDataSource)(nil)

func SErvicegroupLbmonitorBindingDataSource() datasource.DataSource {
	return &ServicegroupLbmonitorBindingDataSource{}
}

type ServicegroupLbmonitorBindingDataSource struct {
	client *service.NitroClient
}

func (d *ServicegroupLbmonitorBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_lbmonitor_binding"
}

func (d *ServicegroupLbmonitorBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ServicegroupLbmonitorBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ServicegroupLbmonitorBindingDataSourceSchema()
}

func (d *ServicegroupLbmonitorBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServicegroupLbmonitorBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	monitorname_Name := data.MonitorName

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Servicegroup_lbmonitor_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "servicegroup_lbmonitor_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check monitor_name
		if val, ok := v["monitor_name"].(string); ok {
			if monitorname_Name.IsNull() || val != monitorname_Name.ValueString() {
				match = false
				continue
			}
		} else if !monitorname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("servicegroup_lbmonitor_binding with monitor_name %s not found", monitorname_Name))
		return
	}

	servicegroup_lbmonitor_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

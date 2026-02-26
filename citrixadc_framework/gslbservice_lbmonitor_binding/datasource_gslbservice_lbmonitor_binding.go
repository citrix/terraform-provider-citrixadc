package gslbservice_lbmonitor_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*GslbserviceLbmonitorBindingDataSource)(nil)

func GSlbserviceLbmonitorBindingDataSource() datasource.DataSource {
	return &GslbserviceLbmonitorBindingDataSource{}
}

type GslbserviceLbmonitorBindingDataSource struct {
	client *service.NitroClient
}

func (d *GslbserviceLbmonitorBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice_lbmonitor_binding"
}

func (d *GslbserviceLbmonitorBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *GslbserviceLbmonitorBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = GslbserviceLbmonitorBindingDataSourceSchema()
}

func (d *GslbserviceLbmonitorBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GslbserviceLbmonitorBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicename_Name := data.Servicename.ValueString()
	monitorname_Name := data.MonitorName

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Gslbservice_lbmonitor_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice_lbmonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "gslbservice_lbmonitor_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("gslbservice_lbmonitor_binding with monitor_name %s not found", monitorname_Name))
		return
	}

	gslbservice_lbmonitor_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

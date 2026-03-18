package gslbservicegroup_lbmonitor_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*GslbservicegroupLbmonitorBindingDataSource)(nil)

func GSlbservicegroupLbmonitorBindingDataSource() datasource.DataSource {
	return &GslbservicegroupLbmonitorBindingDataSource{}
}

type GslbservicegroupLbmonitorBindingDataSource struct {
	client *service.NitroClient
}

func (d *GslbservicegroupLbmonitorBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_lbmonitor_binding"
}

func (d *GslbservicegroupLbmonitorBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *GslbservicegroupLbmonitorBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = GslbservicegroupLbmonitorBindingDataSourceSchema()
}

func (d *GslbservicegroupLbmonitorBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	monitorname_Name := data.MonitorName
	port_Name := data.Port

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Gslbservicegroup_lbmonitor_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_lbmonitor_binding returned empty array.")
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

		// Check port
		if val, ok := v["port"]; ok {
			val, _ = utils.ConvertToInt64(val)
			if port_Name.IsNull() || val != port_Name.ValueInt64() {
				match = false
				continue
			}
		} else if !port_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("gslbservicegroup_lbmonitor_binding with monitor_name %s not found", monitorname_Name))
		return
	}

	gslbservicegroup_lbmonitor_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

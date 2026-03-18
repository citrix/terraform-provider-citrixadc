package servicegroup_servicegroupmember_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ServicegroupServicegroupmemberBindingDataSource)(nil)

func SErvicegroupServicegroupmemberBindingDataSource() datasource.DataSource {
	return &ServicegroupServicegroupmemberBindingDataSource{}
}

type ServicegroupServicegroupmemberBindingDataSource struct {
	client *service.NitroClient
}

func (d *ServicegroupServicegroupmemberBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_servicegroupmember_binding"
}

func (d *ServicegroupServicegroupmemberBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ServicegroupServicegroupmemberBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ServicegroupServicegroupmemberBindingDataSourceSchema()
}

func (d *ServicegroupServicegroupmemberBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	ip_Name := data.Ip
	port_Name := data.Port
	servername_Name := data.Servername

	if data.Ip.IsNull() && data.Servername.IsNull() {
		resp.Diagnostics.AddError("Configuration Error", "Either 'ip' or 'servername' must be specified to read servicegroup_servicegroupmember_binding resource")
		return
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Servicegroup_servicegroupmember_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_servicegroupmember_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "servicegroup_servicegroupmember_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ip
		if !ip_Name.IsNull() {
			if val, ok := v["ip"].(string); ok {
				if ip_Name.IsNull() || val != ip_Name.ValueString() {
					match = false
					continue
				}
			} else if !ip_Name.IsNull() {
				match = false
				continue
			}
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

		// Check servername
		if !servername_Name.IsNull() {
			if val, ok := v["servername"].(string); ok {
				if servername_Name.IsNull() || val != servername_Name.ValueString() {
					match = false
					continue
				}
			} else if !servername_Name.IsNull() {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("servicegroup_servicegroupmember_binding with ip %s not found", ip_Name))
		return
	}

	servicegroup_servicegroupmember_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

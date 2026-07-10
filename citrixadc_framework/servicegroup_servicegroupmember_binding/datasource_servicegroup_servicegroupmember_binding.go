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

	// The effective member key is servername if supplied, else ip (the ADC names
	// the server after the bound ip, so GET always returns servername == ip for
	// ip-based members). Only filter on a value the caller actually supplied.
	lookupKey := ""
	if !servername_Name.IsNull() {
		lookupKey = servername_Name.ValueString()
	} else if !ip_Name.IsNull() {
		lookupKey = ip_Name.ValueString()
	}

	// Iterate through results to find the matching member.
	foundIndex := -1
	for i, v := range dataArr {
		if lookupKey != "" {
			servernameVal, _ := v["servername"].(string)
			if servernameVal != lookupKey {
				continue
			}
		}
		if !port_Name.IsNull() {
			if pv, ok := v["port"]; ok {
				portVal, _ := utils.ConvertToInt64(pv)
				if portVal != port_Name.ValueInt64() {
					continue
				}
			} else {
				continue
			}
		}
		foundIndex = i
		break
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("servicegroup_servicegroupmember_binding with key %s not found", lookupKey))
		return
	}

	servicegroup_servicegroupmember_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

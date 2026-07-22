package gslbservicegroup_gslbservicegroupmember_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*GslbservicegroupGslbservicegroupmemberBindingDataSource)(nil)

func GSlbservicegroupGslbservicegroupmemberBindingDataSource() datasource.DataSource {
	return &GslbservicegroupGslbservicegroupmemberBindingDataSource{}
}

type GslbservicegroupGslbservicegroupmemberBindingDataSource struct {
	client *service.NitroClient
}

func (d *GslbservicegroupGslbservicegroupmemberBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_gslbservicegroupmember_binding"
}

func (d *GslbservicegroupGslbservicegroupmemberBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *GslbservicegroupGslbservicegroupmemberBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = GslbservicegroupGslbservicegroupmemberBindingDataSourceSchema()
}

func (d *GslbservicegroupGslbservicegroupmemberBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel
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
		ResourceType:             service.Gslbservicegroup_gslbservicegroupmember_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "gslbservicegroup_gslbservicegroupmember_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// ip and servername are a mutually-exclusive lookup choice: only filter on the one the
	// user actually supplied in config. (When binding by servername the ADC still echoes the
	// server's ip back, so unconditionally filtering on ip would never match a servername lookup.)
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ip (only when supplied in config)
		if !ip_Name.IsNull() {
			if val, ok := v["ip"].(string); ok {
				if val != ip_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check port
		if !port_Name.IsNull() {
			if val, ok := v["port"]; ok {
				val, _ = utils.ConvertToInt64(val)
				if val != port_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check servername (only when supplied in config)
		if !servername_Name.IsNull() {
			if val, ok := v["servername"].(string); ok {
				if val != servername_Name.ValueString() {
					match = false
					continue
				}
			} else {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("gslbservicegroup_gslbservicegroupmember_binding with ip %s not found", ip_Name))
		return
	}

	gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

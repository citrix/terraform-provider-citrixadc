package lbvserver_service_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LbvserverServiceBindingDataSource)(nil)

func LBvserverServiceBindingDataSource() datasource.DataSource {
	return &LbvserverServiceBindingDataSource{}
}

type LbvserverServiceBindingDataSource struct {
	client *service.NitroClient
}

func (d *LbvserverServiceBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_service_binding"
}

func (d *LbvserverServiceBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LbvserverServiceBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LbvserverServiceBindingDataSourceSchema()
}

func (d *LbvserverServiceBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LbvserverServiceBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	servicegroupname_Name := data.Servicegroupname
	servicename_Name := data.Servicename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lbvserver_service_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_service_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lbvserver_service_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the supplied filter
	// attributes (servicename and/or servicegroupname).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check servicegroupname when supplied in config
		if !servicegroupname_Name.IsNull() {
			if val, ok := v["servicegroupname"].(string); !ok || val != servicegroupname_Name.ValueString() {
				match = false
			}
		}

		// Check servicename when supplied in config
		if match && !servicename_Name.IsNull() {
			if val, ok := v["servicename"].(string); !ok || val != servicename_Name.ValueString() {
				match = false
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lbvserver_service_binding for vserver %s not found with the provided filter attributes", name_Name))
		return
	}

	lbvserver_service_bindingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

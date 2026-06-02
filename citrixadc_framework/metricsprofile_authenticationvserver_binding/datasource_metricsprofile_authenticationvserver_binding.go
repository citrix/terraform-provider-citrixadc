package metricsprofile_authenticationvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*MetricsprofileAuthenticationvserverBindingDataSource)(nil)

func MEtricsprofileAuthenticationvserverBindingDataSource() datasource.DataSource {
	return &MetricsprofileAuthenticationvserverBindingDataSource{}
}

type MetricsprofileAuthenticationvserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *MetricsprofileAuthenticationvserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metricsprofile_authenticationvserver_binding"
}

func (d *MetricsprofileAuthenticationvserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *MetricsprofileAuthenticationvserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = MetricsprofileAuthenticationvserverBindingDataSourceSchema()
}

func (d *MetricsprofileAuthenticationvserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MetricsprofileAuthenticationvserverBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	name_Name := data.Name.ValueString()
	entityname_Name := data.Entityname
	entitytype_Name := data.Entitytype

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Metricsprofile_authenticationvserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read metricsprofile_authenticationvserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "metricsprofile_authenticationvserver_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check entityname
		if val, ok := v["entityname"].(string); ok {
			if entityname_Name.IsNull() || val != entityname_Name.ValueString() {
				match = false
				continue
			}
		} else if !entityname_Name.IsNull() {
			match = false
			continue
		}

		// Check entitytype
		if val, ok := v["entitytype"].(string); ok {
			if entitytype_Name.IsNull() || val != entitytype_Name.ValueString() {
				match = false
				continue
			}
		} else if !entitytype_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("metricsprofile_authenticationvserver_binding with entityname %s not found", entityname_Name))
		return
	}

	metricsprofile_authenticationvserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
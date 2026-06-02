package metricsprofile_uservserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*MetricsprofileUservserverBindingDataSource)(nil)

func MEtricsprofileUservserverBindingDataSource() datasource.DataSource {
	return &MetricsprofileUservserverBindingDataSource{}
}

type MetricsprofileUservserverBindingDataSource struct {
	client *service.NitroClient
}

func (d *MetricsprofileUservserverBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metricsprofile_uservserver_binding"
}

func (d *MetricsprofileUservserverBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *MetricsprofileUservserverBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = MetricsprofileUservserverBindingDataSourceSchema()
}

func (d *MetricsprofileUservserverBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data MetricsprofileUservserverBindingResourceModel
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
		ResourceType:             service.Metricsprofile_uservserver_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read metricsprofile_uservserver_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "metricsprofile_uservserver_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("metricsprofile_uservserver_binding with entityname %s not found", entityname_Name))
		return
	}

	metricsprofile_uservserver_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
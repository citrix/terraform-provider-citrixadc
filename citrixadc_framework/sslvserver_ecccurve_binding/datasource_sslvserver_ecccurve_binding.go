package sslvserver_ecccurve_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslvserverEcccurveBindingDataSource)(nil)

func SSlvserverEcccurveBindingDataSource() datasource.DataSource {
	return &SslvserverEcccurveBindingDataSource{}
}

type SslvserverEcccurveBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslvserverEcccurveBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_ecccurve_binding"
}

func (d *SslvserverEcccurveBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslvserverEcccurveBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslvserverEcccurveBindingDataSourceSchema()
}

func (d *SslvserverEcccurveBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslvserverEcccurveBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vservername_Name := data.Vservername.ValueString()
	ecccurvename_Name := data.Ecccurvename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_ecccurve_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_ecccurve_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslvserver_ecccurve_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ecccurvename
		if val, ok := v["ecccurvename"].(string); ok {
			if ecccurvename_Name.IsNull() || val != ecccurvename_Name.ValueString() {
				match = false
				continue
			}
		} else if !ecccurvename_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslvserver_ecccurve_binding with ecccurvename %s not found", ecccurvename_Name))
		return
	}

	sslvserver_ecccurve_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

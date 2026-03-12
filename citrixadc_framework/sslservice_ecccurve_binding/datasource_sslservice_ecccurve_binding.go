package sslservice_ecccurve_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslserviceEcccurveBindingDataSource)(nil)

func SSlserviceEcccurveBindingDataSource() datasource.DataSource {
	return &SslserviceEcccurveBindingDataSource{}
}

type SslserviceEcccurveBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslserviceEcccurveBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservice_ecccurve_binding"
}

func (d *SslserviceEcccurveBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslserviceEcccurveBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslserviceEcccurveBindingDataSourceSchema()
}

func (d *SslserviceEcccurveBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslserviceEcccurveBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicename_Name := data.Servicename.ValueString()
	ecccurvename_Name := data.Ecccurvename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslservice_ecccurve_binding.Type(),
		ResourceName:             servicename_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslservice_ecccurve_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslservice_ecccurve_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslservice_ecccurve_binding with ecccurvename %s not found", ecccurvename_Name))
		return
	}

	sslservice_ecccurve_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package lbmetrictable_metric_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*LbmetrictableMetricBindingDataSource)(nil)

func LBmetrictableMetricBindingDataSource() datasource.DataSource {
	return &LbmetrictableMetricBindingDataSource{}
}

type LbmetrictableMetricBindingDataSource struct {
	client *service.NitroClient
}

func (d *LbmetrictableMetricBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmetrictable_metric_binding"
}

func (d *LbmetrictableMetricBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LbmetrictableMetricBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LbmetrictableMetricBindingDataSourceSchema()
}

func (d *LbmetrictableMetricBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LbmetrictableMetricBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	metrictable_Name := data.Metrictable.ValueString()
	metric_Name := data.Metric

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lbmetrictable_metric_binding.Type(),
		ResourceName:             metrictable_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lbmetrictable_metric_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lbmetrictable_metric_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check metric
		if val, ok := v["metric"].(string); ok {
			if metric_Name.IsNull() || val != metric_Name.ValueString() {
				match = false
				continue
			}
		} else if !metric_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("lbmetrictable_metric_binding with metric %s not found", metric_Name))
		return
	}

	lbmetrictable_metric_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package policytracing

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*PolicytracingDataSource)(nil)

func POlicytracingDataSource() datasource.DataSource {
	return &PolicytracingDataSource{}
}

type PolicytracingDataSource struct {
	client *service.NitroClient
}

func (d *PolicytracingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policytracing"
}

func (d *PolicytracingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *PolicytracingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = PolicytracingDataSourceSchema()
}

// Read backs the datasource with the NITRO get(all) endpoint. transactionid,
// detail and nodeid are optional GET filters; the first policy-trace record
// matching every supplied filter is returned.
func (d *PolicytracingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicytracingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Policytracing.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read policytracing, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "policytracing returned empty array")
		return
	}

	// Iterate through results to find the first one matching every supplied filter.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if !data.Detail.IsNull() {
			if val, ok := v["detail"].(string); !ok || val != data.Detail.ValueString() {
				match = false
			}
		}

		if match && !data.Transactionid.IsNull() {
			if val, ok := v["transactionid"].(string); !ok || val != data.Transactionid.ValueString() {
				match = false
			}
		}

		if match && !data.Nodeid.IsNull() {
			if intVal, err := utils.ConvertToInt64(v["nodeid"]); err != nil || intVal != data.Nodeid.ValueInt64() {
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
		resp.Diagnostics.AddError("Client Error", "no policytracing record matched the provided filters")
		return
	}

	policytracingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// The datasource has no Create; set the synthetic ID here.
	data.Id = types.StringValue("policytracing-query")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

package snmptrap

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SnmptrapDataSource)(nil)

func SNmptrapDataSource() datasource.DataSource {
	return &SnmptrapDataSource{}
}

type SnmptrapDataSource struct {
	client *service.NitroClient
}

func (d *SnmptrapDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap"
}

func (d *SnmptrapDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SnmptrapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SnmptrapDataSourceSchema()
}

func (d *SnmptrapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SnmptrapResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	td_Name := data.Td.ValueInt64()

	trapclass_Name := data.Trapclass.ValueString()

	trapdestination_Name := data.Trapdestination.ValueString()

	version_Name := data.Version.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "snmptrap",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "snmptrap returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		match := true

		// Handle td comparison - it may come as string or float64
		var tdValue int64
		switch td := v["td"].(type) {
		case float64:
			tdValue = int64(td)
		case string:
			// Try to parse string to int64
			var err error
			tdValue, err = strconv.ParseInt(td, 10, 64)
			if err != nil {
				tdValue = 0
			}
		case int:
			tdValue = int64(td)
		case int64:
			tdValue = td
		default:
			tdValue = 0
		}

		if tdValue != td_Name {
			match = false
		}

		if v["trapclass"].(string) != trapclass_Name {
			match = false
		}

		if v["trapdestination"].(string) != trapdestination_Name {
			match = false
		}

		if v["version"].(string) != version_Name {
			match = false
		}

		if match {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("snmptrap with td %s not found", td_Name))
		return
	}

	snmptrapSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

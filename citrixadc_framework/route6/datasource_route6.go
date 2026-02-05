package route6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Route6DataSource)(nil)

func ROute6DataSource() datasource.DataSource {
	return &Route6DataSource{}
}

type Route6DataSource struct {
	client *service.NitroClient
}

func (d *Route6DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route6"
}

func (d *Route6DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Route6DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Route6DataSourceSchema()
}

func (d *Route6DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Route6ResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	network_Name := data.Network.ValueString()

	td_Name := fmt.Sprintf("%d", data.Td.ValueInt64())

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "route6",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read route6, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "route6 returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		match := true

		if v["network"].(string) != network_Name {
			match = false
		}

		// Handle td comparison - it could be a number or string
		tdVal := ""
		if v["td"] != nil {
			switch td := v["td"].(type) {
			case string:
				tdVal = td
			case float64:
				tdVal = fmt.Sprintf("%.0f", td)
			case int:
				tdVal = fmt.Sprintf("%d", td)
			default:
				tdVal = "0"
			}
		} else {
			tdVal = "0"
		}

		if tdVal != td_Name {
			match = false
		}

		if match {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("route6 with network %s and td %s not found", network_Name, td_Name))
		return
	}

	route6SetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

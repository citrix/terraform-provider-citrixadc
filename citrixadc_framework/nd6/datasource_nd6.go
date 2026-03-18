package nd6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Nd6DataSource)(nil)

func ND6DataSource() datasource.DataSource {
	return &Nd6DataSource{}
}

type Nd6DataSource struct {
	client *service.NitroClient
}

func (d *Nd6DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6"
}

func (d *Nd6DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Nd6DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Nd6DataSourceSchema()
}

func (d *Nd6DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Nd6ResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	neighbor_Name := data.Neighbor.ValueString()

	td_Name := fmt.Sprintf("%d", data.Td.ValueInt64())

	nodeid_Name := fmt.Sprintf("%d", data.Nodeid.ValueInt64())

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "nd6",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nd6, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nd6 returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		match := true

		if v["neighbor"].(string) != neighbor_Name {
			match = false
		}

		// Handle td - it might be nil or a string
		tdVal := "0"
		if v["td"] != nil {
			tdVal = v["td"].(string)
		}
		if tdVal != td_Name {
			match = false
		}

		// Handle nodeid - it might be nil or a string
		nodeidVal := "0"
		if v["nodeid"] != nil {
			nodeidVal = v["nodeid"].(string)
		}
		if nodeidVal != nodeid_Name {
			match = false
		}

		if match {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nd6 with neighbor %s not found", neighbor_Name))
		return
	}

	nd6SetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

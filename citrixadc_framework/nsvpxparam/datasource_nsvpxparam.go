package nsvpxparam

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsvpxparamDataSource)(nil)

func NSvpxparamDataSource() datasource.DataSource {
	return &NsvpxparamDataSource{}
}

type NsvpxparamDataSource struct {
	client *service.NitroClient
}

func (d *NsvpxparamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsvpxparam"
}

func (d *NsvpxparamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsvpxparamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsvpxparamDataSourceSchema()
}

func (d *NsvpxparamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsvpxparamResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	ownernode_Name := fmt.Sprintf("%d", data.Ownernode.ValueInt64())

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "nsvpxparam",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsvpxparam, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nsvpxparam returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		// Handle ownernode comparison with proper type conversion
		ownernodeVal := int64(0)
		if v["ownernode"] != nil {
			// Convert from float64 (JSON number) to int64
			if floatVal, ok := v["ownernode"].(float64); ok {
				ownernodeVal = int64(floatVal)
			} else if intVal, ok := v["ownernode"].(int64); ok {
				ownernodeVal = intVal
			}
		}

		if fmt.Sprintf("%d", ownernodeVal) == ownernode_Name {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nsvpxparam with ownernode %s not found", ownernode_Name))
		return
	}

	nsvpxparamSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

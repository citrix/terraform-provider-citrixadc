package endpointinfo

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*EndpointinfoDataSource)(nil)

func ENdpointinfoDataSource() datasource.DataSource {
	return &EndpointinfoDataSource{}
}

type EndpointinfoDataSource struct {
	client *service.NitroClient
}

func (d *EndpointinfoDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_endpointinfo"
}

func (d *EndpointinfoDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *EndpointinfoDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = EndpointinfoDataSourceSchema()
}

func (d *EndpointinfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EndpointinfoResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	endpointkind_Name := data.Endpointkind
	endpointname_Name := data.Endpointname

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Endpointinfo.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read endpointinfo, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "endpointinfo returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check endpointkind
		if val, ok := v["endpointkind"].(string); ok {
			if endpointkind_Name.IsNull() || val != endpointkind_Name.ValueString() {
				match = false
				continue
			}
		} else if !endpointkind_Name.IsNull() {
			match = false
			continue
		}

		// Check endpointname
		if val, ok := v["endpointname"].(string); ok {
			if endpointname_Name.IsNull() || val != endpointname_Name.ValueString() {
				match = false
				continue
			}
		} else if !endpointname_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("endpointinfo with endpointkind %s not found", endpointkind_Name))
		return
	}

	endpointinfoSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

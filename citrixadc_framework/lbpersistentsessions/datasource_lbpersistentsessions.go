package lbpersistentsessions

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*LbpersistentsessionsDataSource)(nil)

func LBpersistentsessionsDataSource() datasource.DataSource {
	return &LbpersistentsessionsDataSource{}
}

type LbpersistentsessionsDataSource struct {
	client *service.NitroClient
}

func (d *LbpersistentsessionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbpersistentsessions"
}

func (d *LbpersistentsessionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LbpersistentsessionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LbpersistentsessionsDataSourceSchema()
}

// Read backs the datasource with the NITRO get(all) endpoint. vserver and nodeid
// are optional filters; the first session matching every supplied filter is
// returned. nodeid is a valid GET filter and is honored here.
func (d *LbpersistentsessionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LbpersistentsessionsResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Lbpersistentsessions.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lbpersistentsessions, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "lbpersistentsessions returned empty array")
		return
	}

	// Iterate through results to find the first one matching every supplied filter.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		if !data.Vserver.IsNull() {
			if val, ok := v["vserver"].(string); !ok || val != data.Vserver.ValueString() {
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
		resp.Diagnostics.AddError("Client Error", "no lbpersistentsessions matched the provided filters")
		return
	}

	lbpersistentsessionsSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// The datasource has no Create; set the synthetic ID here.
	data.Id = types.StringValue("lbpersistentsessions-query")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

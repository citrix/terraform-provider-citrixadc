package routerdynamicrouting

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

var _ datasource.DataSource = (*RouterdynamicroutingDataSource)(nil)

func ROuterdynamicroutingDataSource() datasource.DataSource {
	return &RouterdynamicroutingDataSource{}
}

type RouterdynamicroutingDataSource struct {
	client *service.NitroClient
}

// RouterdynamicroutingDataSourceModel describes the datasource data model.
//
// The datasource is a read-only "show command" query and is intentionally
// decoupled from the resource model: the resource is action-only and carries a
// commandlines list, while the datasource queries by a single commandstring.
type RouterdynamicroutingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Commandstring types.String `tfsdk:"commandstring"`
	Nodeid        types.Int64  `tfsdk:"nodeid"`
}

func (d *RouterdynamicroutingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routerdynamicrouting"
}

func (d *RouterdynamicroutingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *RouterdynamicroutingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = RouterdynamicroutingDataSourceSchema()
}

func (d *RouterdynamicroutingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data RouterdynamicroutingDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID
	commandstring_Name := data.Commandstring.ValueString()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "routerdynamicrouting",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read routerdynamicrouting, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "routerdynamicrouting returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if cs, ok := v["commandstring"].(string); ok && cs == commandstring_Name {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("routerdynamicrouting with commandstring %s not found", commandstring_Name))
		return
	}

	routerdynamicroutingSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// routerdynamicroutingSetAttrFromGetForDatasource copies the GET response into
// the datasource model and sets the datasource ID.
func routerdynamicroutingSetAttrFromGetForDatasource(ctx context.Context, data *RouterdynamicroutingDataSourceModel, getResponseData map[string]interface{}) {
	if val, ok := getResponseData["commandstring"]; ok && val != nil {
		data.Commandstring = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}

	// Set ID for the datasource
	data.Id = types.StringValue(data.Commandstring.ValueString())
}

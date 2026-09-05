package nsip

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsipDataSource)(nil)

func NSipDataSource() datasource.DataSource {
	return &NsipDataSource{}
}

type NsipDataSource struct {
	client *service.NitroClient
}

func (d *NsipDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsip"
}

func (d *NsipDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsipDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsipDataSourceSchema()
}

func (d *NsipDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsipDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ipaddressName := data.Ipaddress.ValueString()

	trafficDomain := int64(0)
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		trafficDomain = data.Td.ValueInt64()
	}
	argsMap := map[string]string{
		"td": fmt.Sprintf("%d", trafficDomain),
	}

	findParams := service.FindParams{
		ResourceType:             service.Nsip.Type(),
		ResourceName:             ipaddressName,
		ResourceMissingErrorCode: 258,
		ArgsMap:                  argsMap,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsip, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nsip returned empty array")
		return
	}

	// Iterate through results to find the one with the matching ipaddress
	foundIndex := -1
	for i, v := range dataArr {
		if addr, ok := v["ipaddress"].(string); ok && addr == ipaddressName {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nsip with ipaddress %s not found", ipaddressName))
		return
	}

	nsipDataSourceSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

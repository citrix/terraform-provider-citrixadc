package nsip6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*Nsip6DataSource)(nil)

func NSip6DataSource() datasource.DataSource {
	return &Nsip6DataSource{}
}

type Nsip6DataSource struct {
	client *service.NitroClient
}

func (d *Nsip6DataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsip6"
}

func (d *Nsip6DataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *Nsip6DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = Nsip6DataSourceSchema()
}

func (d *Nsip6DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data Nsip6ResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	ipv6address_Name := data.Ipv6address.ValueString()

	td_Name := data.Td.ValueInt64()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "nsip6",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsip6, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "nsip6 returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		match := true

		if v["ipv6address"].(string) != ipv6address_Name {
			match = false
		}

		// Convert td to int64 safely
		tdVal, err := utils.ConvertToInt64(v["td"])
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to convert td field: %s", err))
			return
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("nsip6 with ipv6address %s not found", ipv6address_Name))
		return
	}

	nsip6SetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

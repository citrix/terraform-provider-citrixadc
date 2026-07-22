package dnscaarec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnscaarecDataSource)(nil)

func DNscaarecDataSource() datasource.DataSource {
	return &DnscaarecDataSource{}
}

type DnscaarecDataSource struct {
	client *service.NitroClient
}

func (d *DnscaarecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnscaarec"
}

func (d *DnscaarecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnscaarecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnscaarecDataSourceSchema()
}

func (d *DnscaarecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnscaarecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Locate the specific record by domain (URL) + recordid.
	domainName := data.Domain.ValueString()
	recordid := data.Recordid.ValueInt64()

	findParams := service.FindParams{
		ResourceType:             service.Dnscaarec.Type(),
		ResourceName:             domainName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnscaarec, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnscaarec returned empty array.")
		return
	}

	// Iterate through results to find the one with the right recordid
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["recordid"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil && intVal == recordid {
				foundIndex = i
				break
			}
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnscaarec with domain %s and recordid %d not found", domainName, recordid))
		return
	}

	// dnscaarecSetAttrFromGet faithfully copies the GET response and sets the
	// composite domain,recordid ID.
	dnscaarecSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

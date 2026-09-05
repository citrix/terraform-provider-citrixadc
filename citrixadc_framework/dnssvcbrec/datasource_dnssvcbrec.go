package dnssvcbrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*DnssvcbrecDataSource)(nil)

func DNssvcbrecDataSource() datasource.DataSource {
	return &DnssvcbrecDataSource{}
}

type DnssvcbrecDataSource struct {
	client *service.NitroClient
}

func (d *DnssvcbrecDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssvcbrec"
}

func (d *DnssvcbrecDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *DnssvcbrecDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DnssvcbrecDataSourceSchema()
}

func (d *DnssvcbrecDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DnssvcbrecResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Array filter without parent ID - fetch all records and match on identity.
	domainName := data.Domain.ValueString()
	targetName := data.Targetname.ValueString()

	// dnssvcbrec's get REQUIRES the svcbtype key arg; a plain get-all returns an
	// empty set. This data source is keyed by domain+targetname (svcbtype is not a
	// required input), so query each svcbtype value and merge the results.
	var dataArr []map[string]interface{}
	for _, st := range []string{"HTTPS", "SVCB"} {
		findParams := service.FindParams{
			ResourceType:             "dnssvcbrec",
			ArgsMap:                  map[string]string{"svcbtype": st},
			ResourceMissingErrorCode: 258,
		}
		arr, err := d.client.FindResourceArrayWithParams(findParams)
		if err != nil {
			if utils.IsNotFoundError(err) {
				continue
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read dnssvcbrec, got error: %s", err))
			return
		}
		dataArr = append(dataArr, arr...)
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "dnssvcbrec returned empty array")
		return
	}

	// Iterate through results to find the one with the right identity
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		if d, ok := v["domain"].(string); !ok || d != domainName {
			match = false
		}
		if t, ok := v["targetname"].(string); !ok || t != targetName {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("dnssvcbrec with domain %s not found", domainName))
		return
	}

	dnssvcbrecSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

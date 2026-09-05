package lldpneighbors

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*LldpneighborsDataSource)(nil)

func LLdpneighborsDataSource() datasource.DataSource {
	return &LldpneighborsDataSource{}
}

type LldpneighborsDataSource struct {
	client *service.NitroClient
}

func (d *LldpneighborsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lldpneighbors"
}

func (d *LldpneighborsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *LldpneighborsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = LldpneighborsDataSourceSchema()
}

// Read backs the datasource with get(all) and filters in memory. lldpneighbors
// is a transient diagnostics table: the neighbor list may legitimately be empty,
// so an empty result MUST NOT be treated as an error. ifnum is an optional
// filter; nodeid is an optional cluster-node GET filter.
func (d *LldpneighborsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LldpneighborsDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lldpneighbors datasource via get(all)")

	all, err := d.client.FindAllResources(service.Lldpneighbors.Type())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read lldpneighbors, got error: %s", err))
		return
	}

	// Optional in-memory filtering by ifnum. Empty result is tolerated.
	ifnumFilter := data.Ifnum.ValueString()
	var match map[string]interface{}
	for _, rec := range all {
		if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() && ifnumFilter != "" {
			if v, ok := rec["ifnum"].(string); ok && v == ifnumFilter {
				match = rec
				break
			}
			continue
		}
		// No ifnum filter supplied: take the first record if present.
		match = rec
		break
	}

	if match != nil {
		lldpneighborsDataSourceSetAttrFromGet(ctx, &data, match)
	} else {
		tflog.Debug(ctx, "No lldpneighbors records found; returning empty result")
	}

	// Datasource sets its own synthetic ID (there is no Create).
	data.Id = types.StringValue("lldpneighbors")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

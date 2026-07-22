package lldpneighbors

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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
		lldpneighborsSetAttrFromGetForDatasourceModel(ctx, &data, match)
	} else {
		tflog.Debug(ctx, "No lldpneighbors records found; returning empty result")
	}

	// Datasource sets its own synthetic ID (there is no Create).
	data.Id = types.StringValue("lldpneighbors")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// lldpneighborsSetAttrFromGetForDatasourceModel faithfully copies the get(all)
// response record into the datasource model, including the full read-only LLDP
// neighbor telemetry. Every telemetry field is a string on the vendored NITRO
// struct, so they map to types.String. Missing fields are left as their
// zero-value (null) so callers can distinguish absent telemetry.
func lldpneighborsSetAttrFromGetForDatasourceModel(ctx context.Context, data *LldpneighborsDataSourceModel, getResponseData map[string]interface{}) *LldpneighborsDataSourceModel {
	tflog.Debug(ctx, "In lldpneighborsSetAttrFromGetForDatasourceModel Function")

	setString := func(key string, dst *types.String) {
		if val, ok := getResponseData[key]; ok && val != nil {
			if s, ok := val.(string); ok {
				*dst = types.StringValue(s)
			}
		}
	}

	setString("ifnum", &data.Ifnum)
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}

	setString("chassisidsubtype", &data.Chassisidsubtype)
	setString("chassisid", &data.Chassisid)
	setString("portidsubtype", &data.Portidsubtype)
	setString("portid", &data.Portid)
	setString("ttl", &data.Ttl)
	setString("portdescription", &data.Portdescription)
	setString("sys", &data.Sys)
	setString("sysdesc", &data.Sysdesc)
	setString("mgmtaddresssubtype", &data.Mgmtaddresssubtype)
	setString("mgmtaddress", &data.Mgmtaddress)
	setString("iftype", &data.Iftype)
	setString("ifnumber", &data.Ifnumber)
	setString("vlan", &data.Vlan)
	setString("vlanid", &data.Vlanid)
	setString("portprotosupported", &data.Portprotosupported)
	setString("portprotoenabled", &data.Portprotoenabled)
	setString("portprotoid", &data.Portprotoid)
	setString("portvlanid", &data.Portvlanid)
	setString("protocolid", &data.Protocolid)
	setString("linkaggrcapable", &data.Linkaggrcapable)
	setString("linkaggrenabled", &data.Linkaggrenabled)
	setString("linkaggrid", &data.Linkaggrid)
	setString("flag", &data.Flag)
	setString("syscapabilities", &data.Syscapabilities)
	setString("syscapenabled", &data.Syscapenabled)
	setString("autonegsupport", &data.Autonegsupport)
	setString("autonegenabled", &data.Autonegenabled)
	setString("autonegadvertised", &data.Autonegadvertised)
	setString("autonegmautype", &data.Autonegmautype)
	setString("mtu", &data.Mtu)

	return data
}

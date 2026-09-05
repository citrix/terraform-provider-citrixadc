package nsmigration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsmigrationDataSourceModel describes the data source data model.
//
// The start/stop/complete actions are modelled as separate action-only resources
// (see resource_nsmigration_start.go / _stop.go / _complete.go). This data source
// reads the live migration singleton via the NITRO get; dumpsession is the only
// read/write field (it cannot be set through any NITRO operation on this object).
type NsmigrationDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Dumpsession types.String `tfsdk:"dumpsession"`
}

func NsmigrationDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dumpsession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Displays the current active migrated session details, if DUMPSESSION option is YES. Possible values = YES, NO",
			},
		},
	}
}

// nsmigrationSetAttrFromGet maps a NITRO get response onto the data source model.
func nsmigrationSetAttrFromGet(ctx context.Context, data *NsmigrationDataSourceModel, getResponseData map[string]interface{}) *NsmigrationDataSourceModel {
	tflog.Debug(ctx, "In nsmigrationSetAttrFromGet Function")

	if val, ok := getResponseData["dumpsession"]; ok && val != nil {
		data.Dumpsession = types.StringValue(val.(string))
	} else {
		data.Dumpsession = types.StringNull()
	}

	// Singleton - static ID.
	data.Id = types.StringValue("nsmigration-config")

	return data
}

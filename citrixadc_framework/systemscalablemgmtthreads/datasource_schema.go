package systemscalablemgmtthreads

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemscalablemgmtthreadsDataSourceModel describes the data source data model.
//
// The enable/disable actions are modelled as separate action-only resources
// (see resource_systemscalablemgmtthreads_enable.go /
// resource_systemscalablemgmtthreads_disable.go). This data source reads the live
// feature state via the NITRO get (configuredstate / effectivestate); nodeid is a
// GET-only cluster-node filter.
type SystemscalablemgmtthreadsDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Nodeid          types.Int64  `tfsdk:"nodeid"`
	Configuredstate types.String `tfsdk:"configuredstate"`
	Effectivestate  types.String `tfsdk:"effectivestate"`
}

func SystemscalablemgmtthreadsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node. Minimum value = 0, Maximum value = 31.",
			},
			"configuredstate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the configured state of the Scalable Management Threads feature. Possible values = ENABLED, DISABLED.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the current running state of the Scalable Management Threads feature. Possible values = ENABLED, DISABLED.",
			},
		},
	}
}

// systemscalablemgmtthreadsSetAttrFromGet maps a NITRO get response onto the data
// source model.
func systemscalablemgmtthreadsSetAttrFromGet(ctx context.Context, data *SystemscalablemgmtthreadsDataSourceModel, getResponseData map[string]interface{}) *SystemscalablemgmtthreadsDataSourceModel {
	tflog.Debug(ctx, "In systemscalablemgmtthreadsSetAttrFromGet Function")

	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		} else {
			data.Nodeid = types.Int64Null()
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["configuredstate"]; ok && val != nil {
		data.Configuredstate = types.StringValue(val.(string))
	} else {
		data.Configuredstate = types.StringNull()
	}
	if val, ok := getResponseData["effectivestate"]; ok && val != nil {
		data.Effectivestate = types.StringValue(val.(string))
	} else {
		data.Effectivestate = types.StringNull()
	}

	// Set ID for the resource. Singleton feature toggle - static ID.
	data.Id = types.StringValue("systemscalablemgmtthreads")

	return data
}

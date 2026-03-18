package lsnrtspalgprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnrtspalgprofileResourceModel describes the resource data model.
type LsnrtspalgprofileResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Rtspalgprofilename    types.String `tfsdk:"rtspalgprofilename"`
	Rtspidletimeout       types.Int64  `tfsdk:"rtspidletimeout"`
	Rtspportrange         types.String `tfsdk:"rtspportrange"`
	Rtsptransportprotocol types.String `tfsdk:"rtsptransportprotocol"`
}

func (r *LsnrtspalgprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnrtspalgprofile resource.",
			},
			"rtspalgprofilename": schema.StringAttribute{
				Required:    true,
				Description: "The name of the RTSPALG Profile.",
			},
			"rtspidletimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(120),
				Description: "Idle timeout for the rtsp sessions in seconds.",
			},
			"rtspportrange": schema.StringAttribute{
				Required:    true,
				Description: "port for the RTSP",
			},
			"rtsptransportprotocol": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("TCP"),
				Description: "RTSP ALG Profile transport protocol type.",
			},
		},
	}
}

func lsnrtspalgprofileGetThePayloadFromtheConfig(ctx context.Context, data *LsnrtspalgprofileResourceModel) lsn.Lsnrtspalgprofile {
	tflog.Debug(ctx, "In lsnrtspalgprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnrtspalgprofile := lsn.Lsnrtspalgprofile{}
	if !data.Rtspalgprofilename.IsNull() {
		lsnrtspalgprofile.Rtspalgprofilename = data.Rtspalgprofilename.ValueString()
	}
	if !data.Rtspidletimeout.IsNull() {
		lsnrtspalgprofile.Rtspidletimeout = utils.IntPtr(int(data.Rtspidletimeout.ValueInt64()))
	}
	if !data.Rtspportrange.IsNull() {
		lsnrtspalgprofile.Rtspportrange = data.Rtspportrange.ValueString()
	}
	if !data.Rtsptransportprotocol.IsNull() {
		lsnrtspalgprofile.Rtsptransportprotocol = data.Rtsptransportprotocol.ValueString()
	}

	return lsnrtspalgprofile
}

func lsnrtspalgprofileSetAttrFromGet(ctx context.Context, data *LsnrtspalgprofileResourceModel, getResponseData map[string]interface{}) *LsnrtspalgprofileResourceModel {
	tflog.Debug(ctx, "In lsnrtspalgprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["rtspalgprofilename"]; ok && val != nil {
		data.Rtspalgprofilename = types.StringValue(val.(string))
	} else {
		data.Rtspalgprofilename = types.StringNull()
	}
	if val, ok := getResponseData["rtspidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rtspidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Rtspidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["rtspportrange"]; ok && val != nil {
		data.Rtspportrange = types.StringValue(val.(string))
	} else {
		data.Rtspportrange = types.StringNull()
	}
	if val, ok := getResponseData["rtsptransportprotocol"]; ok && val != nil {
		data.Rtsptransportprotocol = types.StringValue(val.(string))
	} else {
		data.Rtsptransportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Rtspalgprofilename.ValueString())

	return data
}

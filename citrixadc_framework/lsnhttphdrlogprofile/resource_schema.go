package lsnhttphdrlogprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnhttphdrlogprofileResourceModel describes the resource data model.
type LsnhttphdrlogprofileResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Httphdrlogprofilename types.String `tfsdk:"httphdrlogprofilename"`
	Loghost               types.String `tfsdk:"loghost"`
	Logmethod             types.String `tfsdk:"logmethod"`
	Logurl                types.String `tfsdk:"logurl"`
	Logversion            types.String `tfsdk:"logversion"`
}

func (r *LsnhttphdrlogprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnhttphdrlogprofile resource.",
			},
			"httphdrlogprofilename": schema.StringAttribute{
				Required: true,
				// SDK v2 ForceNew -> RequiresReplace (name change recreates the resource)
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the HTTP header logging Profile.",
			},
			"loghost": schema.StringAttribute{
				// SDK v2: Optional + Computed, no Default (value read back from ADC)
				Optional:    true,
				Computed:    true,
				Description: "Host information is logged if option is enabled.",
			},
			"logmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP method information is logged if option is enabled.",
			},
			"logurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL information is logged if option is enabled.",
			},
			"logversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version information is logged if option is enabled.",
			},
		},
	}
}

func lsnhttphdrlogprofileGetThePayloadFromthePlan(ctx context.Context, data *LsnhttphdrlogprofileResourceModel) lsn.Lsnhttphdrlogprofile {
	tflog.Debug(ctx, "In lsnhttphdrlogprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lsnhttphdrlogprofile := lsn.Lsnhttphdrlogprofile{}
	if !data.Httphdrlogprofilename.IsNull() && !data.Httphdrlogprofilename.IsUnknown() {
		lsnhttphdrlogprofile.Httphdrlogprofilename = data.Httphdrlogprofilename.ValueString()
	}
	if !data.Loghost.IsNull() && !data.Loghost.IsUnknown() {
		lsnhttphdrlogprofile.Loghost = data.Loghost.ValueString()
	}
	if !data.Logmethod.IsNull() && !data.Logmethod.IsUnknown() {
		lsnhttphdrlogprofile.Logmethod = data.Logmethod.ValueString()
	}
	if !data.Logurl.IsNull() && !data.Logurl.IsUnknown() {
		lsnhttphdrlogprofile.Logurl = data.Logurl.ValueString()
	}
	if !data.Logversion.IsNull() && !data.Logversion.IsUnknown() {
		lsnhttphdrlogprofile.Logversion = data.Logversion.ValueString()
	}

	return lsnhttphdrlogprofile
}

func lsnhttphdrlogprofileSetAttrFromGet(ctx context.Context, data *LsnhttphdrlogprofileResourceModel, getResponseData map[string]interface{}) *LsnhttphdrlogprofileResourceModel {
	tflog.Debug(ctx, "In lsnhttphdrlogprofileSetAttrFromGet Function")

	// Convert API response to model.
	// Guard else-branches: only null a value when it is unknown; never clobber a
	// known configured value that NITRO omits from the GET response (omit-on-default trap).
	if val, ok := getResponseData["httphdrlogprofilename"]; ok && val != nil {
		data.Httphdrlogprofilename = types.StringValue(val.(string))
	} else if data.Httphdrlogprofilename.IsUnknown() {
		data.Httphdrlogprofilename = types.StringNull()
	}
	if val, ok := getResponseData["loghost"]; ok && val != nil {
		data.Loghost = types.StringValue(val.(string))
	} else if data.Loghost.IsUnknown() {
		data.Loghost = types.StringNull()
	}
	if val, ok := getResponseData["logmethod"]; ok && val != nil {
		data.Logmethod = types.StringValue(val.(string))
	} else if data.Logmethod.IsUnknown() {
		data.Logmethod = types.StringNull()
	}
	if val, ok := getResponseData["logurl"]; ok && val != nil {
		data.Logurl = types.StringValue(val.(string))
	} else if data.Logurl.IsUnknown() {
		data.Logurl = types.StringNull()
	}
	if val, ok := getResponseData["logversion"]; ok && val != nil {
		data.Logversion = types.StringValue(val.(string))
	} else if data.Logversion.IsUnknown() {
		data.Logversion = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Httphdrlogprofilename.ValueString()))

	return data
}

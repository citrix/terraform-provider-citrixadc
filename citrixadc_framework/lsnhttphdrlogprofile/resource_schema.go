package lsnhttphdrlogprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Required:    true,
				Description: "The name of the HTTP header logging Profile.",
			},
			"loghost": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Host information is logged if option is enabled.",
			},
			"logmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "HTTP method information is logged if option is enabled.",
			},
			"logurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "URL information is logged if option is enabled.",
			},
			"logversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Version information is logged if option is enabled.",
			},
		},
	}
}

func lsnhttphdrlogprofileGetThePayloadFromtheConfig(ctx context.Context, data *LsnhttphdrlogprofileResourceModel) lsn.Lsnhttphdrlogprofile {
	tflog.Debug(ctx, "In lsnhttphdrlogprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnhttphdrlogprofile := lsn.Lsnhttphdrlogprofile{}
	if !data.Httphdrlogprofilename.IsNull() {
		lsnhttphdrlogprofile.Httphdrlogprofilename = data.Httphdrlogprofilename.ValueString()
	}
	if !data.Loghost.IsNull() {
		lsnhttphdrlogprofile.Loghost = data.Loghost.ValueString()
	}
	if !data.Logmethod.IsNull() {
		lsnhttphdrlogprofile.Logmethod = data.Logmethod.ValueString()
	}
	if !data.Logurl.IsNull() {
		lsnhttphdrlogprofile.Logurl = data.Logurl.ValueString()
	}
	if !data.Logversion.IsNull() {
		lsnhttphdrlogprofile.Logversion = data.Logversion.ValueString()
	}

	return lsnhttphdrlogprofile
}

func lsnhttphdrlogprofileSetAttrFromGet(ctx context.Context, data *LsnhttphdrlogprofileResourceModel, getResponseData map[string]interface{}) *LsnhttphdrlogprofileResourceModel {
	tflog.Debug(ctx, "In lsnhttphdrlogprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["httphdrlogprofilename"]; ok && val != nil {
		data.Httphdrlogprofilename = types.StringValue(val.(string))
	} else {
		data.Httphdrlogprofilename = types.StringNull()
	}
	if val, ok := getResponseData["loghost"]; ok && val != nil {
		data.Loghost = types.StringValue(val.(string))
	} else {
		data.Loghost = types.StringNull()
	}
	if val, ok := getResponseData["logmethod"]; ok && val != nil {
		data.Logmethod = types.StringValue(val.(string))
	} else {
		data.Logmethod = types.StringNull()
	}
	if val, ok := getResponseData["logurl"]; ok && val != nil {
		data.Logurl = types.StringValue(val.(string))
	} else {
		data.Logurl = types.StringNull()
	}
	if val, ok := getResponseData["logversion"]; ok && val != nil {
		data.Logversion = types.StringValue(val.(string))
	} else {
		data.Logversion = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Httphdrlogprofilename.ValueString())

	return data
}

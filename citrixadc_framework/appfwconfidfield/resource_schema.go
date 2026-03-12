package appfwconfidfield

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwconfidfieldResourceModel describes the resource data model.
type AppfwconfidfieldResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Fieldname types.String `tfsdk:"fieldname"`
	Isregex   types.String `tfsdk:"isregex"`
	State     types.String `tfsdk:"state"`
	Url       types.String `tfsdk:"url"`
}

func (r *AppfwconfidfieldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwconfidfield resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the form field designation.",
			},
			"fieldname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the form field to designate as confidential.",
			},
			"isregex": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Method of specifying the form field name. Available settings function as follows:\n* REGEX. Form field is a regular expression.\n* NOTREGEX. Form field is a literal string.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the confidential field designation.",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "URL of the web page that contains the web form.",
			},
		},
	}
}

func appfwconfidfieldGetThePayloadFromtheConfig(ctx context.Context, data *AppfwconfidfieldResourceModel) appfw.Appfwconfidfield {
	tflog.Debug(ctx, "In appfwconfidfieldGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwconfidfield := appfw.Appfwconfidfield{}
	if !data.Comment.IsNull() {
		appfwconfidfield.Comment = data.Comment.ValueString()
	}
	if !data.Fieldname.IsNull() {
		appfwconfidfield.Fieldname = data.Fieldname.ValueString()
	}
	if !data.Isregex.IsNull() {
		appfwconfidfield.Isregex = data.Isregex.ValueString()
	}
	if !data.State.IsNull() {
		appfwconfidfield.State = data.State.ValueString()
	}
	if !data.Url.IsNull() {
		appfwconfidfield.Url = data.Url.ValueString()
	}

	return appfwconfidfield
}

func appfwconfidfieldSetAttrFromGet(ctx context.Context, data *AppfwconfidfieldResourceModel, getResponseData map[string]interface{}) *AppfwconfidfieldResourceModel {
	tflog.Debug(ctx, "In appfwconfidfieldSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["fieldname"]; ok && val != nil {
		data.Fieldname = types.StringValue(val.(string))
	} else {
		data.Fieldname = types.StringNull()
	}
	if val, ok := getResponseData["isregex"]; ok && val != nil {
		data.Isregex = types.StringValue(val.(string))
	} else {
		data.Isregex = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Fieldname.ValueString(), data.Url.ValueString()))

	return data
}

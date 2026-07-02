package cacheobject

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CacheobjectResourceModel describes the resource data model.
type CacheobjectResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Action                 types.String `tfsdk:"action"`
	Group                  types.String `tfsdk:"group"`
	Groupname              types.String `tfsdk:"groupname"`
	Host                   types.String `tfsdk:"host"`
	Httpmethod             types.String `tfsdk:"httpmethod"`
	Httpstatus             types.Int64  `tfsdk:"httpstatus"`
	Ignoremarkerobjects    types.String `tfsdk:"ignoremarkerobjects"`
	Includenotreadyobjects types.String `tfsdk:"includenotreadyobjects"`
	Locator                types.Int64  `tfsdk:"locator"`
	Nodeid                 types.Int64  `tfsdk:"nodeid"`
	Port                   types.Int64  `tfsdk:"port"`
	Tosecondary            types.String `tfsdk:"tosecondary"`
	Url                    types.String `tfsdk:"url"`
}

func (r *CacheobjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheobject resource.",
			},
			"action": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("flush"),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The action to perform on the cached object(s). One of: expire, flush, save. Defaults to flush. cacheobject is an action-only runtime object; this is not a persistent configuration.",
			},
			"group": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content group whose objects should be listed.",
			},
			"groupname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the content group to which the object belongs. It will display only the objects belonging to the specified content group. You must also set the Host parameter.",
			},
			"host": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host name of the object. Parameter \"url\" must be specified.",
			},
			"httpmethod": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "HTTP request method that caused the object to be stored.",
			},
			"httpstatus": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "HTTP status of the object.",
			},
			"ignoremarkerobjects": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Ignore marker objects. Marker objects are created when a response exceeds the maximum or minimum response size for the content group or has not yet received the minimum number of hits for the content group.",
			},
			"includenotreadyobjects": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Include responses that have not yet reached a minimum number of hits before being cached.",
			},
			"locator": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cached object.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Host port of the object. You must also set the Host parameter.",
			},
			"tosecondary": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Object will be saved onto Secondary. Applies only to the save action.",
			},
			"url": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL of the particular object whose details is required. Parameter \"host\" must be specified along with the URL.",
			},
		},
	}
}

// cacheobjectGetActionPayload builds the payload for the given action. Per the
// NITRO doc, expire/flush accept locator | (url,host[,port,groupname,httpmethod]),
// and save accepts [locator] [tosecondary]. Fields not valid for the action are
// omitted.
func cacheobjectGetActionPayload(ctx context.Context, data *CacheobjectResourceModel, action string) cache.Cacheobject {
	tflog.Debug(ctx, "In cacheobjectGetActionPayload Function")

	cacheobject := cache.Cacheobject{}

	if action == "save" {
		if !data.Locator.IsNull() && !data.Locator.IsUnknown() {
			cacheobject.Locator = utils.IntPtr(int(data.Locator.ValueInt64()))
		}
		if !data.Tosecondary.IsNull() && !data.Tosecondary.IsUnknown() {
			cacheobject.Tosecondary = data.Tosecondary.ValueString()
		}
		return cacheobject
	}

	// expire / flush
	if !data.Locator.IsNull() && !data.Locator.IsUnknown() {
		cacheobject.Locator = utils.IntPtr(int(data.Locator.ValueInt64()))
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		cacheobject.Url = data.Url.ValueString()
	}
	if !data.Host.IsNull() && !data.Host.IsUnknown() {
		cacheobject.Host = data.Host.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		cacheobject.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		cacheobject.Groupname = data.Groupname.ValueString()
	}
	if !data.Httpmethod.IsNull() && !data.Httpmethod.IsUnknown() {
		cacheobject.Httpmethod = data.Httpmethod.ValueString()
	}

	return cacheobject
}

// cacheobjectSetAttrFromGetForDatasource faithfully copies every read-back field
// from the GET (get-all) response into the model, and sets a synthetic ID. Used
// only by the datasource; the resource does not read back (action-only object).
func cacheobjectSetAttrFromGetForDatasource(ctx context.Context, data *CacheobjectResourceModel, getResponseData map[string]interface{}) *CacheobjectResourceModel {
	tflog.Debug(ctx, "In cacheobjectSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["group"]; ok && val != nil {
		data.Group = types.StringValue(val.(string))
	} else {
		data.Group = types.StringNull()
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["host"]; ok && val != nil {
		data.Host = types.StringValue(val.(string))
	} else {
		data.Host = types.StringNull()
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	} else {
		data.Httpmethod = types.StringNull()
	}
	if val, ok := getResponseData["httpstatus"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpstatus = types.Int64Value(intVal)
		} else {
			data.Httpstatus = types.Int64Null()
		}
	} else {
		data.Httpstatus = types.Int64Null()
	}
	if val, ok := getResponseData["ignoremarkerobjects"]; ok && val != nil {
		data.Ignoremarkerobjects = types.StringValue(val.(string))
	} else {
		data.Ignoremarkerobjects = types.StringNull()
	}
	if val, ok := getResponseData["includenotreadyobjects"]; ok && val != nil {
		data.Includenotreadyobjects = types.StringValue(val.(string))
	} else {
		data.Includenotreadyobjects = types.StringNull()
	}
	if val, ok := getResponseData["locator"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Locator = types.Int64Value(intVal)
		} else {
			data.Locator = types.Int64Null()
		}
	} else {
		data.Locator = types.Int64Null()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		} else {
			data.Nodeid = types.Int64Null()
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		} else {
			data.Port = types.Int64Null()
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["tosecondary"]; ok && val != nil {
		data.Tosecondary = types.StringValue(val.(string))
	} else {
		data.Tosecondary = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}

	return data
}

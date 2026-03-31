package appfwprofile_fileuploadtype_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileFileuploadtypeBindingResourceModel describes the resource data model.
type AppfwprofileFileuploadtypeBindingResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Alertonly                 types.String `tfsdk:"alertonly"`
	AsFileuploadtypesUrl      types.String `tfsdk:"as_fileuploadtypes_url"`
	Comment                   types.String `tfsdk:"comment"`
	Filetype                  types.List   `tfsdk:"filetype"`
	Fileuploadtype            types.String `tfsdk:"fileuploadtype"`
	Isautodeployed            types.String `tfsdk:"isautodeployed"`
	Isnameregex               types.String `tfsdk:"isnameregex"`
	IsregexFileuploadtypesUrl types.String `tfsdk:"isregex_fileuploadtypes_url"`
	Name                      types.String `tfsdk:"name"`
	Resourceid                types.String `tfsdk:"resourceid"`
	State                     types.String `tfsdk:"state"`
}

func (r *AppfwprofileFileuploadtypeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_fileuploadtype_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_fileuploadtypes_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FileUploadTypes action URL.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"filetype": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "FileUploadTypes file types.",
			},
			"fileuploadtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FileUploadTypes to allow/deny.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isnameregex": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOTREGEX"),
				Description: "Is field name a regular expression?",
			},
			"isregex_fileuploadtypes_url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is a regular expression?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_fileuploadtype_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AppfwprofileFileuploadtypeBindingResourceModel) appfw.Appfwprofilefileuploadtypebinding {
	tflog.Debug(ctx, "In appfwprofile_fileuploadtype_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwprofile_fileuploadtype_binding := appfw.Appfwprofilefileuploadtypebinding{}
	if !data.Alertonly.IsNull() {
		appfwprofile_fileuploadtype_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsFileuploadtypesUrl.IsNull() {
		appfwprofile_fileuploadtype_binding.Asfileuploadtypesurl = data.AsFileuploadtypesUrl.ValueString()
	}
	if !data.Comment.IsNull() {
		appfwprofile_fileuploadtype_binding.Comment = data.Comment.ValueString()
	}
	if !data.Fileuploadtype.IsNull() {
		appfwprofile_fileuploadtype_binding.Fileuploadtype = data.Fileuploadtype.ValueString()
	}
	if !data.Isautodeployed.IsNull() {
		appfwprofile_fileuploadtype_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Isnameregex.IsNull() {
		appfwprofile_fileuploadtype_binding.Isnameregex = data.Isnameregex.ValueString()
	}
	if !data.IsregexFileuploadtypesUrl.IsNull() {
		appfwprofile_fileuploadtype_binding.Isregexfileuploadtypesurl = data.IsregexFileuploadtypesUrl.ValueString()
	}
	if !data.Name.IsNull() {
		appfwprofile_fileuploadtype_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() {
		appfwprofile_fileuploadtype_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() {
		appfwprofile_fileuploadtype_binding.State = data.State.ValueString()
	}

	return appfwprofile_fileuploadtype_binding
}

func appfwprofile_fileuploadtype_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileFileuploadtypeBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileFileuploadtypeBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_fileuploadtype_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_fileuploadtypes_url"]; ok && val != nil {
		data.AsFileuploadtypesUrl = types.StringValue(val.(string))
	} else {
		data.AsFileuploadtypesUrl = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["fileuploadtype"]; ok && val != nil {
		data.Fileuploadtype = types.StringValue(val.(string))
	} else {
		data.Fileuploadtype = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["isnameregex"]; ok && val != nil {
		data.Isnameregex = types.StringValue(val.(string))
	} else {
		data.Isnameregex = types.StringNull()
	}
	if val, ok := getResponseData["isregex_fileuploadtypes_url"]; ok && val != nil {
		data.IsregexFileuploadtypesUrl = types.StringValue(val.(string))
	} else {
		data.IsregexFileuploadtypesUrl = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	dataFiletype := ""
	if val, ok := getResponseData["filetype"]; ok && val != nil {
		if filetypeSlice, ok := val.([]interface{}); ok {
			dataFiletype = strings.Join(utils.ToStringList(filetypeSlice), ";")
			stringList := utils.ToStringList(filetypeSlice)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Filetype = listValue
		} else {
			data.Filetype = types.ListNull(types.StringType)
		}
	} else {
		data.Filetype = types.ListNull(types.StringType)
	}
	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_fileuploadtypes_url:%s", utils.UrlEncode(fmt.Sprintf("%v", data.AsFileuploadtypesUrl.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("filetype:%s", utils.UrlEncode(fmt.Sprintf("%v", dataFiletype))))
	idParts = append(idParts, fmt.Sprintf("fileuploadtype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Fileuploadtype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

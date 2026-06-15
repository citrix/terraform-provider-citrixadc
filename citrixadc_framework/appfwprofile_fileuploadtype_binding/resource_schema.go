package appfwprofile_fileuploadtype_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	Ruletype                  types.String `tfsdk:"ruletype"`
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "FileUploadTypes action URL.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"filetype": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "FileUploadTypes file types.",
			},
			"fileuploadtype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "FileUploadTypes to allow/deny.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isnameregex": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is field name a regular expression?",
			},
			"isregex_fileuploadtypes_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is a regular expression?",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_fileuploadtype_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileFileuploadtypeBindingResourceModel) appfw.Appfwprofilefileuploadtypebinding {
	tflog.Debug(ctx, "In appfwprofile_fileuploadtype_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_fileuploadtype_binding := appfw.Appfwprofilefileuploadtypebinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsFileuploadtypesUrl.IsNull() && !data.AsFileuploadtypesUrl.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Asfileuploadtypesurl = data.AsFileuploadtypesUrl.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Comment = data.Comment.ValueString()
	}
	if !data.Filetype.IsNull() && !data.Filetype.IsUnknown() {
		var filetypeList []string
		data.Filetype.ElementsAs(ctx, &filetypeList, false)
		appfwprofile_fileuploadtype_binding.Filetype = filetypeList
	}
	if !data.Fileuploadtype.IsNull() && !data.Fileuploadtype.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Fileuploadtype = data.Fileuploadtype.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Isnameregex.IsNull() && !data.Isnameregex.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Isnameregex = data.Isnameregex.ValueString()
	}
	if !data.IsregexFileuploadtypesUrl.IsNull() && !data.IsregexFileuploadtypesUrl.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Isregexfileuploadtypesurl = data.IsregexFileuploadtypesUrl.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_fileuploadtype_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
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
	if val, ok := getResponseData["filetype"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Filetype = listValue
		} else {
			data.Filetype = types.ListNull(types.StringType)
		}
	} else {
		data.Filetype = types.ListNull(types.StringType)
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
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	// filetype is a list of strings; it is encoded as a ';'-joined string to match
	// the legacy SDK v2 composite ID format (name,fileuploadtype,as_fileuploadtypes_url,filetype).
	data.Id = types.StringValue(appfwprofile_fileuploadtype_bindingComposeId(data))

	return data
}

// appfwprofile_fileuploadtype_bindingComposeId builds the new key:value composite ID.
// The filetype list is joined with ';' before URL-encoding to preserve the legacy
// SDK v2 semantics where filetype was a semicolon-separated string within the ID.
func appfwprofile_fileuploadtype_bindingComposeId(data *AppfwprofileFileuploadtypeBindingResourceModel) string {
	filetypeJoined := strings.Join(appfwprofile_fileuploadtype_bindingFiletypeList(data), ";")
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("as_fileuploadtypes_url:%s", utils.UrlEncode(data.AsFileuploadtypesUrl.ValueString())))
	idParts = append(idParts, fmt.Sprintf("filetype:%s", utils.UrlEncode(filetypeJoined)))
	idParts = append(idParts, fmt.Sprintf("fileuploadtype:%s", utils.UrlEncode(data.Fileuploadtype.ValueString())))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	return strings.Join(idParts, ",")
}

// appfwprofile_fileuploadtype_bindingFiletypeList extracts the filetype list as []string.
func appfwprofile_fileuploadtype_bindingFiletypeList(data *AppfwprofileFileuploadtypeBindingResourceModel) []string {
	var filetypeList []string
	if !data.Filetype.IsNull() && !data.Filetype.IsUnknown() {
		data.Filetype.ElementsAs(context.Background(), &filetypeList, false)
	}
	return filetypeList
}

package filteraction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/filter"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilteractionResource{}
var _ resource.ResourceWithConfigure = (*FilteractionResource)(nil)
var _ resource.ResourceWithImportState = (*FilteractionResource)(nil)

func NewFilteractionResource() resource.Resource {
	return &FilteractionResource{}
}

// FilteractionResource models the NITRO `filteraction` config object (part of the
// deprecated `filter` feature). Unlike the action-only custom resources, this is a
// genuine CONFIG-CRUD resource: the NITRO object supports add/get/update/delete BY
// NAME, so Create=AddResource, Read=FindResource (clearing state when the object is
// gone), Update=UpdateResource, Delete=DeleteResource, plus ImportState passthrough.
//
// This is a backward-compatible migration of the legacy SDKv2
// citrixadc_filteraction resource: the resource type name, every schema attribute
// (names/types/optionality/ForceNew) and the id scheme (id = name) are preserved
// exactly.
type FilteractionResource struct {
	client *service.NitroClient
}

// FilteractionResourceModel describes the resource data model. Every schema
// attribute has a matching tfsdk field, mirroring the SDKv2 schema exactly.
type FilteractionResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Qual        types.String `tfsdk:"qual"`
	Page        types.String `tfsdk:"page"`
	Respcode    types.Int64  `tfsdk:"respcode"`
	Servicename types.String `tfsdk:"servicename"`
	Value       types.String `tfsdk:"value"`
}

func (r *FilteractionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filteraction"
}

func (r *FilteractionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FilteractionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// id == name, so a plain passthrough (mirrors the SDKv2
	// schema.ImportStatePassthroughContext importer).
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FilteractionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the filteraction resource. Equals the filter action name.",
			},
			// SDKv2: Required, ForceNew (the name cannot be changed after creation).
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the filtering action.",
			},
			// SDKv2: Required (updateable in-place; not ForceNew).
			"qual": schema.StringAttribute{
				Required:    true,
				Description: "Qualifier, which is the action to be performed.",
			},
			// SDKv2: Optional + Computed.
			"page": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTML page to return for HTTP requests (for use with the ERRORCODE qualifier).",
			},
			// SDKv2: Optional + Computed (TypeInt).
			"respcode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response code to be returned for HTTP requests (for use with the ERRORCODE qualifier).",
			},
			// SDKv2: Optional + Computed.
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to which to forward HTTP requests. Required if the qualifier is FORWARD.",
			},
			// SDKv2: Optional + Computed.
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String containing the header_name and header_value.",
			},
		},
	}
}

func (r *FilteractionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FilteractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating filteraction resource")

	filteractionName := data.Name.ValueString()
	payload := filteractionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	_, err := r.client.AddResource(service.Filteraction.Type(), filteractionName, &payload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create filteraction, got error: %s", err))
		return
	}

	// id = name (mirrors SDKv2 d.SetId(name))
	data.Id = types.StringValue(filteractionName)

	tflog.Trace(ctx, "Created filteraction resource")

	// Read the created state back
	r.readFilteractionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilteractionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FilteractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filteractionName := data.Id.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Reading filteraction state %s", filteractionName))

	getResponseData, err := r.client.FindResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		// Object is gone on the ADC; clear it from state (mirrors SDKv2 d.SetId("")).
		tflog.Warn(ctx, fmt.Sprintf("Clearing filteraction state %s", filteractionName))
		resp.State.RemoveResource(ctx)
		return
	}

	filteractionSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilteractionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FilteractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filteractionName := data.Name.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Updating filteraction resource %s", filteractionName))

	payload := filteractionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	_, err := r.client.UpdateResource(service.Filteraction.Type(), filteractionName, &payload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update filteraction %s, got error: %s", filteractionName, err))
		return
	}

	// Preserve/refresh id (id = name)
	data.Id = types.StringValue(filteractionName)

	tflog.Trace(ctx, "Updated filteraction resource")

	// Read the updated state back
	r.readFilteractionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilteractionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilteractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filteractionName := data.Id.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Deleting filteraction resource %s", filteractionName))

	// Make API call
	err := r.client.DeleteResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete filteraction %s, got error: %s", filteractionName, err))
		return
	}

	tflog.Trace(ctx, "Deleted filteraction resource")
}

// readFilteractionFromApi fetches the object by name and populates the model. It is
// used for the read-back after Create/Update; a failure here is surfaced as an error
// (the object was just written and must be present).
func (r *FilteractionResource) readFilteractionFromApi(ctx context.Context, data *FilteractionResourceModel, diags *diag.Diagnostics) {
	filteractionName := data.Id.ValueString()
	getResponseData, err := r.client.FindResource(service.Filteraction.Type(), filteractionName)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read filteraction %s, got error: %s", filteractionName, err))
		return
	}

	filteractionSetAttrFromGet(ctx, data, getResponseData)
}

// filteractionGetThePayloadFromthePlan builds the NITRO request body from the model.
// Mirrors the SDKv2 create/update payload construction, including the pointer int for
// respcode so an unset value is omitted from the JSON.
func filteractionGetThePayloadFromthePlan(ctx context.Context, data *FilteractionResourceModel) filter.Filteraction {
	tflog.Debug(ctx, "In filteractionGetThePayloadFromthePlan Function")

	filteraction := filter.Filteraction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		filteraction.Name = data.Name.ValueString()
	}
	if !data.Qual.IsNull() && !data.Qual.IsUnknown() {
		filteraction.Qual = data.Qual.ValueString()
	}
	if !data.Page.IsNull() && !data.Page.IsUnknown() {
		filteraction.Page = data.Page.ValueString()
	}
	if !data.Respcode.IsNull() && !data.Respcode.IsUnknown() {
		filteraction.Respcode = utils.IntPtr(int(data.Respcode.ValueInt64()))
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		filteraction.Servicename = data.Servicename.ValueString()
	}
	if !data.Value.IsNull() && !data.Value.IsUnknown() {
		filteraction.Value = data.Value.ValueString()
	}

	return filteraction
}

// filteractionSetAttrFromGet converts a NITRO GET response into the model, mirroring
// the SDKv2 readFilteractionFunc. id is kept equal to name.
func filteractionSetAttrFromGet(ctx context.Context, data *FilteractionResourceModel, getResponseData map[string]interface{}) *FilteractionResourceModel {
	tflog.Debug(ctx, "In filteractionSetAttrFromGet Function")

	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["qual"]; ok && val != nil {
		data.Qual = types.StringValue(val.(string))
	} else {
		data.Qual = types.StringNull()
	}
	if val, ok := getResponseData["page"]; ok && val != nil {
		data.Page = types.StringValue(val.(string))
	} else {
		data.Page = types.StringNull()
	}
	if val, ok := getResponseData["respcode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Respcode = types.Int64Value(intVal)
		} else {
			data.Respcode = types.Int64Null()
		}
	} else {
		data.Respcode = types.Int64Null()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["value"]; ok && val != nil {
		data.Value = types.StringValue(val.(string))
	} else {
		data.Value = types.StringNull()
	}

	// id = name (mirrors SDKv2 d.SetId(name))
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

package filterpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/filter"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilterpolicyResource{}
var _ resource.ResourceWithConfigure = (*FilterpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*FilterpolicyResource)(nil)

func NewFilterpolicyResource() resource.Resource {
	return &FilterpolicyResource{}
}

// FilterpolicyResource models the NetScaler `filterpolicy` configuration object
// (part of the deprecated `filter` feature).
//
// This is a real config-CRUD resource: it maps directly to the NITRO
// `filterpolicy` object and performs Add/Find/Update/Delete by name. It is a
// backward-compatible migration of the legacy SDKv2 citrixadc_filterpolicy
// resource, preserving the same resource type name, schema attribute
// names/types/optionality, CRUD behavior, and id scheme (id == name).
type FilterpolicyResource struct {
	client *service.NitroClient
}

// FilterpolicyResourceModel describes the resource data model. Every schema
// attribute has a matching tfsdk field. Mirrors the legacy SDKv2 schema exactly.
type FilterpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Resaction types.String `tfsdk:"resaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *FilterpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filterpolicy"
}

func (r *FilterpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FilterpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the filterpolicy resource. Equals the policy name.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the filtering action.",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to be performed on requests that match the policy.",
			},
			"resaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The action to be performed on the response.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC classic expression specifying the type of connections that match this policy.",
			},
		},
	}
}

func (r *FilterpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FilterpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating filterpolicy resource")

	filterpolicyName := data.Name.ValueString()
	filterpolicy := filterpolicyGetThePayloadFromthePlan(ctx, &data)

	_, err := r.client.AddResource(service.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create filterpolicy, got error: %s", err))
		return
	}

	// id == name (mirrors the legacy SDKv2 d.SetId(name)).
	data.Id = types.StringValue(filterpolicyName)

	// Read the object back to populate Optional+Computed attributes.
	r.readFilterpolicy(ctx, filterpolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FilterpolicyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filterpolicyName := data.Id.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Reading filterpolicy state %s", filterpolicyName))

	found := r.readFilterpolicy(ctx, filterpolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Object no longer exists on the target; drop it from state (mirrors the
		// legacy SDKv2 d.SetId("") on a failed Find).
		tflog.Warn(ctx, fmt.Sprintf("Clearing filterpolicy state %s", filterpolicyName))
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FilterpolicyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filterpolicyName := data.Name.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Updating filterpolicy %s", filterpolicyName))

	filterpolicy := filterpolicyGetThePayloadFromthePlan(ctx, &data)

	_, err := r.client.UpdateResource(service.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update filterpolicy %s, got error: %s", filterpolicyName, err))
		return
	}

	data.Id = types.StringValue(filterpolicyName)

	// Read the object back to refresh Optional+Computed attributes.
	r.readFilterpolicy(ctx, filterpolicyName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilterpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FilterpolicyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	filterpolicyName := data.Id.ValueString()
	tflog.Debug(ctx, fmt.Sprintf("Deleting filterpolicy %s", filterpolicyName))

	err := r.client.DeleteResource(service.Filterpolicy.Type(), filterpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete filterpolicy %s, got error: %s", filterpolicyName, err))
		return
	}
}

func (r *FilterpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// id == name; passthrough import (mirrors the legacy SDKv2
	// schema.ImportStatePassthroughContext).
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// readFilterpolicy fetches the filterpolicy by name and populates the model.
// Returns false (without error) when the object is not found so callers can drop
// it from state, matching the legacy SDKv2 read behavior.
func (r *FilterpolicyResource) readFilterpolicy(ctx context.Context, name string, data *FilterpolicyResourceModel, diags *diag.Diagnostics) bool {
	found, err := r.client.FindResource(service.Filterpolicy.Type(), name)
	if err != nil {
		// The legacy SDKv2 read treated any Find error as "not present" and
		// cleared the id. Preserve that behavior.
		return false
	}

	data.Name = types.StringValue(found["name"].(string))
	data.Id = types.StringValue(found["name"].(string))
	filterpolicySetAttrFromGet(data, found)
	return true
}

// filterpolicySetAttrFromGet copies the Optional+Computed attributes from the
// NITRO GET response into the model.
func filterpolicySetAttrFromGet(data *FilterpolicyResourceModel, found map[string]interface{}) {
	if v, ok := found["reqaction"]; ok && v != nil {
		data.Reqaction = types.StringValue(v.(string))
	}
	if v, ok := found["resaction"]; ok && v != nil {
		data.Resaction = types.StringValue(v.(string))
	}
	if v, ok := found["rule"]; ok && v != nil {
		data.Rule = types.StringValue(v.(string))
	}
}

// filterpolicyGetThePayloadFromthePlan builds the NITRO request body from the
// plan. Mirrors the legacy SDKv2 payload construction (only non-null attributes
// are sent; the vendored struct uses omitempty).
func filterpolicyGetThePayloadFromthePlan(ctx context.Context, data *FilterpolicyResourceModel) filter.Filterpolicy {
	tflog.Debug(ctx, "In filterpolicyGetThePayloadFromthePlan Function")

	filterpolicy := filter.Filterpolicy{
		Name: data.Name.ValueString(),
	}
	if !data.Reqaction.IsNull() && !data.Reqaction.IsUnknown() {
		filterpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Resaction.IsNull() && !data.Resaction.IsUnknown() {
		filterpolicy.Resaction = data.Resaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		filterpolicy.Rule = data.Rule.ValueString()
	}

	return filterpolicy
}

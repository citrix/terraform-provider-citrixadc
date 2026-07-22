package appfwlearningdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwlearningdataExportResource{}
var _ resource.ResourceWithConfigure = (*AppfwlearningdataExportResource)(nil)

func NewAppfwlearningdataExportResource() resource.Resource {
	return &AppfwlearningdataExportResource{}
}

// AppfwlearningdataExportResource defines the resource implementation.
type AppfwlearningdataExportResource struct {
	client *service.NitroClient
}

// AppfwlearningdataExportResourceModel describes the resource data model.
//
// This resource models the NITRO appfwlearningdata `?action=export` action, which
// exports the learned data for a given profile/security check to a target file.
// export is a one-shot side-effect action with no GET endpoint and no inverse
// API, so Read/Update/Delete are no-ops. The export payload carries exactly three
// attributes: profilename (mandatory), securitycheck (mandatory) and target
// (optional) — confirmed against the NITRO export payload and the NetScaler CLI
// `export appfw learningdata <profileName> <securityCheck> [-target <string>]`.
type AppfwlearningdataExportResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Profilename   types.String `tfsdk:"profilename"`
	Securitycheck types.String `tfsdk:"securitycheck"`
	Target        types.String `tfsdk:"target"`
}

func (r *AppfwlearningdataExportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningdata_export"
}

func (r *AppfwlearningdataExportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwlearningdataExportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	stringReplace := []planmodifier.String{stringplanmodifier.RequiresReplace()}
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwlearningdata_export resource.",
			},
			"profilename": schema.StringAttribute{
				// Mandatory per NITRO export payload (red/bold) and the CLI
				// (mandatory positional argument).
				Required:      true,
				PlanModifiers: stringReplace,
				Description:   "Name of the profile.",
			},
			"securitycheck": schema.StringAttribute{
				// Mandatory per NITRO export payload (red/bold) and the CLI
				// (mandatory positional argument).
				Required:      true,
				PlanModifiers: stringReplace,
				Description:   "Name of the security check. Possible values = startURL, cookieConsistency, fieldConsistency, crossSiteScripting, SQLInjection, fieldFormat, CSRFtag, XMLDoSCheck, XMLWSICheck, XMLAttachmentCheck, TotalXMLRequests, creditCardNumber, ContentType.",
			},
			"target": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: stringReplace,
				Description:   "Target filename for data to be exported.",
			},
		},
	}
}

func (r *AppfwlearningdataExportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwlearningdataExportResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Exporting appfwlearningdata (action-only resource)")
	payload := appfwlearningdata_exportGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes export as POST ?action=export. Use ActOnResource with the
	// case-sensitive "export" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Appfwlearningdata.Type(), &payload, "export")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to export appfwlearningdata, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered appfwlearningdata export")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("appfwlearningdata_export")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataExportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// export is a one-shot action. NITRO has no GET endpoint that reports
	// export-state, so Read is a pure preserve-state no-op.
	var data AppfwlearningdataExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for appfwlearningdata_export; NITRO has no query endpoint for export state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataExportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for export; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state AppfwlearningdataExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwlearningdata_export; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningdataExportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// export is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-export"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for appfwlearningdata_export; NITRO has no inverse of the export action")
}

func appfwlearningdata_exportGetThePayloadFromthePlan(ctx context.Context, data *AppfwlearningdataExportResourceModel) appfw.Appfwlearningdata {
	tflog.Debug(ctx, "In appfwlearningdata_exportGetThePayloadFromthePlan Function")

	// The export action accepts exactly profilename, securitycheck and target.
	appfwlearningdata := appfw.Appfwlearningdata{}
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		appfwlearningdata.Profilename = data.Profilename.ValueString()
	}
	if !data.Securitycheck.IsNull() && !data.Securitycheck.IsUnknown() {
		appfwlearningdata.Securitycheck = data.Securitycheck.ValueString()
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		appfwlearningdata.Target = data.Target.ValueString()
	}

	return appfwlearningdata
}

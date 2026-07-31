package opoption

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"
	"github.com/citrix/adc-nitro-go/service"
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
var _ resource.Resource = &OpoptionResource{}
var _ resource.ResourceWithConfigure = (*OpoptionResource)(nil)
var _ resource.ResourceWithImportState = (*OpoptionResource)(nil)

func NewOpoptionResource() resource.Resource {
	return &OpoptionResource{}
}

// OpoptionResource is the Plugin Framework implementation of citrixadc_opoption.
//
// It is a backward-compatible replacement for the legacy SDKv2 resource, which
// registered "citrixadc_opoption" => resourceCitrixAdcSnmpoption(). In other
// words citrixadc_opoption has always been an ALIAS of citrixadc_snmpoption:
// same schema, same NITRO object (snmp/snmpoption), same CRUD. Only the
// Terraform resource type name differs ("_opoption" vs "_snmpoption").
//
// snmpoption is an unnamed singleton (a global SNMP option object): there is no
// name/key and no NITRO ADD/DELETE. It is configured with an UpdateUnnamedResource
// (HTTP PUT) and read with FindResource(type, ""). Delete therefore only removes
// the object from Terraform state (it cannot be removed on the appliance), exactly
// as the SDKv2 resource did.
type OpoptionResource struct {
	client *service.NitroClient
}

// OpoptionResourceModel describes the resource data model. Field set, tfsdk names
// and types are identical to the SDKv2 snmpoption schema (all TypeString), so
// existing state and configuration remain valid.
type OpoptionResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Severityinfointrap   types.String `tfsdk:"severityinfointrap"`
	Partitionnameintrap  types.String `tfsdk:"partitionnameintrap"`
	Snmpset              types.String `tfsdk:"snmpset"`
	Snmptraplogging      types.String `tfsdk:"snmptraplogging"`
	Snmptraplogginglevel types.String `tfsdk:"snmptraplogginglevel"`
}

func (r *OpoptionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_opoption"
}

func (r *OpoptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *OpoptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *OpoptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				// The unnamed singleton has no natural key; Create assigns a static
				// synthetic id. Carry that id through updates so it is never planned
				// as unknown on a subsequent apply (Update does not re-derive it),
				// which otherwise yields "Provider returned invalid result object
				// after apply". Matches SDKv2, which preserved d.Id across updates.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The ID of the opoption resource.",
			},
			// Each attribute below mirrors the SDKv2 snmpoption schema exactly:
			// TypeString, Optional + Computed, no default. Computed lets NITRO
			// supply/echo values the user omits, matching the legacy behavior and
			// avoiding perpetual diffs.
			"severityinfointrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By default, the severity level info of the trap is not mentioned in the trap message. Enable this option to send severity level of trap as one of the varbind in the trap message.",
			},
			"partitionnameintrap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send partition name as a varbind in traps. By default the partition names are not sent as a varbind.",
			},
			"snmpset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept SNMP SET requests sent to the Citrix ADC, and allow SNMP managers to write values to MIB objects that are configured for write access.",
			},
			"snmptraplogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log any SNMP trap events (for SNMP alarms in which logging is enabled) even if no trap listeners are configured. With the default setting, SNMP trap events are logged if at least one trap listener is configured on the appliance.",
			},
			"snmptraplogginglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level of SNMP trap logs. The default value is INFORMATIONAL.",
			},
		},
	}
}

func (r *OpoptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OpoptionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating opoption resource")

	// Build the payload and configure the singleton via UpdateUnnamedResource
	// (HTTP PUT), exactly as the SDKv2 createSnmpoptionFunc did.
	opoption := opoptionGetThePayloadFromthePlan(ctx, &data)
	err := r.client.UpdateUnnamedResource(service.Snmpoption.Type(), &opoption)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create opoption, got error: %s", err))
		return
	}

	// Synthetic ID for the unnamed singleton. The SDKv2 resource assigned a
	// unique id here; the object is a global singleton so the id value is
	// immaterial to lookup (Read/Update/Delete operate on the unnamed object).
	// A stable static id keeps applies deterministic; imported legacy state keeps
	// whatever id it already holds and continues to function.
	data.Id = types.StringValue("tf-opoption")

	tflog.Trace(ctx, "Created opoption resource")

	// Read the configured state back from the appliance.
	r.readOpoptionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpoptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OpoptionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading opoption resource")

	// Mirror the SDKv2 readSnmpoptionFunc: if the object can no longer be read,
	// clear it from state rather than erroring out.
	getResponseData, err := r.client.FindResource(service.Snmpoption.Type(), "")
	if err != nil {
		tflog.Warn(ctx, "Clearing opoption state")
		resp.State.RemoveResource(ctx)
		return
	}

	opoptionSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpoptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OpoptionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating opoption resource")

	// Re-apply the full desired configuration via UpdateUnnamedResource. NITRO's
	// unnamed PUT is a set operation, so re-sending the desired values yields the
	// same end state the SDKv2 updateSnmpoptionFunc produced by sending only the
	// changed fields.
	opoption := opoptionGetThePayloadFromthePlan(ctx, &data)
	err := r.client.UpdateUnnamedResource(service.Snmpoption.Type(), &opoption)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update opoption, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated opoption resource")

	// Read the updated state back from the appliance.
	r.readOpoptionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OpoptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// snmpoption is an unnamed singleton and has no NITRO DELETE operation. As in
	// the SDKv2 deleteSnmpoptionFunc, deletion only removes the object from
	// Terraform state; the appliance configuration is left untouched.
	tflog.Debug(ctx, "Deleting opoption resource from state (no NITRO DELETE for the singleton)")
}

// readOpoptionFromApi reads the singleton back from the appliance and populates
// the model (id is preserved by the caller).
func (r *OpoptionResource) readOpoptionFromApi(ctx context.Context, data *OpoptionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpoption.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read opoption, got error: %s", err))
		return
	}

	opoptionSetAttrFromGet(ctx, data, getResponseData)
}

// opoptionGetThePayloadFromthePlan builds the snmpoption NITRO body from the plan.
func opoptionGetThePayloadFromthePlan(ctx context.Context, data *OpoptionResourceModel) snmp.Snmpoption {
	tflog.Debug(ctx, "In opoptionGetThePayloadFromthePlan Function")

	opoption := snmp.Snmpoption{}
	if !data.Severityinfointrap.IsNull() && !data.Severityinfointrap.IsUnknown() {
		opoption.Severityinfointrap = data.Severityinfointrap.ValueString()
	}
	if !data.Partitionnameintrap.IsNull() && !data.Partitionnameintrap.IsUnknown() {
		opoption.Partitionnameintrap = data.Partitionnameintrap.ValueString()
	}
	if !data.Snmpset.IsNull() && !data.Snmpset.IsUnknown() {
		opoption.Snmpset = data.Snmpset.ValueString()
	}
	if !data.Snmptraplogging.IsNull() && !data.Snmptraplogging.IsUnknown() {
		opoption.Snmptraplogging = data.Snmptraplogging.ValueString()
	}
	if !data.Snmptraplogginglevel.IsNull() && !data.Snmptraplogginglevel.IsUnknown() {
		opoption.Snmptraplogginglevel = data.Snmptraplogginglevel.ValueString()
	}

	return opoption
}

// opoptionSetAttrFromGet maps the NITRO GET response onto the model. It does not
// touch the id, which is owned by Create/import and preserved across reads.
func opoptionSetAttrFromGet(ctx context.Context, data *OpoptionResourceModel, getResponseData map[string]interface{}) *OpoptionResourceModel {
	tflog.Debug(ctx, "In opoptionSetAttrFromGet Function")

	if val, ok := getResponseData["severityinfointrap"]; ok && val != nil {
		data.Severityinfointrap = types.StringValue(val.(string))
	} else {
		data.Severityinfointrap = types.StringNull()
	}
	if val, ok := getResponseData["partitionnameintrap"]; ok && val != nil {
		data.Partitionnameintrap = types.StringValue(val.(string))
	} else {
		data.Partitionnameintrap = types.StringNull()
	}
	if val, ok := getResponseData["snmpset"]; ok && val != nil {
		data.Snmpset = types.StringValue(val.(string))
	} else {
		data.Snmpset = types.StringNull()
	}
	if val, ok := getResponseData["snmptraplogging"]; ok && val != nil {
		data.Snmptraplogging = types.StringValue(val.(string))
	} else {
		data.Snmptraplogging = types.StringNull()
	}
	if val, ok := getResponseData["snmptraplogginglevel"]; ok && val != nil {
		data.Snmptraplogginglevel = types.StringValue(val.(string))
	} else {
		data.Snmptraplogginglevel = types.StringNull()
	}

	return data
}

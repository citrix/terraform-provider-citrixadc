package nsacls

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsaclsResourceModel describes the resource data model.
//
// nsacls is a *convenience block* resource (mirrors the SDK v2
// citrixadc_nsacls): it manages a set of individual `nsacl` rules keyed by a
// synthetic handle (aclsname) and then applies them with a POST
// ?action=apply on the `nsacls` object. This is the exact same class of
// resource as the sibling rnat_clear migration and is implemented the same way.
type NsaclsResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Type             types.String `tfsdk:"type"`
	Aclsname         types.String `tfsdk:"aclsname"`
	AclsApplyTrigger types.String `tfsdk:"acls_apply_trigger"`
	Acl              types.Set    `tfsdk:"acl"`
}

// NsaclEntryModel describes a single extended ACL rule inside the `acl` set.
// The field set matches the SDK v2 nested schema exactly (names + types).
type NsaclEntryModel struct {
	Aclaction       types.String `tfsdk:"aclaction"`
	Aclname         types.String `tfsdk:"aclname"`
	Destipop        types.String `tfsdk:"destipop"`
	Destipval       types.String `tfsdk:"destipval"`
	Destportop      types.String `tfsdk:"destportop"`
	Destportval     types.String `tfsdk:"destportval"`
	Established     types.Bool   `tfsdk:"established"`
	Icmpcode        types.Int64  `tfsdk:"icmpcode"`
	Icmptype        types.Int64  `tfsdk:"icmptype"`
	Interface       types.String `tfsdk:"interface"`
	Logstate        types.String `tfsdk:"logstate"`
	Priority        types.Int64  `tfsdk:"priority"`
	Protocol        types.String `tfsdk:"protocol"`
	Protocolnumber  types.Int64  `tfsdk:"protocolnumber"`
	Ratelimit       types.Int64  `tfsdk:"ratelimit"`
	Srcipop         types.String `tfsdk:"srcipop"`
	Srcipval        types.String `tfsdk:"srcipval"`
	Srcmac          types.String `tfsdk:"srcmac"`
	Srcportop       types.String `tfsdk:"srcportop"`
	Srcportval      types.String `tfsdk:"srcportval"`
	State           types.String `tfsdk:"state"`
	Td              types.Int64  `tfsdk:"td"`
	Ttl             types.Int64  `tfsdk:"ttl"`
	Vlan            types.Int64  `tfsdk:"vlan"`
	Srcportdataset  types.String `tfsdk:"srcportdataset"`
	Srcipdataset    types.String `tfsdk:"srcipdataset"`
	Destportdataset types.String `tfsdk:"destportdataset"`
	Destipdataset   types.String `tfsdk:"destipdataset"`
}

// nsaclObjectType returns the object type of a single `acl` set element.
// It MUST match the tfsdk tags on NsaclEntryModel exactly.
func nsaclObjectType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aclaction":       types.StringType,
			"aclname":         types.StringType,
			"destipop":        types.StringType,
			"destipval":       types.StringType,
			"destportop":      types.StringType,
			"destportval":     types.StringType,
			"established":     types.BoolType,
			"icmpcode":        types.Int64Type,
			"icmptype":        types.Int64Type,
			"interface":       types.StringType,
			"logstate":        types.StringType,
			"priority":        types.Int64Type,
			"protocol":        types.StringType,
			"protocolnumber":  types.Int64Type,
			"ratelimit":       types.Int64Type,
			"srcipop":         types.StringType,
			"srcipval":        types.StringType,
			"srcmac":          types.StringType,
			"srcportop":       types.StringType,
			"srcportval":      types.StringType,
			"state":           types.StringType,
			"td":              types.Int64Type,
			"ttl":             types.Int64Type,
			"vlan":            types.Int64Type,
			"srcportdataset":  types.StringType,
			"srcipdataset":    types.StringType,
			"destportdataset": types.StringType,
			"destipdataset":   types.StringType,
		},
	}
}

// aclApplyTriggerValidator reproduces the SDK v2 validateAclAction check:
// acls_apply_trigger must be exactly "Yes" or "No" (case-sensitive).
type aclApplyTriggerValidator struct{}

func (v aclApplyTriggerValidator) Description(_ context.Context) string {
	return `must be one of ["Yes" "No"] (case-sensitive)`
}

func (v aclApplyTriggerValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v aclApplyTriggerValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	if value != "Yes" && value != "No" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid acls_apply_trigger value",
			fmt.Sprintf(`%q must be one of ["Yes" "No"] (case-sensitive). Received: %q`, req.Path, value),
		)
	}
}

func (r *NsaclsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsacls resource (the aclsname handle).",
			},
			// SDK v2: type was Optional+Computed+ForceNew with NO Default.
			// Optional+Computed with UseStateForUnknown (stable when unset) and
			// RequiresReplaceIfConfigured (a configured change recreates, matching
			// ForceNew). No schema Default - the effective CLASSIC default is
			// resolved in Create, matching SDK v2's ApplyResource(omitempty) behavior.
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of the acl ,default will be CLASSIC.\nAvailable options as follows:\n* CLASSIC - specifies the regular extended acls.\n* DFD - cluster specific acls,specifies hashmethod for steering of the packet in cluster .",
			},
			// aclsname is Optional+Computed exactly as in SDK v2: when omitted,
			// Create generates a synthetic "tf-nsacl-*" handle. This value is the
			// resource ID.
			"aclsname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name handle for this set of ACLs. If omitted, a unique tf-nsacl-* value is generated.",
			},
			// acls_apply_trigger is a provider-side toggle (Yes/No). SDK v2 reset it
			// to "No" on every Read so a config value of "Yes" always produces a diff
			// and re-applies the ACLs on every run. Optional+Computed with
			// UseStateForUnknown keeps it stable ("No") when unset.
			"acls_apply_trigger": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					aclApplyTriggerValidator{},
				},
				Description: "Set to \"Yes\" to re-apply the ACL set on every run. Reset to \"No\" after each read.",
			},
		},
		Blocks: map[string]schema.Block{
			// The set of extended ACL rules managed by this resource. The SDK v2
			// resource declared `acl` as a TypeSet with Elem: &schema.Resource{},
			// which is consumed with HCL *block* syntax (`acl { ... }`). The
			// backward-compatible framework equivalent is a SetNestedBlock (a block,
			// NOT an attribute): a SetNestedAttribute would force `acl = [ ... ]`
			// assignment syntax and break existing configs (Terraform reports
			// "Unsupported block type"). Blocks cannot be Optional/Computed; a set
			// block with zero elements decodes to an empty set. Nested attributes are
			// Optional (not Computed): Read is a state-preserving no-op (nsacls has no
			// aggregate GET), so a Computed nested attribute would remain unknown
			// after apply.
			"acl": schema.SetNestedBlock{
				Description: "Set of extended ACL rules to configure and apply. Rules removed from this set are deleted on the appliance.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"aclname": schema.StringAttribute{
							Required:    true,
							Description: "Name for the extended ACL rule.",
						},
						"aclaction": schema.StringAttribute{
							Optional:    true,
							Description: "Action to perform on incoming IPv4 packets that match the extended ACL rule. Possible values = ALLOW, BRIDGE, DENY.",
						},
						"destipop": schema.StringAttribute{
							Optional:    true,
							Description: "Either the equals (=) or does not equal (!=) logical operator.",
						},
						"destipval": schema.StringAttribute{
							Optional:    true,
							Description: "IP address or range of IP addresses to match against the destination IP address of an incoming IPv4 packet.",
						},
						"destportop": schema.StringAttribute{
							Optional:    true,
							Description: "Either the equals (=) or does not equal (!=) logical operator.",
						},
						"destportval": schema.StringAttribute{
							Optional:    true,
							Description: "Port number or range of port numbers to match against the destination port number of an incoming IPv4 packet.",
						},
						"established": schema.BoolAttribute{
							Optional:    true,
							Description: "Allow only incoming TCP packets that have the ACK or RST bit set.",
						},
						"icmpcode": schema.Int64Attribute{
							Optional:    true,
							Description: "Code of a particular ICMP message type to match against the ICMP code of an incoming ICMP packet.",
						},
						"icmptype": schema.Int64Attribute{
							Optional:    true,
							Description: "ICMP Message type to match against the message type of an incoming ICMP packet.",
						},
						"interface": schema.StringAttribute{
							Optional:    true,
							Description: "ID of an interface. The Citrix ADC applies the ACL rule only to the incoming packets from the specified interface.",
						},
						"logstate": schema.StringAttribute{
							Optional:    true,
							Description: "Enable or disable logging of events related to the extended ACL rule. Possible values = ENABLED, DISABLED.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Description: "Priority for the extended ACL rule that determines the order in which it is evaluated relative to the other extended ACL rules.",
						},
						"protocol": schema.StringAttribute{
							Optional:    true,
							Description: "Protocol to match against the protocol of an incoming IPv4 packet.",
						},
						"protocolnumber": schema.Int64Attribute{
							Optional:    true,
							Description: "Protocol to match against the protocol of an incoming IPv4 packet.",
						},
						"ratelimit": schema.Int64Attribute{
							Optional:    true,
							Description: "Maximum number of log messages to be generated per second.",
						},
						"srcipop": schema.StringAttribute{
							Optional:    true,
							Description: "Either the equals (=) or does not equal (!=) logical operator.",
						},
						"srcipval": schema.StringAttribute{
							Optional:    true,
							Description: "IP address or range of IP addresses to match against the source IP address of an incoming IPv4 packet.",
						},
						"srcmac": schema.StringAttribute{
							Optional:    true,
							Description: "MAC address to match against the source MAC address of an incoming IPv4 packet.",
						},
						"srcportop": schema.StringAttribute{
							Optional:    true,
							Description: "Either the equals (=) or does not equal (!=) logical operator.",
						},
						"srcportval": schema.StringAttribute{
							Optional:    true,
							Description: "Port number or range of port numbers to match against the source port number of an incoming IPv4 packet.",
						},
						"state": schema.StringAttribute{
							Optional:    true,
							Description: "Enable or disable the extended ACL rule. Possible values = ENABLED, DISABLED.",
						},
						"td": schema.Int64Attribute{
							Optional:    true,
							Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity.",
						},
						"ttl": schema.Int64Attribute{
							Optional:    true,
							Description: "Number of seconds, in multiples of four, after which the extended ACL rule expires.",
						},
						"vlan": schema.Int64Attribute{
							Optional:    true,
							Description: "ID of the VLAN. The Citrix ADC applies the ACL rule only to the incoming packets of the specified VLAN.",
						},
						"srcportdataset": schema.StringAttribute{
							Optional:    true,
							Description: "Policy dataset which can have multiple port ranges bound to it.",
						},
						"srcipdataset": schema.StringAttribute{
							Optional:    true,
							Description: "Policy dataset which can have multiple IP ranges bound to it.",
						},
						"destportdataset": schema.StringAttribute{
							Optional:    true,
							Description: "Policy dataset which can have multiple port ranges bound to it.",
						},
						"destipdataset": schema.StringAttribute{
							Optional:    true,
							Description: "Policy dataset which can have multiple IP ranges bound to it.",
						},
					},
				},
			},
		},
	}
}

// nsaclsBuildNsaclPayload builds the NITRO ns.Nsacl object for a single ACL rule
// from its model, reproducing the SDK v2 createSingleAcl semantics: the
// destip/destport/srcip/srcport booleans are derived from whether the
// corresponding val/dataset is set, and op-without-val combinations are rejected.
func nsaclsBuildNsaclPayload(ctx context.Context, m *NsaclEntryModel) (ns.Nsacl, error) {
	tflog.Debug(ctx, "In nsaclsBuildNsaclPayload Function")

	nsaclName := m.Aclname.ValueString()

	destip := strSet(m.Destipval) || strSet(m.Destipdataset)
	destport := strSet(m.Destportval) || strSet(m.Destportdataset)
	srcip := strSet(m.Srcipval) || strSet(m.Srcipdataset)
	srcport := strSet(m.Srcportval) || strSet(m.Srcportdataset)

	if strSet(m.Destipop) && !strSet(m.Destipval) {
		return ns.Nsacl{}, fmt.Errorf("Error in nsacl spec %s cannot have destipop without destipval", nsaclName)
	}
	if strSet(m.Destportop) && !strSet(m.Destportval) {
		return ns.Nsacl{}, fmt.Errorf("Error in nsacl spec %s cannot have destportop without destportval", nsaclName)
	}
	if strSet(m.Srcipop) && !strSet(m.Srcipval) {
		return ns.Nsacl{}, fmt.Errorf("Error in nsacl spec %s cannot have srcipop without srcipval", nsaclName)
	}
	if strSet(m.Srcportop) && !strSet(m.Srcportval) {
		return ns.Nsacl{}, fmt.Errorf("Error in nsacl spec %s cannot have srcportop without srcportval", nsaclName)
	}

	nsacl := ns.Nsacl{
		Aclname:  nsaclName,
		Destip:   destip,
		Destport: destport,
		Srcip:    srcip,
		Srcport:  srcport,
	}

	if strSet(m.Aclaction) {
		nsacl.Aclaction = m.Aclaction.ValueString()
	}
	if strSet(m.Destipop) {
		nsacl.Destipop = m.Destipop.ValueString()
	}
	if strSet(m.Destipval) {
		nsacl.Destipval = m.Destipval.ValueString()
	}
	if strSet(m.Destipdataset) {
		nsacl.Destipdataset = m.Destipdataset.ValueString()
	}
	if strSet(m.Destportop) {
		nsacl.Destportop = m.Destportop.ValueString()
	}
	if strSet(m.Destportval) {
		nsacl.Destportval = m.Destportval.ValueString()
	}
	if strSet(m.Destportdataset) {
		nsacl.Destportdataset = m.Destportdataset.ValueString()
	}
	if !m.Established.IsNull() && !m.Established.IsUnknown() {
		nsacl.Established = m.Established.ValueBool()
	}
	if strSet(m.Interface) {
		nsacl.Interface = m.Interface.ValueString()
	}
	if strSet(m.Logstate) {
		nsacl.Logstate = m.Logstate.ValueString()
	}
	if strSet(m.Protocol) {
		nsacl.Protocol = m.Protocol.ValueString()
	}
	if strSet(m.Srcipop) {
		nsacl.Srcipop = m.Srcipop.ValueString()
	}
	if strSet(m.Srcipval) {
		nsacl.Srcipval = m.Srcipval.ValueString()
	}
	if strSet(m.Srcipdataset) {
		nsacl.Srcipdataset = m.Srcipdataset.ValueString()
	}
	if strSet(m.Srcmac) {
		nsacl.Srcmac = m.Srcmac.ValueString()
	}
	if strSet(m.Srcportop) {
		nsacl.Srcportop = m.Srcportop.ValueString()
	}
	if strSet(m.Srcportval) {
		nsacl.Srcportval = m.Srcportval.ValueString()
	}
	if strSet(m.Srcportdataset) {
		nsacl.Srcportdataset = m.Srcportdataset.ValueString()
	}
	if strSet(m.State) {
		nsacl.State = m.State.ValueString()
	}

	if intSet(m.Icmpcode) {
		nsacl.Icmpcode = utils.IntPtr(int(m.Icmpcode.ValueInt64()))
	}
	if intSet(m.Icmptype) {
		nsacl.Icmptype = utils.IntPtr(int(m.Icmptype.ValueInt64()))
	}
	if intSet(m.Priority) {
		nsacl.Priority = utils.IntPtr(int(m.Priority.ValueInt64()))
	}
	if intSet(m.Protocolnumber) {
		nsacl.Protocolnumber = utils.IntPtr(int(m.Protocolnumber.ValueInt64()))
	}
	if intSet(m.Ratelimit) {
		nsacl.Ratelimit = utils.IntPtr(int(m.Ratelimit.ValueInt64()))
	}
	if intSet(m.Td) {
		nsacl.Td = utils.IntPtr(int(m.Td.ValueInt64()))
	}
	if intSet(m.Ttl) {
		nsacl.Ttl = utils.IntPtr(int(m.Ttl.ValueInt64()))
	}
	if intSet(m.Vlan) {
		nsacl.Vlan = utils.IntPtr(int(m.Vlan.ValueInt64()))
	}

	return nsacl, nil
}

// strSet reports whether a types.String is a non-null, non-unknown, non-empty value.
func strSet(v types.String) bool {
	return !v.IsNull() && !v.IsUnknown() && v.ValueString() != ""
}

// intSet reports whether a types.Int64 has been explicitly set.
func intSet(v types.Int64) bool {
	return !v.IsNull() && !v.IsUnknown()
}

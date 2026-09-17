# Install an SSL certificate with write-only (ephemeral) content

End-to-end example of the recommended pattern for installing an SSL cert/key on a
NetScaler when the material comes from a secret store (Vault, etc.) and **must not
land in Terraform state**:

```
secret store / sensitive variable
        │  (cert_pem, key_pem, key_passphrase)
        ▼
citrixadc_systemfile.filecontent_wo   ── uploads the files, content never stored in state
        │  (filecontent_wo_version = rotation counter)
        ▼
citrixadc_sslcertkey                  ── installs the pair; cert_hash/key_hash drive a
                                         graceful in-place reload on rotation
```

## Why write-only

`citrixadc_systemfile.filecontent` is stored in Terraform state. For a private key
that is unacceptable. `filecontent_wo` is a [write-only argument](https://developer.hashicorp.com/terraform/language/resources/ephemeral/write-only):
its value is supplied at apply time and is **never** written to state, plan output,
or logs. The passphrase uses the same mechanism via `sslcertkey.passplain_wo`.

> **Scope:** this keeps the material out of Terraform **state**. The appliance still
> stores the file and returns it on a NITRO GET, so this is state hygiene, not
> end-to-end secrecy.

## Requirements

- **Terraform >= 1.11** (write-only arguments).
- A `citrix/citrixadc` provider build that includes `systemfile.filecontent_wo`.
- Network access from the machine running Terraform to the NetScaler NITRO endpoint.

## Run it

```sh
# 1. Generate a self-signed pair for testing (or fetch real material):
openssl req -x509 -newkey rsa:2048 -nodes \
  -keyout server.key -out server.crt -days 365 -subj "/CN=example.test"

# 2. Provide inputs (env keeps secrets off disk):
export TF_VAR_adc_endpoint='http://<nsip>'
export TF_VAR_adc_password='<password>'
export TF_VAR_cert_pem="$(cat server.crt)"
export TF_VAR_key_pem="$(cat server.key)"
# export TF_VAR_key_passphrase='...'   # only if the key is encrypted

# 3. Apply:
terraform init
terraform apply
```

Confirm the key content is **not** in state:

```sh
terraform show -json | grep -c filecontent_wo   # -> 0 (write-only, absent from state)
```

## Rotating the certificate

`systemfile` has no in-place update, so a content change is a re-upload
(destroy+recreate of the file). Rotation is driven by a single counter,
`cert_material_version`:

1. Update `cert_pem` / `key_pem` with the new material.
2. Increment `cert_material_version` (e.g. `1` -> `2`) **in the same apply**.
3. `terraform apply`.

What happens: the `systemfile` files are re-uploaded, and because
`sslcertkey.cert_hash`/`key_hash` reference `filecontent_wo_version`, the hashes
change too, so `sslcertkey` reloads the certificate **in place** (NITRO
`?action=update`) — bindings on vservers are preserved. The resource references
also order the reload after the re-upload.

### Why `cert_hash = tostring(...filecontent_wo_version)`?

`cert_hash`/`key_hash` only need a value that **changes** when the content changes;
they are not the secret, so they live in state safely (as `sslcertkey` already does
with `key_hash`). You **cannot** derive them from `filecontent_wo` — a write-only
argument is null in state and cannot be referenced in expressions. Two valid sources:

- `tostring(citrixadc_systemfile.cert.filecontent_wo_version)` — used here. One bump
  drives both the re-upload and the reload, and creates the dependency edge.
- `sha256(var.cert_pem)` — hash the original source directly. If you use this, you
  must still bump `filecontent_wo_version` on rotation, or the file won't re-upload
  and `sslcertkey` would reload stale content.

## Sourcing from HashiCorp Vault

Uncomment the `vault` provider in `provider.tf` and the `vault_kv_secret_v2` data
source in `resources.tf`, then feed its outputs into the resources and use
`metadata.version` as the rotation counter so rotation is automatic:

```hcl
data "vault_kv_secret_v2" "tls" {
  mount = "secret"
  name  = "netscaler/tls"
}

resource "citrixadc_systemfile" "cert" {
  filename               = local.cert_filename
  filelocation           = local.ssl_dir
  filecontent_wo         = data.vault_kv_secret_v2.tls.data["certificate"]
  filecontent_wo_version = data.vault_kv_secret_v2.tls.metadata.version
}
# ...same for the key, and set cert_hash = tostring(...filecontent_wo_version)
```

## Notes / caveats

- **Exactly one** of `filecontent` or `filecontent_wo` may be set per `systemfile`.
- **Do not `terraform import`** a `systemfile` managed with `filecontent_wo`: the
  write-only content cannot be recovered into state. Import is for the plaintext
  `filecontent` path.
- For an unencrypted key, leave `key_passphrase` empty; `password`/`passplain_wo`
  are then omitted automatically.

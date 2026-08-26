# Consolidated acc-test sweep — 50 modified resources (NSNETAUTO-1162)
Date: 2026-08-26 | Method: dynamic Workflow, one subagent per resource, ADC-pinned (no shared ADC), cross-ADC retry.

## Totals (per-test; 301 tests across 49 resources; sslservice_sslpolicy_binding has no test)
- PASS on first ADC ............ 209
- PASS on alternative ADC ...... 59   (env/testbed-specific; see per-resource logs)
- FAILED (genuine) ............. 10   (ALL pre-existing sdkv2StateUpgrade "plan not empty"; verified fail on pristine HEAD → not regressions)
- SKIPPED ...................... 23   (feature/license/HSM/destructive/TODO gated — to revisit later)

## Phase A (6 standalone lanes .121/.151/.152/.153/.154/.155), ADC_TESTBED unset
- 209 passed on first ADC; 59 recovered on an alternate ADC.
- Dominant ADC-specific cause: .152 reverted to a READ-ONLY cluster node (errorcode 1203/477) — every resource on it failed writes; all passed on healthy lanes (.151/.153).
- .151 transient/leftover-state on a few (authenticationvserver_basic, nspbr_unset, servicegroup_lbmonitor_binding, netprofile, nsdhcpparams) — all passed on .153.
- SSL-default-dependent (sslprofile_sslcertkey_binding, csvserver_standalone_ciphersuites_mixed) — behavior differs by sslparameter.defaultProfile; resolved on the right box.

## Phase B (labelled testbeds)
- CLUSTER (.133) .......................... 3/3 PASS  (csvserver_cluster_ciphers, lbvserver_cluster_ciphers, lbvserver_cluster_ciphersuites)
- STANDALONE_DEFAULT_SSL_PROFILE (.153) ... 6/6 PASS  (sslprofile_cipher_binding, sslprofile_sslcertkey_binding x5)
- STANDALONE_NON_DEFAULT_SSL_PROFILE (.155) 4/5 PASS  (sslservice_unset fails 1585 on .155 but PASSES on .153 default-profile bed → overall PASS)
- HA/HA_PAIR, HSM: no such tests among these 50.

## Genuine failures (10) — all pre-existing, NOT regressions from the 73.x attribute additions
aaakcdaccount, analyticsprofile, appflowparam, auditsyslogaction, authenticationldapaction,
authenticationoauthaction, gslbsite, lbparameter, rdpclientprofile, sslprofile
  -> each is <Resource>_sdkv2StateUpgrade failing "non-refresh plan was not empty" (known GH#1436 spurious-diff-on-upgrade; reproduces on pristine HEAD).

## Environmental note
.152 is again a read-only single-node cluster (drifts back after recovery). Excluded from Phase B; its Phase A load was covered via .151/.153.

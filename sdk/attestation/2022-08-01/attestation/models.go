package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "home/runner/work/kermit/kermit/sdk/attestation/2022-08-01/attestation"

            // AttestOpenEnclaveRequest attestation request for Intel SGX enclaves
            type AttestOpenEnclaveRequest struct {
            // Report - OpenEnclave report from the enclave to be attested (a URL-encoded base64 string)
            Report *string `json:"report,omitempty"`
            // RuntimeData - Runtime data provided by the enclave at the time of report generation. The MAA will verify that the first 32 bytes of the report_data field of the quote contains the SHA256 hash of the decoded "data" field of the runtime data.
            RuntimeData *RuntimeData `json:"runtimeData,omitempty"`
            // InitTimeData - Base64Url encoded "InitTime data". The MAA will verify that the init data was known to the enclave. Note that InitTimeData is invalid for CoffeeLake processors.
            InitTimeData *InitTimeData `json:"initTimeData,omitempty"`
            // DraftPolicyForAttestation - Attest against the provided draft policy. Note that the resulting token cannot be validated.
            DraftPolicyForAttestation *string `json:"draftPolicyForAttestation,omitempty"`
            // Nonce - Nonce for incoming request - emitted in the generated attestation token
            Nonce *string `json:"nonce,omitempty"`
            }

            // AttestSevSnpVMRequest attestation request for AMD SEV SNP Virtual Machine
            type AttestSevSnpVMRequest struct {
            // Report - Hardware rooted report of the virtual machine being attested along with the signing certificate chain and optionally, additional endorsements
            Report *string `json:"report,omitempty"`
            // RuntimeData - Runtime data provided by the enclave at the time of report generation. The MAA will verify that the run time data is known to the attestation target.
            RuntimeData *RuntimeData `json:"runtimeData,omitempty"`
            // InitTimeData - Initialization data provided by the enclave at the time of report generation. The MAA will verify that the init time data is known to the attestation target.
            InitTimeData *InitTimeData `json:"initTimeData,omitempty"`
            // DraftPolicyForAttestation - Attest against the provided draft policy. Note that the resulting token cannot be validated.
            DraftPolicyForAttestation *string `json:"draftPolicyForAttestation,omitempty"`
            // Nonce - Nonce for incoming request - emitted in the generated attestation token
            Nonce *string `json:"nonce,omitempty"`
            }

            // AttestSgxEnclaveRequest attestation request for Intel SGX enclaves
            type AttestSgxEnclaveRequest struct {
            // Quote - Quote of the enclave to be attested (a URL-encoded base64 string)
            Quote *string `json:"quote,omitempty"`
            // RuntimeData - Runtime data provided by the enclave at the time of quote generation. The MAA will verify that the first 32 bytes of the report_data field of the quote contains the SHA256 hash of the decoded "data" field of the runtime data.
            RuntimeData *RuntimeData `json:"runtimeData,omitempty"`
            // InitTimeData - Initialization data provided when the enclave is created. MAA will verify that the init data was known to the enclave. Note that InitTimeData is invalid for CoffeeLake processors.
            InitTimeData *InitTimeData `json:"initTimeData,omitempty"`
            // DraftPolicyForAttestation - Attest against the provided draft policy. Note that the resulting token cannot be validated.
            DraftPolicyForAttestation *string `json:"draftPolicyForAttestation,omitempty"`
            // Nonce - Nonce for incoming request - emitted in the generated attestation token
            Nonce *string `json:"nonce,omitempty"`
            }

            // CertificateManagementBody the body of the JWT used for the PolicyCertificates APIs
            type CertificateManagementBody struct {
            // PolicyCertificate - RFC 7517 Json Web Key describing the certificate.
            PolicyCertificate *JSONWebKey `json:"policyCertificate,omitempty"`
            }

            // CloudError an error response from Attestation.
            type CloudError struct {
            Error *CloudErrorBody `json:"error,omitempty"`
            }

            // CloudErrorBody an error response from Attestation.
            type CloudErrorBody struct {
            // Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
            Code *string `json:"code,omitempty"`
            // Message - A message describing the error, intended to be suitable for displaying in a user interface.
            Message *string `json:"message,omitempty"`
            }

            // InitTimeData initialization time data are a conduit for any configuration information that is unknown
            // when building the Trusted Execution Environment (TEE) and is defined at TEE launch time. This data can
            // be used with confidential container or VM scenarios to capture configuration settings such as disk
            // volume content, network configuration, etc.
            type InitTimeData struct {
            // Data - Initialization time data are passed into the Trusted Execution Environment (TEE) when it is created. For an Icelake SGX quote, the SHA256 hash of the InitTimeData must match the lower 32 bytes of the quote's "config id" attribute. For a SEV-SNP quote, the SHA256 hash of the InitTimeData must match the quote's "host data" attribute. (a URL-encoded base64 string)
            Data *string `json:"data,omitempty"`
            // DataType - The type of data contained within the "data" field. Possible values include: 'DataTypeBinary', 'DataTypeJSON'
            DataType DataType `json:"dataType,omitempty"`
            }

            // JSONWebKey ...
            type JSONWebKey struct {
            // Alg - The "alg" (algorithm) parameter identifies the algorithm intended for
            // use with the key.  The values used should either be registered in the
            // IANA "JSON Web Signature and Encryption Algorithms" registry
            // established by [JWA] or be a value that contains a Collision-
            // Resistant Name.
            Alg *string `json:"alg,omitempty"`
            // Crv - The "crv" (curve) parameter identifies the curve type
            Crv *string `json:"crv,omitempty"`
            // D - RSA private exponent or ECC private key
            D *string `json:"d,omitempty"`
            // Dp - RSA Private Key Parameter
            Dp *string `json:"dp,omitempty"`
            // Dq - RSA Private Key Parameter
            Dq *string `json:"dq,omitempty"`
            // E - RSA public exponent, in Base64
            E *string `json:"e,omitempty"`
            // K - Symmetric key
            K *string `json:"k,omitempty"`
            // Kid - The "kid" (key ID) parameter is used to match a specific key.  This
            // is used, for instance, to choose among a set of keys within a JWK Set
            // during key rollover.  The structure of the "kid" value is
            // unspecified.  When "kid" values are used within a JWK Set, different
            // keys within the JWK Set SHOULD use distinct "kid" values.  (One
            // example in which different keys might use the same "kid" value is if
            // they have different "kty" (key type) values but are considered to be
            // equivalent alternatives by the application using them.)  The "kid"
            // value is a case-sensitive string.
            Kid *string `json:"kid,omitempty"`
            // Kty - The "kty" (key type) parameter identifies the cryptographic algorithm
            // family used with the key, such as "RSA" or "EC". "kty" values should
            // either be registered in the IANA "JSON Web Key Types" registry
            // established by [JWA] or be a value that contains a Collision-
            // Resistant Name.  The "kty" value is a case-sensitive string.
            Kty *string `json:"kty,omitempty"`
            // N - RSA modulus, in Base64
            N *string `json:"n,omitempty"`
            // P - RSA secret prime
            P *string `json:"p,omitempty"`
            // Q - RSA secret prime, with p < q
            Q *string `json:"q,omitempty"`
            // Qi - RSA Private Key Parameter
            Qi *string `json:"qi,omitempty"`
            // Use - Use ("public key use") identifies the intended use of
            // the public key. The "use" parameter is employed to indicate whether
            // a public key is used for encrypting data or verifying the signature
            // on data. Values are commonly "sig" (signature) or "enc" (encryption).
            Use *string `json:"use,omitempty"`
            // X - X coordinate for the Elliptic Curve point
            X *string `json:"x,omitempty"`
            // X5c - The "x5c" (X.509 certificate chain) parameter contains a chain of one
            // or more PKIX certificates [RFC5280].  The certificate chain is
            // represented as a JSON array of certificate value strings.  Each
            // string in the array is a base64-encoded (Section 4 of [RFC4648] --
            // not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value.
            // The PKIX certificate containing the key value MUST be the first
            // certificate.
            X5c *[]string `json:"x5c,omitempty"`
            // Y - Y coordinate for the Elliptic Curve point
            Y *string `json:"y,omitempty"`
            }

            // JSONWebKeySet ...
            type JSONWebKeySet struct {
            autorest.Response `json:"-"`
            // Keys - The value of the "keys" parameter is an array of JWK values.  By
            // default, the order of the JWK values within the array does not imply
            // an order of preference among them, although applications of JWK Sets
            // can choose to assign a meaning to the order for their purposes, if
            // desired.
            Keys *[]JSONWebKey `json:"keys,omitempty"`
            }

            // OpenIDConfigurationResponse the response to the OpenID metadata description document API
            type OpenIDConfigurationResponse struct {
            autorest.Response `json:"-"`
            // ResponseTypesSupported - Types supported in the OpenID metadata API
            ResponseTypesSupported *[]string `json:"response_types_supported,omitempty"`
            // IDTokenSigningAlgValuesSupported - List of the supported signing algorithms
            IDTokenSigningAlgValuesSupported *[]string `json:"id_token_signing_alg_values_supported,omitempty"`
            // RevocationEndpoint - Revocation endpoint
            RevocationEndpoint *string `json:"revocation_endpoint,omitempty"`
            // Issuer - Issuer tenant base endpoint
            Issuer *string `json:"issuer,omitempty"`
            // JwksURI - The URI to retrieve the signing keys
            JwksURI *string `json:"jwks_uri,omitempty"`
            // ClaimsSupported - Set of claims supported by the OpenID metadata endpoint
            ClaimsSupported *[]string `json:"claims_supported,omitempty"`
            }

            // PolicyCertificatesModificationResult the result of a policy certificate modification
            type PolicyCertificatesModificationResult struct {
            // CertificateThumbprint - Hex encoded SHA1 Hash of the binary representation certificate which was added or removed
            CertificateThumbprint *string `json:"x-ms-certificate-thumbprint,omitempty"`
            // CertificateResolution - The result of the operation. Possible values include: 'CertificateModificationIsPresent', 'CertificateModificationIsAbsent'
            CertificateResolution CertificateModification `json:"x-ms-policycertificates-result,omitempty"`
            }

            // PolicyCertificatesModifyResponse the response to an attestation policy management API
            type PolicyCertificatesModifyResponse struct {
            autorest.Response `json:"-"`
            // Token - An RFC7519 JSON Web Token structure whose body is a PolicyCertificatesModificationResult object.
            Token *string `json:"token,omitempty"`
            }

            // PolicyCertificatesResponse the response to an attestation policy management API
            type PolicyCertificatesResponse struct {
            autorest.Response `json:"-"`
            // Token - An RFC7519 JSON Web Token structure containing a PolicyCertificatesResults object which contains the certificates used to validate policy changes
            Token *string `json:"token,omitempty"`
            }

            // PolicyCertificatesResult the result of a call to retrieve policy certificates.
            type PolicyCertificatesResult struct {
            // PolicyCertificates - SHA256 Hash of the binary representation certificate which was added or removed
            PolicyCertificates *JSONWebKeySet `json:"x-ms-policy-certificates,omitempty"`
            }

            // PolicyResponse the response to an attestation policy operation
            type PolicyResponse struct {
            autorest.Response `json:"-"`
            // Token - An RFC7519 JSON Web Token structure whose body is an PolicyResult object.
            Token *string `json:"token,omitempty"`
            }

            // PolicyResult the result of a policy certificate modification
            type PolicyResult struct {
            // PolicyResolution - The result of the operation. Possible values include: 'PolicyModificationUpdated', 'PolicyModificationRemoved'
            PolicyResolution PolicyModification `json:"x-ms-policy-result,omitempty"`
            // PolicyTokenHash - The SHA256 hash of the policy object modified (a URL-encoded base64 string)
            PolicyTokenHash *string `json:"x-ms-policy-token-hash,omitempty"`
            // PolicySigner - The certificate used to sign the policy object, if specified
            PolicySigner *JSONWebKey `json:"x-ms-policy-signer,omitempty"`
            // Policy - A JSON Web Token containing a StoredAttestationPolicy object with the attestation policy
            Policy *string `json:"x-ms-policy,omitempty"`
            }

            // Response the result of an attestation operation
            type Response struct {
            autorest.Response `json:"-"`
            // Token - An RFC 7519 JSON Web Token, the body of which is an AttestationResult object.
            Token *string `json:"token,omitempty"`
            }

            // Result a Microsoft Azure Attestation response token body - the body of a response token issued by MAA
            type Result struct {
            // Jti - Unique Identifier for the token
            Jti *string `json:"jti,omitempty"`
            // Iss - The Principal who issued the token
            Iss *string `json:"iss,omitempty"`
            // Iat - The time at which the token was issued, in the number of seconds since 1970-01-0T00:00:00Z UTC
            Iat *float64 `json:"iat,omitempty"`
            // Exp - The expiration time after which the token is no longer valid, in the number of seconds since 1970-01-0T00:00:00Z UTC
            Exp *float64 `json:"exp,omitempty"`
            // Nbf - The not before time before which the token cannot be considered valid, in the number of seconds since 1970-01-0T00:00:00Z UTC
            Nbf *float64 `json:"nbf,omitempty"`
            // Cnf - An RFC 7800 Proof of Possession Key
            Cnf interface{} `json:"cnf,omitempty"`
            // Nonce - The Nonce input to the attestation request, if provided.
            Nonce *string `json:"nonce,omitempty"`
            // Version - The Schema version of this structure. Current Value: 1.0
            Version *string `json:"x-ms-ver,omitempty"`
            // RuntimeClaims - Runtime Claims
            RuntimeClaims interface{} `json:"x-ms-runtime,omitempty"`
            // InittimeClaims - Inittime Claims
            InittimeClaims interface{} `json:"x-ms-inittime,omitempty"`
            // PolicyClaims - Policy Generated Claims
            PolicyClaims interface{} `json:"x-ms-policy,omitempty"`
            // VerifierType - The Attestation type being attested.
            VerifierType *string `json:"x-ms-attestation-type,omitempty"`
            // PolicySigner - The certificate used to sign the policy object, if specified.
            PolicySigner *JSONWebKey `json:"x-ms-policy-signer,omitempty"`
            // PolicyHash - The SHA256 hash of the BASE64URL encoded policy text used for attestation (a URL-encoded base64 string)
            PolicyHash *string `json:"x-ms-policy-hash,omitempty"`
            // IsDebuggable - True if the enclave is debuggable, false otherwise
            IsDebuggable *bool `json:"x-ms-sgx-is-debuggable,omitempty"`
            // ProductID - The SGX Product ID for the enclave.
            ProductID *float64 `json:"x-ms-sgx-product-id,omitempty"`
            // MrEnclave - The HEX encoded SGX MRENCLAVE value for the enclave.
            MrEnclave *string `json:"x-ms-sgx-mrenclave,omitempty"`
            // MrSigner - The HEX encoded SGX MRSIGNER value for the enclave.
            MrSigner *string `json:"x-ms-sgx-mrsigner,omitempty"`
            // Svn - The SGX SVN value for the enclave.
            Svn *float64 `json:"x-ms-sgx-svn,omitempty"`
            // EnclaveHeldData - A copy of the RuntimeData specified as an input to the attest call. (a URL-encoded base64 string)
            EnclaveHeldData *string `json:"x-ms-sgx-ehd,omitempty"`
            // SgxCollateral - The SGX SVN value for the enclave.
            SgxCollateral interface{} `json:"x-ms-sgx-collateral,omitempty"`
            // DeprecatedVersion - DEPRECATED: Private Preview version of x-ms-ver claim.
            DeprecatedVersion *string `json:"ver,omitempty"`
            // DeprecatedIsDebuggable - DEPRECATED: Private Preview version of x-ms-sgx-is-debuggable claim.
            DeprecatedIsDebuggable *bool `json:"is-debuggable,omitempty"`
            // DeprecatedSgxCollateral - DEPRECATED: Private Preview version of x-ms-sgx-collateral claim.
            DeprecatedSgxCollateral interface{} `json:"maa-attestationcollateral,omitempty"`
            // DeprecatedEnclaveHeldData - DEPRECATED: Private Preview version of x-ms-sgx-ehd claim. (a URL-encoded base64 string)
            DeprecatedEnclaveHeldData *string `json:"aas-ehd,omitempty"`
            // DeprecatedEnclaveHeldData2 - DEPRECATED: Private Preview version of x-ms-sgx-ehd claim. (a URL-encoded base64 string)
            DeprecatedEnclaveHeldData2 *string `json:"maa-ehd,omitempty"`
            // DeprecatedProductID - DEPRECATED: Private Preview version of x-ms-sgx-product-id
            DeprecatedProductID *float64 `json:"product-id,omitempty"`
            // DeprecatedMrEnclave - DEPRECATED: Private Preview version of x-ms-sgx-mrenclave.
            DeprecatedMrEnclave *string `json:"sgx-mrenclave,omitempty"`
            // DeprecatedMrSigner - DEPRECATED: Private Preview version of x-ms-sgx-mrsigner.
            DeprecatedMrSigner *string `json:"sgx-mrsigner,omitempty"`
            // DeprecatedSvn - DEPRECATED: Private Preview version of x-ms-sgx-svn.
            DeprecatedSvn *float64 `json:"svn,omitempty"`
            // DeprecatedTee - DEPRECATED: Private Preview version of x-ms-tee.
            DeprecatedTee *string `json:"tee,omitempty"`
            // DeprecatedPolicySigner - DEPRECATED: Private Preview version of x-ms-policy-signer
            DeprecatedPolicySigner *JSONWebKey `json:"policy_signer,omitempty"`
            // DeprecatedPolicyHash - DEPRECATED: Private Preview version of x-ms-policy-hash (a URL-encoded base64 string)
            DeprecatedPolicyHash *string `json:"policy_hash,omitempty"`
            // DeprecatedRpData - DEPRECATED: Private Preview version of nonce
            DeprecatedRpData *string `json:"rp_data,omitempty"`
            }

            // RuntimeData runtime data are a conduit for any information defined by the Trusted Execution Environment
            // (TEE) when actually running.
            type RuntimeData struct {
            // Data - Runtime data are generated by the Trusted Execution Environment (TEE). For an SGX quote (Coffeelake or Icelake), the SHA256 hash of the RuntimeData must match the lower 32 bytes of the quote's "report data" attribute. For a SEV-SNP quote, the SHA256 hash of the RuntimeData must match the quote's "report data" attribute. (a URL-encoded base64 string)
            Data *string `json:"data,omitempty"`
            // DataType - The type of data contained within the "data" field. Possible values include: 'DataTypeBinary', 'DataTypeJSON'
            DataType DataType `json:"dataType,omitempty"`
            }

            // StoredAttestationPolicy ...
            type StoredAttestationPolicy struct {
            // AttestationPolicy - Policy text to set as a sequence of UTF-8 encoded octets. (a URL-encoded base64 string)
            AttestationPolicy *string `json:"AttestationPolicy,omitempty"`
            }

            // TpmAttestationRequest attestation request for Trusted Platform Module (TPM) attestation.
            type TpmAttestationRequest struct {
            // Data - Protocol data containing artifacts for attestation. (a URL-encoded base64 string)
            Data *string `json:"data,omitempty"`
            }

            // TpmAttestationResponse attestation response for Trusted Platform Module (TPM) attestation.
            type TpmAttestationResponse struct {
            autorest.Response `json:"-"`
            // Data - Protocol data containing attestation service response. (a URL-encoded base64 string)
            Data *string `json:"data,omitempty"`
            }


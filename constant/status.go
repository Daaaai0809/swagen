package constant

const (
	CONTINUE_STATUS_CODE            = "100"
	SWITCHING_PROTOCOLS_STATUS_CODE = "101"
	PROCESSING_STATUS_CODE          = "102"
	EARLY_HINTS_STATUS_CODE         = "103"

	SUCESS_STATUS_CODE                        = "200"
	CREATED_STATUS_CODE                       = "201"
	ACCEPTED_STATUS_CODE                      = "202"
	NON_AUTHORITATIVE_INFORMATION_STATUS_CODE = "203"
	NO_CONTENT_STATUS_CODE                    = "204"
	RESET_CONTENT_STATUS_CODE                 = "205"
	PARTIAL_CONTENT_STATUS_CODE               = "206"
	MULTI_STATUS_STATUS_CODE                  = "207"
	ALREADY_REPORTED_STATUS_CODE              = "208"
	IM_USED_STATUS_CODE                       = "226"

	MULTIPLE_CHOICES_STATUS_CODE   = "300"
	MOVED_PERMANENTLY_STATUS_CODE  = "301"
	FOUND_STATUS_CODE              = "302"
	SEE_OTHER_STATUS_CODE          = "303"
	NOT_MODIFIED_STATUS_CODE       = "304"
	USE_PROXY_STATUS_CODE          = "305"
	TEMPORARY_REDIRECT_STATUS_CODE = "307"
	PERMANENT_REDIRECT_STATUS_CODE = "308"

	BAD_REQUEST_STATUS_CODE                     = "400"
	UNAUTHORIZED_STATUS_CODE                    = "401"
	PAYMENT_REQUIRED_STATUS_CODE                = "402"
	FORBIDDEN_STATUS_CODE                       = "403"
	NOT_FOUND_STATUS_CODE                       = "404"
	METHOD_NOT_ALLOWED_STATUS_CODE              = "405"
	NOT_ACCEPTABLE_STATUS_CODE                  = "406"
	PROXY_AUTHENTICATION_REQUIRED_STATUS_CODE   = "407"
	REQUEST_TIMEOUT_STATUS_CODE                 = "408"
	CONFLICT_STATUS_CODE                        = "409"
	GONE_STATUS_CODE                            = "410"
	LENGTH_REQUIRED_STATUS_CODE                 = "411"
	PRECONDITION_FAILED_STATUS_CODE             = "412"
	PAYLOAD_TOO_LARGE_STATUS_CODE               = "413"
	URI_TOO_LONG_STATUS_CODE                    = "414"
	UNSUPPORTED_MEDIA_TYPE_STATUS_CODE          = "415"
	RANGE_NOT_SATISFIABLE_STATUS_CODE           = "416"
	EXPECTATION_FAILED_STATUS_CODE              = "417"
	I_AM_A_TEAPOT_STATUS_CODE                   = "418"
	MISDIRECTED_REQUEST_STATUS_CODE             = "421"
	UNPROCESSABLE_ENTITY_STATUS_CODE            = "422"
	LOCKED_STATUS_CODE                          = "423"
	FAILED_DEPENDENCY_STATUS_CODE               = "424"
	UPGRADE_REQUIRED_STATUS_CODE                = "426"
	PRECONDITION_REQUIRED_STATUS_CODE           = "428"
	TOO_MANY_REQUESTS_STATUS_CODE               = "429"
	REQUEST_HEADER_FIELDS_TOO_LARGE_STATUS_CODE = "431"
	UNAVAILABLE_FOR_LEGAL_REASONS_STATUS_CODE   = "451"

	INTERNAL_SERVER_ERROR_STATUS_CODE           = "500"
	NOT_IMPLEMENTED_STATUS_CODE                 = "501"
	BAD_GATEWAY_STATUS_CODE                     = "502"
	SERVICE_UNAVAILABLE_STATUS_CODE             = "503"
	GATEWAY_TIMEOUT_STATUS_CODE                 = "504"
	HTTP_VERSION_NOT_SUPPORTED_STATUS_CODE      = "505"
	VARIANT_ALSO_NEGOTIATES_STATUS_CODE         = "506"
	INSUFFICIENT_STORAGE_STATUS_CODE            = "507"
	LOOP_DETECTED_STATUS_CODE                   = "508"
	NOT_EXTENDED_STATUS_CODE                    = "510"
	NETWORK_AUTHENTICATION_REQUIRED_STATUS_CODE = "511"

	DEFAULT_STATUS = "default"
)

var StatusCodesList = []string{
	CONTINUE_STATUS_CODE,
	SWITCHING_PROTOCOLS_STATUS_CODE,
	PROCESSING_STATUS_CODE,
	EARLY_HINTS_STATUS_CODE,
	SUCESS_STATUS_CODE,
	CREATED_STATUS_CODE,
	ACCEPTED_STATUS_CODE,
	NON_AUTHORITATIVE_INFORMATION_STATUS_CODE,
	NO_CONTENT_STATUS_CODE,
	RESET_CONTENT_STATUS_CODE,
	PARTIAL_CONTENT_STATUS_CODE,
	MULTI_STATUS_STATUS_CODE,
	ALREADY_REPORTED_STATUS_CODE,
	IM_USED_STATUS_CODE,
	MULTIPLE_CHOICES_STATUS_CODE,
	MOVED_PERMANENTLY_STATUS_CODE,
	FOUND_STATUS_CODE,
	SEE_OTHER_STATUS_CODE,
	NOT_MODIFIED_STATUS_CODE,
	USE_PROXY_STATUS_CODE,
	TEMPORARY_REDIRECT_STATUS_CODE,
	PERMANENT_REDIRECT_STATUS_CODE,
	BAD_REQUEST_STATUS_CODE,
	UNAUTHORIZED_STATUS_CODE,
	PAYMENT_REQUIRED_STATUS_CODE,
	FORBIDDEN_STATUS_CODE,
	NOT_FOUND_STATUS_CODE,
	METHOD_NOT_ALLOWED_STATUS_CODE,
	NOT_ACCEPTABLE_STATUS_CODE,
	PROXY_AUTHENTICATION_REQUIRED_STATUS_CODE,
	REQUEST_TIMEOUT_STATUS_CODE,
	CONFLICT_STATUS_CODE,
	GONE_STATUS_CODE,
	LENGTH_REQUIRED_STATUS_CODE,
	PRECONDITION_FAILED_STATUS_CODE,
	PAYLOAD_TOO_LARGE_STATUS_CODE,
	URI_TOO_LONG_STATUS_CODE,
	UNSUPPORTED_MEDIA_TYPE_STATUS_CODE,
	RANGE_NOT_SATISFIABLE_STATUS_CODE,
	EXPECTATION_FAILED_STATUS_CODE,
	I_AM_A_TEAPOT_STATUS_CODE,
	MISDIRECTED_REQUEST_STATUS_CODE,
	UNPROCESSABLE_ENTITY_STATUS_CODE,
	LOCKED_STATUS_CODE,
	FAILED_DEPENDENCY_STATUS_CODE,
	UPGRADE_REQUIRED_STATUS_CODE,
	PRECONDITION_REQUIRED_STATUS_CODE,
	TOO_MANY_REQUESTS_STATUS_CODE,
	REQUEST_HEADER_FIELDS_TOO_LARGE_STATUS_CODE,
	UNAVAILABLE_FOR_LEGAL_REASONS_STATUS_CODE,
	INTERNAL_SERVER_ERROR_STATUS_CODE,
	NOT_IMPLEMENTED_STATUS_CODE,
	BAD_GATEWAY_STATUS_CODE,
	SERVICE_UNAVAILABLE_STATUS_CODE,
	GATEWAY_TIMEOUT_STATUS_CODE,
	HTTP_VERSION_NOT_SUPPORTED_STATUS_CODE,
	VARIANT_ALSO_NEGOTIATES_STATUS_CODE,
	INSUFFICIENT_STORAGE_STATUS_CODE,
	LOOP_DETECTED_STATUS_CODE,
	NOT_EXTENDED_STATUS_CODE,
	NETWORK_AUTHENTICATION_REQUIRED_STATUS_CODE,
	DEFAULT_STATUS,
}
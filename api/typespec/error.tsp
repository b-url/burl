import "@typespec/http";

using TypeSpec.Http;

enum ErrorType {
    API_ERROR: "api_error",
    INVALID_REQUEST_ERROR : "invalid_request_error",
}

enum ErrorCode {
    PARAMETER_INVALID: "parameter_invalid",
    PARAMETER_MISSING: "parameter_missing",
    PROCESSING_ERROR: "processing_error",
    RESOURCE_ALREADY_EXISTS: "resource_already_exists",
    RESOURCE_MISSING: "resource_missing",
}

@doc("Error is the response model when an API call is unsuccessful.")
@error
model Error {
    type: ErrorType;
    code: ErrorCode;
    message: string;
    details: Record<unknown>;
}
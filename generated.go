// Package pf CODE GENERATED. DO NOT MODIFY
package pf

import "fmt"

// ContinueSlug ..
const ContinueSlug = "continue"

// Continue error with message
func Continue(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ContinueSlug,
		Code:    100,
		Message: message,
	}
}

// Continuef error with message
func Continuef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ContinueSlug,
		Code:    100,
		Message: message,
	}
}

// ContinueDetail error with detail lines
func ContinueDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ContinueSlug,
		Code:   100,
		Detail: detail,
	}
}

// SwitchingProtocolsSlug ..
const SwitchingProtocolsSlug = "switching_protocols"

// SwitchingProtocols error with message
func SwitchingProtocols(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    SwitchingProtocolsSlug,
		Code:    101,
		Message: message,
	}
}

// SwitchingProtocolsf error with message
func SwitchingProtocolsf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    SwitchingProtocolsSlug,
		Code:    101,
		Message: message,
	}
}

// SwitchingProtocolsDetail error with detail lines
func SwitchingProtocolsDetail(detail []string) *Fail {
	return &Fail{
		Slug:   SwitchingProtocolsSlug,
		Code:   101,
		Detail: detail,
	}
}

// ProcessingSlug ..
const ProcessingSlug = "processing"

// Processing error with message
func Processing(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ProcessingSlug,
		Code:    102,
		Message: message,
	}
}

// Processingf error with message
func Processingf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ProcessingSlug,
		Code:    102,
		Message: message,
	}
}

// ProcessingDetail error with detail lines
func ProcessingDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ProcessingSlug,
		Code:   102,
		Detail: detail,
	}
}

// OkSlug ..
const OkSlug = "ok"

// Ok error with message
func Ok(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    OkSlug,
		Code:    200,
		Message: message,
	}
}

// Okf error with message
func Okf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    OkSlug,
		Code:    200,
		Message: message,
	}
}

// OkDetail error with detail lines
func OkDetail(detail []string) *Fail {
	return &Fail{
		Slug:   OkSlug,
		Code:   200,
		Detail: detail,
	}
}

// CreatedSlug ..
const CreatedSlug = "created"

// Created error with message
func Created(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    CreatedSlug,
		Code:    201,
		Message: message,
	}
}

// Createdf error with message
func Createdf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    CreatedSlug,
		Code:    201,
		Message: message,
	}
}

// CreatedDetail error with detail lines
func CreatedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   CreatedSlug,
		Code:   201,
		Detail: detail,
	}
}

// AcceptedSlug ..
const AcceptedSlug = "accepted"

// Accepted error with message
func Accepted(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    AcceptedSlug,
		Code:    202,
		Message: message,
	}
}

// Acceptedf error with message
func Acceptedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    AcceptedSlug,
		Code:    202,
		Message: message,
	}
}

// AcceptedDetail error with detail lines
func AcceptedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   AcceptedSlug,
		Code:   202,
		Detail: detail,
	}
}

// NonAuthoritativeInformationSlug ..
const NonAuthoritativeInformationSlug = "non_authoritative_information"

// NonAuthoritativeInformation error with message
func NonAuthoritativeInformation(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NonAuthoritativeInformationSlug,
		Code:    203,
		Message: message,
	}
}

// NonAuthoritativeInformationf error with message
func NonAuthoritativeInformationf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NonAuthoritativeInformationSlug,
		Code:    203,
		Message: message,
	}
}

// NonAuthoritativeInformationDetail error with detail lines
func NonAuthoritativeInformationDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NonAuthoritativeInformationSlug,
		Code:   203,
		Detail: detail,
	}
}

// NoContentSlug ..
const NoContentSlug = "no_content"

// NoContent error with message
func NoContent(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NoContentSlug,
		Code:    204,
		Message: message,
	}
}

// NoContentf error with message
func NoContentf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NoContentSlug,
		Code:    204,
		Message: message,
	}
}

// NoContentDetail error with detail lines
func NoContentDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NoContentSlug,
		Code:   204,
		Detail: detail,
	}
}

// ResetContentSlug ..
const ResetContentSlug = "reset_content"

// ResetContent error with message
func ResetContent(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ResetContentSlug,
		Code:    205,
		Message: message,
	}
}

// ResetContentf error with message
func ResetContentf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ResetContentSlug,
		Code:    205,
		Message: message,
	}
}

// ResetContentDetail error with detail lines
func ResetContentDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ResetContentSlug,
		Code:   205,
		Detail: detail,
	}
}

// PartialContentSlug ..
const PartialContentSlug = "partial_content"

// PartialContent error with message
func PartialContent(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PartialContentSlug,
		Code:    206,
		Message: message,
	}
}

// PartialContentf error with message
func PartialContentf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PartialContentSlug,
		Code:    206,
		Message: message,
	}
}

// PartialContentDetail error with detail lines
func PartialContentDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PartialContentSlug,
		Code:   206,
		Detail: detail,
	}
}

// MultiStatusSlug ..
const MultiStatusSlug = "multi_status"

// MultiStatus error with message
func MultiStatus(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    MultiStatusSlug,
		Code:    207,
		Message: message,
	}
}

// MultiStatusf error with message
func MultiStatusf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    MultiStatusSlug,
		Code:    207,
		Message: message,
	}
}

// MultiStatusDetail error with detail lines
func MultiStatusDetail(detail []string) *Fail {
	return &Fail{
		Slug:   MultiStatusSlug,
		Code:   207,
		Detail: detail,
	}
}

// AlreadyReportedSlug ..
const AlreadyReportedSlug = "already_reported"

// AlreadyReported error with message
func AlreadyReported(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    AlreadyReportedSlug,
		Code:    208,
		Message: message,
	}
}

// AlreadyReportedf error with message
func AlreadyReportedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    AlreadyReportedSlug,
		Code:    208,
		Message: message,
	}
}

// AlreadyReportedDetail error with detail lines
func AlreadyReportedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   AlreadyReportedSlug,
		Code:   208,
		Detail: detail,
	}
}

// ImUsedSlug ..
const ImUsedSlug = "im_used"

// ImUsed error with message
func ImUsed(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ImUsedSlug,
		Code:    226,
		Message: message,
	}
}

// ImUsedf error with message
func ImUsedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ImUsedSlug,
		Code:    226,
		Message: message,
	}
}

// ImUsedDetail error with detail lines
func ImUsedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ImUsedSlug,
		Code:   226,
		Detail: detail,
	}
}

// MultipleChoicesSlug ..
const MultipleChoicesSlug = "multiple_choices"

// MultipleChoices error with message
func MultipleChoices(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    MultipleChoicesSlug,
		Code:    300,
		Message: message,
	}
}

// MultipleChoicesf error with message
func MultipleChoicesf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    MultipleChoicesSlug,
		Code:    300,
		Message: message,
	}
}

// MultipleChoicesDetail error with detail lines
func MultipleChoicesDetail(detail []string) *Fail {
	return &Fail{
		Slug:   MultipleChoicesSlug,
		Code:   300,
		Detail: detail,
	}
}

// MovedPermanentlySlug ..
const MovedPermanentlySlug = "moved_permanently"

// MovedPermanently error with message
func MovedPermanently(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    MovedPermanentlySlug,
		Code:    301,
		Message: message,
	}
}

// MovedPermanentlyf error with message
func MovedPermanentlyf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    MovedPermanentlySlug,
		Code:    301,
		Message: message,
	}
}

// MovedPermanentlyDetail error with detail lines
func MovedPermanentlyDetail(detail []string) *Fail {
	return &Fail{
		Slug:   MovedPermanentlySlug,
		Code:   301,
		Detail: detail,
	}
}

// FoundSlug ..
const FoundSlug = "found"

// Found error with message
func Found(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    FoundSlug,
		Code:    302,
		Message: message,
	}
}

// Foundf error with message
func Foundf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    FoundSlug,
		Code:    302,
		Message: message,
	}
}

// FoundDetail error with detail lines
func FoundDetail(detail []string) *Fail {
	return &Fail{
		Slug:   FoundSlug,
		Code:   302,
		Detail: detail,
	}
}

// SeeOtherSlug ..
const SeeOtherSlug = "see_other"

// SeeOther error with message
func SeeOther(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    SeeOtherSlug,
		Code:    303,
		Message: message,
	}
}

// SeeOtherf error with message
func SeeOtherf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    SeeOtherSlug,
		Code:    303,
		Message: message,
	}
}

// SeeOtherDetail error with detail lines
func SeeOtherDetail(detail []string) *Fail {
	return &Fail{
		Slug:   SeeOtherSlug,
		Code:   303,
		Detail: detail,
	}
}

// NotModifiedSlug ..
const NotModifiedSlug = "not_modified"

// NotModified error with message
func NotModified(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NotModifiedSlug,
		Code:    304,
		Message: message,
	}
}

// NotModifiedf error with message
func NotModifiedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NotModifiedSlug,
		Code:    304,
		Message: message,
	}
}

// NotModifiedDetail error with detail lines
func NotModifiedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NotModifiedSlug,
		Code:   304,
		Detail: detail,
	}
}

// UseProxySlug ..
const UseProxySlug = "use_proxy"

// UseProxy error with message
func UseProxy(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UseProxySlug,
		Code:    305,
		Message: message,
	}
}

// UseProxyf error with message
func UseProxyf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UseProxySlug,
		Code:    305,
		Message: message,
	}
}

// UseProxyDetail error with detail lines
func UseProxyDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UseProxySlug,
		Code:   305,
		Detail: detail,
	}
}

// TemporaryRedirectSlug ..
const TemporaryRedirectSlug = "temporary_redirect"

// TemporaryRedirect error with message
func TemporaryRedirect(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    TemporaryRedirectSlug,
		Code:    307,
		Message: message,
	}
}

// TemporaryRedirectf error with message
func TemporaryRedirectf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    TemporaryRedirectSlug,
		Code:    307,
		Message: message,
	}
}

// TemporaryRedirectDetail error with detail lines
func TemporaryRedirectDetail(detail []string) *Fail {
	return &Fail{
		Slug:   TemporaryRedirectSlug,
		Code:   307,
		Detail: detail,
	}
}

// PermanentRedirectSlug ..
const PermanentRedirectSlug = "permanent_redirect"

// PermanentRedirect error with message
func PermanentRedirect(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PermanentRedirectSlug,
		Code:    308,
		Message: message,
	}
}

// PermanentRedirectf error with message
func PermanentRedirectf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PermanentRedirectSlug,
		Code:    308,
		Message: message,
	}
}

// PermanentRedirectDetail error with detail lines
func PermanentRedirectDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PermanentRedirectSlug,
		Code:   308,
		Detail: detail,
	}
}

// BadRequestSlug ..
const BadRequestSlug = "bad_request"

// BadRequest error with message
func BadRequest(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    BadRequestSlug,
		Code:    400,
		Message: message,
	}
}

// BadRequestf error with message
func BadRequestf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    BadRequestSlug,
		Code:    400,
		Message: message,
	}
}

// BadRequestDetail error with detail lines
func BadRequestDetail(detail []string) *Fail {
	return &Fail{
		Slug:   BadRequestSlug,
		Code:   400,
		Detail: detail,
	}
}

// UnauthorizedSlug ..
const UnauthorizedSlug = "unauthorized"

// Unauthorized error with message
func Unauthorized(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UnauthorizedSlug,
		Code:    401,
		Message: message,
	}
}

// Unauthorizedf error with message
func Unauthorizedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UnauthorizedSlug,
		Code:    401,
		Message: message,
	}
}

// UnauthorizedDetail error with detail lines
func UnauthorizedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UnauthorizedSlug,
		Code:   401,
		Detail: detail,
	}
}

// PaymentRequiredSlug ..
const PaymentRequiredSlug = "payment_required"

// PaymentRequired error with message
func PaymentRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PaymentRequiredSlug,
		Code:    402,
		Message: message,
	}
}

// PaymentRequiredf error with message
func PaymentRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PaymentRequiredSlug,
		Code:    402,
		Message: message,
	}
}

// PaymentRequiredDetail error with detail lines
func PaymentRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PaymentRequiredSlug,
		Code:   402,
		Detail: detail,
	}
}

// ForbiddenSlug ..
const ForbiddenSlug = "forbidden"

// Forbidden error with message
func Forbidden(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ForbiddenSlug,
		Code:    403,
		Message: message,
	}
}

// Forbiddenf error with message
func Forbiddenf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ForbiddenSlug,
		Code:    403,
		Message: message,
	}
}

// ForbiddenDetail error with detail lines
func ForbiddenDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ForbiddenSlug,
		Code:   403,
		Detail: detail,
	}
}

// NotFoundSlug ..
const NotFoundSlug = "not_found"

// NotFound error with message
func NotFound(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NotFoundSlug,
		Code:    404,
		Message: message,
	}
}

// NotFoundf error with message
func NotFoundf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NotFoundSlug,
		Code:    404,
		Message: message,
	}
}

// NotFoundDetail error with detail lines
func NotFoundDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NotFoundSlug,
		Code:   404,
		Detail: detail,
	}
}

// MethodNotAllowedSlug ..
const MethodNotAllowedSlug = "method_not_allowed"

// MethodNotAllowed error with message
func MethodNotAllowed(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    MethodNotAllowedSlug,
		Code:    405,
		Message: message,
	}
}

// MethodNotAllowedf error with message
func MethodNotAllowedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    MethodNotAllowedSlug,
		Code:    405,
		Message: message,
	}
}

// MethodNotAllowedDetail error with detail lines
func MethodNotAllowedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   MethodNotAllowedSlug,
		Code:   405,
		Detail: detail,
	}
}

// NotAcceptableSlug ..
const NotAcceptableSlug = "not_acceptable"

// NotAcceptable error with message
func NotAcceptable(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NotAcceptableSlug,
		Code:    406,
		Message: message,
	}
}

// NotAcceptablef error with message
func NotAcceptablef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NotAcceptableSlug,
		Code:    406,
		Message: message,
	}
}

// NotAcceptableDetail error with detail lines
func NotAcceptableDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NotAcceptableSlug,
		Code:   406,
		Detail: detail,
	}
}

// ProxyAuthenticationRequiredSlug ..
const ProxyAuthenticationRequiredSlug = "proxy_authentication_required"

// ProxyAuthenticationRequired error with message
func ProxyAuthenticationRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ProxyAuthenticationRequiredSlug,
		Code:    407,
		Message: message,
	}
}

// ProxyAuthenticationRequiredf error with message
func ProxyAuthenticationRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ProxyAuthenticationRequiredSlug,
		Code:    407,
		Message: message,
	}
}

// ProxyAuthenticationRequiredDetail error with detail lines
func ProxyAuthenticationRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ProxyAuthenticationRequiredSlug,
		Code:   407,
		Detail: detail,
	}
}

// RequestTimeoutSlug ..
const RequestTimeoutSlug = "request_timeout"

// RequestTimeout error with message
func RequestTimeout(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    RequestTimeoutSlug,
		Code:    408,
		Message: message,
	}
}

// RequestTimeoutf error with message
func RequestTimeoutf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    RequestTimeoutSlug,
		Code:    408,
		Message: message,
	}
}

// RequestTimeoutDetail error with detail lines
func RequestTimeoutDetail(detail []string) *Fail {
	return &Fail{
		Slug:   RequestTimeoutSlug,
		Code:   408,
		Detail: detail,
	}
}

// ConflictSlug ..
const ConflictSlug = "conflict"

// Conflict error with message
func Conflict(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ConflictSlug,
		Code:    409,
		Message: message,
	}
}

// Conflictf error with message
func Conflictf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ConflictSlug,
		Code:    409,
		Message: message,
	}
}

// ConflictDetail error with detail lines
func ConflictDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ConflictSlug,
		Code:   409,
		Detail: detail,
	}
}

// GoneSlug ..
const GoneSlug = "gone"

// Gone error with message
func Gone(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    GoneSlug,
		Code:    410,
		Message: message,
	}
}

// Gonef error with message
func Gonef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    GoneSlug,
		Code:    410,
		Message: message,
	}
}

// GoneDetail error with detail lines
func GoneDetail(detail []string) *Fail {
	return &Fail{
		Slug:   GoneSlug,
		Code:   410,
		Detail: detail,
	}
}

// LengthRequiredSlug ..
const LengthRequiredSlug = "length_required"

// LengthRequired error with message
func LengthRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    LengthRequiredSlug,
		Code:    411,
		Message: message,
	}
}

// LengthRequiredf error with message
func LengthRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    LengthRequiredSlug,
		Code:    411,
		Message: message,
	}
}

// LengthRequiredDetail error with detail lines
func LengthRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   LengthRequiredSlug,
		Code:   411,
		Detail: detail,
	}
}

// PreconditionFailedSlug ..
const PreconditionFailedSlug = "precondition_failed"

// PreconditionFailed error with message
func PreconditionFailed(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PreconditionFailedSlug,
		Code:    412,
		Message: message,
	}
}

// PreconditionFailedf error with message
func PreconditionFailedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PreconditionFailedSlug,
		Code:    412,
		Message: message,
	}
}

// PreconditionFailedDetail error with detail lines
func PreconditionFailedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PreconditionFailedSlug,
		Code:   412,
		Detail: detail,
	}
}

// PayloadTooLargeSlug ..
const PayloadTooLargeSlug = "payload_too_large"

// PayloadTooLarge error with message
func PayloadTooLarge(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PayloadTooLargeSlug,
		Code:    413,
		Message: message,
	}
}

// PayloadTooLargef error with message
func PayloadTooLargef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PayloadTooLargeSlug,
		Code:    413,
		Message: message,
	}
}

// PayloadTooLargeDetail error with detail lines
func PayloadTooLargeDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PayloadTooLargeSlug,
		Code:   413,
		Detail: detail,
	}
}

// RequestURITooLongSlug ..
const RequestURITooLongSlug = "request_uri_too_long"

// RequestURITooLong error with message
func RequestURITooLong(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    RequestURITooLongSlug,
		Code:    414,
		Message: message,
	}
}

// RequestURITooLongf error with message
func RequestURITooLongf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    RequestURITooLongSlug,
		Code:    414,
		Message: message,
	}
}

// RequestURITooLongDetail error with detail lines
func RequestURITooLongDetail(detail []string) *Fail {
	return &Fail{
		Slug:   RequestURITooLongSlug,
		Code:   414,
		Detail: detail,
	}
}

// UnsupportedMediaTypeSlug ..
const UnsupportedMediaTypeSlug = "unsupported_media_type"

// UnsupportedMediaType error with message
func UnsupportedMediaType(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UnsupportedMediaTypeSlug,
		Code:    415,
		Message: message,
	}
}

// UnsupportedMediaTypef error with message
func UnsupportedMediaTypef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UnsupportedMediaTypeSlug,
		Code:    415,
		Message: message,
	}
}

// UnsupportedMediaTypeDetail error with detail lines
func UnsupportedMediaTypeDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UnsupportedMediaTypeSlug,
		Code:   415,
		Detail: detail,
	}
}

// RequestedRangeNotSatisfiableSlug ..
const RequestedRangeNotSatisfiableSlug = "requested_range_not_satisfiable"

// RequestedRangeNotSatisfiable error with message
func RequestedRangeNotSatisfiable(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    RequestedRangeNotSatisfiableSlug,
		Code:    416,
		Message: message,
	}
}

// RequestedRangeNotSatisfiablef error with message
func RequestedRangeNotSatisfiablef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    RequestedRangeNotSatisfiableSlug,
		Code:    416,
		Message: message,
	}
}

// RequestedRangeNotSatisfiableDetail error with detail lines
func RequestedRangeNotSatisfiableDetail(detail []string) *Fail {
	return &Fail{
		Slug:   RequestedRangeNotSatisfiableSlug,
		Code:   416,
		Detail: detail,
	}
}

// ExpectationFailedSlug ..
const ExpectationFailedSlug = "expectation_failed"

// ExpectationFailed error with message
func ExpectationFailed(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ExpectationFailedSlug,
		Code:    417,
		Message: message,
	}
}

// ExpectationFailedf error with message
func ExpectationFailedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ExpectationFailedSlug,
		Code:    417,
		Message: message,
	}
}

// ExpectationFailedDetail error with detail lines
func ExpectationFailedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ExpectationFailedSlug,
		Code:   417,
		Detail: detail,
	}
}

// IMATeapotSlug ..
const IMATeapotSlug = "i_m_a_teapot"

// IMATeapot error with message
func IMATeapot(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    IMATeapotSlug,
		Code:    418,
		Message: message,
	}
}

// IMATeapotf error with message
func IMATeapotf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    IMATeapotSlug,
		Code:    418,
		Message: message,
	}
}

// IMATeapotDetail error with detail lines
func IMATeapotDetail(detail []string) *Fail {
	return &Fail{
		Slug:   IMATeapotSlug,
		Code:   418,
		Detail: detail,
	}
}

// MisdirectedRequestSlug ..
const MisdirectedRequestSlug = "misdirected_request"

// MisdirectedRequest error with message
func MisdirectedRequest(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    MisdirectedRequestSlug,
		Code:    421,
		Message: message,
	}
}

// MisdirectedRequestf error with message
func MisdirectedRequestf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    MisdirectedRequestSlug,
		Code:    421,
		Message: message,
	}
}

// MisdirectedRequestDetail error with detail lines
func MisdirectedRequestDetail(detail []string) *Fail {
	return &Fail{
		Slug:   MisdirectedRequestSlug,
		Code:   421,
		Detail: detail,
	}
}

// UnprocessableEntitySlug ..
const UnprocessableEntitySlug = "unprocessable_entity"

// UnprocessableEntity error with message
func UnprocessableEntity(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UnprocessableEntitySlug,
		Code:    422,
		Message: message,
	}
}

// UnprocessableEntityf error with message
func UnprocessableEntityf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UnprocessableEntitySlug,
		Code:    422,
		Message: message,
	}
}

// UnprocessableEntityDetail error with detail lines
func UnprocessableEntityDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UnprocessableEntitySlug,
		Code:   422,
		Detail: detail,
	}
}

// LockedSlug ..
const LockedSlug = "locked"

// Locked error with message
func Locked(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    LockedSlug,
		Code:    423,
		Message: message,
	}
}

// Lockedf error with message
func Lockedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    LockedSlug,
		Code:    423,
		Message: message,
	}
}

// LockedDetail error with detail lines
func LockedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   LockedSlug,
		Code:   423,
		Detail: detail,
	}
}

// FailedDependencySlug ..
const FailedDependencySlug = "failed_dependency"

// FailedDependency error with message
func FailedDependency(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    FailedDependencySlug,
		Code:    424,
		Message: message,
	}
}

// FailedDependencyf error with message
func FailedDependencyf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    FailedDependencySlug,
		Code:    424,
		Message: message,
	}
}

// FailedDependencyDetail error with detail lines
func FailedDependencyDetail(detail []string) *Fail {
	return &Fail{
		Slug:   FailedDependencySlug,
		Code:   424,
		Detail: detail,
	}
}

// UpgradeRequiredSlug ..
const UpgradeRequiredSlug = "upgrade_required"

// UpgradeRequired error with message
func UpgradeRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UpgradeRequiredSlug,
		Code:    426,
		Message: message,
	}
}

// UpgradeRequiredf error with message
func UpgradeRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UpgradeRequiredSlug,
		Code:    426,
		Message: message,
	}
}

// UpgradeRequiredDetail error with detail lines
func UpgradeRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UpgradeRequiredSlug,
		Code:   426,
		Detail: detail,
	}
}

// PreconditionRequiredSlug ..
const PreconditionRequiredSlug = "precondition_required"

// PreconditionRequired error with message
func PreconditionRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    PreconditionRequiredSlug,
		Code:    428,
		Message: message,
	}
}

// PreconditionRequiredf error with message
func PreconditionRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    PreconditionRequiredSlug,
		Code:    428,
		Message: message,
	}
}

// PreconditionRequiredDetail error with detail lines
func PreconditionRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   PreconditionRequiredSlug,
		Code:   428,
		Detail: detail,
	}
}

// TooManyRequestsSlug ..
const TooManyRequestsSlug = "too_many_requests"

// TooManyRequests error with message
func TooManyRequests(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    TooManyRequestsSlug,
		Code:    429,
		Message: message,
	}
}

// TooManyRequestsf error with message
func TooManyRequestsf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    TooManyRequestsSlug,
		Code:    429,
		Message: message,
	}
}

// TooManyRequestsDetail error with detail lines
func TooManyRequestsDetail(detail []string) *Fail {
	return &Fail{
		Slug:   TooManyRequestsSlug,
		Code:   429,
		Detail: detail,
	}
}

// RequestHeaderFieldsTooLargeSlug ..
const RequestHeaderFieldsTooLargeSlug = "request_header_fields_too_large"

// RequestHeaderFieldsTooLarge error with message
func RequestHeaderFieldsTooLarge(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    RequestHeaderFieldsTooLargeSlug,
		Code:    431,
		Message: message,
	}
}

// RequestHeaderFieldsTooLargef error with message
func RequestHeaderFieldsTooLargef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    RequestHeaderFieldsTooLargeSlug,
		Code:    431,
		Message: message,
	}
}

// RequestHeaderFieldsTooLargeDetail error with detail lines
func RequestHeaderFieldsTooLargeDetail(detail []string) *Fail {
	return &Fail{
		Slug:   RequestHeaderFieldsTooLargeSlug,
		Code:   431,
		Detail: detail,
	}
}

// ConnectionClosedWithoutResponseSlug ..
const ConnectionClosedWithoutResponseSlug = "connection_closed_without_response"

// ConnectionClosedWithoutResponse error with message
func ConnectionClosedWithoutResponse(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ConnectionClosedWithoutResponseSlug,
		Code:    444,
		Message: message,
	}
}

// ConnectionClosedWithoutResponsef error with message
func ConnectionClosedWithoutResponsef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ConnectionClosedWithoutResponseSlug,
		Code:    444,
		Message: message,
	}
}

// ConnectionClosedWithoutResponseDetail error with detail lines
func ConnectionClosedWithoutResponseDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ConnectionClosedWithoutResponseSlug,
		Code:   444,
		Detail: detail,
	}
}

// UnavailableForLegalReasonsSlug ..
const UnavailableForLegalReasonsSlug = "unavailable_for_legal_reasons"

// UnavailableForLegalReasons error with message
func UnavailableForLegalReasons(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    UnavailableForLegalReasonsSlug,
		Code:    451,
		Message: message,
	}
}

// UnavailableForLegalReasonsf error with message
func UnavailableForLegalReasonsf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    UnavailableForLegalReasonsSlug,
		Code:    451,
		Message: message,
	}
}

// UnavailableForLegalReasonsDetail error with detail lines
func UnavailableForLegalReasonsDetail(detail []string) *Fail {
	return &Fail{
		Slug:   UnavailableForLegalReasonsSlug,
		Code:   451,
		Detail: detail,
	}
}

// ClientClosedRequestSlug ..
const ClientClosedRequestSlug = "client_closed_request"

// ClientClosedRequest error with message
func ClientClosedRequest(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ClientClosedRequestSlug,
		Code:    499,
		Message: message,
	}
}

// ClientClosedRequestf error with message
func ClientClosedRequestf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ClientClosedRequestSlug,
		Code:    499,
		Message: message,
	}
}

// ClientClosedRequestDetail error with detail lines
func ClientClosedRequestDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ClientClosedRequestSlug,
		Code:   499,
		Detail: detail,
	}
}

// InternalServerErrorSlug ..
const InternalServerErrorSlug = "internal_server_error"

// InternalServerError error with message
func InternalServerError(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    InternalServerErrorSlug,
		Code:    500,
		Message: message,
	}
}

// InternalServerErrorf error with message
func InternalServerErrorf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    InternalServerErrorSlug,
		Code:    500,
		Message: message,
	}
}

// InternalServerErrorDetail error with detail lines
func InternalServerErrorDetail(detail []string) *Fail {
	return &Fail{
		Slug:   InternalServerErrorSlug,
		Code:   500,
		Detail: detail,
	}
}

// NotImplementedSlug ..
const NotImplementedSlug = "not_implemented"

// NotImplemented error with message
func NotImplemented(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NotImplementedSlug,
		Code:    501,
		Message: message,
	}
}

// NotImplementedf error with message
func NotImplementedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NotImplementedSlug,
		Code:    501,
		Message: message,
	}
}

// NotImplementedDetail error with detail lines
func NotImplementedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NotImplementedSlug,
		Code:   501,
		Detail: detail,
	}
}

// BadGatewaySlug ..
const BadGatewaySlug = "bad_gateway"

// BadGateway error with message
func BadGateway(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    BadGatewaySlug,
		Code:    502,
		Message: message,
	}
}

// BadGatewayf error with message
func BadGatewayf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    BadGatewaySlug,
		Code:    502,
		Message: message,
	}
}

// BadGatewayDetail error with detail lines
func BadGatewayDetail(detail []string) *Fail {
	return &Fail{
		Slug:   BadGatewaySlug,
		Code:   502,
		Detail: detail,
	}
}

// ServiceUnavailableSlug ..
const ServiceUnavailableSlug = "service_unavailable"

// ServiceUnavailable error with message
func ServiceUnavailable(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    ServiceUnavailableSlug,
		Code:    503,
		Message: message,
	}
}

// ServiceUnavailablef error with message
func ServiceUnavailablef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    ServiceUnavailableSlug,
		Code:    503,
		Message: message,
	}
}

// ServiceUnavailableDetail error with detail lines
func ServiceUnavailableDetail(detail []string) *Fail {
	return &Fail{
		Slug:   ServiceUnavailableSlug,
		Code:   503,
		Detail: detail,
	}
}

// GatewayTimeoutSlug ..
const GatewayTimeoutSlug = "gateway_timeout"

// GatewayTimeout error with message
func GatewayTimeout(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    GatewayTimeoutSlug,
		Code:    504,
		Message: message,
	}
}

// GatewayTimeoutf error with message
func GatewayTimeoutf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    GatewayTimeoutSlug,
		Code:    504,
		Message: message,
	}
}

// GatewayTimeoutDetail error with detail lines
func GatewayTimeoutDetail(detail []string) *Fail {
	return &Fail{
		Slug:   GatewayTimeoutSlug,
		Code:   504,
		Detail: detail,
	}
}

// HTTPVersionNotSupportedSlug ..
const HTTPVersionNotSupportedSlug = "http_version_not_supported"

// HTTPVersionNotSupported error with message
func HTTPVersionNotSupported(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    HTTPVersionNotSupportedSlug,
		Code:    505,
		Message: message,
	}
}

// HTTPVersionNotSupportedf error with message
func HTTPVersionNotSupportedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    HTTPVersionNotSupportedSlug,
		Code:    505,
		Message: message,
	}
}

// HTTPVersionNotSupportedDetail error with detail lines
func HTTPVersionNotSupportedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   HTTPVersionNotSupportedSlug,
		Code:   505,
		Detail: detail,
	}
}

// VariantAlsoNegotiatesSlug ..
const VariantAlsoNegotiatesSlug = "variant_also_negotiates"

// VariantAlsoNegotiates error with message
func VariantAlsoNegotiates(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    VariantAlsoNegotiatesSlug,
		Code:    506,
		Message: message,
	}
}

// VariantAlsoNegotiatesf error with message
func VariantAlsoNegotiatesf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    VariantAlsoNegotiatesSlug,
		Code:    506,
		Message: message,
	}
}

// VariantAlsoNegotiatesDetail error with detail lines
func VariantAlsoNegotiatesDetail(detail []string) *Fail {
	return &Fail{
		Slug:   VariantAlsoNegotiatesSlug,
		Code:   506,
		Detail: detail,
	}
}

// InsufficientStorageSlug ..
const InsufficientStorageSlug = "insufficient_storage"

// InsufficientStorage error with message
func InsufficientStorage(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    InsufficientStorageSlug,
		Code:    507,
		Message: message,
	}
}

// InsufficientStoragef error with message
func InsufficientStoragef(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    InsufficientStorageSlug,
		Code:    507,
		Message: message,
	}
}

// InsufficientStorageDetail error with detail lines
func InsufficientStorageDetail(detail []string) *Fail {
	return &Fail{
		Slug:   InsufficientStorageSlug,
		Code:   507,
		Detail: detail,
	}
}

// LoopDetectedSlug ..
const LoopDetectedSlug = "loop_detected"

// LoopDetected error with message
func LoopDetected(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    LoopDetectedSlug,
		Code:    508,
		Message: message,
	}
}

// LoopDetectedf error with message
func LoopDetectedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    LoopDetectedSlug,
		Code:    508,
		Message: message,
	}
}

// LoopDetectedDetail error with detail lines
func LoopDetectedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   LoopDetectedSlug,
		Code:   508,
		Detail: detail,
	}
}

// NotExtendedSlug ..
const NotExtendedSlug = "not_extended"

// NotExtended error with message
func NotExtended(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NotExtendedSlug,
		Code:    510,
		Message: message,
	}
}

// NotExtendedf error with message
func NotExtendedf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NotExtendedSlug,
		Code:    510,
		Message: message,
	}
}

// NotExtendedDetail error with detail lines
func NotExtendedDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NotExtendedSlug,
		Code:   510,
		Detail: detail,
	}
}

// NetworkAuthenticationRequiredSlug ..
const NetworkAuthenticationRequiredSlug = "network_authentication_required"

// NetworkAuthenticationRequired error with message
func NetworkAuthenticationRequired(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NetworkAuthenticationRequiredSlug,
		Code:    511,
		Message: message,
	}
}

// NetworkAuthenticationRequiredf error with message
func NetworkAuthenticationRequiredf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NetworkAuthenticationRequiredSlug,
		Code:    511,
		Message: message,
	}
}

// NetworkAuthenticationRequiredDetail error with detail lines
func NetworkAuthenticationRequiredDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NetworkAuthenticationRequiredSlug,
		Code:   511,
		Detail: detail,
	}
}

// NetworkConnectTimeoutErrorSlug ..
const NetworkConnectTimeoutErrorSlug = "network_connect_timeout_error"

// NetworkConnectTimeoutError error with message
func NetworkConnectTimeoutError(msg ...interface{}) *Fail {
	message := formatMessage(msg...)
	return &Fail{
		Slug:    NetworkConnectTimeoutErrorSlug,
		Code:    599,
		Message: message,
	}
}

// NetworkConnectTimeoutErrorf error with message
func NetworkConnectTimeoutErrorf(format string, args ...interface{}) *Fail {
	message := fmt.Sprintf(format, args...)
	return &Fail{
		Slug:    NetworkConnectTimeoutErrorSlug,
		Code:    599,
		Message: message,
	}
}

// NetworkConnectTimeoutErrorDetail error with detail lines
func NetworkConnectTimeoutErrorDetail(detail []string) *Fail {
	return &Fail{
		Slug:   NetworkConnectTimeoutErrorSlug,
		Code:   599,
		Detail: detail,
	}
}

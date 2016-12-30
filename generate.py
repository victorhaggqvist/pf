
# generated with
# var links = document.querySelectorAll("li a"); var out = "";for (var i =0; i < links.length; i++) {out += links[i].innerText+"\n";} console.log(out);
# on https://httpstatuses.com/

codes = """100 Continue
101 Switching Protocols
102 Processing
200 OK
201 Created
202 Accepted
203 Non-authoritative Information
204 No Content
205 Reset Content
206 Partial Content
207 Multi-Status
208 Already Reported
226 IM Used
300 Multiple Choices
301 Moved Permanently
302 Found
303 See Other
304 Not Modified
305 Use Proxy
307 Temporary Redirect
308 Permanent Redirect
400 Bad Request
401 Unauthorized
402 Payment Required
403 Forbidden
404 Not Found
405 Method Not Allowed
406 Not Acceptable
407 Proxy Authentication Required
408 Request Timeout
409 Conflict
410 Gone
411 Length Required
412 Precondition Failed
413 Payload Too Large
414 Request-URI Too Long
415 Unsupported Media Type
416 Requested Range Not Satisfiable
417 Expectation Failed
418 I'm a teapot
421 Misdirected Request
422 Unprocessable Entity
423 Locked
424 Failed Dependency
426 Upgrade Required
428 Precondition Required
429 Too Many Requests
431 Request Header Fields Too Large
444 Connection Closed Without Response
451 Unavailable For Legal Reasons
499 Client Closed Request
500 Internal Server Error
501 Not Implemented
502 Bad Gateway
503 Service Unavailable
504 Gateway Timeout
505 HTTP Version Not Supported
506 Variant Also Negotiates
507 Insufficient Storage
508 Loop Detected
510 Not Extended
511 Network Authentication Required
599 Network Connect Timeout Error"""

from pprint import pprint
from slugify import slugify

def to_camel_case(snake_str):
    components = snake_str.split('_')
    return "".join(x.title() for x in components)

def template_code(obj):
    camel_slug = to_camel_case(obj['slug'])
    camel_slug = camel_slug.replace('Uri', 'URI')
    camel_slug = camel_slug.replace('Http', 'HTTP')

    tpl = """
// {camel}Slug ..
const {camel}Slug = "{slug}"

// {camel} error with message
func {camel}(msg ...interface{{}}) *Fail {{
    message := formatMessage(msg...)
    return &Fail{{
        Slug: {camel}Slug,
        Code: {code},
        Message: message,
    }}
}}

// {camel}f error with message
func {camel}f(format string, args ...interface{{}}) *Fail {{
    message := fmt.Sprintf(format, args...)
    return &Fail{{
        Slug: {camel}Slug,
        Code: {code},
        Message: message,
    }}
}}

// {camel}Detail error with detail lines
func {camel}Detail(detail []string) *Fail {{
    return &Fail{{
        Slug: {camel}Slug,
        Code: {code},
        Detail: detail,
    }}
}}
"""
    return tpl.format(camel=camel_slug, slug=obj['slug'], code=obj['code'])

codes_abs = []
for c in codes.split('\n'):
    obj = {'code': int(c[0:3]), 'msg': c[4:]}
    obj['slug'] = slugify(obj['msg'], separator='_')
    codes_abs.append(obj)



# pprint(codes_abs)

header = """// Package pf CODE GENERATED. DO NOT MODIFY
package pf

import "fmt"
"""

buff = ""
buff += header
# print(header)
for c in codes_abs:
    buff += template_code(c)
    # print(template_code(c))

with open('generated.go', 'w') as f:
    f.write(buff)

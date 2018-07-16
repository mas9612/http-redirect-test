# HTTP Redirect Test

This repository is created to test the behavior of some browsers when the Web server returns status code 30x (redirect).

## Tested Status codes
* 301 Moved Permanently
* 302 Found
* 307 Temporary Redirect
* 308 Permanent Redirect

## Tested clients
* curl 7.43.0
* Safari 11.1.2
* Google Chrome 67.0.3396.99
* Firefox 59.0.2

## Result
### Method GET
| client  | 301 | 302 | 307 | 308 |
|:--------|:----|:----|:----|:----|
| curl    | GET | GET | GET | GET |
| Safari  | GET | GET | GET | GET |
| Chrome  | GET | GET | GET | GET |
| Firefox | GET | GET | GET | GET |

### Method POST
| client  | 301  | 302  | 307  | 308  |
|:--------|:-----|:-----|:-----|:-----|
| curl    | POST | POST | POST | POST |
| Safari  | GET  | GET  | POST | POST |
| Chrome  | GET  | GET  | POST | POST |
| Firefox | GET  | GET  | POST | POST |

# Utils modules

Common utils that need in every development,
it provides a built-in golang function which needs to be wrapped in utils.

## Modules
* conversion
  * String
    * `IntToString(num int) string`
    * `Int8ToString(num int8) string`
    * `Int16ToString(num int16) string`
    * `Int32ToString(num int32) string`
    * `Int64ToString(num int64) string`
    * `Float32ToString(num float32) string`
    * `Float64ToString(num float64) string`
    * `BoolToString(num bool) string`
    * `EscapeString(text string) string`
    * `UnescapeString(text string) string`
  * Numeric
    * `StringToInt(num string) int`
    * `StringToInt8(num string) int8`
    * `StringToInt16(num string) int16`
    * `StringToInt32(num string) int32`
    * `StringToInt64(num string) int64`
    * `StringToFloat32(num string) float32`
    * `StringToFloat64(num string) float64`
    * `StringToBool(num string) bool`

* hash
  * Base64
    * `Base64Encode(text string) string`
    * `Base64Decode(text string) string`
  * Sha1
    * `Sha1HashString(text string) string`

* url
  * `EscapeStringFromUrl(text string) string`
  * `UnescapeStringFromUrl(text string) string`
  * `IsUrlValid(text string) bool`
  * `GetHostFromUrl(text string) string`
  * `GetSchemeFromUrl(text string) string`
  * `GetPathFromUrl(text string) string`
  * `GetRawPathFromUrl(text string) string`
  * `GetRawQueryFromUrl(text string) string`

* date
  * date
    * `GetCurrentDateString() string`
    * `GetCurrentDatetime() time.Time`
    * `GetCurrentUTCDatetime() time.Time`
    * `GetCurrentUTCDateString() string`
    * `GetCurrentMonthNumber() int`
    * `GetCurrentMonthString() string`
    * `GetCurrentDay() int`
    * `GetFirstDateOfMonth(t time.Time) time.Time`
    * `GetFirstDateCurrentMonth() time.Time`
    * `GetLastDateOfMonth(t time.Time) time.Time`
    * `GetLastDateCurrentMonth() time.Time`
    * `GetFirstDateOfYear(t time.Time) time.Time`
    * `GetFirstDateCurrentYear() time.Time`
    * `GetLastDateOfYear(t time.Time) time.Time`
    * `GetLastDateCurrentYear() time.Time`
  * conversion
    * `StringToDate(text string) time.Time`
    * `StringToUTCDate(text string) time.Time`
  * validation
    * `IsDateIsValid(text string) bool`


## Install module

```shell
  go get github.com/FerdinaKusumah/utils
```


## Example
```go
package main

import (
	"fmt"
	"github.com/FerdinaKusumah/utils/conversion"
	"github.com/FerdinaKusumah/utils/hash"
)

func main(){

	x := conversion.StringToInt("1230")
	fmt.Println(x) // 1230

	y := hash.Base64Encode("hello world")
	fmt.Println(y) // aGVsbG8gd29ybGQ=
}
```


## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/FerdinaKusumah/utils/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/FerdinaKusumah/utils/issues). All bug reports are appreciated, some features have not been tested yet due to lack of free time.

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**utils** released under MIT. See `LICENSE` for more details.
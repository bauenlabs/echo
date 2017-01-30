# log
--
    import "github.com/voiceis/echo/lib/log"

Package log provides a standard set of logging tools.

## Usage

```go
var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
	FatalLog   *log.Logger
)
```

#### func  Error

```go
func Error(err error)
```
Error prints a error log.

#### func  Fatal

```go
func Fatal(err error, statusCode int)
```
Fatal prints a fatal error log.

#### func  Info

```go
func Info(logString string)
```
Info prints a info log.

#### func  Trace

```go
func Trace(logString string)
```
Trace prints a trace log.

#### func  Warning

```go
func Warning(logString string)
```
Warning prints a warning log.

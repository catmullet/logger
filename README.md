# logger
Great way to Log in Golang

## Import it
```golang 
import (

  log "github.com/catmullet/logger"
  
)
```
## Understand the line format
```golang
"[%s][%s] %s %s, (%s)"
```
It goes [App Name (Environment variable 'APP')], [Log Level], Message, json object, file and line number)

## Use it
```golang
// Log Levels Error, Warn, Info, Debug, Fatal
log.Info("Hello")
log.Info("Hello %s", "World")
log.Info("Hello World", object)
```

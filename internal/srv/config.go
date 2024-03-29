package srv

import "time"

const readTimeout = time.Second * 6
const writeTimeout = time.Second * 10
const idleTimeout = 0       // Because it's zero, read timeout will be used.
const readHeaderTimeout = 0 // Because it's zero, read timeout will be used.
const maxHeaderBytes = 0
const addr = ":4000"
const disableGeneralOptionsHandler = false

// connContext = nil
// baseContext = nil
// connState = nil
// tlsNextProto = nil
// tlsConfig = nil

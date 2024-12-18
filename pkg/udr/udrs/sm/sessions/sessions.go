package sessions

//"net"
//"reflect"
//"sync"
//"time"

//"k8s.io/klog"

//"w5gc.io/wipro5gcore/pkg/smf/pdusmsp/sm/nodes"

type SessionId string

type Sessions interface {
	GetSession(sessionId SessionId) bool // Retrieve Session
	AddSession(sessionId SessionId)      //Create Session
	DeleteSession(sessionId SessionId)   // Delete session
	UpdateSession(sessionId SessionId)   // Update session
}

package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"k8s-client/pkg/resource/pod"
	"log"
	"net/http"
	"sync"
	"time"

	"gopkg.in/igm/sockjs-go.v2/sockjs"
	"k8s.io/client-go/kubernetes"
)

const (
	LOG      = "log"
	ATTACH   = "attach"
	TERMINAL = "terminal"
)

type SockSession struct {
	SessionId string `json:"sessionId"` // session id
	Namespace string `json:"namespace"` // namespace
	Name      string `json:"name"`      // pod name
	Container string `json:"container"` // container name
	SockType  string `json:"sockType"`  // sock type
	Client    kubernetes.Interface
	Session   sockjs.Session
	bound     chan error
}

type SockMessage struct {
	Op        string `json:"op"`        // operation
	Data      string `json:"data"`      // data
	SessionId string `json:"sessionId"` // session id
}

type SockSessionMap struct {
	Sessions map[string]SockSession
	Lock     sync.RWMutex
}

func (s *SockSessionMap) Set(sessionId string, session SockSession) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Sessions[sessionId] = session
}

func (s *SockSessionMap) Get(sessionId string) SockSession {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return s.Sessions[sessionId]
}

func (s *SockSessionMap) CloseSession(sessionId string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	sockSession := s.Sessions[sessionId]
	sockSession.Session.Close(1000, "close")
	delete(s.Sessions, sessionId)
}

var SockSessions = SockSessionMap{
	Sessions: make(map[string]SockSession),
}

func genSessionId() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	id := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(id, bytes)
	return string(id), nil
}

func CreateAttachHandler(path string) http.Handler {
	return sockjs.NewHandler(path, sockjs.DefaultOptions, handleSession)
}

func handleSession(session sockjs.Session) {
	var (
		buf      string
		err      error
		msg      SockMessage
		sockSess SockSession
	)

	log.Printf("msg from client: %s", buf)

	if buf, err = session.Recv(); err != nil {
		return
	}

	if err = json.Unmarshal([]byte(buf), &msg); err != nil {
		return
	}

	log.Printf("msg from client: %s", msg)

	if msg.Op != "bind" {
		return
	}

	if sockSess = SockSessions.Get(msg.SessionId); sockSess.SessionId == "" {
		return
	}

	log.Printf("session %v", sockSess)
	sockSess.Session = session
	SockSessions.Set(msg.SessionId, sockSess)
	sockSess.bound <- nil
}

func WaitFor(sessionId string) {
	log.Printf("Wait for session %s", sessionId)
	select {
	case <-SockSessions.Get(sessionId).bound:
		close(SockSessions.Get(sessionId).bound)
		err := process(SockSessions.Get(sessionId))
		if err != nil {
			SockSessions.CloseSession(sessionId)
			return
		}
	case <-time.After(10 * time.Second):
		close(SockSessions.Get(sessionId).bound)
		delete(SockSessions.Sessions, sessionId)
		return
	}
}

func process(session SockSession) error {
	if session.SockType == TERMINAL {
	}

	if session.SockType == ATTACH {
	}

	if session.SockType == LOG {
		return logProcess(session)
	}
	return nil
}

func logProcess(session SockSession) error {
	log.Printf("log process %v", session)
	stream, err := pod.GetLog(session.Client, session.Namespace, session.Name, session.Container)

	if err != nil {
		log.Printf("get log error %v", err)
		return err
	}

	buf := make([]byte, 1024)

	for {
		n, err := stream.Read(buf)
		if err != nil {
			return err
		}
		log.Printf("log %s", string(buf[:n]))
		session.Session.Send(string(buf[:n]))
	}

}

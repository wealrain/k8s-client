package handler

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

const END_OF_TRANSMISSION = "\u0004"

type PtyHandler struct {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

type TerminalMessage struct {
	Op        string `json:"op"`        // operation
	Data      string `json:"data"`      // data
	SessionId string `json:"sessionId"` // session id
	Rows      int    `json:"rows"`      // rows
	Cols      int    `json:"cols"`      // cols
}

type TerminalSession struct {
	id       string
	bound    chan error
	session  sockjs.Session
	sizeChan chan remotecommand.TerminalSize
}

// Next returns the new terminal size after the terminal has been resized. It returns nil when
// monitoring has been stopped.
func (t TerminalSession) Next() *remotecommand.TerminalSize {
	size := <-t.sizeChan
	if size.Height == 0 && size.Width == 0 {
		return nil
	}
	return &size
}

// session 管理器
type SessionMap struct {
	Sessions map[string]TerminalSession
	Lock     sync.RWMutex
}

func (s *SessionMap) Set(sessionId string, session TerminalSession) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.Sessions[sessionId] = session
}

func (s *SessionMap) Get(sessionId string) TerminalSession {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return s.Sessions[sessionId]
}

func (s *SessionMap) CloseSession(sessionId string, status uint32, reason string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	terminalSession := s.Sessions[sessionId]
	if err := terminalSession.session.Close(status, reason); err != nil {
		log.Println(err)
	}
	close(terminalSession.sizeChan)
	delete(s.Sessions, sessionId)
}

var terminalSessions = SessionMap{
	Sessions: make(map[string]TerminalSession),
}

// 处理Session 连接
func HandleTerminalSession(session sockjs.Session) {
	var (
		buf             string
		err             error
		msg             TerminalMessage
		terminalSession TerminalSession
	)

	if buf, err = session.Recv(); err != nil {
		return
	}

	if err = json.Unmarshal([]byte(buf), &msg); err != nil {
		return
	}

	if msg.Op != "bind" {
		return
	}

	if terminalSession = terminalSessions.Get(msg.SessionId); terminalSession.id == "" {
		return
	}

	terminalSession.session = session
	terminalSessions.Set(msg.SessionId, terminalSession)
	terminalSession.bound <- nil
}

// 生成随机sessionId
func generateSessionId() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	id := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(id, bytes)
	return string(id)
}

func HandleExecShell(c *gin.Context) {
	client := managerK8SClient(c)
	sessionId := generateSessionId()

	terminalSessions.Set(sessionId, TerminalSession{
		id:       sessionId,
		bound:    make(chan error),
		sizeChan: make(chan remotecommand.TerminalSize),
	})

	// 启动一个协程进行处理
	go waitForTerminal(client, sessionId, c)

	c.JSON(200, sessionId)
}

func waitForTerminal(client kubernetes.Interface, sessionId string, c *gin.Context) {

}

func startProcess(client *kubernetes.Clientset, c *gin.Context, cmd []string, ptyHandler PtyHandler) error {
	namespace := c.Param("namespace")
	pod := c.Param("pod")
	container := c.Param("container")

	req := client.CoreV1().
		RESTClient().
		Post().
		Resource("pods").
		Name(pod).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Container: container,
		Command:   cmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(managerK8SConig(c), "POST", req.URL())

	if err != nil {
		return err
	}

	return exec.Stream(remotecommand.StreamOptions{
		Stdin:             ptyHandler,
		Stdout:            ptyHandler,
		Stderr:            ptyHandler,
		TerminalSizeQueue: ptyHandler,
	})
}

package client

import (
	"log"
	"os"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ClientManagerMap struct {
	ClientMap map[string]*ClientManager
	sync.RWMutex
}

type ClientManager struct {
	Client       kubernetes.Interface
	Config       *rest.Config
	VerberClient *Verber
}

var ClientManagerMapInstance = ClientManagerMap{ClientMap: make(map[string]*ClientManager)}

func (c *ClientManagerMap) GetClientManager(id string) *ClientManager {
	c.RLock()
	defer c.RUnlock()
	return c.ClientMap[id]
}

func (c *ClientManagerMap) SetClientManager(id string, clientManager *ClientManager) {
	c.Lock()
	defer c.Unlock()
	c.ClientMap[id] = clientManager
}

func (c *ClientManagerMap) DeleteClientManager(id string) {
	c.Lock()
	defer c.Unlock()
	delete(c.ClientMap, id)
}

func NewClientManager(id, config string) (*ClientManager, error) {
	clientManager := &ClientManager{}
	err := clientManager.initClient(id, config)
	if err != nil {
		return nil, err
	}
	ClientManagerMapInstance.SetClientManager(id, clientManager)
	return clientManager, nil
}

func (c *ClientManager) initClient(id, configContent string) error {
	// 判断配置文件是否存在，文件名config-cluster-id
	file, err := os.Create("config-" + id)
	if err != nil {
		return err
	}

	_, err = file.WriteString(configContent)
	if err != nil {
		return err
	}
	file.Close()
	config, err := clientcmd.BuildConfigFromFlags("", "config-"+id)
	if err != nil {
		log.Println(err)
		return err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
		return err
	}
	c.Client = client
	c.Config = config
	c.VerberClient = NewVerber(client, config)
	return nil
}

func (c *ClientManager) InitClient(kubeConfig string) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return err
	}
	log.Printf("config: %v", config)
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	c.Client = client
	c.Config = config
	c.VerberClient = NewVerber(client, config)
	return nil
}

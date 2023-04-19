package cluster

import (
	"k8s-client/config"
	"k8s-client/pkg/errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conf = config.LoadConfig()

func GetDBConnection(database string) *gorm.DB {

	username := Conf.Mysql.Username
	password := Conf.Mysql.Password
	host := Conf.Mysql.Host
	port := Conf.Mysql.Port
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type Cluster struct {
	Id       int    `json:"id" column:"id" gorm:"primary_key;AUTO_INCREMENT"` // 集群ID
	Name     string `json:"name" column:"name"`                               // 集群名称
	Config   string `json:"config" column:"config"`                           // 集群配置文件路径
	Version  string `json:"version" column:"version"`                         // 集群版本
	Password string `json:"password" column:"password"`                       // 集群密码
	Status   string `json:"status" column:"status"`                           // 是否启用
}

func clusterDatabase() *gorm.DB {
	return GetDBConnection("k8s")
}

func (Cluster) TableName() string {
	return "cluster"
}

func GetClusterById(id string) (*Cluster, error) {
	var cluster Cluster
	db := clusterDatabase()
	db.Where("id = ?", id).First(&cluster)
	return &cluster, db.Error
}

func GetClusterList() ([]Cluster, error) {
	var clusters []Cluster
	db := clusterDatabase()
	db.Find(&clusters)
	return clusters, db.Error
}

func AddCluster(cluster *Cluster) error {
	db := clusterDatabase()
	exist, err := CheckClusterExist(cluster.Name)
	if err != nil {
		return err
	}

	if exist {
		return errors.NewClusterError("集群已存在")
	}

	result := db.Create(&cluster)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("cluster: %v", cluster)
	return nil
}

func UpdateCluster(cluster *Cluster) error {
	log.Printf("cluster: %v", cluster)
	db := clusterDatabase()
	db.Save(&cluster)
	return db.Error
}

func DeleteCluster(id string) error {
	db := clusterDatabase()
	db.Where("id = ?", id).Delete(&Cluster{})
	return db.Error
}

func CheckClusterExist(name string) (bool, error) {
	var cluster Cluster
	db := clusterDatabase()
	db.Where("name = ?", name).First(&cluster)
	if cluster.Name == name {
		return true, nil
	}
	return false, db.Error
}

package api

import (
	"net/http"

	"github.com/ch3lo/overlord/api/types"
	"github.com/ch3lo/overlord/util"
	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	/*	bag := manager.ServicesBag()

		var services map[string]types.Service = make(map[string]types.Service, 0)
		for _, srv := range bag {
			var versions map[string]types.ServiceVersion = make(map[string]types.ServiceVersion, 0)
			for _, v := range srv.Versions {
				versions[v.Version] = types.ServiceVersion{}
			}
			services[srv.Id] = types.Service{Versions: versions}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"services": services})*/
}

func PutService(c *gin.Context) {
	var service types.Service

	if err := c.BindJSON(&service); err != nil {
		util.Log.Println(err)
		se := &SerializationError{Message: err.Error()}
		c.JSON(http.StatusOK, gin.H{
			"status":  se.GetStatus(),
			"message": se.GetMessage()})
		return
	}
	/*
		srv := &manager.Service{Id: service.Id}

		if err := manager.RegisterService(srv); err != nil {
			util.Log.Println(err)
			if ce, ok := err.(*util.ElementAlreadyExists); ok {
				c.JSON(http.StatusOK, gin.H{
					"status":  ce.GetStatus(),
					"message": ce.GetMessage()})
			} else {
				ue := &UnknownError{}
				c.JSON(http.StatusOK, gin.H{
					"status":  ue.GetStatus(),
					"message": ue.GetMessage()})
			}
			return
		}
	*/
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"service": service})
}

func GetServiceByServiceId(c *gin.Context) {
	/*	serviceId := c.Param("service_id")

		bag := manager.ServicesBag()
		for _, v := range bag {
			if v.Id == serviceId {
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"service": types.Service{}})
				return
			}
		}*/

	snf := &ServiceNotFound{}
	c.JSON(http.StatusOK, gin.H{
		"status":  snf.GetStatus(),
		"message": snf.GetMessage()})
}

func GetServiceByClusterAndServiceId(c *gin.Context) {
	/*cluster := c.Param("cluster")
	serviceId := c.Param("service_id")

		status, err := manager.GetService(cluster, serviceId)
		if err != nil {
			snf := &ServiceNotFound{}
			c.JSON(http.StatusOK, gin.H{
				"status":  snf.GetStatus(),
				"message": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"service": ""})
}

func PutServiceVersionByServiceId(c *gin.Context) {
	//	serviceId := c.Param("service_id")
	var sv types.ServiceVersion

	if err := c.BindJSON(&sv); err != nil {
		util.Log.Println(err)
		se := &SerializationError{Message: err.Error()}
		c.JSON(http.StatusOK, gin.H{
			"status":  se.GetStatus(),
			"message": se.GetMessage()})
		return
	}
	/*
		managedsv := &manager.ServiceVersion{Version: sv.Version}

		if err := manager.RegisterServiceVersion(serviceId, managedsv); err != nil {
			util.Log.Println(err)
			if ce, ok := err.(*util.ElementAlreadyExists); ok {
				c.JSON(http.StatusOK, gin.H{
					"status":  ce.GetStatus(),
					"message": ce.GetMessage()})
			} else {
				ue := &UnknownError{}
				c.JSON(http.StatusOK, gin.H{
					"status":  ue.GetStatus(),
					"message": ue.GetMessage()})
			}
			return
		}
	*/
	c.JSON(http.StatusOK, gin.H{
		"status":          http.StatusOK,
		"service_version": sv})
}

/*
func ServicesTestGet(c *gin.Context) {

	for i := 0; i < 5; i++ {
		var service manager.Service
		service.Address = fmt.Sprintf("localhost:8%s", i)
		service.Id = strconv.Itoa(i)
		service.Status = "status"
		manager.Register(&service)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   manager.MonitoredServices[0].GetMonitor().Check("asd", "localhost:80")})
}
*/

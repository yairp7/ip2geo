package controllers

import (
	"fmt"
	"sync"

	"github.com/yairp7/ip2geo/internal/common"
)

type BaseController struct {
	name       string
	isActive   bool
	currentOps sync.WaitGroup
	loggerImpl common.Logger
	services   []common.Closer
}

func NewBaseController(name string, loggerImpl common.Logger) BaseController {
	return BaseController{
		name:       name,
		isActive:   true,
		currentOps: sync.WaitGroup{},
		loggerImpl: loggerImpl,
	}
}

func (c *BaseController) Close() {
	c.loggerImpl.Debug(fmt.Sprintf("%s Shutdown\n", c.name))
	c.isActive = false
	c.currentOps.Wait()
	for _, closeableService := range c.services {
		closeableService.Close()
	}
}

func (c *BaseController) RegisterOp() {
	c.currentOps.Add(1)
}

func (c *BaseController) UnregisterOp() {
	c.currentOps.Done()
}

func (c *BaseController) RegisterService(services ...common.Closer) {
	for _, service := range services {
		c.services = append(c.services, service)
	}
}

package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (manager *MyManager) GetName() string {
	fmt.Println("GET NAME MY MANAGER")
	return manager.Name
}

func (manager *MyManager) GetManagerName() string {
	return manager.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (vp *MyVicePresident) GetName() string {
	fmt.Println("GET NAME MY VICE PRESIDENT")
	return vp.Name
}

func (vp *MyVicePresident) GetVicePresidentName() string {
	return vp.Name
}

func TestTypeInheritance(t *testing.T) {
	assert.Equal(t, "Zaki", GetName(&MyManager{Name: "Zaki"}))
	assert.Equal(t, "Zaki", GetName(&MyVicePresident{Name: "Zaki"}))
}

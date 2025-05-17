//go:build wireinject
// +build wireinject

package di

import (
	"io"
	"learning-dependency-injection-golang/simple"
	"os"

	"github.com/google/wire"
)

func InitializedService(isError bool) (*simple.SimpleService, error) {
	wire.Build(simple.NewSimpleRepository, simple.NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *simple.DatabaseRepository {
	wire.Build(simple.NewDatabasePostgreSQL, simple.NewDatabaseMongoDB, simple.NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(simple.NewFooRepository, simple.NewFooService)

var barSet = wire.NewSet(simple.NewBarRepository, simple.NewBarService)

func InitializeFooBarService() *simple.FooBarService {
	wire.Build(fooSet, barSet, simple.NewFooBarService)
	return nil
}

var fooBarSet = wire.NewSet(simple.NewFoo, simple.NewBar)

func InitializedFooBar() *simple.FooBar {
	wire.Build(fooBarSet, wire.Struct(new(simple.FooBar), "Foo", "Bar"))
	return nil
}

var helloSet = wire.NewSet(
	simple.NewSayHelloImpl,
	wire.Bind(new(simple.SayHello), new(*simple.SayHelloImpl)),
)

func InitializedHelloService() *simple.HelloService {
	wire.Build(helloSet, simple.NewHelloService)
	return nil
}

var fooValue = &simple.Foo{}
var barValue = &simple.Bar{}

func InitializedFooBarUsingValue() *simple.FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(simple.FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *simple.Configuration {
	wire.Build(simple.NewApplication, wire.FieldsOf(new(*simple.Application), "Configuration"))
	return nil
}

func InitializedConnection(name string) (*simple.Connection, func()) {
	wire.Build(simple.NewFile, simple.NewConnection)
	return nil, nil
}

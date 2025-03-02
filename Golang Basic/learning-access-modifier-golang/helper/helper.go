package helper

var version = "1.0.0"      // private variable (can access only in the same package)
var Application = "golang" // public variable (can access in the same package or outside package)

// private function (can access only in the same package)
func sayGoodBye(name string) string {
	return "Good Bye " + name
}

// public function (can access in the same package or outside package)
func SayHello(name string) string {
	return "Hello " + name
}

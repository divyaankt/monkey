package object

//Environment ...
/*
In the Monkey interpreter,the environment is the object
that associates values with variable names. Typically it
is a list of frames, and each frame is a list of variable
bindings. This method of implementing an interpreter has
been popularised in the
SICP ( Structure and Interpretation of Computer Programs) book.
*/
type Environment struct {
	store map[string]Object
}

//NewEnvironment ...
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

//Get ...
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

//Set ...
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

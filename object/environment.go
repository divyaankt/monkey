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
	outer *Environment
}

//NewEnvironment ...
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

//Get ...
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

//Set ...
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

//NewEnclosedEnvironment ...
/*This new enclosed environment is for implementing local variables
present inside the function body. Here we are extending the main environment.
Extending the environment means that we create a new instance of object.
Environment with a pointer to the environment it should extend. By doing
that we enclose a fresh and empty environment with an existing one.
When the new environment’s Get method is called and it itself doesn’t have
a value associated with the given name, it calls the Get of the enclosing environment.
That’s the environment it’s extending. And if that enclosing environment can’t find
the value, it calls its own enclosing environment and so on until there is no enclosing
environment anymore and we can safely say that we have an ERROR.
*/
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

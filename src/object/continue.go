package object

const CONTINUE_OBJ = "CONTINUE"

type Continue struct {
	Object
}

func (obj *Continue) Type() ObjectType {
	return CONTINUE_OBJ
}

func (obj *Continue) String() string {
	return CONTINUE_OBJ
}

func (obj *Continue) Method(method string, args []Object) (Object, bool) {
	return nil, false
}

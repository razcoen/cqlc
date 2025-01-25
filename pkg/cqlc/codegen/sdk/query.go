package sdk

type Queries []*Query

type Query struct {
	FuncName    string
	Annotations []string
	Stmt        string
	Params      []string
	Selects     []string
	Table       string
	Keyspace    string
}

type Annotation string

const (
	AnnotationExec  Annotation = "exec"
	AnnotationOne   Annotation = "one"
	AnnotationMany  Annotation = "many"
	AnnotationBatch Annotation = "batch"
)

func ParseAnnotation(s string) (Annotation, bool) {
	m := map[Annotation]bool{AnnotationExec: true, AnnotationOne: true, AnnotationMany: true, AnnotationBatch: true}
	a := Annotation(s)
	_, ok := m[a]
	return a, ok
}

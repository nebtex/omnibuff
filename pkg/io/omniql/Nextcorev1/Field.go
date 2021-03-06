package Nextcorev1

//go:generate mockery -name=FieldReader
//FieldReader ...
type FieldReader interface {
	//Name ...
	Name() string

	//Type ...
	Type() string

	//Documentation ...
	Documentation() (DocumentationReader, error)

	//Default String representation of the default value
	Default() string
}

//VectorFieldReader ...
type VectorFieldReader interface {
	// Returns the current size of this vector
	Len() int

	//Get the item in the position i, if i < Len(),
	//if item does not exist should return the default value for the underlying data type
	//when i > Len() should return an VectorInvalidIndexError
	Get(i int) (item FieldReader, err error)
}

//FieldShard ...
type FieldShard interface {
	Name() ShardHolderAndDisposer
	Type() ShardHolderAndDisposer
	Documentation(func(DocumentationShard)) ShardHolderAndDisposer
	Default() ShardHolderAndDisposer
}

//FieldWildcardShard ...
func FieldWildcardShard(s FieldShard) {

	s.Name()
	s.Type()
	s.Default()
}

//FieldForwardShard ...
func FieldForwardShard(s FieldShard) {

	s.Name()
	s.Type()
	s.Documentation(DocumentationForwardShard)
	s.Default()
}


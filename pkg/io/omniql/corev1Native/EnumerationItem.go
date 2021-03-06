package corev1Native

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
)

//EnumerationItem ...
type EnumerationItem struct {
	Name          string `json:"name"`
	Documentation *Documentation `json:"documentation"`
}

//EnumerationItemReader ...
type EnumerationItemReader struct {
	_enumerationitem *EnumerationItem
	documentation    *DocumentationReader
}

//Name ...
func (i *EnumerationItemReader) Name() (value string) {
	value = i._enumerationitem.Name
	return
}

//Documentation ...
func (i *EnumerationItemReader) Documentation() corev1.DocumentationReader {

	if i.documentation != nil {
		return i.documentation
	}

	return nil
}

func NewEnumerationItemReader(e *EnumerationItem) *EnumerationItemReader {
	if e == nil {
		return nil
	}
	return &EnumerationItemReader{
		_enumerationitem: e,
		documentation:    NewDocumentationReader(e.Documentation),
	}
}

type VectorEnumerationItemReader struct {
	_vector []*EnumerationItemReader
}

func (vi *VectorEnumerationItemReader) Len() (size int) {
	size = len(vi._vector)
	return
}

func (vi *VectorEnumerationItemReader) Get(i int) (item corev1.EnumerationItemReader, err error) {

	if i < 0 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vector)}
		return
	}

	if i > len(vi._vector)-1 {
		err = &hybrids.VectorInvalidIndexError{Index: i, Len: len(vi._vector)}
		return
	}

	item = vi._vector[i]
	return

}

func NewVectorEnumerationItemReader(v []*EnumerationItem) (obj *VectorEnumerationItemReader) {
	if v == nil {
		return nil
	}

	obj = &VectorEnumerationItemReader{}
	obj._vector = make([]*EnumerationItemReader, len(v))

	for i, value := range v {
		obj._vector[i] = NewEnumerationItemReader(value)
	}
	return
}

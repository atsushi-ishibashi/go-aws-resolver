package awsresolver

import (
	"reflect"
	"testing"
)

func TestResolverIface(t *testing.T) {
	resolver := NewResolver("ap-northeast-1")
	iface := reflect.TypeOf((*ResolverIface)(nil)).Elem()
	if reflect.TypeOf(resolver).Implements(iface) == false {
		t.Errorf("Resolver expected to implement ResolverIface")
	}
}

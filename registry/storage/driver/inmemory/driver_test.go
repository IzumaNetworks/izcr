package inmemory

import (
	"testing"

	storagedriver "go.izuma.io/izcr/registry/storage/driver"
	"go.izuma.io/izcr/registry/storage/driver/testsuites"
)

func newDriverConstructor() (storagedriver.StorageDriver, error) {
	return New(), nil
}

func TestInMemoryDriverSuite(t *testing.T) {
	testsuites.Driver(t, newDriverConstructor, false)
}

func BenchmarkInMemoryDriverSuite(b *testing.B) {
	testsuites.BenchDriver(b, newDriverConstructor)
}

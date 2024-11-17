package checks

import (
	"testing"

	"github.com/caas-team/sparrow/test"
	"github.com/stretchr/testify/assert"
)

func TestBase_Shutdown(t *testing.T) {
	test.MarkAsShort(t)

	tests := []struct {
		name string
		b    *Base
	}{
		{
			name: "shutdown",
			b: &Base{
				Done: make(chan struct{}, 1),
			},
		},
		{
			name: "already shutdown",
			b: &Base{
				Done:   make(chan struct{}, 1),
				closed: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.b.closed {
				close(tt.b.Done)
			}
			tt.b.Shutdown()

			if !tt.b.closed {
				t.Error("Base.Shutdown() should close Base.Done")
			}

			assert.Panics(t, func() {
				tt.b.Done <- struct{}{}
			}, "Base.Done should be closed")
		})
	}
}

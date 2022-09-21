package tests

import (
	"playground/handlers"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/stretchr/testify/assert"
)

func TestAtomicity(t *testing.T) {
	handlers.Setup()
	type (
		testCase struct {
			name      string
			checkFunc func(*testing.T, interface{})
		}
	)

	var atomicityTests = []*testCase{
		{
			name: "get value first",
			checkFunc: func(t *testing.T, i interface{}) {
				var res handlers.GetValueResp
				b, _ := json.Marshal(i)
				json.Unmarshal(b, &res)

				assert.NotNil(t, res)
			},
		}
	}

	for _,test:=range atomicityTests{
		
	}

}

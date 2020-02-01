package obj_pool_test

import (
	"ch32/obj_pool"
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := obj_pool.NewObjPool(10)
	if err := pool.ReleaseObj(&obj_pool.ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}

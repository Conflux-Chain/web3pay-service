package util

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKMutex(t *testing.T) {
	kmutex := NewKMutex()
	lockKey := MutexKey("kmutex")

	testNum := int64(0)
	testTimes := int64(2000)

	wg := &sync.WaitGroup{}

	for i := int64(0); i < testTimes; i++ {
		wg.Add(1)

		go func() {
			kmutex.Lock(lockKey)
			defer kmutex.Unlock(lockKey)

			amt := time.Duration(rand.Intn(50))
			time.Sleep(time.Millisecond * amt)

			testNum = testNum + 1

			wg.Done()
		}()
	}

	wg.Wait()

	assert.EqualValues(t, testTimes, testNum)
}

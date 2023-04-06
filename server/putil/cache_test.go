package putil

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/util"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCacheMiddleware_Disabled(t *testing.T) {
	mfs := afero.NewMemMapFs()
	m := NewCacheMiddleware(CacheConfig{Disabled: true, FS: mfs})
	e := echo.New()
	e.Use(m.Middleware())
	e.GET("/aaa", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	r := httptest.NewRequest("GET", "/aaa", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "hello", w.Body.String())
}

func TestCacheMiddleware(t *testing.T) {
	res := ""
	called := 0
	mfs := afero.NewMemMapFs()
	m := NewCacheMiddleware(CacheConfig{FS: mfs})
	e := echo.New()
	g := e.Group("/api")
	g.Use(m.Middleware())
	g.GET("/aaa", func(c echo.Context) error {
		called++
		return c.String(http.StatusOK, res)
	})
	g.GET("/bbb", func(c echo.Context) error {
		return c.String(http.StatusNotFound, res)
	})

	// 1st: first request
	res = "res1"
	r := httptest.NewRequest("GET", "/api/aaa", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "res1", w.Body.String())
	assert.Equal(t, "res1", string(lo.Must(afero.ReadFile(mfs, "_api_aaa"))))
	assert.Equal(t, 1, called)

	// 2nd: cached request
	res = "res2"
	r = httptest.NewRequest("GET", "/api/aaa", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "res1", w.Body.String())
	assert.Equal(t, "res1", string(lo.Must(afero.ReadFile(mfs, "_api_aaa"))))
	assert.Equal(t, 1, called) // handler not called

	// 3rd: cache exprires
	res = "res2"
	m.now = func() time.Time { return util.Now().Add(3 * time.Minute) } // advance time
	r = httptest.NewRequest("GET", "/api/aaa", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "res2", w.Body.String())
	assert.Equal(t, "res2", string(lo.Must(afero.ReadFile(mfs, "_api_aaa"))))
	assert.Equal(t, 2, called) // handler not called

	// 4th: error
	res = "res3"
	m.now = util.Now
	r = httptest.NewRequest("GET", "/api/bbb", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "res3", w.Body.String())
	assert.Equal(t, "res3", string(lo.Must(afero.ReadFile(mfs, "_api_bbb")))) //

	// 4th: error ignores cache
	res = "res4"
	m.now = util.Now
	r = httptest.NewRequest("GET", "/api/bbb", nil)
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, "res4", w.Body.String())
	assert.Equal(t, "res4", string(lo.Must(afero.ReadFile(mfs, "_api_bbb")))) //
}

func TestCacheMiddlewareAsync(t *testing.T) {
	res := ""
	called := atomic.Int32{}
	l := sync.Mutex{}
	mfs := afero.NewMemMapFs()
	m := NewCacheMiddleware(CacheConfig{FS: mfs})
	e := echo.New()
	e.Use(m.Middleware())
	e.GET("/aaa", func(c echo.Context) error {
		called.Add(1)
		time.Sleep(time.Millisecond * 100)
		return c.String(http.StatusOK, res)
	})

	res = "res1"
	a := []string{}
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		r := httptest.NewRequest("GET", "/aaa", nil)
		w := httptest.NewRecorder()

		l.Lock()
		a = append(a, "load")
		l.Unlock()

		e.ServeHTTP(w, r)

		l.Lock()
		a = append(a, "done")
		l.Unlock()

		assert.Equal(t, "res1", w.Body.String())
		wg.Done()
	}()

	go func() {
		r := httptest.NewRequest("GET", "/aaa", nil)
		w := httptest.NewRecorder()

		l.Lock()
		a = append(a, "load")
		l.Unlock()

		e.ServeHTTP(w, r)

		l.Lock()
		a = append(a, "done")
		l.Unlock()

		assert.Equal(t, "res1", w.Body.String())
		wg.Done()
	}()

	wg.Wait()
	assert.Equal(t, []string{"load", "load", "done", "done"}, a)
	assert.Equal(t, "res1", string(lo.Must(afero.ReadFile(mfs, "_aaa"))))
	assert.Equal(t, int32(1), called.Load()) // lock works
}

func TestCacheEntry_Active(t *testing.T) {
	now := util.Now()
	expires := now.Add(defaultCacheTTL)
	assert.True(t, cacheEntry{Expires: expires}.Active(now))
	assert.True(t, cacheEntry{Expires: expires}.Active(now.Add(-defaultCacheTTL)))
	assert.True(t, cacheEntry{Expires: expires}.Active(now.Add(defaultCacheTTL).Add(-time.Second)))
	assert.False(t, cacheEntry{Expires: expires}.Active(now.Add(defaultCacheTTL)))
}

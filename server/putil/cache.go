package putil

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/reearth/reearthx/log"
	"github.com/reearth/reearthx/util"
	"github.com/spf13/afero"
)

const defaultCacheTTL = 3 * time.Minute
const cacheBasePath = "cache"

type CacheConfig struct {
	Disabled     bool
	TTL          time.Duration
	FS           afero.Fs
	CacheControl bool
}

type CacheMiddleware struct {
	cfg  CacheConfig
	lock *KeyLock[string]
	m    *util.SyncMap[string, cacheEntry]
	now  func() time.Time
}

func NewCacheMiddleware(cfg CacheConfig) *CacheMiddleware {
	if cfg.TTL == 0 {
		cfg.TTL = defaultCacheTTL
	}
	if cfg.FS == nil {
		osfs := afero.NewOsFs()
		_ = osfs.MkdirAll(cacheBasePath, os.FileMode(0755))
		cfg.FS = afero.NewBasePathFs(osfs, cacheBasePath)
	}

	return &CacheMiddleware{
		cfg:  cfg,
		lock: NewKeyLock[string](),
		m:    util.NewSyncMap[string, cacheEntry](),
		now:  util.Now,
	}
}

func (m *CacheMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if m.cfg.Disabled {
				return next(c)
			}

			key := m.key(c)
			if key == "" {
				return next(c)
			}

			if err := m.load(c, key); err != nil {
				return err
			} else if c.Response().Committed {
				return nil
			}

			// Lock only when a new cache needs to be created.
			m.lock.Lock(key)
			defer m.lock.Unlock(key)

			// If someone has already locked the cache, someone is in the process of creating a cache,
			// so when the lock is released, try loading the cache again.
			if err := m.load(c, key); err != nil {
				return err
			} else if c.Response().Committed {
				return nil
			}

			writing := false
			if f, err := m.cfg.FS.Create(key); err == nil {
				defer func() { _ = f.Close() }()
				writing = true
				c.Response().Writer = &responseWriter{Writer: f, ResponseWriter: c.Response().Writer}
				log.Debugf("cache: new: key=%s", key)
			} else {
				log.Errorf("cache: failed to create a file: key=%s, err=%v", key, err)
			}

			if err := next(c); err != nil {
				return err
			}

			if !writing || key == "" || c.Response().Status != http.StatusOK {
				return nil
			}

			e := m.save(key, c)
			maxAge := m.setCacheControl(c, m.cfg.TTL)
			log.Debugf("cache: created: key=%s, expires_at=%d, content_type=%s, max-age=%d", key, e.Expires.Unix(), e.ContentType, maxAge)
			return nil
		}
	}
}

type cacheEntry struct {
	Expires     time.Time
	ContentType string
}

func (c cacheEntry) Active(now time.Time) bool {
	return c.Expires.After(now)
}

func (m *CacheMiddleware) key(c echo.Context) string {
	return strings.ReplaceAll(c.Request().URL.Path, "/", "_")
}

func (m *CacheMiddleware) load(c echo.Context, key string) error {
	if e, ok := m.m.Load(key); ok && e.Active(m.now()) {
		r, err := m.readFile(key)
		if r != nil {
			defer func() { _ = r.Close() }()
			maxAge := m.setCacheControl(c, e.Expires.Sub(util.Now()))
			log.Debugf("cache: hit: key=%s, expires_at=%d, content_type=%s, max-age=%d", key, e.Expires.Unix(), e.ContentType, maxAge)
			return c.Stream(http.StatusOK, e.ContentType, r)
		} else {
			log.Errorf("cache: failed to load a file: key=%s, err=%v", key, err)
		}
	}
	return nil
}

func (m *CacheMiddleware) setCacheControl(c echo.Context, d time.Duration) int {
	maxAge := -1
	if m.cfg.CacheControl {
		maxAge2 := int(d.Seconds())
		if maxAge2 < 10 {
			maxAge2 = 1 // 1s
		}
		c.Response().Header().Set(echo.HeaderCacheControl, fmt.Sprintf("public, max-age=%d", maxAge2))
		maxAge = maxAge2
	}
	return maxAge
}

func (m *CacheMiddleware) save(k string, c echo.Context) cacheEntry {
	e := cacheEntry{
		Expires:     util.Now().Add(m.cfg.TTL),
		ContentType: c.Response().Header().Get(echo.HeaderContentType),
	}
	m.m.Store(k, e)
	return e
}

func (m *CacheMiddleware) readFile(k string) (io.ReadCloser, error) {
	f, err := m.cfg.FS.Open(k)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	return f, nil
}

type responseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *responseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if b, err := w.ResponseWriter.Write(b); err != nil {
		return b, err
	}
	return w.Writer.Write(b)
}

func (w *responseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

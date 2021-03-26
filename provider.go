package journal

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/shestakovda/typex"

	"github.com/shestakovda/journal/crash"
)

/*
	NewProvider - конструктор сборщика журнала.

	* max - максимальный разрешенный уровень логирования для записей
	* crp - провайдер отчетов об ошибках
	* drv - реализация драйвера для сохранения записи журнала
	* log - реализация логгера для записи журнала в консольку или файл
	* srv - наименование сервиса, инициатора записей в журнале
*/
func NewProvider(max int, crp crash.Provider, drv Driver, log Logger, srv string) Provider {
	if log == nil {
		log = new(GlogLogger)
	}

	p := &provider{
		max:   max,
		drv:   drv,
		crp:   crp,
		log:   log,
		srv:   srv,
		lvl:   0,
		start: time.Now(),
		chain: make([]*Stage, 0, 16),
	}
	p.host, _ = os.Hostname()
	p.point = p.start
	return p
}

type provider struct {
	sync.RWMutex

	crash bool
	point time.Time
	start time.Time
	chain []*Stage
	host  string

	drv Driver
	log Logger
	srv string
	max int
	lvl int
	opt bool
	crp crash.Provider
	dbg map[string]string
}

func (p *provider) V(lvl int) bool {
	if lvl > p.max {
		return false
	}

	p.lvl = lvl
	return true
}

func (p *provider) Print(txt string, args ...interface{}) {
	p.stage(&Stage{
		Text: fmt.Sprintf(txt, args...),
	})
}

func (p *provider) Model(mtp ModelType, mid string, txt string, args ...interface{}) {
	p.stage(&Stage{
		EnID: mid,
		Type: mtp.ID(),
		Text: fmt.Sprintf(txt, args...),
	})
}

func (p *provider) Crash(err error) (r *crash.Report) {
	if r = p.crp.Report(err); r != nil {
		p.stage(&Stage{Fail: r})
	}
	return r
}

func (p *provider) Close() *Entry {
	var err error

	e := &Entry{
		ID:      typex.NewUUID().Hex(),
		Debug:   p.dbg,
		Service: p.srv,
		Host:    p.host,
		Chain:   p.chain,
		Start:   p.start.UTC(),
		Total:   time.Since(p.start),
	}

	// Нужно сохранять, если есть драйвер, есть ошибка или обязательно надо
	if p.drv != nil && (p.crash || !p.opt) {
		if err = p.drv.InsertEntry(e); err != nil {
			p.Crash(err)
			e.Chain = p.chain
			e.Total = time.Since(p.start)
		}
	}

	e.Debug = nil // убираем дебаг - он нужен только для ошибок

	if p.crash {
		p.log.Error("%s", e)
	} else {
		p.log.Print("%s", e)
	}

	return e
}

func (p *provider) Clone() Provider {
	prv := NewProvider(p.max, p.crp, p.drv, p.log, p.srv)
	prv.SaveOnlyError(p.opt)
	return prv
}

func (p *provider) SaveOnlyError(opt bool) { p.opt = opt }

func (p *provider) Debug(dbg map[string]string) {
	if len(p.dbg) == 0 {
		p.dbg = dbg
	} else {
		for i := range dbg {
			p.dbg[i] = dbg[i]
		}
	}
}

func (p *provider) stage(s *Stage) {
	if s.Fail != nil {
		s.EnID = s.Fail.ID
		s.Type = ModelTypeCrash.ID()
		s.Text = fmt.Sprintf("[ %d ] %s", s.Fail.Status, s.Fail.Title)
	}

	if s.EnID == "" {
		s.Type = ModelTypeUnknown.ID()
	}

	p.Lock()
	defer p.Unlock()

	s.Verb = p.lvl
	s.Wait = time.Since(p.point)
	p.point = time.Now()
	p.crash = p.crash || s.Type == ModelTypeCrash.ID()
	p.chain = append(p.chain, s)
	p.lvl = 0
}

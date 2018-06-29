package movers

import (
	"github.com/wieku/danser/bmath/curves"
	//"osubot/io"
	"github.com/wieku/danser/beatmap/objects"
	"math"
	"github.com/wieku/danser/render"
	"log"
	"github.com/wieku/danser/settings"
)

type CircularMover struct {
	ca curves.Curve
	beginTime, endTime int64
	invert float64
}

func NewCircularMover() Mover {
	cm := &CircularMover{invert:-1}
	return cm
}

func (bm *CircularMover) Reset() {
	bm.invert = -1
}

func (bm *CircularMover) SetObjects(end, start objects.BaseObject) {
	endPos := end.GetBasicData().EndPos
	startPos := start.GetBasicData().StartPos
	bm.endTime = end.GetBasicData().EndTime
	bm.beginTime = start.GetBasicData().StartTime

	if settings.Dance.HalfCircle.StreamTrigger < 0 || (bm.beginTime - bm.endTime) < settings.Dance.HalfCircle.StreamTrigger {
		bm.invert = -1 * bm.invert
	}

	if endPos == startPos {
		bm.ca = curves.NewLinear(endPos, startPos)
		return
	}

	point := endPos.Mid(startPos)
	p := point.Sub(endPos).Rotate(bm.invert*math.Pi/2).Scl(settings.Dance.HalfCircle.RadiusMultiplier).Add(point)
	log.Println(point.Dst(endPos), p.Dst(point))
	bm.ca = curves.NewCirArc(endPos, p, startPos)
}

func (bm CircularMover) Update(time int64, cursor *render.Cursor) {
	cursor.SetPos(bm.ca.PointAt(float64(time - bm.endTime)/float64(bm.beginTime - bm.endTime)))
}

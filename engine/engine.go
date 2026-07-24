package engine

type Engine struct {
	copier Copier
}

func New(
	copier Copier,
) *Engine {

	return &Engine{
		copier: copier,
	}
}

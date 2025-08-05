package decoder

import "sync"

type DecoderModuleDeps struct{}

type DecoderModule struct {
	*DecoderService
}

func NewDecoderModule(deps DecoderModuleDeps) *DecoderModule {
	service := NewDecoderService(DecoderServiceDeps{})
	return &DecoderModule{
		DecoderService: service,
	}
}

type DecoderModuleConstructor func(DecoderModuleDeps) *DecoderModule

type DecoderServiceDeps struct{}

type DecoderService struct {
	typeCache    sync.Map // reflect.Type -> *StructMeta
	dbFieldCache sync.Map // reflect.Type -> map[string]string
}

func NewDecoderService(deps DecoderServiceDeps) *DecoderService {
	return &DecoderService{
		typeCache:    sync.Map{},
		dbFieldCache: sync.Map{},
	}
}

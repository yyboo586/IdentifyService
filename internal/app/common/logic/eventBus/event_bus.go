package eventBus

import (
	"IdentifyService/internal/app/common/service"

	"github.com/asaskevich/EventBus"
)

func init() {
	service.RegisterEventBus(EventBus.New())
}

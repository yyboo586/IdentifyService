package logic

import (
	_ "IdentifyService/internal/app/common/logic/cache"
	_ "IdentifyService/internal/app/common/logic/captcha"
	_ "IdentifyService/internal/app/common/logic/eventBus"
	_ "IdentifyService/internal/app/common/logic/middleware"
	_ "IdentifyService/internal/app/common/logic/snowIDGen"
	_ "IdentifyService/internal/app/common/logic/sysAttachment"
	_ "IdentifyService/internal/app/common/logic/sysConfig"
	_ "IdentifyService/internal/app/common/logic/sysDictData"
	_ "IdentifyService/internal/app/common/logic/sysDictType"
	_ "IdentifyService/internal/app/common/logic/upload"
)

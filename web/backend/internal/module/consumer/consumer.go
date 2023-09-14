package consumer

import (
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/global/mq"
	"envoy-go-fliter-hub/tools"
)

func Init() {
	logger := logs.NameSpace("Consumer").Sugar()
	tools.PanicIfErr(updateConsumer())

	go func() {
		for {
			mq.MQ.ListenUpdateSignal()

			err := updateConsumer()
			if err != nil {
				logger.Error("updateConsumer error: ", err)
				return
			}

		}
	}()

}

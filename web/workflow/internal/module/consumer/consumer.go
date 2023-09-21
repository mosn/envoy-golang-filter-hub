package consumer

import (
	"envoy-go-fliter-hub/internal/global/logs"
	"envoy-go-fliter-hub/internal/global/mq"
)

func Init() {
	logger := logs.NameSpace("Consumer").Sugar()
	err := updateConsumer()
	if err != nil {
		panic(err)
	}

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

package transfer

import "mnc-rest-api/internal/domain"

type TransferController struct {
	transferUsecase domain.TransferUsecase
}

func New(transferUsecase domain.TransferUsecase) TransferController {
	return TransferController{
		transferUsecase: transferUsecase,
	}
}

package responses

import (
	"github.com/MKW-Limitless-team/limitless-types/ltrc"
	"github.com/MKW-Limitless-team/limitless-types/wwfc"
)

type PlayerInfoResponse struct {
	Status     Status           `json:"status,omitempty"`
	PlayerData *ltrc.PlayerData `json:"playerdata,omitempty"`
	User       *wwfc.User       `json:"user,omitempty"`
	Message    string           `json:"message,omitempty"`
}

func FailureResponse(message string) *PlayerInfoResponse {
	return &PlayerInfoResponse{Status: Failure, Message: message}
}

func SuccessResponse(data PlayerInfoResponse) *PlayerInfoResponse {
	return &PlayerInfoResponse{Status: Success, PlayerData: data.PlayerData, User: data.User}
}

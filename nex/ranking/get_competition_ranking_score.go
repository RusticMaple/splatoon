package nex_ranking_splaton

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/splatoon"
	"github.com/PretendoNetwork/splatoon/globals"
	ranking_splatoon_types "github.com/PretendoNetwork/splatoon/nex/ranking/types"
)

func GetCompetitionRankingScore(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) {
	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	parameters := packet.RMCMessage().Parameters
	parametersStream := nex.NewByteStreamIn(parameters, globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	params := ranking_splatoon_types.NewCompetitionRankingGetParam()

	err = params.ExtractFrom(parametersStream)
	if err != nil {
		common_globals.Logger.Error("Failed to extract param on call to GetCompetitionRankingScore.")
		common_globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	// todo: actually implement this function
	retVal := types.NewList[ranking_splatoon_types.CompetitionRankingScoreInfo]()

	retVal.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = ranking.ProtocolID
	rmcResponse.MethodID = ranking.MethodGetCompetitionRankingScore
	rmcResponse.CallID = callID

	return rmcResponse, nil

}

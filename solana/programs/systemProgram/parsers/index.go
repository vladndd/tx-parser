package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func Router(result *types.ParsedResult, i int) (types.Action, error) {
	instruction := result.RawTx.Transaction.Message.Instructions[i]
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := binary.LittleEndian.Uint32(decode[:4])

	switch discriminator {
	case systemProgram.TransferDiscriminator:
		return TransferParser(result, i, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     systemProgram.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}

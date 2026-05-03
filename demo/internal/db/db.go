package db

import (
	"context"
	"fmt"
	"os"
	"sofa-demo/internal/calc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DATABASE_URL = "postgresql://postgres:dbpassword@host.docker.internal:5432/sofadb"

func CreateDBPool() (*pgxpool.Pool, error) {
	status := "No error"

	// Parse the url string into a config
	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		status = fmt.Sprintf("Error parsing URL: %v", err)
		return nil, err
	}

	// dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// }

	// Create a connection pool with the config
	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		status = fmt.Sprintf("Error creating connection pool: %v", err)
		fmt.Fprintf(os.Stderr, "%v\n", status)
		return nil, err
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		status = fmt.Sprintf("Error pinging DB: %v", err)
		fmt.Fprintf(os.Stderr, "%v\n", status)
		return nil, err
	}

	return dbPool, err
}

func CloseDBPool(pool *pgxpool.Pool) {
	pool.Close()
}

func InsertOneTrailerInput(pool *pgxpool.Pool, r []string) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO trailer_input VALUES ("+
			"$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, "+
			"$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, "+
			"$21, $22, $23, $24, $25, $26, $27, $28, $29, $30, "+
			"$31, $32, $33, $34, $35, $36, $37, $38, $39, $40, "+
			"$41, $42, $43, $44, $45, $46, $47, $48, $49, $50, "+
			"$51, $52, $53, $54, $55, $56, $57, $58, $59, $60, "+
			"$61, $62, $63, $64, $65, $66, $67, $68, $69, $70, "+
			"$71, $72, $73, $74, $75, $76, $77, $78, $79, $80, "+
			"$81, $82, $83, $84, $85, $86, $87, $88, $89, $90, "+
			"$91, $92, $93, $94, $95, $96, $97, $98, $99, $100, "+
			"$101, $102, $103, $104, $105, $106, $107, $108, $109, $110, "+
			"$111, $112, $113, $114, $115, $116, $117, $118, $119, $120, "+
			"$121)",
		r[0], r[1], r[2], r[3], r[4], r[5], r[6], r[7], r[8], r[9],
		r[10], r[11], r[12], r[13], r[14], r[15], r[16], r[17], r[18], r[19],
		r[20], r[21], r[22], r[23], r[24], r[25], r[26], r[27], r[28], r[29],
		r[30], r[31], r[32], r[33], r[34], r[35], r[36], r[37], r[38], r[39],
		r[40], r[41], r[42], r[43], r[44], r[45], r[46], r[47], r[48], r[49],
		r[50], r[51], r[52], r[53], r[54], r[55], r[56], r[57], r[58], r[59],
		r[60], r[61], r[62], r[63], r[64], r[65], r[66], r[67], r[68], r[69],
		r[70], r[71], r[72], r[73], r[74], r[75], r[76], r[77], r[78], r[79],
		r[80], r[81], r[82], r[83], r[84], r[85], r[86], r[87], r[88], r[89],
		r[90], r[91], r[92], r[93], r[94], r[95], r[96], r[97], r[98], r[99],
		r[100], r[101], r[102], r[103], r[104], r[105], r[106], r[107], r[108], r[109],
		r[110], r[111], r[112], r[113], r[114], r[115], r[116], r[117], r[118], r[119],
		r[120],
	)

	return err
}

func InsertOneTrailerLoad(pool *pgxpool.Pool, t *calc.TrailerLoad) error {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO trailer_loads VALUES ("+
			"$1, $2, $3, "+
			"$4, $5, $6, $7, $8, $9, "+
			"$10, $11, $12, $13, $14, $15, "+
			"$16, $17, $18, $19, $20, $21, "+
			"$22, $23, $24, $25, $26, $27, "+
			"$28, $29, $30, $31, $32, $33, $34, "+
			"$35, $36, $37, $38, $39, $40, "+
			"$41, $42, $43, $44, $45, $46, "+
			"$47, $48, $49, $50, $51, $52, "+
			"$53, $54, $55, $56, $57, $58, "+
			"$59, $60, $61, $62, $63, $64, "+
			"$65, $66, $67, $68, "+
			"$69, $70, $71, $72, "+
			"$73, $74, $75, $76, $77, $78, $79, "+
			"$80, $81, $82, $83, $84, $85, "+
			"$86, $87, $88, $89, $90, $91, "+
			"$92, $93, $94, $95, "+
			"$96, $97, $98, $99, $100, $101, $102, "+
			"$103, $104, $105, $106, $107, $108, $109, "+
			"$110, $111, $112, $113, $114, "+
			"$115, $116, $117, $118, $119)",
		t.Timestamp,        // 1
		t.VIN,              // 2
		t.StatusLoaded,     // 3
		t.FAirAxlex1LN,     // 4
		t.FAirAxlex1RN,     // 5
		t.FAirAxlex2LN,     // 6
		t.FAirAxlex2RN,     // 7
		t.FAirAxlex3LN,     // 8
		t.FAirAxlex3RN,     // 9
		t.KAx1LNPerm,       // 10
		t.KAx1RNPerm,       // 11
		t.KAx2LNPerm,       // 12
		t.KAx2RNPerm,       // 13
		t.KAx3LNPerm,       // 14
		t.KAx3RNPerm,       // 15
		t.FInertiaAx1LN,    // 16
		t.FInertiaAx1RN,    // 17
		t.FInertiaAx2LN,    // 18
		t.FInertiaAx2RN,    // 19
		t.FInertiaAx3LN,    // 20
		t.FInertiaAx3RN,    // 21
		t.XAx1Lm,           // 22
		t.XAx1Rm,           // 23
		t.XAx2Lm,           // 24
		t.XAx2Rm,           // 25
		t.XAx3Lm,           // 26
		t.XAx3Rm,           // 27
		t.DeltaT,           // 28
		t.XDotAx1LMps,      // 29
		t.XDotAx1RMps,      // 30
		t.XDotAx2LMps,      // 31
		t.XDotAx2RMps,      // 32
		t.XDotAx3LMps,      // 33
		t.XDotAx3RMps,      // 34
		t.XDotFAx1LMps,     // 35
		t.XDotFAx1RMps,     // 36
		t.XDotFAx2LMps,     // 37
		t.XDotFAx2RMps,     // 38
		t.XDotFAx3LMps,     // 39
		t.XDotFAx3RMps,     // 40
		t.FdAx1LN,          // 41
		t.FdAx1RN,          // 42
		t.FdAx2LN,          // 43
		t.FdAx2RN,          // 44
		t.FdAx3LN,          // 45
		t.FdAx3RN,          // 46
		t.FzAx1LN,          // 47
		t.FzAx1RN,          // 48
		t.FzAx2LN,          // 49
		t.FzAx2RN,          // 50
		t.FzAx3LN,          // 51
		t.FzAx3RN,          // 52
		t.FzTotalN,         // 53
		t.FzLTotalN,        // 54
		t.FzRTotalN,        // 55
		t.FzAx1TotalN,      // 56
		t.FzAx2TotalN,      // 57
		t.FzAx3TotalN,      // 58
		t.FzAx1AbsLN,       // 59
		t.FzAx1AbsRN,       // 60
		t.FzAx2AbsLN,       // 61
		t.FzAx2AbsRN,       // 62
		t.FzAx3AbsLN,       // 63
		t.FzAx3AbsRN,       // 64
		t.LTRGroup,         // 65
		t.DeltaFzLatAx1,    // 66
		t.DeltaFzLatAx2,    // 67
		t.DeltaFzLatAx3,    // 68
		t.UMps,             // 69
		t.RKpRps,           // 70
		t.AyG,              // 71
		t.MRollNm,          // 72
		t.BPAxBar,          // 73
		t.BPAx1LBar,        // 74
		t.BPAx1RBar,        // 75
		t.BPAx2LBar,        // 76
		t.BPAx2RBar,        // 77
		t.BPAx3LBar,        // 78
		t.BPAx3RBar,        // 79
		t.TWheelAx1L,       // 80
		t.TWheelAx1R,       // 81
		t.TWheelAx2L,       // 82
		t.TWheelAx2R,       // 83
		t.TWheelAx3L,       // 84
		t.TWheelAx3R,       // 85
		t.FxBrakeAx1L,      // 86
		t.FxBrakeAx1R,      // 87
		t.FxBrakeAx2L,      // 88
		t.FxBrakeAx2R,      // 89
		t.FxBrakeAx3L,      // 90
		t.FxBrakeAx3R,      // 91
		t.FxBrakeAx1Total,  // 92
		t.FxBrakeAx2Total,  // 93
		t.FxBrakeAx3Total,  // 94
		t.FxBrakeTotal,     // 95
		t.BetaBodyRad,      // 96
		t.AlphaAx1LRad,     // 97
		t.AlphaAx1RRad,     // 98
		t.AlphaAx2LRad,     // 99
		t.AlphaAx2RRad,     // 100
		t.AlphaAx3LRad,     // 101
		t.AlphaAx3RRad,     // 102
		t.FyAx1LN,          // 103
		t.FyAx1RN,          // 104
		t.FyAx2LN,          // 105
		t.FyAx2RN,          // 106
		t.FyAx3LN,          // 107
		t.FyAx3RN,          // 108
		t.FyTotalN,         // 109
		t.MzYawNm,          // 110
		t.MassEffective,    // 111
		t.FxKingpinN,       // 112
		t.FyKingpinN,       // 113
		t.FzKingpinN,       // 114
		t.MassEstRearkg,    // 115
		t.RollTransferGain, // 116
		t.DeflectBiasAx1,   // 117
		t.DeflectBiasAx2,   // 118
		t.DeflectBiasAx3,   // 119
	)

	return err
}

func GetTrailerInfo(pool *pgxpool.Pool) map[string]calc.Trailer {
	// Slice of trailer VINs
	trailerVINs := []string{
		"VIN_FMC_0004",
		"VIN_FMC_0005",
		"VIN_FMC_0006",
		"VIN_FMC_0007",
		"VIN_FMC_TEST",
	}

	trailers := make(map[string]calc.Trailer)

	for _, t := range trailerVINs {
		// Query the trailers table to get the current trailer
		rows, _ := pool.Query(context.Background(), "SELECT * FROM trailers WHERE VIN=$1", t)
		trailer, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[calc.Trailer])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Get trailer failed: %v\n", err)
		}

		trailers[t] = trailer
	}
	return trailers
}

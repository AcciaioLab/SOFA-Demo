package db

import (
	"context"
	"fmt"
	"os"

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

func InsertOneRecord(pool *pgxpool.Pool, r []string) {
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
	if err != nil {
		fmt.Printf("Error inserting record into DB: %v.\n", err)
	}
}

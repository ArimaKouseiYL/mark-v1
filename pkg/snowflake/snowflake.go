package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func GetSnowFlakeId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalln(err)
	}
	return node.Generate().Int64()
}

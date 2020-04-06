package pingcapsqlparsersample

import (
	"fmt"
	"github.com/pingcap/parser"
	"testing"
	_ "github.com/pingcap/parser/test_driver"
	//_ "github.com/pingcap/tidb/types/parser_driver"
)

func TestParser(t *testing.T) {

	p := parser.New()
	stmtNodes, _, err := p.Parse("select * from tbl where id = 1", "", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range stmtNodes {
		fmt.Println(node)

	}
	ttt := fmt.Sprintf("这是一个：%s", "ceshi")
	fmt.Println(ttt)
	//fmt.Println(stmtNodes[0], err)

}

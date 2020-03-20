package tests

import (
	"github.com/sadlil/gologger"
	"testing"
)

var logger gologger.GoLogger

func init()  {

	logger = gologger.GetLogger(gologger.ELASTICSEARCH, "http://172.16.5.65:9200/golog")
	logger.Info("123")

}

func TestLogger(t *testing.T)  {
	logger.Info("456")
}
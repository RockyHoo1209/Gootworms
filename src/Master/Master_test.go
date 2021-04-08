package Master

import (
	"main/src/Utils/DBUtils/MongoUtil"
	"testing"
)


func TestMaster(*testing.T){
	MongoUtil.InitMongo()
	InitMaster().RunMaster()
}
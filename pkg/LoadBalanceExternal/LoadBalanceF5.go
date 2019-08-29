package LoadBalanceExternal


type F5VIP struct {
	PoolName string `json:"PoolName"  bson:"PoolName"`
	F5PoolMembers F5PoolMembers
}

type F5PoolMembers struct {
	PoolMemberName string `json:"PoolMemberName"  bson:"PoolMemberName"`
	PoolMemberDescription string `json:"PoolMemberDesc"  bson:"PoolMemberDesc"`


}



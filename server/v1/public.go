package v1

import (
	"../gctfConfig"
	"../model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
[
{
"score": "15",
"username": "cug2"
},
{
"score": "10",
"username": "cug1"
}
]
 */
func GetUsersRank(c *gin.Context) {
	data, err := model.GctfDataManage.Query("select `username`,`score` from gctf_user limit 50 ")
	var userName_scores []map[string]string

	for x := range data {
		userName_score := make(map[string]string)
		userName_score["username"] = string(data[x]["username"])
		userName_score["score"] = string(data[x]["score"])
		userName_scores = append(userName_scores, userName_score)
	}
	//d,_:=json.Marshal(userName_scores)
	c.JSON(http.StatusOK, userName_scores)
	if gctfConfig.GCTF_DEBUG && err != nil {
		log.Fatal("error to get userscore", data)
	}

}

func GetTeamsRank(c *gin.Context) {
}

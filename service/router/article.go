package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/iceEvening/me/service/common"
	"github.com/iceEvening/me/service/middleware/jwt"
)

//RegisterArticleRouter is a func which register article router
func (r *Router) RegisterArticleRouter() {
	article := r.E.Group("/article")
	article.GET("/:id", r.article)
	article.Use(jwt.JWTAuth())
	{
		article.POST("/new", r.newArticle)
		article.POST("/edit", r.editArticle)
		article.POST("/delete", r.delArticle)
	}

	articles := r.E.Group("/articles")
	articles.GET("/:id", r.articles)
}

type newArticleReq struct {
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
	HTML     string `json:"html"`
}

func (r *Router) newArticle(c *gin.Context) {
	var na newArticleReq
	if c.BindJSON(&na) == nil {
		logrus.Debugln(na)
		CheckEditAuth(c, uint(Atoi(c, na.UserID)))
		err := r.model.NewArticle(uint(Atoi(c, na.UserID)), na.Nickname, na.Title, na.Markdown, na.HTML)
		if err != nil {
			logrus.Errorf("new article error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("new article  error: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "new article success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

type editArticleReq struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
	HTML     string `json:"html"`
}

func (r *Router) editArticle(c *gin.Context) {
	var ear editArticleReq
	if c.BindJSON(&ear) == nil {
		logrus.Debugln(ear)
		CheckEditAuth(c, uint(Atoi(c, ear.UserID)))
		err := r.model.EditArticle(uint(Atoi(c, ear.UserID)), uint(Atoi(c, ear.ID)), ear.Nickname, ear.Title, ear.Markdown, ear.HTML)
		if err != nil {
			logrus.Errorf("edit article error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("edit article  error: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "edit article success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

type delArticleReq struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (r *Router) delArticle(c *gin.Context) {
	var dar delArticleReq
	if c.BindJSON(&dar) == nil {
		logrus.Debugln(dar)
		CheckEditAuth(c, uint(Atoi(c, dar.UserID)))
		err := r.model.DelArticle(uint(Atoi(c, dar.UserID)), uint(Atoi(c, dar.ID)))
		if err != nil {
			logrus.Errorf("delete article error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("delete article  error: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "delete article success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

func (r *Router) article(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		logrus.Debugln(id)
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "invalid article id",
			})
		}

		article, err := r.model.Article(uint(idInt))
		if err != nil {
			logrus.Errorf("article failed: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("failed to display article: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      0,
			"msg":         "article success",
			"nickname":    article.NickName,
			"title":       article.Title,
			"html":        article.HTML,
			"markdown":    article.Markdown,
			"create_time": article.CreatedAt.Format(common.ANSIC),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "no article id",
		})
	}
}

func (r *Router) articles(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		logrus.Debugln(id)
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "invalid user id",
			})
		}

		articles, err := r.model.Articles(uint(idInt))
		if err != nil {
			logrus.Errorf("get articles failed: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("user id not found: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"msg":      "article success",
			"articles": *articles,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "no user id",
		})
	}
}

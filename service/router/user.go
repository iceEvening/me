package router

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/iceEvening/me/service/middleware/jwt"
)

//RegisterUserRouter is a func which include user's APIs
func (r *Router) RegisterUserRouter() {
	r.E.POST("/signup", r.signUp)
	r.E.POST("/signin", r.signIn)

	user := r.E.Group("/user")
	user.GET("careers/:id", r.getCareers)
	user.GET("educations/:id", r.getEducations)
	user.GET("profile/:id", r.getUser)
	user.Use(jwt.JWTAuth())
	{
		user.GET("ping", func(c *gin.Context) {
			c.String(http.StatusOK, getWhat())
		})
		user.POST("careers", r.editCareers)
		user.POST("delcareer", r.delCareer)
		user.POST("educations", r.editEducations)
		user.POST("deleducation", r.delEducation)
		user.POST("image", r.editUserImg)
		user.POST("edituser", r.editUser)
	}
}

type signUpReq struct {
	Email    string `json:"email"`
	NickName string `json:"nickname"`
	Passwd   string `json:"password"`
}

func (r *Router) signUp(c *gin.Context) {
	var su signUpReq
	if c.BindJSON(&su) == nil {
		logrus.Debugln(su)
		err := r.model.SignUp(su.Email, su.NickName, su.Passwd)
		if err != nil {
			logrus.Errorf("sign up error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("sign up error: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"nickname": su.NickName,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Sign up finished ===============

type signInReq struct {
	Email  string `json:"email"`
	Passwd string `json:"password"`
}

type signInResp struct {
	email string
	token string
}

func (r *Router) signIn(c *gin.Context) {
	var si signInReq
	if c.BindJSON(&si) == nil {
		logrus.Debugln(si)
		user, err := r.model.SignIn(si.Email, si.Passwd)
		if err != nil {
			logrus.Errorf("sign in error: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("sign in error: %v", err),
			})
			return
		}
		token, err := generateToken(c, user.ID)
		if err != nil {
			logrus.Errorf("sign in error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "sign in error: unknow error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"msg":      "sign in success",
			"email":    user.Email,
			"user_id":  user.ID,
			"nickname": user.Nickname,
			"token":    token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Sign in finished ===============

type editCareer struct {
	ID          string    `json:"id"`
	Company     string    `json:"company"`
	Department  string    `json:"department"`
	Post        string    `json:"post"`
	Rank        string    `json:"rank"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	IsCurrent   bool      `json:"isCurrent"`
}

type editCareersReq struct {
	ID   string       `json:"id"`
	List []editCareer `json:"list"`
}

func (r *Router) editCareers(c *gin.Context) {
	var ecr editCareersReq
	if c.BindJSON(&ecr) == nil {
		logrus.Debugln(ecr)
		CheckEditAuth(c, uint(Atoi(c, ecr.ID)))
		var err error
		for _, ec := range ecr.List {
			err = r.model.EditCareers(
				uint(Atoi(c, ecr.ID)),
				uint(Atoi(c, ec.ID)),
				ec.Company,
				ec.Department,
				ec.Post,
				ec.Rank,
				ec.Description,
				ec.Start,
				ec.End,
				ec.IsCurrent,
			)
		}
		if err != nil {
			logrus.Errorf("edit careers error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("failed to edit careers: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "edit careers success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Edit careers finished ===============

type delCareerReq struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (r *Router) delCareer(c *gin.Context) {
	var dcr delCareerReq
	if c.BindJSON(&dcr) == nil {
		logrus.Debugln(dcr)
		CheckEditAuth(c, uint(Atoi(c, dcr.UserID)))
		err := r.model.DelCareer(uint(Atoi(c, dcr.UserID)), uint(Atoi(c, dcr.ID)))
		if err != nil {
			logrus.Errorf("del career error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("failed to del career: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "del career success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Delete career finished ===============

func (r *Router) getCareers(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		logrus.Debugln(id)
		idInt := Atoi(c, id)
		careers, err := r.model.GetCareers(int64(idInt))
		if err != nil {
			logrus.Errorf("get careers failed: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("failed to get careers: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"msg":     "get careers success",
			"careers": *careers,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "invalid user id",
		})
	}
}

//=============== Get careers finished ===============

type editEducation struct {
	ID        string    `json:"id"`
	School    string    `json:"school"`
	Major     string    `json:"major"`
	Degree    string    `json:"degree"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	IsCurrent bool      `json:"isCurrent"`
}

type editEducationsReq struct {
	ID   string          `json:"id"`
	List []editEducation `json:"list"`
}

func (r *Router) editEducations(c *gin.Context) {
	var eer editEducationsReq
	if c.BindJSON(&eer) == nil {
		logrus.Debugln(eer)
		CheckEditAuth(c, uint(Atoi(c, eer.ID)))
		var err error
		for _, ee := range eer.List {
			err = r.model.EditEducations(
				uint(Atoi(c, eer.ID)),
				uint(Atoi(c, ee.ID)),
				ee.School,
				ee.Major,
				ee.Degree,
				ee.Start,
				ee.End,
				ee.IsCurrent,
			)
		}
		if err != nil {
			logrus.Errorf("edit educations error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("failed to edit educations: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "edit educations success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Edit educations finished ===============

type delEducationReq struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}

func (r *Router) delEducation(c *gin.Context) {
	var der delEducationReq
	if c.BindJSON(&der) == nil {
		logrus.Debugln(der)
		CheckEditAuth(c, uint(Atoi(c, der.UserID)))
		err := r.model.DelEducation(uint(Atoi(c, der.UserID)), uint(Atoi(c, der.ID)))
		if err != nil {
			logrus.Errorf("del education error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("failed to del education: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "del education success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Delete educations finished ===============

func (r *Router) getEducations(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		logrus.Debugln(id)
		idInt := Atoi(c, id)
		educations, err := r.model.GetEducations(int64(idInt))
		if err != nil {
			logrus.Errorf("get educations failed: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("failed to get educations: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":     0,
			"msg":        "get educations success",
			"educations": *educations,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "invalid user id",
		})
	}
}

//=============== Get educations finished ===============

func (r *Router) editUserImg(c *gin.Context) {
	idInt := Atoi(c, c.PostForm("id"))
	CheckEditAuth(c, uint(idInt))
	img, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    fmt.Errorf("invalid image: %v", err),
		})
		return
	}
	f, err := img.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    fmt.Errorf("read image failed: %v", err),
		})
		return
	}
	imgByte := make([]byte, img.Size)
	for {
		_, err := f.Read(imgByte)
		if err == io.EOF {
			break
		}
	}

	err = r.model.EditUser(uint(idInt), imgByte, "", "", "", 0, "", time.Now(), "", "")
	if err != nil {
		logrus.Errorf("edit image error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    fmt.Errorf("failed to edit image: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "edit image success",
	})
}

//=============== Edit user image finished ===============

func (r *Router) getUser(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		logrus.Debugln(id)
		idInt := Atoi(c, id)
		user, err := r.model.GetUser(uint(idInt))
		if err != nil {
			logrus.Errorf("get user profile failed: %v", err)
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    fmt.Sprintf("failed to get user profile: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   0,
			"msg":      "get user profile success",
			"img":      user.Img,
			"me":       user.Me,
			"nickname": user.Nickname,
			"name":     user.Name,
			"email":    user.Email,
			"gender":   user.Gender,
			"birthday": user.Birthday,
			"hometown": user.Hometown,
			"city":     user.City,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "invalid user id",
		})
	}
}

//=============== Edit get profile finished ===============

type editUserReq struct {
	ID       string    `json:"id"`
	Nickname string    `json:"nickname"`
	Me       string    `json:"me"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Hometown string    `json:"hometown"`
	City     string    `json:"city"`
}

func (r *Router) editUser(c *gin.Context) {
	var eur editUserReq
	if c.BindJSON(&eur) == nil {
		logrus.Debugln(eur)
		CheckEditAuth(c, uint(Atoi(c, eur.ID)))
		var gender uint
		if eur.Gender != "" {
			gender = uint(Atoi(c, eur.Gender))
		}
		var err error
		err = r.model.EditUser(
			uint(Atoi(c, eur.ID)),
			nil,
			eur.Nickname,
			eur.Me,
			eur.Name,
			gender,
			eur.Email,
			eur.Birthday,
			eur.Hometown,
			eur.City,
		)
		if err != nil {
			logrus.Errorf("edit user failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    fmt.Errorf("failed to edit user: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "edit user success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
			"msg":    "fail to analysis form data",
		})
	}
}

//=============== Edit user profile finished ===============

func getWhat() string {
	return "What happend?"
}

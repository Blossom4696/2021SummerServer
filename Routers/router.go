package Routers

import (
	"github.com/bigby/project/Config"
	"github.com/bigby/project/Controllers"
	"github.com/bigby/project/Middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	// 使用 session(cookie-based)
	router.Use(sessions.Sessions("LoginSessionKey", store))

	// 习题路由组
	v1 := router.Group("app/exercise")
	v1.Use(Middlewares.JWT(Config.TeacherType)) //需要校验权限大于教师级别

	v1.GET("/get_all_exercise", Controllers.ExerciseQuery)
	v1.POST("/insert_exercise", Controllers.ExerciseInsert)
	v1.PUT("/update_exercise/:id", Controllers.ExerciseUpdate)
	v1.DELETE("/delete_exercise/:id", Controllers.ExerciseDelete)
	v1.GET("/get_exercise", Controllers.ExerciseQueryByEid)

	v1.GET("/get_sorted_exercise", Controllers.ExerciseQueryBySorted)

	// 学生路由组
	v2 := router.Group("app/student")

	v2.GET("/get_student", Controllers.StudentQueryFromStudent)
	v2.POST("/register", Controllers.StudentRegister)
	v2.POST("/update_student", Controllers.StudentUpdateFromStudent)
	v2.POST("/insert_order_teacher", Controllers.OrderTeacherInsertFromStudent)
	v2.GET("/get_learn_situation_list", Controllers.LearnSituationQueryFromStudent)
	v2.POST("/insert_ask_question", Controllers.AskQuestionInsertFromStudent)

	// 文件路由组
	v3 := router.Group("app/file")

	v3.GET("/get_image", Controllers.ImageDownload)
	v3.GET("/get_audio", Controllers.AudioDownload)
	v3.GET("/get_pdf", Controllers.PdfDownload)
	v3.POST("/upload_file", Controllers.FileUpload)

	// 校区路由组
	v4 := router.Group("/app/campus")
	v4.GET("/get_all_campus", Controllers.CampusQuery)

	// 教师路由组
	v5 := router.Group("/app/teacher")
	v5.Use(Middlewares.JWT(Config.TeacherType)) //需要校验权限大于教师级别

	v5.GET("/get_student_list", Controllers.StudentListQueryFromTeacher)
	v5.GET("/get_student_sorted_list", Controllers.StudentListQuerySortedFromTeacher)
	v5.GET("/get_student", Controllers.StudentQueryFromTeacher)
	v5.POST("/insert_today_exercise", Controllers.TodayExerciseInsertFromTeacher)
	v5.GET("/get_teacher", Controllers.TeacherQueryFromTeacher)
	v5.POST("/update_teacher", Controllers.TeacherUpdateFromTeacher)
	v5.GET("/get_learn_situation_list", Controllers.LearnSituationQueryFromTeacher)
	v5.POST("/insert_learn_situation", Controllers.LearnSituationInsertFromTeacher)
	v5.PUT("/update_learn_situation", Controllers.LearnSituationUpdateFromTeacher)
	v5.DELETE("/delete_learn_situation/:id", Controllers.LearnSituationDeleteFromTeacher)
	v5.GET("/get_ask_question_list", Controllers.AskQuestionListQueryFromTeacher)
	v5.GET("/get_ask_question", Controllers.AskQuestionQueryFromTeacher)
	v5.GET("/generate_wrong_problem_pdf", Controllers.GenerateWPPdfFromTeacher)
	v5.GET("/generate_exercise_pdf", Controllers.GenerateExercisePdfFromTeacher)

	// 今日习题路由组
	v6 := router.Group("/app/today")
	v6.GET("/get_today_exercise", Controllers.TodayExerciseQuery)
	v6.GET("/judge_today_exercise", Controllers.JudgeTodayExercise)

	// 登录路由组
	v7 := router.Group("app/login/")
	v7.GET("teacher_login", Controllers.TeacherLogin)
	v7.GET("student_login", Controllers.StudentLogin)
	v7.GET("admin_login", Controllers.AdminLogin)

	// 错题路由组
	v8 := router.Group("app/wrong_problem/")
	v8.GET("get_wrong_problem", Controllers.WrongProblemQueryByWid)
	v8.GET("get_wrong_problem_list", Controllers.WrongProblemQuery)
	v8.GET("get_sorted_wrong_problem", Controllers.WrongProblemQueryBySorted)
	v8.POST("insert_wrong_problem", Controllers.WrongProblemInsert)
	v8.DELETE("delete_wrong_problem/:id", Controllers.WrongProblemDelete)

	// 用户公用的方法
	v9 := router.Group("app/user/common")
	v9.Use(Middlewares.JWT(Config.StudentType)) //需要校验权限大于学生级别
	v9.GET("/get_exercise_field", Controllers.ExerciseGetField)
	v9.GET("/search_student_from_admin", Controllers.SearchStudentFromAdmin)
	v9.GET("/search_student_from_teacher", Controllers.SearchStudentFromTeacher)
	v9.GET("/search_teacher", Controllers.SearchTeacher)
	v9.GET("/search_exercise", Controllers.SearchExercise)
	v9.GET("/search_wrong_problem", Controllers.SearchWrongProblem)

	// 测试
	v10 := router.Group("app/test")
	v10.GET("/login", Controllers.LoginTest)

	// 管理员
	vadmin := router.Group("app/admin")
	vadmin.Use(Middlewares.JWT(Config.AdminType)) //需要校验权限大于教师级别
	vadmin.GET("/get_student_list", Controllers.StudentListQueryFromAdmin)
	vadmin.GET("/get_student_sorted_list", Controllers.StudentListQuerySortedFromAdmin)
	vadmin.GET("/get_student", Controllers.StudentQueryFromAdmin)
	vadmin.GET("/get_teacher_list", Controllers.TeacherListQueryFromAdmin)
	vadmin.GET("/get_teacher_sorted_list", Controllers.TeacherListQuerySortedFromAdmin)
	vadmin.GET("/get_teacher", Controllers.TeacherQueryFromAdmin)
	vadmin.GET("/get_campus_list", Controllers.CampusListQueryFromAdmin)
	vadmin.POST("/insert_campus", Controllers.CampusInsertFromAdmin)
	vadmin.PUT("/update_campus", Controllers.CampusUpdateFromAdmin)
	vadmin.DELETE("/delete_campus/:id", Controllers.CampusDeleteFromAdmin)
	vadmin.POST("/insert_student", Controllers.StudentInsertFromAdmin)
	vadmin.POST("/insert_teacher", Controllers.TeacherInsertFromAdmin)
	vadmin.POST("/update_student", Controllers.StudentUpdateFromAdmin)
	vadmin.POST("/update_teacher", Controllers.TeacherUpdateFromAdmin)
	vadmin.GET("/get_order_teacher_list", Controllers.OrderTeacherListQueryByAdmin)
	vadmin.GET("/get_order_teacher", Controllers.OrderTeacherQueryByAdmin)

	// 添加swagger的路由
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

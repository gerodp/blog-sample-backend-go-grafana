package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/gerodp/simpleBlogApp/controller"
	"github.com/gerodp/simpleBlogApp/model"
	"github.com/gerodp/simpleBlogApp/repository"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func buildHandler(promRespTimeDuration *prometheus.HistogramVec, method string, path string, handler func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		handler(c)
		duration := time.Since(start).Seconds()
		promRespTimeDuration.With(prometheus.Labels{"method": method, "path": path}).Observe(duration)
	}
}

func createTestUser(userRepo model.UserRepository) error {
	_, err := userRepo.FindByUsername("testint1")

	if err == nil {
		return nil
	} else {
		user := model.User{
			ID:        1,
			Username:  "testint1",
			Password:  "$2a$14$oemuupbL/xOA3d.jS3CBOeLf0SfCt.cWrqWBqr4jC71CxuF.mrrCa",
			Email:     "test@integration.int",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		_, err2 := userRepo.Save(&user)
		return err2
	}
}

func createTestPosts(postCount int, postRepo model.PostRepository) error {

	posts, err := postRepo.Find()
	if err == nil && posts != nil && len(posts) > 0 {
		return nil
	} else {
		for i := 0; i < postCount; i++ {
			post := model.Post{
				Title:     fmt.Sprintf("The wood and the Mountain %d", i),
				Body:      "When half way through the journey of our life I found that I was in a gloomy wood, because the path which led aright was lost. And ah, how hard it is to say just what this wild and rough and stubborn woodland was, the very thought of which renews my fear! So bitter ’t is, that death is little worse; but of the good to treat which there I found, I ’ll speak of what I else discovered there. I cannot well say how I entered it, so full of slumber was I at the moment when I forsook the pathway of the truth; but after I had reached a mountain’s foot, where that vale ended which had pierced my heart with fear, I looked on high, and saw its shoulders mantled already with that planet’s rays which leadeth one aright o’er every path. Then quieted a little was the fear, which in the lake-depths of my heart had lasted throughout the night I passed so piteously. And even as he who, from the deep emerged with sorely troubled breath upon the shore, turns round, and gazes at the dangerous water; even so my mind, which still was fleeing on, turned back to look again upon the pass which ne’er permitted any one to live. When I had somewhat eased my weary body, o’er the lone slope I so resumed my way, that e’er the lower was my steady foot. Then lo, not far from where the ascent began, a Leopard which, exceeding light and swift, was covered over with a spotted hide, and from my presence did not move away; nay, rather, she so hindered my advance, that more than once I turned me to go back. Some time had now from early morn elapsed, and with those very stars the sun was rising that in his escort were, when Love Divine in the beginning moved those beauteous things; I therefore had as cause for hoping well of that wild beast with gaily mottled skin, the hour of daytime and the year’s sweet season; but not so, that I should not fear the sight, which next appeared before me, of a Lion, — against me this one seemed to be advancing with head erect and with such raging hunger, that even the air seemed terrified thereby — and of a she-Wolf, which with every lust seemed in her leanness laden, and had caused many ere now to lead unhappy lives. The latter so oppressed me with the fear that issued from her aspect, that I lost the hope I had of winning to the top. And such as he is, who is glad to gain, and who, when times arrive that make him lose, weeps and is saddened in his every thought; such did that peaceless animal make me, which, ’gainst me coming, pushed me, step by step, back to the place where silent is the sun. While toward the lowland I was falling fast, the sight of one was offered to mine eyes, who seemed, through long continued silence, weak. When him in that vast wilderness I saw, “Have pity on me,” I cried out to him, “whate’er thou be, or shade, or very man!” “Not man,” he answered, “I was once a man; and both my parents were of Lombardy, and Mantuans with respect to fatherland. ’Neath Julius was I born, though somewhat late, and under good Augustus’ rule I lived in Rome, in days of false and lying gods. I was a poet, and of that just man, Anchises’ son, I sang, who came from Troy after proud Ilion had been consumed. But thou, to such sore trouble why return? Why climbst thou not the Mountain of Delight, which is of every joy the source and cause?” “Art thou that Virgil, then, that fountain-head which poureth forth so broad a stream of speech?” I answered him with shame upon my brow. “O light and glory of the other poets, let the long study, and the ardent love which made me con thy book, avail me now. Thou art my teacher and authority; thou only art the one from whom I took the lovely manner which hath done me honor. Behold the beast on whose account I turned; from her protect me, O thou famous Sage, for she makes both my veins and pulses tremble!” “A different course from this must thou pursue,” he answered, when he saw me shedding tears, “if from this wilderness thou wouldst escape; for this wild beast, on whose account thou criest, alloweth none to pass along her way, but hinders him so greatly, that she kills; and is by nature so malign and guilty, that never doth she sate her greedy lust, but after food is hungrier than before. Many are the animals with which she mates, and still more will there be, until the Hound shall come, and bring her to a painful death. He shall not feed on either land or wealth, but wisdom, love and power shall be his food, and ’tween two Feltros shall his birth take place. Of that low Italy he ’ll be the savior, for which the maid Camilla died of wounds, with Turnus, Nisus and Eurỳalus. And he shall drive her out of every town, till he have put her back again in Hell, from which the earliest envy sent her forth. I therefore think and judge it best for thee to follow me; and I shall be thy guide, and lead thee hence through an eternal place, where thou shalt hear the shrieks of hopelessness of those tormented spirits of old times, each one of whom bewails the second death; then those shalt thou behold who, though in fire, contented are, because they hope to come, whene’er it be, unto the blessèd folk; to whom, thereafter, if thou wouldst ascend, there ’ll be for that a worthier soul than I. With her at my departure I shall leave thee, because the Emperor who rules up there, since I was not obedient to His law, wills none shall come into His town through me. He rules as emperor everywhere, and there as king; there is His town and lofty throne. O happy he whom He thereto elects!” And I to him: “O Poet, I beseech thee, even by the God it was not thine to know, so may I from this ill and worse escape, conduct me thither where thou saidst just now, that I may see Saint Peter’s Gate, and those whom thou describest as so whelmed with woe.” He then moved on, and I behind him kept.",
				AuthorID:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			_, err2 := postRepo.Save(&post)

			if err2 != nil {
				return err2
			}
		}
		return nil
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // include the timestamp in the output

	gin.DefaultWriter = log.Writer()
	gin.DefaultErrorWriter = log.Writer()
	log.Println("Backend server started!")

	promRespTimeDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "A histogram of the HTTP request durations in seconds.",
		Buckets: prometheus.ExponentialBuckets(0.001, 2, 6),
	}, []string{"method", "path"})

	prometheus.MustRegister(promRespTimeDuration)

	addr := fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))
	log.Printf("DB address =%s\n", addr)

	repo, err := repository.NewRepository(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), addr, os.Getenv("DATABASE_NAME"))

	if err != nil {
		log.Fatalln("Error creating repository", err)
		os.Exit(1)
	}

	//Populates the DB with the Test Data
	errUser := createTestUser(repo.Users)

	if errUser != nil {
		log.Fatalln("Error creating test user", err)
		os.Exit(1)
	}

	errPosts := createTestPosts(10, repo.Posts)

	if errPosts != nil {
		log.Fatalln("Error creating test posts", err)
		os.Exit(1)
	}
	//End popullate

	userCon := controller.NewUserController(repo.Users)

	postCon := controller.NewPostController(repo.Posts)

	authMdw, amErr := controller.NewAuthMiddleware(repo.Users)

	if amErr != nil {
		log.Fatalln("Error creating Auth Middleware. Exiting...")
		os.Exit(1)
	}

	r := gin.Default()
	pprof.Register(r)

	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	r.POST("/login", buildHandler(promRespTimeDuration, "POST", "/login", authMdw.LoginHandler))

	auth := r.Group("/auth")

	r.POST("/refresh_token", buildHandler(promRespTimeDuration, "POST", "/refresh_token", authMdw.RefreshHandler))

	auth.Use(authMdw.MiddlewareFunc())
	{
		auth.POST("/user", buildHandler(promRespTimeDuration, "POST", "/auth/user", userCon.CreateUser))

		auth.GET("/user", buildHandler(promRespTimeDuration, "GET", "/auth/user", userCon.Find))

		auth.POST("/post", buildHandler(promRespTimeDuration, "POST", "/auth/post", postCon.CreatePost))

		auth.GET("/post", buildHandler(promRespTimeDuration, "GET", "/auth/post", postCon.Find))

		auth.DELETE("/post/:id", buildHandler(promRespTimeDuration, "DELETE", "auth/post/:id", postCon.DeletePost))
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint not found"})
	})

	s := &http.Server{
		Addr:     os.Getenv("BACKEND_ADDRESS"),
		Handler:  r,
		ErrorLog: log.Default(),
	}
	s.ListenAndServe()
}

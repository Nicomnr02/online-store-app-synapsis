package main

import (
	"fmt"
	"log"
	"net/http"
	"online_app_store/api"
	"online_app_store/middleware"
	"online_app_store/repositories"
	"online_app_store/services"
	"online_app_store/utils"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	os.Setenv("DATABASE_URL", "postgres://syn_cha:1@localhost:5432/app-store-db-synapsis")

	err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}
	db := utils.GetDBConnection()
	log.Default().Println(db) //! print

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		mux := http.NewServeMux()

		mux = RunServer(db, mux)

		logrus.Info("Server is running on port 8000")
		err = http.ListenAndServe(":8000", mux)
		if err != nil {

			panic(err)
		}

	}()

	wg.Wait()

}

func RunServer(db *gorm.DB, mux *http.ServeMux) *http.ServeMux {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userAPIHandler := api.NewUserAPI(userService)

	credsRepo := repositories.NewCredentialRepository(db)
	credsService := services.NewCredentialService(userRepo, credsRepo)
	credsAPIHandler := api.NewCredentialAPI(credsService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productAPIHandler := api.NewProductAPI(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, productRepo)
	categoryAPIHandler := api.NewCategoryAPI(categoryService)

	MuxRoute(mux, "GET", "/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("_Synapsis Challenge_"))
	}))

	//! debug jwt token
	MuxRoute(mux, "GET", "/test-jwt", middleware.CheckSession(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("jwt passed"))
	})))

	MuxRoute(mux, "POST", "/v1/user/register", middleware.POST(http.HandlerFunc(userAPIHandler.Register)))
	MuxRoute(mux, "POST", "/v1/user/login", middleware.POST(http.HandlerFunc(credsAPIHandler.Login)))

	MuxRoute(mux, "GET", "/v1/user/show/products", middleware.GET(middleware.CheckSession(http.HandlerFunc(productAPIHandler.GetAllProducts))))
	MuxRoute(mux, "GET", "/v1/user/show/product/on", middleware.GET(middleware.CheckSession(http.HandlerFunc(productAPIHandler.GetProductsByID))), "?product_id=")
	MuxRoute(mux, "GET", "/v1/user/show/productsByCategory/on", middleware.GET(middleware.CheckSession(http.HandlerFunc(productAPIHandler.GetProductsByCategoryID))), "?category_id=")
	MuxRoute(mux, "POST", "/v1/admin/send/products", middleware.POST(middleware.CheckSession(http.HandlerFunc(productAPIHandler.StoreManyProducts))))

	MuxRoute(mux, "POST", "/v1/admin/send/categories", middleware.POST(middleware.CheckSession(http.HandlerFunc(categoryAPIHandler.StoreManyCategories))))
	MuxRoute(mux, "GET", "/v1/user/show/categories", middleware.GET(middleware.CheckSession(http.HandlerFunc(categoryAPIHandler.GetAllCategories))))
	MuxRoute(mux, "GET", "/v1/user/show/categoriesWithProducts", middleware.GET(middleware.CheckSession(http.HandlerFunc(categoryAPIHandler.GetAllCategoriesWithProducts))))

	return mux
}

func MuxRoute(mux *http.ServeMux, method string, path string, handler http.Handler, opt ...string) {
	if len(opt) > 0 {
		fmt.Printf("[%s]: %s %v \n", method, path, opt)
	} else {
		fmt.Printf("[%s]: %s \n", method, path)
	}

	mux.Handle(path, handler)
}

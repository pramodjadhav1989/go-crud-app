package router

import ("net/http")

var RegisterRout=func(){
	//Product
	http.HandleFunc("/Product/Manage",controllers.ManageProduct)
	http.HandleFunc("/Product/GetProduct/{ID}",controllers.GetProduct)
	http.HandleFunc("/Product/GetProducts",controllers.GetProducts)
}
package model

import (
"encoding/json"
"/pkg/config"
"fmt"
)
type Product struct{
	ID int `json:id` 
  PRODUCT_CODE string `json:productCode`
  NAME string `json:name`
  SHORTNAME string`json:shortName`
  CREATEDBY string`json:createdBy`
  CREATEDON string`json:createdOn`
  MODIFYBY string`json:modifyBy`
  MODIFYON string`json:modifyOn`
  ISDELETE int `json:isDelete`
}

var ProductList[] Product

Block{
	Try: func ManageProduct(p Product,process string) {
		
	},
	Catch: func(e Exception) {
		fmt.Printf("Caught %v\n", e)
	},
	Finally: func() {
		fmt.Println("Finally...")
	},
}.Do()




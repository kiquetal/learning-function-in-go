package main 

import ("fmt"
         "strings")
func processAndReplaceBadWords(body string) string {

	var badWords = []string{"kerfuffle", "sharbert", "fornax"}
	var newBody string
	fmt.Printf("Original word:%v\n",body)
	for _, badWord := range badWords {
                if len(newBody) < 1 {
		newBody = strings.Replace(body,badWord,"****",-1)
               } else {
               newBody = strings.Replace(newBody,badWord,"****",-1)

	       }

	}
	fmt.Printf("Altered :%v",newBody)
	return newBody

}

func main() {
  processAndReplaceBadWords("I kerfuffle because I fornax y need sharbert")
  fmt.Println("---")
    processAndReplaceBadWords("I drfuffle because I gfornamx y need sharbgert")

}

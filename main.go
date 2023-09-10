package main

import (
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Produto struct{
	nome, link string
	preco, avaliacao string
}

func main() {

	var produtos []Produto

	site := "https://www.amazon.com.br/s?k=suplementos&sprefix=suplem%2Caps%2C369&ref=nb_sb_ss_ts-doa-p_1_6"

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Deu merda, erro:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Deu merda! Código:", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		fmt.Println("Merda:", err)
	}

	doc.Find("div.a-section.a-spacing-small.puis-padding-left-small.puis-padding-right-small").Each(func(i int, s *goquery.Selection) {
		item := Produto{}

		item.nome= "aa"
		produtos = append(produtos, item)

		/* nome := s.Find("div.a-section.a-spacing-none.a-spacing-top-small.s-title-instructions-style").First()
		nome = nome.Find("h2")
		nome = nome.Find("a")
		nome = nome.Find("span")
		
		item.nome = nome.Text()

		//Nota
		nota := s.Find("div.a-section.a-spacing-none.a-spacing-top-micro")
		nota = nota.Find("div")
		nota = nota.Find("span").First()
		nota = nota.Find("span").First()

		
		item.avaliacao = nota.Text()

		//Link
		complemento := "https://www.amazon.com.br"
		endereco := s.Find("div.a-section.a-spacing-none.a-spacing-top-small.s-title-instructions-style")
 */
	})

}
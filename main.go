package main

import (
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Produto struct {
	nome, link, preco string
}

func pegarNome(pai *goquery.Selection) string {
	nome := pai.Find("div.a-section.a-spacing-none.a-spacing-top-small.s-title-instructions-style").First()
	nome = nome.Find("h2")
	nome = nome.Find("a")
	nome = nome.Find("span")

	valorNome := nome.Text()
	return valorNome
}

func pegarPreco(pai *goquery.Selection) (string, string) {
	precoContent := pai.Find("s-price-instructions-style")
	precoContent = precoContent.Find("a")
	linkProd, _ := precoContent.Attr("href")
	precoContent = precoContent.Find("span").First()
	precoContent = precoContent.Find("span")
	precoProd := precoContent.Text()

	return precoProd, linkProd

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
		//Nome
		item.nome = pegarNome(s)
		//FIXME
		item.preco, item.link = pegarPreco(s)

		produtos = append(produtos, item)

		fmt.Println("Informações produto:")
		fmt.Println("Nome:", item.nome)
		fmt.Println("Preço:", item.preco)
		fmt.Println("Link:", item.link)
	})

}

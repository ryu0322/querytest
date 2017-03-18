package main

import (
	"strconv"
	"os"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ChimeraCoder/anaconda"
	"github.com/peterh/liner"
	"github.com/PuerkitoBio/goquery"
)

type tweetinfo struct {
	conskey string		// ConsumerKey
	secretkey string	// ConsumerSecret
	accesskey string	// Access key
	accseckey string	// Access Secret
}

type yujiros struct {
	serifu string
}
func main() {
	app := cli.NewApp()
	app.Name = "tweet yujiro!"
	app.Usage = "yujiro`s meigen tweet!"
	app.HideHelp = true
	app.Flags = []cli.Flag{
			cli.IntFlag{Name:"number n", Value: 10, Usage: "number of output lines"},
	}

	app.Action = func(c *cli.Context) {
		err := mainRoutine(c.Int("number"))
		if err != nil {
			fmt.Printf("%s\r\n", err)
			return
		}
	}

	app.Run(os.Args)
}

func mainRoutine(lineval int)(error) {

	yujiro, err := readSerifu()

	if err != nil {
		return err
	}

	err = gotweet(yujiro)

	if err != nil {
		return err
	}

	return nil
}

func readSerifu()([]*yujiros, error) {
	yujiroarr := []*yujiros{}
	que, err := goquery.NewDocument("http://meigenkakugen.net/%E7%AF%84%E9%A6%AC%E5%8B%87%E6%AC%A1%E9%83%8E/")

	if err != nil {
		fmt.Println("%s", err)
		return nil, err
	}

	que.Find("div.graybox").Each(func(_ int, sel *goquery.Selection) {
			yuji := yujiros{ serifu: sel.Text(), }

			yujiroarr = append(yujiroarr, &yuji)
	})

	return yujiroarr, nil
}

func gotweet(yujiro []*yujiros) error {
	
	for cnt, yuji := range yujiro {
		outtext := ""

		if cnt > 10 {
			outtext = fmt.Sprintf("%d | %s", (cnt + 1 ), yuji.serifu )
		} else {
			outtext = fmt.Sprintf("%d  | %s", (cnt + 1 ), yuji.serifu )
		}

		fmt.Println(outtext)
	}

	line := liner.NewLiner()
	defer line.Close()

	li, err := line.Prompt("Tweet No?：")

	if err != nil {
		fmt.Println("input error[%s]", err)
		return err
	}

	idx, err := strconv.ParseInt(li, 10, 0)

	if err != nil {
		fmt.Println("input validation error!")
		return err
	}

	tweinf := tweetinfo{
		conskey: "kZO0WttiCtlaUnssDyVHrCE4F",
		secretkey: "wm3IQBUyd8AUE096RVzyWahtk84YzmtnAJW4paFuoErHmafuaL",
		accesskey: "784050923626598401-of3mQXAU6v3rs3jHafz4h8SNVGbxFxL",
		accseckey: "T7jsb8bfQ1YUNtNPPJRRhaQg8US7ntJTgtWhQ0zWf7lre",
	}

	anaconda.SetConsumerKey(tweinf.conskey)
	anaconda.SetConsumerSecret(tweinf.secretkey)
	tweapi := anaconda.NewTwitterApi(tweinf.accesskey, tweinf.accseckey)

	tweet, err := tweapi.PostTweet(yujiro[idx-1].serifu, nil)

	if err != nil {
		fmt.Println("tweet error!! [%s]", err)
		return err
	}

	fmt.Println(tweet)

	return nil
}

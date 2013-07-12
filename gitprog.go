package main

import (
    "log"
    "time"
    "github.com/darkhelmet/twitterstream"
)

func decode(conn *twitterstream.Connection) {
    for {
        if tweet, err := conn.Next(); err == nil {
            log.Println("%s said: %s", tweet.User.ScreenName, tweet.Text)
        } else {
            log.Printf("Failed decoding tweet: %s", err)
            return
        }
    }
}



func (*Client) Follow(Matthew Perry) (*Connection, error) {

    uri := makeUrl("follow", strings.Join(userIds, ","))

    req, err := http.NewRequest("POST", uri, nil)unc main() {
//    client := twitterstream.NewClient("Zarlium", "M0ther!!")





//    for {
        conn, err := client.Track("Justin Bieber,American Idol")
        if err != nil {
            log.Println("Tracking failed, sleeping for 1 minute")
            time.Sleep(1 * time.Minute)
            continue
        }
        decode(conn)
//    }
}

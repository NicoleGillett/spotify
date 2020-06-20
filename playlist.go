package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var userID = flag.String("user", "nicolegillett", "the Spotify user ID to look up")

func main() {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	client := spotify.Authenticator{}.NewClient(token)
	playlists, err := client.GetPlaylistsForUser("nicolegillett")
	if err != nil {
		log.Fatal(err)
	}
	firstPlaylistID := playlists.Playlists[0].ID
	tracks, err := client.GetPlaylistTracks(firstPlaylistID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tracks.Tracks[0].Track.Name)
	fmt.Println(tracks.Tracks[0].Track.Artists[0].Name)
}

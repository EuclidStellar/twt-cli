
package main

import (
    "bufio"
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

const (
    keysFile = "twitter_keys.json"
)

type TwitterKeys struct {
    ConsumerKey       string `json:"consumer_key"`
    ConsumerSecret    string `json:"consumer_secret"`
    AccessToken       string `json:"access_token"`
    AccessTokenSecret string `json:"access_token_secret"`
}

func main() {
    var keys TwitterKeys

    
    if _, err := os.Stat(keysFile); os.IsNotExist(err) {
        fmt.Println("No Twitter API keys found. Please enter your keys:")
        keys = getTwitterKeys()
       
        saveKeys(keys)
    } else {
        
        keys = loadKeys()
    }

  
    fmt.Println("\nEnter your tweet:")
    tweet := getUserInput()

   
    response := postTweet(keys, tweet)
    fmt.Println("\nResponse from Twitter API:")
    fmt.Println(response)
}

func getTwitterKeys() TwitterKeys {
    var keys TwitterKeys

    fmt.Println("Enter your Consumer Key:")
    keys.ConsumerKey = getUserInput()

    fmt.Println("Enter your Consumer Secret:")
    keys.ConsumerSecret = getUserInput()

    fmt.Println("Enter your Access Token:")
    keys.AccessToken = getUserInput()

    fmt.Println("Enter your Access Token Secret:")
    keys.AccessTokenSecret = getUserInput()

    return keys
}

func postTweet(keys TwitterKeys, tweet string) string {
    url := "https://api.twitter.com/2/tweets"
    method := "POST"

    payload := strings.NewReader(fmt.Sprintf(`{"text": "%s"}`, tweet))

    client := &http.Client{}

    req, err := http.NewRequest(method, url, payload)
    if err != nil {
        return fmt.Sprintf("Error creating request: %v", err)
    }

    req.Header.Add("Content-Type", "application/json")

    oauthHeader := generateOAuthHeader(method, url, keys.ConsumerKey, keys.ConsumerSecret, keys.AccessToken, keys.AccessTokenSecret)
    req.Header.Add("Authorization", oauthHeader)

    res, err := client.Do(req)
    if err != nil {
        return fmt.Sprintf("Error sending request: %v", err)
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return fmt.Sprintf("Error reading response body: %v", err)
    }

    return string(body)
}

func generateOAuthHeader(method, apiUrl, consumerKey, consumerSecret, accessToken, accessTokenSecret string) string {
    oauthNonce := generateNonce()
    oauthTimestamp := fmt.Sprintf("%d", time.Now().Unix())

    oauthParams := map[string]string{
        "oauth_consumer_key":     consumerKey,
        "oauth_nonce":            oauthNonce,
        "oauth_signature_method": "HMAC-SHA1",
        "oauth_timestamp":        oauthTimestamp,
        "oauth_token":            accessToken,
        "oauth_version":          "1.0",
    }


    allParams := url.Values{}
    for key, val := range oauthParams {
        allParams.Set(key, val)
    }

    baseString := method + "&" + url.QueryEscape(apiUrl) + "&" + url.QueryEscape(allParams.Encode())

    signingKey := url.QueryEscape(consumerSecret) + "&" + url.QueryEscape(accessTokenSecret)

    hmacHash := hmac.New(sha1.New, []byte(signingKey))
    hmacHash.Write([]byte(baseString))
    signature := base64.StdEncoding.EncodeToString(hmacHash.Sum(nil))

    oauthHeader := "OAuth "
    for key, val := range oauthParams {
        oauthHeader += key + "=\"" + val + "\", "
    }
    oauthHeader += "oauth_signature=\"" + url.QueryEscape(signature) + "\""

    return oauthHeader
}

func generateNonce() string {
    const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    nonce := make([]byte, 32)
    for i := range nonce {
        nonce[i] = letters[rand.Intn(len(letters))]
    }
    return string(nonce)
}

func getUserInput() string {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    return input
}

func saveKeys(keys TwitterKeys) {
    data, err := json.MarshalIndent(keys, "", "  ")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    err = ioutil.WriteFile(keysFile, data, 0644)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Twitter API keys saved successfully.")
}

func loadKeys() TwitterKeys {
    var keys TwitterKeys
    data, err := ioutil.ReadFile(keysFile)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    err = json.Unmarshal(data, &keys)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    return keys
}
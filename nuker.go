package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "net/url"
    "os"
    "strconv"
    "strings"
    "sync"
    "time"
)

// getProxy selects a random proxy from a list in a file
func getProxy() string {
    data, _ := ioutil.ReadFile("proxies.txt")
    proxies := strings.Split(string(data), "\n")
    return proxies[rand.Intn(len(proxies))]
}

// spam sends messages to a webhook, retrying on failure
func spam(webhook string, msg string, sleep time.Duration, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        proxy := getProxy()
        proxyUrl, _ := url.Parse("http://" + proxy)

        msgMap := map[string]string{"content": msg}
        msgJson, _ := json.Marshal(msgMap)

        proxyReq, _ := http.NewRequest("POST", webhook, bytes.NewBuffer(msgJson))
        proxyReq.Header.Add("Content-Type", "application/json")

        client := &http.Client{
            Transport: &http.Transport{
                Proxy: http.ProxyURL(proxyUrl),
            },
        }
        var resp *http.Response
        var err error
        for i := 0; i < 3; i++ {
            resp, err = client.Do(proxyReq)
            if err == nil {
                break
            }
            time.Sleep(time.Second * time.Duration(i+1))
        }
        if err != nil {
            fmt.Println("Error:", err)
            fmt.Println("\033[31mBad webhook or proxy, retrying...\033[0m")
        } else if resp.StatusCode == 204 {
            fmt.Println("\033[32mSent Message \033[0m")
        }
        time.Sleep(sleep)
    }
}

// main function to run the spam function with user inputs
func main() {
    fmt.Println("\033[31m" + `
    /$$   /$$                               /$$$$$$    /$$    /$$$$$$    /$$    /$$$$$$   /$$$$$$ 
    | $$  | $$                              /$$__  $$ /$$$$   /$$__  $$ /$$$$   /$$__  $$ /$$__  $$
    | $$  | $$  /$$$$$$$  /$$$$$$   /$$$$$$|__/  \ $$|_  $$  | $$  \ $$|_  $$  | $$  \ $$|__/  \ $$
    | $$  | $$ /$$_____/ /$$__  $$ /$$__  $$  /$$$$$/  | $$  |  $$$$$$$  | $$  |  $$$$$$/   /$$$$$/
    | $$  | $$|  $$$$$$ | $$$$$$$$| $$  \__/ |___  $$  | $$   \____  $$  | $$   >$$__  $$  |___  $$
    | $$  | $$ \____  $$| $$_____/| $$      /$$  \ $$  | $$   /$$  \ $$  | $$  | $$  \ $$ /$$  \ $$
    |  $$$$$$/ /$$$$$$$/|  $$$$$$$| $$     |  $$$$$$/ /$$$$$$|  $$$$$$/ /$$$$$$|  $$$$$$/|  $$$$$$/
    \______/ |_______/  \_______/|__/      \______/ |______/ \______/ |______/ \______/  \______/ 
                                                                                                
                                                                                                
    ╔═══════════════════════════════════════════════╗
    ║ User319183 | discord.gg/KHJjX3y2B4            ║
    ║ Discord Webhook Spammer                       ║
    ╚═══════════════════════════════════════════════╝
    ` + "\033[0m")
    reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[31mPlease Insert webhook URL: \033[0m")
	webhook, _ := reader.ReadString('\n')
	fmt.Print("\033[31mPlease Insert webhook Spam Message: \033[0m")
	msg, _ := reader.ReadString('\n')
	fmt.Print("\033[31mPlease Insert Threads (It's Recommended To Use 10): \033[0m")
	threadsStr, _ := reader.ReadString('\n')
	threads, _ := strconv.Atoi(strings.TrimSpace(threadsStr))
	fmt.Print("\033[31mPlease Insert Sleep (It's Recommended To Use 2): \033[0m")
	sleepStr, _ := reader.ReadString('\n')
	sleep, _ := strconv.Atoi(strings.TrimSpace(sleepStr))
	fmt.Println("\033[31mStarting...\033[0m")
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go spam(strings.TrimSpace(webhook), strings.TrimSpace(msg), time.Duration(sleep)*time.Second, &wg)
	}
	wg.Wait()
}
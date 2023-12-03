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
    "os/exec"
    "strconv"
    "strings"
    "sync"
    "time"
)

var proxies []string
var startTime time.Time

func init() {
    data, _ := ioutil.ReadFile("proxies.txt")
    proxies = strings.Split(string(data), "\n")
}

func getProxy() string {
    return proxies[rand.Intn(len(proxies))]
}

func setTitle(title string) {
    cmd := exec.Command("cmd", "/c", "title "+title)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func spam(webhook string, msg string, tts bool, username string, avatarURL string, sleep time.Duration, wg *sync.WaitGroup, success *int, fail *int, errors *int) {
    defer wg.Done()
    for {
        proxy := getProxy()
        proxyUrl, _ := url.Parse("http://" + proxy)

        msgMap := map[string]interface{}{"content": msg, "tts": tts}
        if username != "" {
            msgMap["username"] = username
        }
        if avatarURL != "" {
            msgMap["avatar_url"] = avatarURL
        }
        msgJson, _ := json.Marshal(msgMap)

        proxyReq, _ := http.NewRequest("POST", webhook, bytes.NewBuffer(msgJson))
        proxyReq.Header.Add("Content-Type", "application/json")

        client := &http.Client{
            Transport: &http.Transport{
                Proxy: http.ProxyURL(proxyUrl),
            },
        }
        resp, err := client.Do(proxyReq)
        if err != nil {
            *errors++
            fmt.Println("\033[31mError | " + err.Error() + "\033[0m")
        } else if resp.StatusCode == 204 {
            *success++
            fmt.Println("\033[32mSent Message | " + resp.Status + "\033[0m")
        } else if resp.StatusCode == 429 {
            retryAfter, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
            fmt.Println("\033[33mRate limited. Sleeping for " + strconv.Itoa(retryAfter) + "ms\033[0m")
            time.Sleep(time.Duration(retryAfter) * time.Millisecond)
        } else {
            *fail++
            fmt.Println("\033[31mFailed To Send Message | " + resp.Status + "\033[0m")
        }
        total := *success + *fail + *errors
        percent := float64(*success) / float64(total) * 100
        elapsed := time.Since(startTime).Minutes()
        rate := float64(*success) / elapsed
        rateRounded := fmt.Sprintf("%.1f", rate)
        setTitle(fmt.Sprintf("Success: %d, Fail: %d, Errors: %d, Success Rate: %s per minute @ %.2f%%, Elapsed Time: %.2f minutes", *success, *fail, *errors, rateRounded, percent, elapsed))
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
    ║ The fastest Discord Webhook Spammer           ║
    ╚═══════════════════════════════════════════════╝
    ` + "\033[0m")
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\033[31mPlease Insert webhook URLs (one per line, press enter when done): \033[0m")
    var webhooks []string
    for {
        webhook, _ := reader.ReadString('\n')
        webhook = strings.TrimSpace(webhook)
        if webhook == "" {
            break
        }
        webhooks = append(webhooks, webhook)
    }

    fmt.Print("\033[31mPlease Insert Custom Username (press enter to skip): \033[0m")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("\033[31mPlease Insert Avatar URL (press enter to skip): \033[0m")
    avatarURL, _ := reader.ReadString('\n')
    avatarURL = strings.TrimSpace(avatarURL)
    
    fmt.Print("\033[31mPlease Insert Message: \033[0m")
    msg, _ := reader.ReadString('\n')

    fmt.Print("\033[31mDo you want to send TTS messages? (yes/no): \033[0m")
    ttsStr, _ := reader.ReadString('\n')
    tts := strings.TrimSpace(ttsStr) == "yes"
    fmt.Print("\033[31mPlease Insert Threads (It's Recommended To Use 1): \033[0m")
    threadsStr, _ := reader.ReadString('\n')
    threads, _ := strconv.Atoi(strings.TrimSpace(threadsStr))
    fmt.Print("\033[31mPlease Insert Sleep in seconds (It's Recommended To Use 1.2): \033[0m")
    sleepStr, _ := reader.ReadString('\n')
    sleep, _ := strconv.Atoi(strings.TrimSpace(sleepStr))
    fmt.Println("\033[31mStarting...\033[0m")
    startTime = time.Now()
    var wg sync.WaitGroup
    success, fail, errors := 0, 0, 0
    for i := 0; i < threads; i++ {
        for _, webhook := range webhooks {
            wg.Add(1)
            go spam(webhook, strings.TrimSpace(msg), tts, username, avatarURL, time.Duration(sleep)*time.Second, &wg, &success, &fail, &errors)
        }
    }

    // New goroutine to update the title every second
    go func() {
        for {
            total := success + fail + errors
            percent := float64(success) / float64(total) * 100
            elapsed := time.Since(startTime).Minutes()
            rate := float64(success) / elapsed
            rateRounded := fmt.Sprintf("%.1f", rate)
            setTitle(fmt.Sprintf("Success: %d, Fail: %d, Errors: %d, Success Rate: %s per minute @ %.2f%%, Elapsed Time: %.2f minutes", success, fail, errors, rateRounded, percent, elapsed))
            time.Sleep(1 * time.Second)
        }
    }()

    wg.Wait()
}

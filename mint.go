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

const layout = "2006-1-2"

// 执行系统命令
func system(s string) {
    cmd := exec.Command("/bin/sh", "-c", s)
    var out bytes.Buffer

    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        fmt.Println("err", err)
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

// 将字符串转换成Time
func parser(str string) string {
    t, err := time.Parse(layout, str)
    if err != nil {
        fmt.Println(err)
        return time.Now().Format(time.RFC1123Z)
    }
    return t.Format(time.RFC1123Z)
}

func main() {
    date := flag.String("date", today(), "date")
    flag.Parse()

    var list []string
    list = strings.Split(*date, ",")
    fmt.Println(list)

    for _, item := range list {
        cur := parser(item)
        num := rand.Intn(10)
        for i := 0; i < num; i ++ {
            str := "echo '" + cur + strconv.Itoa(rand.Intn(1000000)) + "' > realwork.txt; git add realwork.txt; GIT_AUTHOR_DATE='" + cur + "' GIT_COMMITTER_DATE='" + cur + "' git commit -m 'update'"
            system(str)
        }
    }
    system("git push")
    system("git rm realwork.txt; git commit -m 'delete'; git push;")
}

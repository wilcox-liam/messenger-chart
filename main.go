package main

//go:generate go run main.go

import (
    "os"
    "fmt"
    "io/ioutil"
    "github.com/wcharczuk/go-chart"
    "encoding/json"  
    "time" 
)

const messagesPath = "/mnt/chromeos/removable/LaCie/messages/"

type Messages struct {
    Messages []Message `json:"messages"`
}

type Message struct {
    SenderName string `json:"sender_name"`
    Timestampms int64 `json:"timestamp_ms"`
}

func main() {
    var messages []Message
    messages = append (messages, loadJSON(messagesPath + "message_1.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_2.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_3.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_4.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_5.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_6.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_7.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_8.json")...)
    messages = append (messages, loadJSON(messagesPath + "message_9.json")...)

    //Count of messages send per date
    //var data  = make(map[string]int)
    //COunt of messages send per sender name
    var count = make(map[string]int)

    for i := 0; i < 2; i++ {
        count[messages[i].SenderName] += 1
        tm := time.Unix(messages[i].Timestampms, 0).UTC
        fmt.Println(tm.String())

        //data[messages[i].Timestampms] += 1
    }

    fmt.Println(count)

    graph := chart.BarChart{
        Title: "Test Bar Chart",
        Background: chart.Style{
            Padding: chart.Box{
                Top: 40,
            },
        },
        Height:   512,
        BarWidth: 60,
        Bars: []chart.Value{
            {Value: 5.25, Label: "Blue"},
            {Value: 4.88, Label: "Green"},
            {Value: 4.74, Label: "Gray"},
            {Value: 3.22, Label: "Orange"},
            {Value: 3, Label: "Test"},
            {Value: 2.27, Label: "??"},
            {Value: 1, Label: "!!"},
        },
    }

    f, _ := os.Create("output.png")
    defer f.Close()
    graph.Render(chart.PNG, f)
}

func loadJSON (fileName string) []Message {
        // Open our jsonFile
    jsonFile, err := os.Open(fileName)
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var messages Messages

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    err = json.Unmarshal(byteValue, &messages)

    return messages.Messages
}
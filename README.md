A simple Discord webhook sender for golang


**USAGE**

`goWebhook.CreateWebhook()`  
will return a webhook object that can then be changed to your liking

`goWebhook.AddField(title,value,inline)` 
takes the arguments of *title* (string), *value*(string), *inline* (boolean) 
will add a field to the webhook

`goWebhook.SendWebhook(webhookurl)` takes the arguments of *url*(string)
will return a bool dependent on success


**EXAMPLES**

```
package main

import "github.com/alexabrahall/goWebhook"

func main(){
  hook := goWebhook.CreateWebhook()
  hook.AddField("New Title","New Value,true)
  hook.SendWebook("https://discordapp.com/api/yourwebhook") // no tests to check if webhook was successful
  // or
  webhookReq, err := hook.SendWebook("https://discordapp.com/api/yourwebhook")    // use variable to check if post request was successful
  
  if err !=nil{
  fmt.Println("Unexpected error sending webhook!")      
  
  if webhookReq.StatusCode == 204 { //204 status is successful webhook post
  fmt.Println("Webhook sent")
  }else{
  fmt.Println("Webhook failed")
  fmt.Println(webhookReq.StatusCode)

}
```


**TODO**

-Improve error handling




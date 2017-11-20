# Chatbot Golang SPA 
[![Golang website](https://avatars1.githubusercontent.com/u/4314092?s=200&v=4)](https://golang.org/)
###### Martin Fennell, G00291266
## Introduction
This repo is for my project for the 3rd year module Data Representation and Quering, which I have been using Google [Go](https://golang.org/) to design a `single page webapp chatbot` to show my understanding of how data is passed around on the internet. The project will show my learning outcome and include technologies such as [HTTP](https://developer.mozilla.org/en-US/docs/Web/HTTP), [client/server]() and [AJAX](). The project requirements can be found [here](https://data-representation.github.io/problems/project.html).
#### Aims
The aims of the project are for me to learn and show my understanding of how to write and adhere to the standards and protocols of the web by using different data models and architectures used in modern web applications.
#### What is a chatbot
A [chatbot](https://en.wikipedia.org/wiki/Chatbot) is usualy a computer program which converses with a user either via audio or text. They are often designed to convincingly simulate how a human behaves in a conversation. Chatbots are typically used in dialog systems for various practical purposes including customer service or information acquisition. Some chatbots use sophisticated natural language processing systems, but many simpler systems scan for keywords within the input, then pull a reply with the most matching keywords, or the most similar wording pattern, from a database.
#### How to use the application
- Download Go from [`here`](https://golang.org/) and follow install instructions.
- Download a zip of repo or use `git clone`.
- Open `terminal`/`powershell` and `cd` to directory where the repo is downloaded.
- Type `go run main.go` in the command line.
- Open a web browser such as `chrome` and type `localhost:8080` in URL bar and enter.
- In the text box located at the footer type in a statement such as `hello`.
- To continue dialouge with eliza use some of the examples provided in the table beneath these instructions.
- When you want to stop the program from running type `bye`/`quit` to stop the program.

| User Input Examples | Responses Examples |
| ------ | ------ |
| Hello | Hi, could I start by getting your name? |
| My name is Martin | It's nice to meet you martin. How can I help? |
| I am not feeling myself | Did you come to me because you are not feeling yourself? |
| I'm sick | Why do you think you're sick? |
| Why can't I get better | Why don't you tell me the reason why can't you get better? |
| Because I'm not able to get better | Why do you think you're not able to get better? |
| Can you help | Why do you ask if I can help? |
| I would feel like I'm better but then I panic | Why do you think you're better but then you panic? |
| I feel like i can't be happy | How do you know you can't be happy? |
| I need something to help me through the day | Are you sure you need something to help you through the day? |
| My father was crule to me as a child | How do you feel about your father? |
| My mother has always favoured my other siblings | Tell me more about your mother. |
| I was bullied as a child | How do you think your childhood experiences relate to your feelings today? |
| Sorry for being an emotional wreck | There are many times when no apology is needed. |
|  | Am I getting the silent treatment? |
| Bye | Thank you, that will be $300. Have a good day! |

#### To view more info using curl (optional)
- Download and install `curl` from [here](https://curl.haxx.se/) and follow instruction on how to use it [here](https://curl.haxx.se/docs/manpage.html)
- Open `terminal` and `cd` to repo then type `curl -v localhost:8080`
- Use other curl commands to investigate more.

## Technologies
Below are a list of technologies that I have used for this project.
#### Go
Go is an open source programming language developed by in 2009 by Robert Griesemer, Rob Pike, Ken Thompson and many contributors from the open source community. It is a statically typed language and very similar to C except it has garbage collection and concurrent programming features added among other features. Below is an example of Golang code used in this project.
```Golang
func getRandomAnswer(answers []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(answers))
	return answers[index]
}
```
#### SPA
A single-page application (SPA) is a web application or web site that interacts with the user by dynamically rewriting the current page rather than loading entire new pages from a server. This approach avoids interruption of the user experience between successive pages, making the application behave more like a desktop application. In a SPA, either all necessary code such as HTML, JavaScript, and CSS is retrieved with a single page load or the appropriate resources are dynamically loaded and added to the page as necessary, usually in response to user actions. This means that the page doesnt have to be reloaded as the user interacts with it. Below is an example of how the SPA is served using Go and how the HTML page is dynamicly updated. 
```Golang
func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)
	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8080", nil)
}
```
The HTML unorderlist is empty at application start and has items added to is as the user converses with eliza.
```html
<ul class="list-group bg-light" id="output-area"></ul>
```
The List above this line is updated by The Jquery used below.
```javascript
$('#output-area').append('<li class="list-group-item">' + 'User: ' + $('#user-input').val() + '</li>');

$.get('/user-input', {value:$('#user-input').val()}).done(function(data) { 
    setTimeout(function(){
        $('#output-area').append('<li class="list-group-item">'+ 'Eliza: ' + data + '</li>');
    }, 500);
})
```
#### Client / Server
The client–server model is a distributed application structure that partitions tasks or workloads between the providers of a resource or service, called servers, and service requesters, called clients. Often clients and servers communicate over a computer network on separate hardware, A server host runs one or more server programs which share their resources with clients. A client does not share any of its resources, but requests a server's content or service function. Clients therefore initiate communication sessions with servers which await incoming requests.

![client/server](https://upload.wikimedia.org/wikipedia/commons/c/c9/Client-server-model.svg)
#### AJAX
AJAX stands for Asynchronous JavaScript And XML. In a nutshell, it is the use of the DOM to communicate with servers. It can send and receive information in various formats, including JSON, XML, HTML, and text files. AJAX’s most appealing characteristic is its `asynchronous` nature, which means it can communicate with the server, exchange data, and update the page without having to refresh the page. The two major features of AJAX allow you to do the following:
- Make requests to the server without reloading the page
- Receive and work with data from the server

Below is an example of how I used jQuery to use AJAX in the project for the conversation with the eliza clone,
```javascript
$("#user-input-form").submit(
    function(event) {
    event.preventDefault();
          
    if (document.forms[0].elements['user-input'].value=='') {
        $('#output-area').append('<li class="list-group-item">'+'User: '+$('#user-input').val()+'</li>');
        setTimeout(function(){
            $('#output-area').append('<li class="list-group-item">'+'Eliza: Am I getting the silent treatment?'+'</li>');
        }, 500);
    } else {
        $('#output-area').append('<li class="list-group-item">' +'User: '+$('#user-input').val() + '</li>');
            
        $.get('/user-input', {value:$('#user-input').val()}).done(function(data){
            setTimeout(function(){
                $('#output-area').append('<li class="list-group-item">'+'Eliza: '+data+ '</li>');
            }, 500);
        })
    }
    });
```
## Technologies not implimented
```
Json might be able to get this working
```
```
Herouku need to look into this further
```
```
HTML Template more work and wasnt exactly showing best use of AJAX
```
```
Firebase hoping to change/host the file in NoSQL db but may not work
```
## References
#### Ref Name
- [Ref]()
- [Ref]()
- [Ref]()

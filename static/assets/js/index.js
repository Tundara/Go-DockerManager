var successLoad = "<span>Successfully Loaded</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>";
var inProgress = "</i><span>In Progress </span><i class='fa-solid fa-cog fa-spin'></i></br></br>";

function Loaded() {
    document.getElementById("return_str").innerHTML = successLoad
}

function CreateContainer(Id) {
    //splitstr = imgName.split("/")
    document.getElementById("return_str").innerHTML = inProgress
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/container/create/"+Id);
    xhr.responseType = 'json';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        if (responseObj["status"] == "Created") {
            document.getElementById("return_str").innerHTML = "<span>Container successfully created</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
            window.onload
            GetContainersList()
            } else {
                document.getElementById("return_str").innerHTML = "<span>Container not started, image exist ?</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
            }
    };
}

function DeleteAllImages() {
    document.getElementById("return_str").innerHTML = inProgress
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/images/delete");
    xhr.responseType = 'json';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        if (xhr.status == 200) {
            ListAllImages()
            document.getElementById("return_str").innerHTML = "<span>All images deleted</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
        } else {
            document.getElementById("return_str").innerHTML = "<span>No image found or container still running</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
        }
    };
}

function DeleteAllContainers() {
    document.getElementById("return_str").innerHTML = inProgress
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/containers/delete");
    xhr.responseType = 'json';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        if (xhr.status == 200) {
            document.getElementById("return_str").innerHTML = "<span>All Containers deleted</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
            GetContainersList()
        } else {
            document.getElementById("return_str").innerHTML = "<span>No container running</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
        }
    };
}

function DeleteContainer(id) {
    document.getElementById("return_str").innerHTML = inProgress
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/container/delete/"+id);
    xhr.responseType = 'json';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        if (xhr.status == 200) {
            document.getElementById("return_str").innerHTML = "<span>Container Deleted</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
            GetContainersList()
        } else {
            document.getElementById("return_str").innerHTML = "<span>No container running</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
        }
    };
}

function GetContainersList() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/containers/list");
    xhr.responseType = 'text';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        var parse = JSON.parse(responseObj);
        document.getElementById("containers-list").innerHTML = "";
        if (xhr.status == 200) {
            console.log(parse)
            for (let i = 0; i < parse.data.length; i++) {
                var parsed = parse.data[i];
                document.getElementById("containers-list").innerHTML += "<div value="+parsed.Id+" class='listcontaineritem'>"+
                "<strong> Image Name : "+parsed.Names+"</strong></br>" +
                "<strong>"+parsed.Status+"</strong></br>"+
                "<button class='btnitems' onclick=OpenContainerInATerm('"+parsed.Id+"') >Open</button>&nbsp<button class='btnitems' onclick=DeleteContainer('"+parsed.Id+"') >Delete</button></div>"
            }
        } else {
            document.getElementById("containers-list").innerHTML += "<span>"+parse.data+"</span>"
        }
    };
}

function OpenContainerInATerm(id) {
    //contName = document.getElementById("inputCont").value
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/container/exec/"+id);
    xhr.responseType = 'text';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        if (xhr.status == 200) {
            window.open("/redirectToTerm", "_blank")
        }
    };
}

function PullImage() {
    document.getElementById("return_str").innerHTML = inProgress
    var imageName = document.getElementById("inputPull").value
    var split = imageName.split("/");
    console.log(split);
    if (split[1]) {
        let xhr = new XMLHttpRequest();
        xhr.open("GET", "/container/pull/"+imageName);
        xhr.responseType = 'text';
        xhr.send();
        xhr.onload = function() {
            let responseObj = xhr.response;
            if (xhr.status == 200) { 
                ListAllImages()
                document.getElementById("return_str").innerHTML = "<span>Successfully Pulled</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
            } else {
                document.getElementById("return_str").innerHTML = "<span>Not Pulled</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
            }
        };
    } else {
        document.getElementById("return_str").innerHTML = "<span>Bad syntax, please correct syntax, exemple : library/ubuntu</span></br></br>"
    }
}

function DeleteImageById(id) {
    document.getElementById("return_str").innerHTML = inProgress
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "/image/delete/"+id);
    xhr.responseType = 'text';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        console.log(responseObj)
        if (xhr.status == 200) {
            document.getElementById("return_str").innerHTML = "<span>Image deleted</span>&nbsp;<i class='fa-solid fa-check'></i></br></br>"
            ListAllImages()
        } else {
            document.getElementById("return_str").innerHTML = "<span>Delete container before delete image</span>&nbsp;<i class='fa-solid fa-xmark'></i></br></br>"
        }
    };
}

function GetAvatarImg(name) {
    var Names = name
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "/images/getavatar/"+name);
    xhr.responseType = 'text';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        console.log(responseObj)
        if (xhr.status == 200) {
            if (responseObj) {
                document.getElementById("img-"+Names).src = responseObj
            } else {
                document.getElementById("img-"+Names).src = "/assets/img/No_image_available.svg.png"
            }
        }
    };
}

function ListAllImages() {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/images/list");
    xhr.responseType = 'text';
    xhr.send();
    xhr.onload = function() {
        let responseObj = xhr.response;
        var parse = JSON.parse(responseObj);
        document.getElementById("images-list").innerHTML = "";
        if (xhr.status == 200) {
            console.log(parse)
            for (let i = 0; i < parse.data.length; i++) {
                var parsed = parse.data[i];
                var tags = parsed.RepoTags[0];
                var tostr = tags;
                var cut = tostr.split("/");
                var recut = cut[0].split(':')
                document.getElementById("images-list").innerHTML += "<div id="+parsed.Id+" class='listimgitem'>"+
                "<img id='img-"+recut[0]+"' src="+GetAvatarImg(recut[0])+"></img>"+
                "<strong>"+parsed.RepoTags+"</strong></br>" +
                "<button id='btnitems' onclick=DeleteImageById('"+parsed.Id+"') class='btnitems'>Delete</button>&nbsp<button id='btnstart' onclick=CreateContainer('"+parsed.Id+"') class='btnitems'>Start</button></div>"
            }   
        } else {
            document.getElementById("images-list").innerHTML += "<span>"+parse.data+"</span>"
        }
    };
}

GetContainersList()
ListAllImages()
Loaded()
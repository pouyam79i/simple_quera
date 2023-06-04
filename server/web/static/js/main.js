// Update this when you receive a token
myToken = null
myEmail = null
codeFile = null
inputFile = null

// Import files in html
function importText(eID, type) {
    var data = null
    let input = document.createElement('input');
    input.type = 'file';
    input.multiple = false;
    input.accept = type;
    input.onchange = _ => {
      // you can use this method to get file and perform respective operations
            var fr = new FileReader();
            fr.onload=function(){  
                displayItem(eID, fr.result);
            }
            fr.readAsText(input.files[0]);
            if(eID=='editor1'){
                codeFile = input.files[0];
            }else if (eID =='editor2'){
                inputFile = input.files[0];
            }
        };
    input.click();
}

function displayItem(eID, text){
    console.log("Loading data on ID:",eID,"\n", text);
    document.getElementById(eID).innerText = `${text}`;
}

// login form action here:
function singInUP() {

    console.log("Submit clicked")

    // reading email password
    let username = document.getElementById("email").value.trim().toLowerCase();
    let password = document.getElementById("psw").value.trim();

    if (username == "" || password == "") {
       console.log("Plz insert email:password properly")
    } else {
        if (!username.endsWith("@gmail.com")){
            console.log("Use a gmail to sign in")
            alert("Use a gmail to sign in!")
            return
        }
        if (username.length > 512){
            alert("Gmail with more than 512 char is not valid!")
            return
        }
        if (password.length > 512){
            alert("Password with more than 512 char is not valid!")
            return
        }
        console.log("Email:", username, "Password:", password)
        postSingInData(username, password);

    } 
}

// Data sender
async function postSingInData(username, password){
    try{

        let response = await fetch("http://localhost:8085/signin", {
            method: 'POST',
            headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin' : '*', 
            'Access-Control-Allow-Credentials' : true
            },
            body: `{
            "email": "${username}",
            "password": "${password}"
            }`,
            }).then(response=>response.json())
            .then(data=>{ 
                console.log(data.result);
                console.log(data.token);
                console.log(data.info);
                if(data.result){
                    myEmail = username;
                    myToken = data.token;
                    document.getElementById("close_sing_in").click();
                    singin_btn = document.getElementById("singin_btn");
                    singin_btn.style.pointerEvents="none";
                    singin_btn.innerText = username;
                    // TODO: add and enable sing out mechanism
                }else{
                    alert(data.info)
                    myToken = null
                    myEmail = null
                }
            
            });

    }catch(e){
        alert("Failed to Connect to Server")
        console.log("Failed to post: reason: ", e)
    }

}

// Upload file mechanism
function upload(){

    if (myToken == null){
        alert("Please sing in first!")
        return;
    }

    if (codeFile == null || inputFile == null){
        alert("Please import files properly!");
        return;
    }
    // uploading files
    crypto.randomUUID()
    uploadFiles()
}

async function uploadFiles(){

    if (myToken == null){
        alert("Please sing in before uploading!");
        return;
    }

    try{

        const formData = new FormData();
        formData.append("files", codeFile);
        formData.append("files", inputFile);

        let response = await fetch("http://localhost:8085/upload", {
            method: 'POST',
            headers: {
            'Content-Type' : 'multipart/form-data',
            'X-PM-TOKEN' : myToken
            },
            body: formData,
            }).then(response=>response.json())
            .then(data=>{ 
                console.log(data)
            });

    }catch(e){
        alert("Failed to Connect to Server")
        console.log("Failed to post. reason: ", e)
    }


}

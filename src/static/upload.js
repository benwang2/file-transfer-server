let fileInput = document.getElementById("file")
let urlInput = document.getElementById("url")
let dataInput = document.getElementById("data")
let data = {}

// let f_url = document.getElementById("f-url")
//       let f_data = document.getElementById("f-data")
//       upd = ()=>{
//         if (f_url.value != ""){
//           f_data.setAttribute('value','{"type":"link","destination":"'+f_url.value+'"}')
//         } else {
//           f_data.setAttribute('value','{"type":"file"}')
//         }
//       f_url.addEventListener('change', upd)
//       upd()

fileInput.addEventListener("change", (ev) => {
    if (ev.target.files.length > 0){
        let file = ev.target.files[0]
        data["type"] = "file"
        urlInput.setAttribute("placeholder",file["name"])
        dataInput.setAttribute("value", JSON.stringify(data))
        console.log(data)
    }
})

urlInput.addEventListener("change", (ev) => {
    urlInput.setAttribute("placeholder","http://example.org")
    data["type"] = "link"
    data["destination"] = urlInput.value
    dataInput.setAttribute("value", JSON.stringify(data))
    console.log(data)
})

data["destination"] = urlInput.value
dataInput.setAttribute("value", JSON.stringify(data))